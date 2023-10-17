// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GoogleworkspaceProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (g *jsiiProxy_GoogleworkspaceProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateGoogleworkspaceProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateGoogleworkspaceProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGoogleworkspaceProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateGoogleworkspaceProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func validateNewGoogleworkspaceProviderParameters(scope constructs.Construct, id *string, config *GoogleworkspaceProviderConfig) error {
	return nil
}

