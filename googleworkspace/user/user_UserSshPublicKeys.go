package user


type UserSshPublicKeys struct {
	// An SSH public key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#key User#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// An expiration time in microseconds since epoch.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#expiration_time_usec User#expiration_time_usec}
	ExpirationTimeUsec *string `field:"optional" json:"expirationTimeUsec" yaml:"expirationTimeUsec"`
}

