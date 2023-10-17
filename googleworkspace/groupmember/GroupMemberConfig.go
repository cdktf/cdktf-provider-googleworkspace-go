// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupmember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupMemberConfig struct {
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
	// The member's email address.
	//
	// A member can be a user or another group. This property is required when adding a member to a group. The email must be unique and cannot be an alias of another group. If the email address is changed, the API automatically reflects the email address changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_member#email GroupMember#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// Identifies the group in the API request.
	//
	// The value can be the group's email address, group alias, or the unique group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_member#group_id GroupMember#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// Defaults to `ALL_MAIL`.
	//
	// Defines mail delivery preferences of member. Acceptable values are:
	// 	- `ALL_MAIL`: All messages, delivered as soon as they arrive.
	// 	- `DAILY`: No more than one message a day.
	// 	- `DIGEST`: Up to 25 messages bundled into a single message.
	// 	- `DISABLED`: Remove subscription.
	// 	- `NONE`: No messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_member#delivery_settings GroupMember#delivery_settings}
	DeliverySettings *string `field:"optional" json:"deliverySettings" yaml:"deliverySettings"`
	// Defaults to `MEMBER`.
	//
	// The member's role in a group. The API returns an error for cycles in group memberships. For example, if group1 is a member of group2, group2 cannot be a member of group1. Acceptable values are:
	// 	- `MANAGER`: This role is only available if the Google Groups for Business is enabled using the Admin Console. A `MANAGER` role can do everything done by an `OWNER` role except make a member an `OWNER` or delete the group. A group can have multiple `MANAGER` members.
	// 	- `MEMBER`: This role can subscribe to a group, view discussion archives, and view the group's membership list.
	// 	- `OWNER`: This role can send messages to the group, add or remove members, change member roles, change group's settings, and delete the group. An OWNER must be a member of the group. A group can have more than one OWNER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_member#role GroupMember#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_member#timeouts GroupMember#timeouts}
	Timeouts *GroupMemberTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Defaults to `USER`.
	//
	// The type of group member. Acceptable values are:
	// 	- `CUSTOMER`: The member represents all users in a domain. An email address is not returned and the ID returned is the customer ID.
	// 	- `GROUP`: The member is another group.
	// 	- `USER`: The member is a user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_member#type GroupMember#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

