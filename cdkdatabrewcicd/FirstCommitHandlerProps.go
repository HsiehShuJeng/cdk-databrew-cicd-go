package cdkdatabrewcicd


type FirstCommitHandlerProps struct {
	// The branch name used in the CodeCommit repo.
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// The ARN of the CodeCommit repository.
	CodeCommitRepoArn *string `field:"required" json:"codeCommitRepoArn" yaml:"codeCommitRepoArn"`
	// The name of the CodeCommit repo.
	RepoName *string `field:"required" json:"repoName" yaml:"repoName"`
	// The name of the Lambda function which deals with first commit via AWS CodeCommit.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// The name of the IAM role for the Lambda function which deals with first commit via AWS CodeCommit.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
}

