// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package chromepolicy

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ChromePolicyPoliciesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ChromePolicyPoliciesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ChromePolicyPoliciesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ChromePolicyPoliciesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ChromePolicyPoliciesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ChromePolicyPoliciesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewChromePolicyPoliciesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

