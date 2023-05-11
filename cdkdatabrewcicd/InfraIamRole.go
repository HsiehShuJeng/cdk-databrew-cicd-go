package cdkdatabrewcicd

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-databrew-cicd-go/cdkdatabrewcicd/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-databrew-cicd-go/cdkdatabrewcicd/v2/internal"
)

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

	if err := validateNewInfraIamRoleParameters(scope, name, props); err != nil {
		panic(err)
	}
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

	if err := validateInfraIamRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

