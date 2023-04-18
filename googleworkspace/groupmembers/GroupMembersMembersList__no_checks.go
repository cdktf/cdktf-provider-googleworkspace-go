//go:build no_runtime_type_checking

package groupmembers

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GroupMembersMembersList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GroupMembersMembersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GroupMembersMembersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GroupMembersMembersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GroupMembersMembersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GroupMembersMembersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGroupMembersMembersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

