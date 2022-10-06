package schema


type SchemaTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/schema#create Schema#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/schema#delete Schema#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/schema#update Schema#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

