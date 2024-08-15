package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2024"
	"github.com/rocktavious/autopilot/v2023"
)

func TestFormatErrorsWorks(t *testing.T) {
	// Arrange
	errs := []ol.OpsLevelErrors{
		{Message: "can't be blank", Path: []string{"resource", "id"}},
		{Message: "is not a valid input", Path: []string{"id"}},
	}
	// Act
	output := ol.FormatErrors(errs)
	// Assert
	autopilot.Equals(t, `OpsLevel API Errors:
	- 'resource.id' can't be blank
	- 'id' is not a valid input`,
		output.Error())
}

func TestFormatErrorsNoPath(t *testing.T) {
	// Arrange
	errs := []ol.OpsLevelErrors{
		{Message: "can't be blank", Path: []string{"base"}},
		{Message: "is not a valid input", Path: []string{""}},
	}
	// Act
	output := ol.FormatErrors(errs)
	// Assert
	autopilot.Equals(t, `OpsLevel API Errors:
	- '' can't be blank
	- '' is not a valid input`,
		output.Error())
}
