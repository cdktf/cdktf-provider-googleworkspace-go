package schema

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SchemaConfig struct {
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
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/schema#fields Schema#fields}
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// The schema's name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/schema#schema_name Schema#schema_name}
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
	// Display name for the schema.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/schema#display_name Schema#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/schema#timeouts Schema#timeouts}
	Timeouts *SchemaTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

