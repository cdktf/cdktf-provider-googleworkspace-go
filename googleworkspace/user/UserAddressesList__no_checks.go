// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package user

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserAddressesList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserAddressesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserAddressesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_UserAddressesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserAddressesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserAddressesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserAddressesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

