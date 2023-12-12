package opslevel

import (
	"encoding/json"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

var structValidator = validator.New(validator.WithRequiredStructEnabled())

// Validates resource's `validate:""` struct tags
func IsResourceValidate[T any](opslevelResource T) error {
	return structValidator.Struct(opslevelResource)
}

// Apply resource's `default:""` struct tags
func SetDefaultsFor[T any](opslevelResource *T) {
	validator.New(validator.WithRequiredStructEnabled())
	if err := defaults.Set(opslevelResource); err != nil {
		panic(err)
	}
}

// Make new OpsLevel resource with defaults set
func NewExampleOf[T any]() T {
	var opslevelResource T
	SetDefaultsFor[T](&opslevelResource)
	return opslevelResource
}

// Get JSON formatted string of an OpsLevel resource - also sets defaults
func JsonOf[T any](opslevelResource T) string {
	SetDefaultsFor[T](&opslevelResource)
	out, err := json.Marshal(opslevelResource)
	if err != nil {
		panic(err)
	}
	return string(out)
}

// Generate yaml formatted string of an OpsLevel resource - also sets defaults
func YamlOf[T any](opslevelResource T) string {
	SetDefaultsFor[T](&opslevelResource)
	out, err := yaml.Marshal(opslevelResource)
	if err != nil {
		panic(err)
	}
	return string(out)
}
