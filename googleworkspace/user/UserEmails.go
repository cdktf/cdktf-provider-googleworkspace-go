// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package user


type UserEmails struct {
	// The type of the email account. Acceptable values: `custom`, `home`, `other`, `work`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#type User#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The user's email address.
	//
	// Also serves as the email ID. This value can be the user's primary email address or an alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#address User#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// If the value of type is custom, this property contains the custom type string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#custom_type User#custom_type}
	CustomType *string `field:"optional" json:"customType" yaml:"customType"`
	// Defaults to `false`. Indicates if this is the user's primary email. Only one entry can be marked as primary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#primary User#primary}
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
}

