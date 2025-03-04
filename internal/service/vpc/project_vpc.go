package vpc

import (
	"context"
	"log"
	"time"

	"github.com/aiven/aiven-go-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
)

var aivenProjectVPCSchema = map[string]*schema.Schema{
	"project": schemautil.CommonSchemaProjectReference,

	"cloud_name": {
		ForceNew:    true,
		Required:    true,
		Type:        schema.TypeString,
		Description: schemautil.Complex("Defines where the cloud provider and region where the service is hosted in. See the Service resource for additional information.").ForceNew().Build(),
	},
	"network_cidr": {
		ForceNew:    true,
		Required:    true,
		Type:        schema.TypeString,
		Description: "Network address range used by the VPC like 192.168.0.0/24",
	},
	"state": {
		Computed:    true,
		Type:        schema.TypeString,
		Description: schemautil.Complex("State of the VPC.").PossibleValues("APPROVED", "ACTIVE", "DELETING", "DELETED").Build(),
	},
}

func ResourceProjectVPC() *schema.Resource {
	return &schema.Resource{
		Description:   "The Project VPC resource allows the creation and management of Aiven Project VPCs.",
		CreateContext: resourceProjectVPCCreate,
		ReadContext:   resourceProjectVPCRead,
		DeleteContext: resourceProjectVPCDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: schemautil.DefaultResourceTimeouts(),

		Schema: aivenProjectVPCSchema,
	}
}

func resourceProjectVPCCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)
	projectName := d.Get("project").(string)
	vpc, err := client.VPCs.Create(
		projectName,
		aiven.CreateVPCRequest{
			CloudName:   d.Get("cloud_name").(string),
			NetworkCIDR: d.Get("network_cidr").(string),
		},
	)
	if err != nil {
		return diag.FromErr(err)
	}

	// Make sure the VPC is active before returning it because service creation, moving
	// service to VPC, and some other operations will fail unless the VPC is active
	waiter := ProjectVPCActiveWaiter{
		Client:  client,
		Project: projectName,
		VPCID:   vpc.ProjectVPCID,
	}

	_, err = waiter.Conf(d.Timeout(schema.TimeoutCreate)).WaitForStateContext(ctx)
	if err != nil {
		return diag.Errorf("error waiting for Aiven project VPC to be ACTIVE: %s", err)
	}

	d.SetId(schemautil.BuildResourceID(projectName, vpc.ProjectVPCID))

	return resourceProjectVPCRead(ctx, d, m)
}

func resourceProjectVPCRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	projectName, vpcID, err := schemautil.SplitResourceID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	vpc, err := client.VPCs.Get(projectName, vpcID)
	if err != nil {
		return diag.FromErr(schemautil.ResourceReadHandleNotFound(err, d))
	}

	err = copyVPCPropertiesFromAPIResponseToTerraform(d, vpc, projectName)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceProjectVPCDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	projectName, vpcID, err := schemautil.SplitResourceID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	waiter := ProjectVPCDeleteWaiter{
		Client:  client,
		Project: projectName,
		VPCID:   vpcID,
	}

	timeout := d.Timeout(schema.TimeoutDelete)
	_, err = waiter.Conf(timeout).WaitForStateContext(ctx)
	if err != nil {
		return diag.Errorf("error waiting for Aiven project VPC to be DELETED: %s", err)
	}

	return nil
}

func copyVPCPropertiesFromAPIResponseToTerraform(d *schema.ResourceData, vpc *aiven.VPC, project string) error {
	if err := d.Set("project", project); err != nil {
		return err
	}
	if err := d.Set("cloud_name", vpc.CloudName); err != nil {
		return err
	}
	if err := d.Set("network_cidr", vpc.NetworkCIDR); err != nil {
		return err
	}
	if err := d.Set("state", vpc.State); err != nil {
		return err
	}

	return nil
}

// ProjectVPCActiveWaiter is used to wait for VPC to enter active state. This check needs to be
// performed before creating a service that has a project VPC to ensure there has been sufficient
// time for other actions that update the state to have been completed
type ProjectVPCActiveWaiter struct {
	Client  *aiven.Client
	Project string
	VPCID   string
}

// RefreshFunc will call the Aiven client and refresh it's state.
func (w *ProjectVPCActiveWaiter) RefreshFunc() resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		vpc, err := w.Client.VPCs.Get(w.Project, w.VPCID)
		if err != nil {
			return nil, "", err
		}

		log.Printf("[DEBUG] Got %s state while waiting for VPC connection to be ACTIVE.", vpc.State)

		return vpc, vpc.State, nil
	}
}

// Conf sets up the configuration to refresh.
func (w *ProjectVPCActiveWaiter) Conf(timeout time.Duration) *resource.StateChangeConf {
	log.Printf("[DEBUG] Active waiter timeout %.0f minutes", timeout.Minutes())

	return &resource.StateChangeConf{
		Pending:    []string{"APPROVED", "DELETING", "DELETED"},
		Target:     []string{"ACTIVE"},
		Refresh:    w.RefreshFunc(),
		Delay:      10 * time.Second,
		Timeout:    timeout,
		MinTimeout: 2 * time.Second,
	}
}

// ProjectVPCDeleteWaiter is used to wait for VPC been deleted.
type ProjectVPCDeleteWaiter struct {
	Client  *aiven.Client
	Project string
	VPCID   string
}

// RefreshFunc will call the Aiven client and refresh it's state.
func (w *ProjectVPCDeleteWaiter) RefreshFunc() resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		vpc, err := w.Client.VPCs.Get(w.Project, w.VPCID)
		if err != nil {
			// might be already gone after deletion
			if aiven.IsNotFound(err) {
				return &aiven.VPC{}, "DELETED", nil
			}

			return nil, "", err
		}

		if vpc.State != "DELETING" && vpc.State != "DELETED" {
			err := w.Client.VPCs.Delete(w.Project, w.VPCID)
			if err != nil {
				if aiven.IsNotFound(err) {
					return vpc, "DELETED", nil
				}

				// VPC cannot be deleted while there are services migrating from
				// it or service deletion is still in progress
				if err.(aiven.Error).Status != 409 {
					return nil, "", err
				}
			}
		}

		log.Printf("[DEBUG] Got %s state while waiting for VPC connection to be DELETED.", vpc.State)

		return vpc, vpc.State, nil
	}
}

// Conf sets up the configuration to refresh.
func (w *ProjectVPCDeleteWaiter) Conf(timeout time.Duration) *resource.StateChangeConf {
	log.Printf("[DEBUG] Delete waiter timeout %.0f minutes", timeout.Minutes())

	return &resource.StateChangeConf{
		Pending:    []string{"APPROVED", "DELETING", "ACTIVE"},
		Target:     []string{"DELETED"},
		Refresh:    w.RefreshFunc(),
		Delay:      10 * time.Second,
		Timeout:    timeout,
		MinTimeout: 2 * time.Second,
	}
}
