package cdkdatabrewcicd

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-databrew-cicd-go/cdkdatabrewcicd/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-databrew-cicd-go/cdkdatabrewcicd/v2/internal"
)

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

	if err := validateNewDataBrewCodePipelineParameters(scope, name, props); err != nil {
		panic(err)
	}
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

	if err := validateDataBrewCodePipeline_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

