// A construct for AWS Glue DataBrew wtih CICD
package cdkdatabrewcicd


type PreProductionLambdaProps struct {
	// The ARN of the S3 bucket for the DataBrew CICD pipeline.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The ARN of the IAM role in the pre-production account.
	PreproductionIamRoleArn *string `field:"required" json:"preproductionIamRoleArn" yaml:"preproductionIamRoleArn"`
	// The Lambda funciton name for the pre-production account.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// The name of the IAM role for the pre-produciton Lambda function.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
}

