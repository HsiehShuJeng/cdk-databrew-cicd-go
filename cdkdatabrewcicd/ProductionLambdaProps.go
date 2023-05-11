package cdkdatabrewcicd


type ProductionLambdaProps struct {
	// The ARN of the S3 bucket for the DataBrew CICD pipeline.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The ARN of the IAM role in the production account.
	ProductionIamRoleArn *string `field:"required" json:"productionIamRoleArn" yaml:"productionIamRoleArn"`
	// The Lambda funciton name for the production account.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// The name of the IAM role for the produciton Lambda function.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
}

