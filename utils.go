package opslevel

import (
	"github.com/go-playground/validator/v10"
	"github.com/opslevel/moredefaults"
)

var structValidator = validator.New(validator.WithRequiredStructEnabled())

// IsResourceValid runs validator.Validate on all `validate` struct tags
func IsResourceValid[T any](opslevelResource T) error {
	return structValidator.Struct(opslevelResource)
}

// SetDefaultsFor applies all `default` struct tags
func SetDefaultsFor[T any](opslevelResource *T) {
	validator.New(validator.WithRequiredStructEnabled())
	if err := moredefaults.Set(opslevelResource); err != nil {
		panic(err)
	}
}

// SetExamplesFor applies all `example` struct tags
func SetExamplesFor[T any](opslevelResource *T) {
	validator.New(validator.WithRequiredStructEnabled())
	if err := moredefaults.Set(opslevelResource, "example"); err != nil {
		panic(err)
	}
}

// NewExampleOf makes a new OpsLevel resource with all `example` struct tags applied
func NewExampleOf[T any]() T {
	var opslevelResource T
	SetExamplesFor[T](&opslevelResource)
	return opslevelResource
}
