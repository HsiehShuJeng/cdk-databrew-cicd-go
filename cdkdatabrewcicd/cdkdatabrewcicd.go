// A construct for AWS Glue DataBrew wtih CICD
package cdkdatabrewcicd

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-databrew-cicd-go/cdkdatabrewcicd/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-databrew-cicd-go/cdkdatabrewcicd/internal"
)

type CodePipelineIamRole interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// The representative of the IAM role for the CodePipeline CICD pipeline.
	Role() awsiam.Role
	// The ARN of the IAM role for the CodePipeline CICD pipeline.
	RoleArn() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CodePipelineIamRole
type jsiiProxy_CodePipelineIamRole struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CodePipelineIamRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipelineIamRole) Role() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipelineIamRole) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}


func NewCodePipelineIamRole(scope constructs.Construct, name *string, props *CodePipelineIamRoleProps) CodePipelineIamRole {
	_init_.Initialize()

	j := jsiiProxy_CodePipelineIamRole{}

	_jsii_.Create(
		"cdk-databrew-cicd.CodePipelineIamRole",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

func NewCodePipelineIamRole_Override(c CodePipelineIamRole, scope constructs.Construct, name *string, props *CodePipelineIamRoleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-databrew-cicd.CodePipelineIamRole",
		[]interface{}{scope, name, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CodePipelineIamRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-databrew-cicd.CodePipelineIamRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodePipelineIamRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

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

type DataBrewCodePipeline interface {
	constructs.Construct
	// The name of the branch that will trigger the DataBrew CICD pipeline.
	BranchName() *string
	// The ARN of the S3 bucket for the CodePipeline DataBrew CICD pipeline.
	BucketArn() *string
	// The ARN of the CodeCommit repository.
	CodeCommitRepoArn() *string
	// The ARN of the DataBrew CICD pipeline.
	CodePipelineArn() *string
	// the (required) name of the Artifact at the first stage.
	FirstStageArtifactName() *string
	// The tree node.
	Node() constructs.Node
	// The ARN of the Lambda function for the pre-production account.
	PreproductionFunctionArn() *string
	// The ARN of the Lambda function for the production account.
	ProductionFunctionArn() *string
	// The name of the CodeCommit repositroy for the DataBrew CICD pipeline.
	RepoName() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for DataBrewCodePipeline
type jsiiProxy_DataBrewCodePipeline struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_DataBrewCodePipeline) BranchName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataBrewCodePipeline) BucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataBrewCodePipeline) CodeCommitRepoArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeCommitRepoArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataBrewCodePipeline) CodePipelineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codePipelineArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataBrewCodePipeline) FirstStageArtifactName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstStageArtifactName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataBrewCodePipeline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataBrewCodePipeline) PreproductionFunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preproductionFunctionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataBrewCodePipeline) ProductionFunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productionFunctionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataBrewCodePipeline) RepoName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoName",
		&returns,
	)
	return returns
}


func NewDataBrewCodePipeline(scope constructs.Construct, name *string, props *DataBrewCodePipelineProps) DataBrewCodePipeline {
	_init_.Initialize()

	j := jsiiProxy_DataBrewCodePipeline{}

	_jsii_.Create(
		"cdk-databrew-cicd.DataBrewCodePipeline",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

func NewDataBrewCodePipeline_Override(d DataBrewCodePipeline, scope constructs.Construct, name *string, props *DataBrewCodePipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-databrew-cicd.DataBrewCodePipeline",
		[]interface{}{scope, name, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DataBrewCodePipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-databrew-cicd.DataBrewCodePipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataBrewCodePipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

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

type FirstCommitHandler interface {
	constructs.Construct
	// The representative of Lambda function which deals with first commit via AWS CodeCommit.
	Function() awslambda.IFunction
	// The name of the Lambda function which deals with first commit via AWS CodeCommit.
	FunctionName() *string
	// The tree node.
	Node() constructs.Node
	// The name of the IAM role for the Lambda function which deals with first commit via AWS CodeCommit.
	RoleName() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for FirstCommitHandler
type jsiiProxy_FirstCommitHandler struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_FirstCommitHandler) Function() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"function",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirstCommitHandler) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirstCommitHandler) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirstCommitHandler) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}


