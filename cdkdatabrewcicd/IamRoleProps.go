package cdkdatabrewcicd


type IamRoleProps struct {
	// The ID of your infrastructure account.
	AccountID *string `field:"required" json:"accountID" yaml:"accountID"`
	// 'preproduction' or 'production'.
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// The role name.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
}

