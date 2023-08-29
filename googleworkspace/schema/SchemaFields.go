// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema


type SchemaFields struct {
	// The name of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/schema#field_name Schema#field_name}
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// The type of the field. Acceptable values are:  - `BOOL` - `DATE` - `DOUBLE` - `EMAIL` - `INT64` - `PHONE` - `STRING`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/schema#field_type Schema#field_type}
	FieldType *string `field:"required" json:"fieldType" yaml:"fieldType"`
	// Display Name of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/schema#display_name Schema#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Defaults to `true`. Boolean specifying whether the field is indexed or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/schema#indexed Schema#indexed}
	Indexed interface{} `field:"optional" json:"indexed" yaml:"indexed"`
	// Defaults to `false`. A boolean specifying whether this is a multi-valued field or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/schema#multi_valued Schema#multi_valued}
	MultiValued interface{} `field:"optional" json:"multiValued" yaml:"multiValued"`
	// numeric_indexing_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/schema#numeric_indexing_spec Schema#numeric_indexing_spec}
	NumericIndexingSpec *SchemaFieldsNumericIndexingSpec `field:"optional" json:"numericIndexingSpec" yaml:"numericIndexingSpec"`
	// Defaults to `ALL_DOMAIN_USERS`.
	//
	// Specifies who can view values of this field. See Retrieve users as a non-administrator for more information. Acceptable values are:
	// - `ADMINS_AND_SELF`
	// - `ALL_DOMAIN_USERS`
	// Note: It may take up to 24 hours for changes to this field to be reflected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/schema#read_access_type Schema#read_access_type}
	ReadAccessType *string `field:"optional" json:"readAccessType" yaml:"readAccessType"`
}

