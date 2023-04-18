package user


type UserOrganizations struct {
	// The type of organization. Acceptable values: `domain_only`, `school`, `unknown`, `work`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#type User#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The cost center of the user's organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#cost_center User#cost_center}
	CostCenter *string `field:"optional" json:"costCenter" yaml:"costCenter"`
	// If the value of type is custom, this property contains the custom value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#custom_type User#custom_type}
	CustomType *string `field:"optional" json:"customType" yaml:"customType"`
	// Specifies the department within the organization, such as sales or engineering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#department User#department}
	Department *string `field:"optional" json:"department" yaml:"department"`
	// The description of the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#description User#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The domain the organization belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#domain User#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The full-time equivalent millipercent within the organization (100000 = 100%).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#full_time_equivalent User#full_time_equivalent}
	FullTimeEquivalent *float64 `field:"optional" json:"fullTimeEquivalent" yaml:"fullTimeEquivalent"`
	// The physical location of the organization. This does not need to be a fully qualified address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#location User#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The name of the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#name User#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Indicates if this is the user's primary organization. A user may only have one primary organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#primary User#primary}
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// Text string symbol of the organization. For example, the text symbol for Google is GOOG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#symbol User#symbol}
	Symbol *string `field:"optional" json:"symbol" yaml:"symbol"`
	// The user's title within the organization. For example, member or engineer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#title User#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

