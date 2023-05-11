package cdkdatabrewcicd

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-databrew-cicd-go/cdkdatabrewcicd/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-databrew-cicd-go/cdkdatabrewcicd/v2/internal"
)

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

	if err := validateNewFirstCommitHandlerParameters(scope, name, props); err != nil {
		panic(err)
	}
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

	if err := validateFirstCommitHandler_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

