package role

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RoleConfig struct {
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
	// Name of the role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/role#name Role#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// privileges block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/role#privileges Role#privileges}
	Privileges interface{} `field:"required" json:"privileges" yaml:"privileges"`
	// A short description of the role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/role#description Role#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

