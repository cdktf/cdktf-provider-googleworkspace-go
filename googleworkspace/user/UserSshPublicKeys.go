package user


type UserSshPublicKeys struct {
	// An SSH public key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#key User#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// An expiration time in microseconds since epoch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#expiration_time_usec User#expiration_time_usec}
	ExpirationTimeUsec *string `field:"optional" json:"expirationTimeUsec" yaml:"expirationTimeUsec"`
}

