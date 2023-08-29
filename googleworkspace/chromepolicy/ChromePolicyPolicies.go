// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chromepolicy


type ChromePolicyPolicies struct {
	// The full qualified name of the policy schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/chrome_policy#schema_name ChromePolicy#schema_name}
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
	// JSON encoded map that represents key/value pairs that correspond to the given schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/chrome_policy#schema_values ChromePolicy#schema_values}
	SchemaValues *map[string]*string `field:"required" json:"schemaValues" yaml:"schemaValues"`
}

