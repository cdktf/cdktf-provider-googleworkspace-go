package role


type RolePrivileges struct {
	// The name of the privilege.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/role#privilege_name Role#privilege_name}
	PrivilegeName *string `field:"required" json:"privilegeName" yaml:"privilegeName"`
	// The obfuscated ID of the service this privilege is for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/role#service_id Role#service_id}
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
}

