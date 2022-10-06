package roleassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RoleAssignmentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The unique ID of the user this role is assigned to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/role_assignment#assigned_to RoleAssignment#assigned_to}
	AssignedTo *string `field:"required" json:"assignedTo" yaml:"assignedTo"`
	// The ID of the role that is assigned.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/role_assignment#role_id RoleAssignment#role_id}
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// If the role is restricted to an organization unit, this contains the ID for the organization unit the exercise of this role is restricted to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/role_assignment#org_unit_id RoleAssignment#org_unit_id}
	OrgUnitId *string `field:"optional" json:"orgUnitId" yaml:"orgUnitId"`
	// Defaults to `CUSTOMER`. The scope in which this role is assigned. Valid values are : - `CUSTOMER` - `ORG_UNIT`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/role_assignment#scope_type RoleAssignment#scope_type}
	ScopeType *string `field:"optional" json:"scopeType" yaml:"scopeType"`
}

