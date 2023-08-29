// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package domainalias

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DomainAliasConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/domain_alias#domain_alias_name DomainAlias#domain_alias_name}
	DomainAliasName *string `field:"required" json:"domainAliasName" yaml:"domainAliasName"`
	// The parent domain name that the domain alias is associated with.
	//
	// This can either be a primary or secondary domain name within a customer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/domain_alias#parent_domain_name DomainAlias#parent_domain_name}
	ParentDomainName *string `field:"optional" json:"parentDomainName" yaml:"parentDomainName"`
}

