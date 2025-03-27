package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2025"
	"github.com/rocktavious/autopilot/v2023"
)

func TestHasAPIErrorsWorks(t *testing.T) {
	// Arrange
	errs := []ol.Error{
		{Message: "can't be blank", Path: []string{"resource", "id"}},
		{Message: "is not a valid input", Path: []string{"id"}},
	}
	// Act
	output := ol.HasAPIErrors(errs)
	// Assert
	autopilot.Equals(t, true, ol.ErrIs(output, ol.ErrorAPIError))
	autopilot.Equals(t, `OpsLevel API Errors:
	- 'resource.id' can't be blank
	- 'id' is not a valid input`,
		output.Error())
}

func TestHasAPIErrorsNoPath(t *testing.T) {
	// Arrange
	errs := []ol.Error{
		{Message: "can't be blank", Path: []string{"base"}},
		{Message: "is not a valid input", Path: []string{""}},
	}
	// Act
	output := ol.HasAPIErrors(errs)
	// Assert
	autopilot.Equals(t, true, ol.ErrIs(output, ol.ErrorAPIError))
	autopilot.Equals(t, `OpsLevel API Errors:
	- '' can't be blank
	- '' is not a valid input`,
		output.Error())
}

func TestIsResourceFoundError(t *testing.T) {
	// Arrange
	err1 := ol.NewClientError(ol.ErrorResourceNotFound, "resource 'Example' not found")
	err2 := ol.NewClientError(ol.ErrorAPIError, "resource 'Example' not found")
	// Act
	// Assert
	autopilot.Equals(t, true, ol.ErrIs(err1, ol.ErrorResourceNotFound))
	autopilot.Equals(t, false, ol.ErrIs(err2, ol.ErrorResourceNotFound))
}
