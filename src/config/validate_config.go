package config

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

func ValidateConfig() {
	schemaLoader := gojsonschema.NewReferenceLoader("file://environment/config_validation.json")
	configLoader := gojsonschema.NewReferenceLoader("file://environment/config.json")
	result, err := gojsonschema.Validate(schemaLoader, configLoader)
	if err != nil {
		panic("config files ar not in place")
	}
	if !result.Valid() {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, err := range result.Errors() {
			fmt.Printf("- %s\n", err)
		}
		panic("document is not valid")
	}
}
