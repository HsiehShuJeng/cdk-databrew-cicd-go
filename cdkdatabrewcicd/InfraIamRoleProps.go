package cdkdatabrewcicd


type InfraIamRoleProps struct {
	// The role name for the infrastructure account.
	// Default: 'CrossAccountRepositoryContributorRole'.
	//
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
}

