package cdkdatabrewcicd


type CodePipelineIamRoleProps struct {
	// The ARN of the S3 bucket where you store your artifacts.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The ARN of the Lambda function for the pre-production account.
	PreproductionLambdaArn *string `field:"required" json:"preproductionLambdaArn" yaml:"preproductionLambdaArn"`
	// The ARN of the Lambda function for the production account.
	ProductionLambdaArn *string `field:"required" json:"productionLambdaArn" yaml:"productionLambdaArn"`
	// The role name for the CodePipeline CICD pipeline.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
}

