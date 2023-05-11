//go:build no_runtime_type_checking

package cdkdatabrewcicd

// Building without runtime type checking enabled, so all the below just return nil

func validateDataBrewCodePipeline_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDataBrewCodePipelineParameters(scope constructs.Construct, name *string, props *DataBrewCodePipelineProps) error {
	return nil
}

