// A construct for AWS Glue DataBrew wtih CICD
package cdkdatabrewcicd


type InfraIamRoleProps struct {
	// The role name for the infrastructure account.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
}
