// A construct for AWS Glue DataBrew wtih CICD
package cdkdatabrewcicd

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-databrew-cicd-go/cdkdatabrewcicd/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-databrew-cicd-go/cdkdatabrewcicd/v2/internal"
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

