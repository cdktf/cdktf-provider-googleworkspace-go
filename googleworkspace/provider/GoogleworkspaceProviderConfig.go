package provider


type GoogleworkspaceProviderConfig struct {
	// A temporary [OAuth 2.0 access token] obtained from the Google Authorization server, i.e. the `Authorization: Bearer` token used to authenticate HTTP requests to Google Admin SDK APIs. This is an alternative to `credentials`, and ignores the `oauth_scopes` field. If both are specified, `access_token` will be used over the `credentials` field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs#access_token GoogleworkspaceProvider#access_token}
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs#alias GoogleworkspaceProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Either the path to or the contents of a service account key file in JSON format you can manage key files using the Cloud Console).
	//
	// If not provided, the application default credentials will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs#credentials GoogleworkspaceProvider#credentials}
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
	// The customer id provided with your Google Workspace subscription. It is found in the admin console under Account Settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs#customer_id GoogleworkspaceProvider#customer_id}
	CustomerId *string `field:"optional" json:"customerId" yaml:"customerId"`
	// The impersonated user's email with access to the Admin APIs can access the Admin SDK Directory API.
	//
	// `impersonated_user_email` is required for all services except group and user management.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs#impersonated_user_email GoogleworkspaceProvider#impersonated_user_email}
	ImpersonatedUserEmail *string `field:"optional" json:"impersonatedUserEmail" yaml:"impersonatedUserEmail"`
	// The list of the scopes required for your application (for a list of possible scopes, see [Authorize requests](https://developers.google.com/admin-sdk/directory/v1/guides/authorizing)).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs#oauth_scopes GoogleworkspaceProvider#oauth_scopes}
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
	// The service account used to create the provided `access_token` if authenticating using the `access_token` method and needing to impersonate a user.
	//
	// This service account will require the GCP role `Service Account Token Creator` if needing to impersonate a user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs#service_account GoogleworkspaceProvider#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
}

