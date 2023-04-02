//go:build no_runtime_type_checking

// A construct for AWS Glue DataBrew wtih CICD
package cdkdatabrewcicd

// Building without runtime type checking enabled, so all the below just return nil

func validatePreProductionLambda_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPreProductionLambdaParameters(scope constructs.Construct, name *string, props *PreProductionLambdaProps) error {
	return nil
}

