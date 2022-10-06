//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package schema

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Schema) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Schema) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Schema) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Schema) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Schema) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Schema) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Schema) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Schema) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Schema) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Schema) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Schema) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Schema) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_Schema) validatePutFieldsParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_Schema) validatePutTimeoutsParameters(value *SchemaTimeouts) error {
	return nil
}

func validateSchema_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Schema) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Schema) validateSetDisplayNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Schema) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Schema) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Schema) validateSetSchemaNameParameters(val *string) error {
	return nil
}

func validateNewSchemaParameters(scope constructs.Construct, id *string, config *SchemaConfig) error {
	return nil
}

