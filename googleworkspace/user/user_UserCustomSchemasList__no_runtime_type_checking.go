//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package user

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserCustomSchemasList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserCustomSchemasList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserCustomSchemasList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_UserCustomSchemasList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserCustomSchemasList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserCustomSchemasList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserCustomSchemasListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

