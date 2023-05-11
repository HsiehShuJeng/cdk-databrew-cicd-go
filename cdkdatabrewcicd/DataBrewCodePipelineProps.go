package cdkdatabrewcicd


type DataBrewCodePipelineProps struct {
	// The ARN of the IAM role in the pre-production account.
	PreproductionIamRoleArn *string `field:"required" json:"preproductionIamRoleArn" yaml:"preproductionIamRoleArn"`
	// The ARN of the IAM role in the production account.
	ProductionIamRoleArn *string `field:"required" json:"productionIamRoleArn" yaml:"productionIamRoleArn"`
	// The name of the branch that will trigger the DataBrew CICD pipeline.
	BranchName *string `field:"optional" json:"branchName" yaml:"branchName"`
	// The name of the S3 bucket for the CodePipeline DataBrew CICD pipeline.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// the (required) name of the Artifact at the first stage.
	FirstStageArtifactName *string `field:"optional" json:"firstStageArtifactName" yaml:"firstStageArtifactName"`
	// The name of the CodePipeline Databrew CICD pipeline.
	PipelineName *string `field:"optional" json:"pipelineName" yaml:"pipelineName"`
	// The name of the CodeCommit repositroy for the DataBrew CICD pipeline.
	RepoName *string `field:"optional" json:"repoName" yaml:"repoName"`
}

