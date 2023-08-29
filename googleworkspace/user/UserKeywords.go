// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package user


type UserKeywords struct {
	// Each entry can have a type which indicates standard type of that entry.
	//
	// For example, keyword could be of type occupation or outlook. In addition to the standard type, an entry can have a custom type and can give it any name. Such types should have the CUSTOM value as type and also have a customType value. Acceptable values: `custom`, `mission`, `occupation`, `outlook`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#type User#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Keyword.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#value User#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Custom Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#custom_type User#custom_type}
	CustomType *string `field:"optional" json:"customType" yaml:"customType"`
}

