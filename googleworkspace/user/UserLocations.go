// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package user


type UserLocations struct {
	// The location type. Acceptable values: `custom`, `default`, `desk`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#type User#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Textual location.
	//
	// This is most useful for display purposes to concisely describe the location. For example, Mountain View, CA or Near Seattle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#area User#area}
	Area *string `field:"optional" json:"area" yaml:"area"`
	// Building identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#building_id User#building_id}
	BuildingId *string `field:"optional" json:"buildingId" yaml:"buildingId"`
	// If the location type is custom, this property contains the custom value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#custom_type User#custom_type}
	CustomType *string `field:"optional" json:"customType" yaml:"customType"`
	// Most specific textual code of individual desk location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#desk_code User#desk_code}
	DeskCode *string `field:"optional" json:"deskCode" yaml:"deskCode"`
	// Floor name/number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#floor_name User#floor_name}
	FloorName *string `field:"optional" json:"floorName" yaml:"floorName"`
	// Floor section.
	//
	// More specific location within the floor. For example, if a floor is divided into sections A, B, and C, this field would identify one of those values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#floor_section User#floor_section}
	FloorSection *string `field:"optional" json:"floorSection" yaml:"floorSection"`
}

