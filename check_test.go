package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
	"github.com/shurcooL/graphql"
)

func TestCreateCheckServiceConfiguration(t *testing.T) {
	// Arrange
	client := ANewClient(t, "check/create_service_configuration")
	// Act
	result, err := client.CreateCheckServiceConfiguration(CheckServiceConfigurationCreateInput{
		Name:     "Hello World",
		Enabled:  true,
		Category: graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
		Level:    graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
		Notes:    "Hello World Check",
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", result.Id)
	autopilot.Equals(t, "Performance", result.Category.Name)
	autopilot.Equals(t, "Bronze", result.Level.Name)
}

func TestUpdateCheckServiceConfiguration(t *testing.T) {
	// Arrange
	client := ANewClient(t, "check/update_service_configuration")
	// Act
	result, err := client.UpdateCheckServiceConfiguration(CheckServiceConfigurationUpdateInput{
		Id:       graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
		Name:     "Hello World",
		Enabled:  true,
		Category: graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
		Level:    graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
		Notes:    "Hello World Check",
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", result.Id)
	autopilot.Equals(t, "Performance", result.Category.Name)
	autopilot.Equals(t, "Bronze", result.Level.Name)
}

func TestCreateCheckServiceOwnership(t *testing.T) {
	// Arrange
	client := ANewClient(t, "check/create_service_ownership")
	// Act
	result, err := client.CreateCheckServiceOwnership(CheckServiceOwnershipCreateInput{
		Name:     "Hello World",
		Enabled:  true,
		Category: graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
		Level:    graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
		Notes:    "Hello World Check",
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", result.Id)
	autopilot.Equals(t, "Performance", result.Category.Name)
	autopilot.Equals(t, "Bronze", result.Level.Name)
}

func TestUpdateCheckServiceOwnership(t *testing.T) {
	// Arrange
	client := ANewClient(t, "check/update_service_ownership")
	// Act
	result, err := client.UpdateCheckServiceOwnership(CheckServiceOwnershipUpdateInput{
		Id:       graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
		Name:     "Hello World",
		Enabled:  true,
		Category: graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
		Level:    graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
		Notes:    "Hello World Check",
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", result.Id)
	autopilot.Equals(t, "Performance", result.Category.Name)
	autopilot.Equals(t, "Bronze", result.Level.Name)
}

func TestCreateCheckServiceProperty(t *testing.T) {
	// Arrange
	client := ANewClient(t, "check/create_service_property")
	// Act
	result, err := client.CreateCheckServiceProperty(CheckServicePropertyCreateInput{
		Name:     "Hello World",
		Enabled:  true,
		Category: graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
		Level:    graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
		Notes:    "Hello World Check",
		Property: ServicePropertyFramework,
		Predicate: PredicateInput{
			Type:  PredicateTypeEquals,
			Value: "postgres",
		},
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", result.Id)
	autopilot.Equals(t, "Performance", result.Category.Name)
	autopilot.Equals(t, "Bronze", result.Level.Name)
}

func TestUpdateCheckServiceProperty(t *testing.T) {
	// Arrange
	client := ANewClient(t, "check/update_service_property")
	// Act
	result, err := client.UpdateCheckServiceProperty(CheckServicePropertyUpdateInput{
		Id:       graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
		Name:     "Hello World",
		Enabled:  true,
		Category: graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
		Level:    graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
		Notes:    "Hello World Check",
		Property: ServicePropertyFramework,
		Predicate: PredicateInput{
			Type:  PredicateTypeEquals,
			Value: "postgres",
		},
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", result.Id)
	autopilot.Equals(t, "Performance", result.Category.Name)
	autopilot.Equals(t, "Bronze", result.Level.Name)
}

func TestGetCheck(t *testing.T) {
	// Arrange
	client := ANewClient(t, "check/get")
	// Act
	result, err := client.GetCheck("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4")
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", result.Id)
	autopilot.Equals(t, "Performance", result.Category.Name)
	autopilot.Equals(t, "Bronze", result.Level.Name)
}

func TestGetMissingCheck(t *testing.T) {
	// Arrange
	client := ANewClient(t, "check/get_missing")
	// Act
	_, err := client.GetCheck("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf")
	// Assert
	autopilot.Assert(t, err != nil, "This test should throw an error.")
}

func TestDeleteCheck(t *testing.T) {
	// Arrange
	client := ANewClient(t, "check/delete")
	// Act
	err := client.DeleteCheck("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpHZW5lcmljLzIxNzI")
	// Assert
	autopilot.Equals(t, nil, err)
}
