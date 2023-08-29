// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagoogleworkspacegroupmembers

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleworkspaceGroupMembersConfig struct {
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
	// Identifies the group in the API request.
	//
	// The value can be the group's email address, group alias, or the unique group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/data-sources/group_members#group_id DataGoogleworkspaceGroupMembers#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// If true, lists indirect group memberships.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/data-sources/group_members#include_derived_membership DataGoogleworkspaceGroupMembers#include_derived_membership}
	IncludeDerivedMembership interface{} `field:"optional" json:"includeDerivedMembership" yaml:"includeDerivedMembership"`
}

