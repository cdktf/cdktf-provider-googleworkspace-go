package schema


type SchemaFieldsNumericIndexingSpec struct {
	// Maximum value of this field.
	//
	// This is meant to be indicative rather than enforced. Values outside this range will still be indexed, but search may not be as performant.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/schema#max_value Schema#max_value}
	MaxValue *float64 `field:"optional" json:"maxValue" yaml:"maxValue"`
	// Minimum value of this field.
	//
	// This is meant to be indicative rather than enforced. Values outside this range will still be indexed, but search may not be as performant.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/schema#min_value Schema#min_value}
	MinValue *float64 `field:"optional" json:"minValue" yaml:"minValue"`
}

