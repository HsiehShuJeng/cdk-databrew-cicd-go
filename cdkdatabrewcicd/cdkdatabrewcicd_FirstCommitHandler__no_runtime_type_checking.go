//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// A construct for AWS Glue DataBrew wtih CICD
package cdkdatabrewcicd

// Building without runtime type checking enabled, so all the below just return nil

func validateFirstCommitHandler_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewFirstCommitHandlerParameters(scope constructs.Construct, name *string, props *FirstCommitHandlerProps) error {
	return nil
}

