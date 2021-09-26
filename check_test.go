package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go"
	"github.com/rocktavious/autopilot"
)

var checkCreateInput = ol.CheckCreateInput{
	Name:     "Hello World",
	Enabled:  true,
	Category: ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
	Level:    ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
	Notes:    "Hello World Check",
}

var checkUpdateInput = ol.CheckUpdateInput{
	Id:       ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
	Name:     "Hello World",
	Enabled:  ol.Bool(true),
	Category: ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
	Level:    ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
	Notes:    "Hello World Check",
}

var testcases = map[string]struct {
	fixture string
	body    func(c *ol.Client) (*ol.Check, error)
}{
	"CreateCustomEvent": {
		fixture: "check/create_custom_event",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckCustomEvent(ol.CheckCustomEventCreateInput{
				CheckCreateInput: checkCreateInput,
				ServiceSelector:  ".metadata.name",
				SuccessCondition: ".metadata.name",
				Message:          "#Hello World",
				Integration:      ol.NewID("Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg"),
			})
		},
	},
	"UpdateCustomEvent": {
		fixture: "check/update_custom_event",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckCustomEvent(ol.CheckCustomEventUpdateInput{
				CheckUpdateInput: checkUpdateInput,
				ServiceSelector:  ".metadata.name",
				SuccessCondition: ".metadata.name",
				Message:          "#Hello World",
				Integration:      ol.NewID("Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg"),
			})
		},
	},
	"CreateManual": {
		fixture: "check/create_manual",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckManual(ol.CheckManualCreateInput{
				CheckCreateInput: checkCreateInput,
				UpdateFrequency:  ol.NewManualCheckFrequencyInput("2021-07-26T20:22:44.427Z", ol.FrequencyTimeScaleWeek, 1),
			})
		},
	},
	"UpdateManual": {
		fixture: "check/update_manual",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckManual(ol.CheckManualUpdateInput{
				CheckUpdateInput: checkUpdateInput,
				UpdateFrequency:  ol.NewManualCheckFrequencyInput("2021-07-26T20:22:44.427Z", ol.FrequencyTimeScaleWeek, 1),
			})
		},
	},
	"CreateRepositoryFile": {
		fixture: "check/create_repo_file",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckRepositoryFile(ol.CheckRepositoryFileCreateInput{
				CheckCreateInput: checkCreateInput,
				DirectorySearch:  true,
				Filepaths:        []string{"/src", "/test"},
				FileContentsPredicate: &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: "postgres",
				},
			})
		},
	},
	"UpdateRepositoryFile": {
		fixture: "check/update_repo_file",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckRepositoryFile(ol.CheckRepositoryFileUpdateInput{
				CheckUpdateInput: checkUpdateInput,
				DirectorySearch:  true,
				Filepaths:        []string{"/src", "/test"},
				FileContentsPredicate: &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: "postgres",
				},
			})
		},
	},
	"CreateRepositoryIntegrated": {
		fixture: "check/create_repo_integrated",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckRepositoryIntegrated(ol.CheckRepositoryIntegratedCreateInput{
				CheckCreateInput: checkCreateInput,
			})
		},
	},
	"UpdateRepositoryIntegrated": {
		fixture: "check/update_repo_integrated",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckRepositoryIntegrated(ol.CheckRepositoryIntegratedUpdateInput{
				CheckUpdateInput: checkUpdateInput,
			})
		},
	},
	"CreateRepositorySearch": {
		fixture: "check/create_repo_search",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckRepositorySearch(ol.CheckRepositorySearchCreateInput{
				CheckCreateInput: checkCreateInput,
				FileExtensions:   []string{"sbt", "py"},
				FileContentsPredicate: ol.PredicateInput{
					Type:  ol.PredicateTypeEnumContains,
					Value: "postgres",
				},
			})
		},
	},
	"UpdateRepositorySearch": {
		fixture: "check/update_repo_search",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckRepositorySearch(ol.CheckRepositorySearchUpdateInput{
				CheckUpdateInput: checkUpdateInput,
				FileExtensions:   []string{"sbt", "py"},
				FileContentsPredicate: &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumContains,
					Value: "postgres",
				},
			})
		},
	},
	"CreateServiceConfiguration": {
		fixture: "check/create_service_configuration",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckServiceConfiguration(ol.CheckServiceConfigurationCreateInput{
				CheckCreateInput: checkCreateInput,
			})
		},
	},
	"UpdateServiceConfiguration": {
		fixture: "check/update_service_configuration",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckServiceConfiguration(ol.CheckServiceConfigurationUpdateInput{
				CheckUpdateInput: checkUpdateInput,
			})
		},
	},
	"CreateServiceOwnership": {
		fixture: "check/create_service_ownership",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckServiceOwnership(ol.CheckServiceOwnershipCreateInput{
				CheckCreateInput: checkCreateInput,
			})
		},
	},
	"UpdateServiceOwnership": {
		fixture: "check/update_service_ownership",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckServiceOwnership(ol.CheckServiceOwnershipUpdateInput{
				CheckUpdateInput: checkUpdateInput,
			})
		},
	},
	"CreateServiceProperty": {
		fixture: "check/create_service_property",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckServiceProperty(ol.CheckServicePropertyCreateInput{
				CheckCreateInput: checkCreateInput,
				Property:         ol.ServicePropertyTypeEnumFramework,
				Predicate: &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: "postgres",
				},
			})
		},
	},
	"UpdateServiceProperty": {
		fixture: "check/update_service_property",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckServiceProperty(ol.CheckServicePropertyUpdateInput{
				CheckUpdateInput: checkUpdateInput,
				Property:         ol.ServicePropertyTypeEnumFramework,
				Predicate: &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: "postgres",
				},
			})
		},
	},
	"CreateTagDefined": {
		fixture: "check/create_tag_defined",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckTagDefined(ol.CheckTagDefinedCreateInput{
				CheckCreateInput: checkCreateInput,
				TagKey:           "app",
				TagPredicate: &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: "postgres",
				},
			})
		},
	},
	"UpdateTagDefined": {
		fixture: "check/update_tag_defined",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckTagDefined(ol.CheckTagDefinedUpdateInput{
				CheckUpdateInput: checkUpdateInput,
				TagKey:           "app",
				TagPredicate: &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: "postgres",
				},
			})
		},
	},
	"CreateToolUsage": {
		fixture: "check/create_tool_usage",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckToolUsage(ol.CheckToolUsageCreateInput{
				CheckCreateInput: checkCreateInput,
				ToolCategory:     ol.ToolCategoryMetrics,
				ToolNamePredicate: &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: "datadog",
				},
				EnvironmentPredicate: &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: "production",
				},
			})
		},
	},
	"UpdateToolUsage": {
		fixture: "check/update_tool_usage",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckToolUsage(ol.CheckToolUsageUpdateInput{
				CheckUpdateInput: checkUpdateInput,
				ToolCategory:     ol.ToolCategoryMetrics,
				ToolNamePredicate: &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: "datadog",
				},
				EnvironmentPredicate: &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: "production",
				},
			})
		},
	},
	"GetCheck": {
		fixture: "check/get",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.GetCheck("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4")
		},
	},
}

func TestChecks(t *testing.T) {
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			// Arrange
			client := ATestClient(t, tc.fixture)
			// Act
			result, err := tc.body(client)
			// Assert
			autopilot.Equals(t, nil, err)
			autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", result.Id)
			autopilot.Equals(t, "Performance", result.Category.Name)
			autopilot.Equals(t, "Bronze", result.Level.Name)
		})
	}
}

func TestListChecks(t *testing.T) {
	// Arrange
	client := ATestClient(t, "check/list")
	// Act
	result, err := client.ListChecks()
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Metrics Tool", result[2].Name)
	autopilot.Equals(t, "Tier 1 Services", result[2].Filter.Name)
}

func TestGetMissingCheck(t *testing.T) {
	// Arrange
	client := ATestClient(t, "check/get_missing")
	// Act
	_, err := client.GetCheck("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf")
	// Assert
	autopilot.Assert(t, err != nil, "This test should throw an error.")
}

func TestDeleteCheck(t *testing.T) {
	// Arrange
	client := ATestClient(t, "check/delete")
	// Act
	err := client.DeleteCheck("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpHZW5lcmljLzIxNzI")
	// Assert
	autopilot.Equals(t, nil, err)
}
