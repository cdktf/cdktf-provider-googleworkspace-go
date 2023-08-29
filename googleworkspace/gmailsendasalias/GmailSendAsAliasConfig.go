// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gmailsendasalias

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GmailSendAsAliasConfig struct {
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
	// User's primary email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/gmail_send_as_alias#primary_email GmailSendAsAlias#primary_email}
	PrimaryEmail *string `field:"required" json:"primaryEmail" yaml:"primaryEmail"`
	// The email address that appears in the 'From:' header for mail sent using this alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/gmail_send_as_alias#send_as_email GmailSendAsAlias#send_as_email}
	SendAsEmail *string `field:"required" json:"sendAsEmail" yaml:"sendAsEmail"`
	// A name that appears in the 'From:' header for mail sent using this alias.
	//
	// For custom 'from' addresses, when this is empty, Gmail will populate the 'From:' header with the name that is used for the primary address associated with the account. If the admin has disabled the ability for users to update their name format, requests to update this field for the primary login will silently fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/gmail_send_as_alias#display_name GmailSendAsAlias#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Whether this address is selected as the default 'From:' address in situations such as composing a new message or sending a vacation auto-reply.
	//
	// Every Gmail account has exactly one default send-as address, so the only legal value that clients may write to this field is true. Changing this from false to true for an address will result in this field becoming false for the other previous default address. Toggling an existing alias' default to false is not possible, another alias must be added/imported and toggled to true to remove the default from an existing alias. To avoid drift with Terraform, please change the previous default's config to false AFTER a new default is applied and perform a refresh to synchronize with remote state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/gmail_send_as_alias#is_default GmailSendAsAlias#is_default}
	IsDefault interface{} `field:"optional" json:"isDefault" yaml:"isDefault"`
	// An optional email address that is included in a 'Reply-To:' header for mail sent using this alias.
	//
	// If this is empty, Gmail will not generate a 'Reply-To:' header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/gmail_send_as_alias#reply_to_address GmailSendAsAlias#reply_to_address}
	ReplyToAddress *string `field:"optional" json:"replyToAddress" yaml:"replyToAddress"`
	// An optional HTML signature that is included in messages composed with this alias in the Gmail web UI.
	//
	// This signature is added to new emails only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/gmail_send_as_alias#signature GmailSendAsAlias#signature}
	Signature *string `field:"optional" json:"signature" yaml:"signature"`
	// smtp_msa block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/gmail_send_as_alias#smtp_msa GmailSendAsAlias#smtp_msa}
	SmtpMsa *GmailSendAsAliasSmtpMsa `field:"optional" json:"smtpMsa" yaml:"smtpMsa"`
	// Defaults to `true`.
	//
	// Whether Gmail should treat this address as an alias for the user's primary email address. This setting only applies to custom 'from' aliases. See https://support.google.com/a/answer/1710338 for help on making this decision
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/gmail_send_as_alias#treat_as_alias GmailSendAsAlias#treat_as_alias}
	TreatAsAlias interface{} `field:"optional" json:"treatAsAlias" yaml:"treatAsAlias"`
}

