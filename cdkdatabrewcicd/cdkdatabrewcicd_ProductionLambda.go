// A construct for AWS Glue DataBrew wtih CICD
package cdkdatabrewcicd

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-databrew-cicd-go/cdkdatabrewcicd/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-databrew-cicd-go/cdkdatabrewcicd/v2/internal"
)

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

	if err := validateNewProductionLambdaParameters(scope, name, props); err != nil {
		panic(err)
	}
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

	if err := validateProductionLambda_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