func NewFirstCommitHandler(scope constructs.Construct, name *string, props *FirstCommitHandlerProps) FirstCommitHandler {
	_init_.Initialize()

	j := jsiiProxy_FirstCommitHandler{}

	_jsii_.Create(
		"cdk-databrew-cicd.FirstCommitHandler",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

func NewFirstCommitHandler_Override(f FirstCommitHandler, scope constructs.Construct, name *string, props *FirstCommitHandlerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-databrew-cicd.FirstCommitHandler",
		[]interface{}{scope, name, props},
		f,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func FirstCommitHandler_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-databrew-cicd.FirstCommitHandler",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirstCommitHandler) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

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

// IAM Role.
//
// Defines an IAM role for pre-production and production AWS accounts.
type IamRole interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// The ARN of the IAM role for pre-production or production.
	RoleArn() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for IamRole
type jsiiProxy_IamRole struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_IamRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamRole) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}


func NewIamRole(scope constructs.Construct, name *string, props *IamRoleProps) IamRole {
	_init_.Initialize()

	j := jsiiProxy_IamRole{}

	_jsii_.Create(
		"cdk-databrew-cicd.IamRole",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

func NewIamRole_Override(i IamRole, scope constructs.Construct, name *string, props *IamRoleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-databrew-cicd.IamRole",
		[]interface{}{scope, name, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func IamRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-databrew-cicd.IamRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IamRoleProps struct {
	// The ID of your infrastructure account.
	AccountID *string `field:"required" json:"accountID" yaml:"accountID"`
	// 'preproduction' or 'production'.
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// The role name.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
}

type InfraIamRole interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// The ARN of the IAM role for the infrastructure account.
	RoleArn() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for InfraIamRole
type jsiiProxy_InfraIamRole struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_InfraIamRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraIamRole) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}


func NewInfraIamRole(scope constructs.Construct, name *string, props *InfraIamRoleProps) InfraIamRole {
	_init_.Initialize()

	j := jsiiProxy_InfraIamRole{}

	_jsii_.Create(
		"cdk-databrew-cicd.InfraIamRole",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

func NewInfraIamRole_Override(i InfraIamRole, scope constructs.Construct, name *string, props *InfraIamRoleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-databrew-cicd.InfraIamRole",
		[]interface{}{scope, name, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func InfraIamRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-databrew-cicd.InfraIamRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InfraIamRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type InfraIamRoleProps struct {
	// The role name for the infrastructure account.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
}

type PreProductionLambda interface {
	constructs.Construct
	// The representative of Lambda function for the pre-production account.
	Function() awslambda.IFunction
	// The Lambda funciton name for the pre-production account.
	FunctionName() *string
	// The tree node.
	Node() constructs.Node
	// The name of the IAM role for the pre-produciton Lambda function.
	RoleName() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for PreProductionLambda
type jsiiProxy_PreProductionLambda struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_PreProductionLambda) Function() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"function",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreProductionLambda) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreProductionLambda) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreProductionLambda) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}


func NewPreProductionLambda(scope constructs.Construct, name *string, props *PreProductionLambdaProps) PreProductionLambda {
	_init_.Initialize()

	j := jsiiProxy_PreProductionLambda{}

	_jsii_.Create(
		"cdk-databrew-cicd.PreProductionLambda",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

func NewPreProductionLambda_Override(p PreProductionLambda, scope constructs.Construct, name *string, props *PreProductionLambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-databrew-cicd.PreProductionLambda",
		[]interface{}{scope, name, props},
		p,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func PreProductionLambda_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-databrew-cicd.PreProductionLambda",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreProductionLambda) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

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

type ProductionLambda interface {
	constructs.Construct
	// The representative of Lambda function for the production account.
	Function() awslambda.IFunction
	// The Lambda funciton name for the production account.
	FunctionName() *string
	// The tree node.
	Node() constructs.Node
	// The name of the IAM role for the produciton Lambda function.
	RoleName() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ProductionLambda
type jsiiProxy_ProductionLambda struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ProductionLambda) Function() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"function",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductionLambda) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductionLambda) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductionLambda) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}


func NewProductionLambda(scope constructs.Construct, name *string, props *ProductionLambdaProps) ProductionLambda {
	_init_.Initialize()

	j := jsiiProxy_ProductionLambda{}

	_jsii_.Create(
		"cdk-databrew-cicd.ProductionLambda",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

func NewProductionLambda_Override(p ProductionLambda, scope constructs.Construct, name *string, props *ProductionLambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-databrew-cicd.ProductionLambda",
		[]interface{}{scope, name, props},
		p,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ProductionLambda_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-databrew-cicd.ProductionLambda",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProductionLambda) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

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

