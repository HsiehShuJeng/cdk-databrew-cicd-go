// A construct for AWS Glue DataBrew wtih CICD
package cdkdatabrewcicd

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-databrew-cicd-go/cdkdatabrewcicd/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-databrew-cicd-go/cdkdatabrewcicd/v2/internal"
)

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

	if err := validateNewPreProductionLambdaParameters(scope, name, props); err != nil {
		panic(err)
	}
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

	if err := validatePreProductionLambda_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

