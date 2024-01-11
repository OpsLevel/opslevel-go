package opslevel

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/taimoorgit/moredefaults"
	"gopkg.in/yaml.v3"
)

var structValidator = validator.New(validator.WithRequiredStructEnabled())

// Validates resource's `validate:""` struct tags
func IsResourceValid[T any](opslevelResource T) error {
	return structValidator.Struct(opslevelResource)
}

// Apply resource's `default:""` struct tags
func SetDefaultsFor[T any](opslevelResource *T, key ...string) {
	validator.New(validator.WithRequiredStructEnabled())
	if err := moredefaults.Set(opslevelResource); err != nil {
		panic(err)
	}
}

// Apply resource's `example:""` struct tags
func SetExamplesFor[T any](opslevelResource *T, key ...string) {
	validator.New(validator.WithRequiredStructEnabled())
	if err := moredefaults.Set(opslevelResource, "example"); err != nil {
		panic(err)
	}
}

// Make new OpsLevel resource with defaults set
func NewExampleOf[T any]() T {
	var opslevelResource T
	SetDefaultsFor[T](&opslevelResource, "example")
	return opslevelResource
}

// Get JSON formatted string of an OpsLevel resource - also sets defaults
func JsonOf[T any](opslevelResource T, key ...string) string {
	switch len(key) {
	case 0:
		SetDefaultsFor[T](&opslevelResource)
	case 1:
		SetDefaultsFor[T](&opslevelResource, key[0])
	default:
		panic("only one 'key' can be passed")
	}
	out, err := json.Marshal(opslevelResource)
	if err != nil {
		panic(err)
	}
	return string(out)
}

// Generate yaml formatted string of an OpsLevel resource - also sets defaults
func YamlOf[T any](opslevelResource T, key ...string) string {
	switch len(key) {
	case 0:
		SetDefaultsFor[T](&opslevelResource)
	case 1:
		SetDefaultsFor[T](&opslevelResource, key[0])
	default:
		panic("only one 'key' can be passed")
	}
	out, err := yaml.Marshal(opslevelResource)
	if err != nil {
		panic(err)
	}
	return string(out)
}
