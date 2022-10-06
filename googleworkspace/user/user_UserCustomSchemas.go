package user


type UserCustomSchemas struct {
	// The name of the schema.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#schema_name User#schema_name}
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
	// JSON encoded map that represents key/value pairs that correspond to the given schema.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#schema_values User#schema_values}
	SchemaValues *map[string]*string `field:"required" json:"schemaValues" yaml:"schemaValues"`
}

