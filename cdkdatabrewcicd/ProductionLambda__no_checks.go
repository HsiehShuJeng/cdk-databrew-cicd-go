//go:build no_runtime_type_checking

package cdkdatabrewcicd

// Building without runtime type checking enabled, so all the below just return nil

func validateProductionLambda_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewProductionLambdaParameters(scope constructs.Construct, name *string, props *ProductionLambdaProps) error {
	return nil
}

