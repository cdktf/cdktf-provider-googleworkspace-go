package datagoogleworkspacedomainalias

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleworkspaceDomainAliasConfig struct {
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
	// The domain alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/d/domain_alias#domain_alias_name DataGoogleworkspaceDomainAlias#domain_alias_name}
	DomainAliasName *string `field:"required" json:"domainAliasName" yaml:"domainAliasName"`
}

