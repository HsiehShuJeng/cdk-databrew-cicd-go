// A construct for AWS Glue DataBrew wtih CICD
package cdkdatabrewcicd

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-databrew-cicd.CodePipelineIamRole",
		reflect.TypeOf((*CodePipelineIamRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CodePipelineIamRole{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-databrew-cicd.CodePipelineIamRoleProps",
		reflect.TypeOf((*CodePipelineIamRoleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-databrew-cicd.DataBrewCodePipeline",
		reflect.TypeOf((*DataBrewCodePipeline)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "branchName", GoGetter: "BranchName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "codeCommitRepoArn", GoGetter: "CodeCommitRepoArn"},
			_jsii_.MemberProperty{JsiiProperty: "codePipelineArn", GoGetter: "CodePipelineArn"},
			_jsii_.MemberProperty{JsiiProperty: "firstStageArtifactName", GoGetter: "FirstStageArtifactName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "preproductionFunctionArn", GoGetter: "PreproductionFunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "productionFunctionArn", GoGetter: "ProductionFunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "repoName", GoGetter: "RepoName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataBrewCodePipeline{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-databrew-cicd.DataBrewCodePipelineProps",
		reflect.TypeOf((*DataBrewCodePipelineProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-databrew-cicd.FirstCommitHandler",
		reflect.TypeOf((*FirstCommitHandler)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "function", GoGetter: "Function"},
			_jsii_.MemberProperty{JsiiProperty: "functionName", GoGetter: "FunctionName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "roleName", GoGetter: "RoleName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_FirstCommitHandler{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-databrew-cicd.FirstCommitHandlerProps",
		reflect.TypeOf((*FirstCommitHandlerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-databrew-cicd.IamRole",
		reflect.TypeOf((*IamRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IamRole{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-databrew-cicd.IamRoleProps",
		reflect.TypeOf((*IamRoleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-databrew-cicd.InfraIamRole",
		reflect.TypeOf((*InfraIamRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_InfraIamRole{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-databrew-cicd.InfraIamRoleProps",
		reflect.TypeOf((*InfraIamRoleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-databrew-cicd.PreProductionLambda",
		reflect.TypeOf((*PreProductionLambda)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "function", GoGetter: "Function"},
			_jsii_.MemberProperty{JsiiProperty: "functionName", GoGetter: "FunctionName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "roleName", GoGetter: "RoleName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PreProductionLambda{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-databrew-cicd.PreProductionLambdaProps",
		reflect.TypeOf((*PreProductionLambdaProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-databrew-cicd.ProductionLambda",
		reflect.TypeOf((*ProductionLambda)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "function", GoGetter: "Function"},
			_jsii_.MemberProperty{JsiiProperty: "functionName", GoGetter: "FunctionName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "roleName", GoGetter: "RoleName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ProductionLambda{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-databrew-cicd.ProductionLambdaProps",
		reflect.TypeOf((*ProductionLambdaProps)(nil)).Elem(),
	)
}
