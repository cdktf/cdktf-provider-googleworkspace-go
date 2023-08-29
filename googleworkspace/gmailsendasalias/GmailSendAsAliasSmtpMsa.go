// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gmailsendasalias


type GmailSendAsAliasSmtpMsa struct {
	// The hostname of the SMTP service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/gmail_send_as_alias#host GmailSendAsAlias#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// The port of the SMTP service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/gmail_send_as_alias#port GmailSendAsAlias#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The password that will be used for authentication with the SMTP service.
	//
	// This is a write-only field that can be specified in requests to create or update SendAs settings; it is never populated in responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/gmail_send_as_alias#password GmailSendAsAlias#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Defaults to `securityModeUnspecified`. The protocol that will be used to secure communication with the SMTP service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/gmail_send_as_alias#security_mode GmailSendAsAlias#security_mode}
	SecurityMode *string `field:"optional" json:"securityMode" yaml:"securityMode"`
	// The username that will be used for authentication with the SMTP service.
	//
	// This is a write-only field that can be specified in requests to create or update SendAs settings; it is never populated in responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/gmail_send_as_alias#username GmailSendAsAlias#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

