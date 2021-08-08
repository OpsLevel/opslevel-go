package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go"
	"github.com/rocktavious/autopilot"
)

var testcases = map[string]struct {
	fixture string
	body    func(c *ol.Client) (*ol.Check, error)
}{
	"CreateCustom": {
		fixture: "check/create_custom",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckCustom(ol.CheckCustomCreateInput{
				Name:     "Hello World",
				Enabled:  true,
				Category: ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
			})
		},
	},
	"UpdateCustom": {
		fixture: "check/update_custom",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckCustom(ol.CheckCustomUpdateInput{
				Id:       ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:     "Hello World",
				Enabled:  ol.Bool(true),
				Category: ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
			})
		},
	},
	"CreateCustomEvent": {
		fixture: "check/create_custom_event",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckCustomEvent(ol.CheckCustomEventCreateInput{
				Name:             "Hello World",
				Enabled:          true,
				Category:         ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:            ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:            "Hello World Check",
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
				Id:               ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:             "Hello World",
				Enabled:          ol.Bool(true),
				Category:         ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:            ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:            "Hello World Check",
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
				Name:            "Hello World",
				Enabled:         true,
				Category:        ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:           ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:           "Hello World Check",
				UpdateFrequency: ol.NewManualCheckFrequencyInput("2021-07-26T20:22:44.427Z", ol.FrequencyTimeScaleWeek, 1),
			})
		},
	},
	"UpdateManual": {
		fixture: "check/update_manual",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckManual(ol.CheckManualUpdateInput{
				Id:              ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:            "Hello World",
				Enabled:         ol.Bool(true),
				Category:        ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:           ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:           "Hello World Check",
				UpdateFrequency: ol.NewManualCheckFrequencyInput("2021-07-26T20:22:44.427Z", ol.FrequencyTimeScaleWeek, 1),
			})
		},
	},
	"CreatePayload": {
		fixture: "check/create_payload",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckPayload(ol.CheckPayloadCreateInput{
				Name:         "Hello World",
				Enabled:      true,
				Category:     ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:        ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:        "Hello World Check",
				JQExpression: ".metadata.name",
				Message:      "#Hello World",
			})
		},
	},
	"UpdatePayload": {
		fixture: "check/update_payload",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckPayload(ol.CheckPayloadUpdateInput{
				Id:           ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:         "Hello World",
				Enabled:      ol.Bool(true),
				Category:     ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:        ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:        "Hello World Check",
				JQExpression: ".metadata.name",
				Message:      "#Hello World",
			})
		},
	},
	"CreateRepositoryFile": {
		fixture: "check/create_repo_file",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckRepositoryFile(ol.CheckRepositoryFileCreateInput{
				Name:            "Hello World",
				Enabled:         true,
				Category:        ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:           ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:           "Hello World Check",
				DirectorySearch: true,
				Filepaths:       []string{"/src", "/test"},
				FileContentsPredicate: ol.PredicateInput{
					Type:  ol.PredicateTypeEquals,
					Value: "postgres",
				},
			})
		},
	},
	"UpdateRepositoryFile": {
		fixture: "check/update_repo_file",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckRepositoryFile(ol.CheckRepositoryFileUpdateInput{
				Id:              ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:            "Hello World",
				Enabled:         ol.Bool(true),
				Category:        ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:           ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:           "Hello World Check",
				DirectorySearch: true,
				Filepaths:       []string{"/src", "/test"},
				FileContentsPredicate: ol.PredicateInput{
					Type:  ol.PredicateTypeEquals,
					Value: "postgres",
				},
			})
		},
	},
	"CreateRepositoryIntegrated": {
		fixture: "check/create_repo_integrated",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckRepositoryIntegrated(ol.CheckRepositoryIntegratedCreateInput{
				Name:     "Hello World",
				Enabled:  true,
				Category: ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
			})
		},
	},
	"UpdateRepositoryIntegrated": {
		fixture: "check/update_repo_integrated",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckRepositoryIntegrated(ol.CheckRepositoryIntegratedUpdateInput{
				Id:       ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:     "Hello World",
				Enabled:  ol.Bool(true),
				Category: ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
			})
		},
	},
	"CreateRepositorySearch": {
		fixture: "check/create_repo_search",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckRepositorySearch(ol.CheckRepositorySearchCreateInput{
				Name:           "Hello World",
				Enabled:        true,
				Category:       ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:          ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:          "Hello World Check",
				FileExtensions: []string{"sbt", "py"},
				FileContentsPredicate: ol.PredicateInput{
					Type:  ol.PredicateTypeContains,
					Value: "postgres",
				},
			})
		},
	},
	"UpdateRepositorySearch": {
		fixture: "check/update_repo_search",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckRepositorySearch(ol.CheckRepositorySearchUpdateInput{
				Id:             ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:           "Hello World",
				Enabled:        ol.Bool(true),
				Category:       ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:          ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:          "Hello World Check",
				FileExtensions: []string{"sbt", "py"},
				FileContentsPredicate: ol.PredicateInput{
					Type:  ol.PredicateTypeContains,
					Value: "postgres",
				},
			})
		},
	},
	"CreateServiceConfiguration": {
		fixture: "check/create_service_configuration",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckServiceConfiguration(ol.CheckServiceConfigurationCreateInput{
				Name:     "Hello World",
				Enabled:  true,
				Category: ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
			})
		},
	},
	"UpdateServiceConfiguration": {
		fixture: "check/update_service_configuration",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckServiceConfiguration(ol.CheckServiceConfigurationUpdateInput{
				Id:       ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:     "Hello World",
				Enabled:  ol.Bool(true),
				Category: ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
			})
		},
	},
	"CreateServiceOwnership": {
		fixture: "check/create_service_ownership",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckServiceOwnership(ol.CheckServiceOwnershipCreateInput{
				Name:     "Hello World",
				Enabled:  true,
				Category: ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
			})
		},
	},
	"UpdateServiceOwnership": {
		fixture: "check/update_service_ownership",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckServiceOwnership(ol.CheckServiceOwnershipUpdateInput{
				Id:       ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:     "Hello World",
				Enabled:  ol.Bool(true),
				Category: ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
			})
		},
	},
	"CreateServiceProperty": {
		fixture: "check/create_service_property",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckServiceProperty(ol.CheckServicePropertyCreateInput{
				Name:     "Hello World",
				Enabled:  true,
				Category: ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
				Property: ol.ServicePropertyFramework,
				Predicate: ol.PredicateInput{
					Type:  ol.PredicateTypeEquals,
					Value: "postgres",
				},
			})
		},
	},
	"UpdateServiceProperty": {
		fixture: "check/update_service_property",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckServiceProperty(ol.CheckServicePropertyUpdateInput{
				Id:       ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:     "Hello World",
				Enabled:  ol.Bool(true),
				Category: ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
				Property: ol.ServicePropertyFramework,
				Predicate: ol.PredicateInput{
					Type:  ol.PredicateTypeEquals,
					Value: "postgres",
				},
			})
		},
	},
	"CreateTagDefined": {
		fixture: "check/create_tag_defined",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckTagDefined(ol.CheckTagDefinedCreateInput{
				Name:     "Hello World",
				Enabled:  true,
				Category: ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
				TagKey:   "app",
				TagPredicate: ol.PredicateInput{
					Type:  ol.PredicateTypeEquals,
					Value: "postgres",
				},
			})
		},
	},
	"UpdateTagDefined": {
		fixture: "check/update_tag_defined",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckTagDefined(ol.CheckTagDefinedUpdateInput{
				Id:       ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:     "Hello World",
				Enabled:  ol.Bool(true),
				Category: ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
				TagKey:   "app",
				TagPredicate: ol.PredicateInput{
					Type:  ol.PredicateTypeEquals,
					Value: "postgres",
				},
			})
		},
	},
	"CreateToolUsage": {
		fixture: "check/create_tool_usage",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.CreateCheckToolUsage(ol.CheckToolUsageCreateInput{
				Name:         "Hello World",
				Enabled:      true,
				Category:     ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:        ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:        "Hello World Check",
				ToolCategory: ol.ToolCategoryMetrics,
				ToolNamePredicate: ol.PredicateInput{
					Type:  ol.PredicateTypeEquals,
					Value: "datadog",
				},
				EnvironmentPredicate: ol.PredicateInput{
					Type:  ol.PredicateTypeEquals,
					Value: "production",
				},
			})
		},
	},
	"UpdateToolUsage": {
		fixture: "check/update_tool_usage",
		body: func(c *ol.Client) (*ol.Check, error) {
			return c.UpdateCheckToolUsage(ol.CheckToolUsageUpdateInput{
				Id:           ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:         "Hello World",
				Enabled:      ol.Bool(true),
				Category:     ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:        ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:        "Hello World Check",
				ToolCategory: ol.ToolCategoryMetrics,
				ToolNamePredicate: ol.PredicateInput{
					Type:  ol.PredicateTypeEquals,
					Value: "datadog",
				},
				EnvironmentPredicate: ol.PredicateInput{
					Type:  ol.PredicateTypeEquals,
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
