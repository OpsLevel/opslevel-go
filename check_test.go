package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
	"github.com/shurcooL/graphql"
)

var testcases = map[string]struct {
	fixture string
	body    func(c *Client) (*Check, error)
}{
	"CreateCustom": {
		fixture: "check/create_custom",
		body: func(c *Client) (*Check, error) {
			return c.CreateCheckCustom(CheckCustomCreateInput{
				Name:     "Hello World",
				Enabled:  true,
				Category: graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
			})
		},
	},
	"UpdateCustom": {
		fixture: "check/update_custom",
		body: func(c *Client) (*Check, error) {
			return c.UpdateCheckCustom(CheckCustomUpdateInput{
				Id:       graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:     "Hello World",
				Enabled:  true,
				Category: graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
			})
		},
	},
	"CreateCustomEvent": {
		fixture: "check/create_custom_event",
		body: func(c *Client) (*Check, error) {
			return c.CreateCheckCustomEvent(CheckCustomEventCreateInput{
				Name:             "Hello World",
				Enabled:          true,
				Category:         graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:            graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:            "Hello World Check",
				ServiceSelector:  ".metadata.name",
				SuccessCondition: ".metadata.name",
				Message:          "#Hello World",
				Integration:      graphql.ID("Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg"),
			})
		},
	},
	"UpdateCustomEvent": {
		fixture: "check/update_custom_event",
		body: func(c *Client) (*Check, error) {
			return c.UpdateCheckCustomEvent(CheckCustomEventUpdateInput{
				Id:               graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:             "Hello World",
				Enabled:          true,
				Category:         graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:            graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:            "Hello World Check",
				ServiceSelector:  ".metadata.name",
				SuccessCondition: ".metadata.name",
				Message:          "#Hello World",
				Integration:      graphql.ID("Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg"),
			})
		},
	},
	"CreateManual": {
		fixture: "check/create_manual",
		body: func(c *Client) (*Check, error) {
			return c.CreateCheckManual(CheckManualCreateInput{
				Name:            "Hello World",
				Enabled:         true,
				Category:        graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:           graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:           "Hello World Check",
				UpdateFrequency: NewManualCheckFrequencyInput("2021-07-26T20:22:44.427Z", FrequencyTimeScaleWeek, 1),
			})
		},
	},
	"UpdateManual": {
		fixture: "check/update_manual",
		body: func(c *Client) (*Check, error) {
			return c.UpdateCheckManual(CheckManualUpdateInput{
				Id:              graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:            "Hello World",
				Enabled:         true,
				Category:        graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:           graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:           "Hello World Check",
				UpdateFrequency: NewManualCheckFrequencyInput("2021-07-26T20:22:44.427Z", FrequencyTimeScaleWeek, 1),
			})
		},
	},
	"CreatePayload": {
		fixture: "check/create_payload",
		body: func(c *Client) (*Check, error) {
			return c.CreateCheckPayload(CheckPayloadCreateInput{
				Name:         "Hello World",
				Enabled:      true,
				Category:     graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:        graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:        "Hello World Check",
				JQExpression: ".metadata.name",
				Message:      "#Hello World",
			})
		},
	},
	"UpdatePayload": {
		fixture: "check/update_payload",
		body: func(c *Client) (*Check, error) {
			return c.UpdateCheckPayload(CheckPayloadUpdateInput{
				Id:           graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:         "Hello World",
				Enabled:      true,
				Category:     graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:        graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:        "Hello World Check",
				JQExpression: ".metadata.name",
				Message:      "#Hello World",
			})
		},
	},
	"CreateRepositoryFile": {
		fixture: "check/create_repo_file",
		body: func(c *Client) (*Check, error) {
			return c.CreateCheckRepositoryFile(CheckRepositoryFileCreateInput{
				Name:            "Hello World",
				Enabled:         true,
				Category:        graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:           graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:           "Hello World Check",
				DirectorySearch: true,
				Filepaths:       []string{"/src", "/test"},
				FileContentsPredicate: PredicateInput{
					Type:  PredicateTypeEquals,
					Value: "postgres",
				},
			})
		},
	},
	"UpdateRepositoryFile": {
		fixture: "check/update_repo_file",
		body: func(c *Client) (*Check, error) {
			return c.UpdateCheckRepositoryFile(CheckRepositoryFileUpdateInput{
				Id:              graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:            "Hello World",
				Enabled:         true,
				Category:        graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:           graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:           "Hello World Check",
				DirectorySearch: true,
				Filepaths:       []string{"/src", "/test"},
				FileContentsPredicate: PredicateInput{
					Type:  PredicateTypeEquals,
					Value: "postgres",
				},
			})
		},
	},
	"CreateRepositoryIntegrated": {
		fixture: "check/create_repo_integrated",
		body: func(c *Client) (*Check, error) {
			return c.CreateCheckRepositoryIntegrated(CheckRepositoryIntegratedCreateInput{
				Name:     "Hello World",
				Enabled:  true,
				Category: graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
			})
		},
	},
	"UpdateRepositoryIntegrated": {
		fixture: "check/update_repo_integrated",
		body: func(c *Client) (*Check, error) {
			return c.UpdateCheckRepositoryIntegrated(CheckRepositoryIntegratedUpdateInput{
				Id:       graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:     "Hello World",
				Enabled:  true,
				Category: graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
			})
		},
	},
	"CreateRepositorySearch": {
		fixture: "check/create_repo_search",
		body: func(c *Client) (*Check, error) {
			return c.CreateCheckRepositorySearch(CheckRepositorySearchCreateInput{
				Name:           "Hello World",
				Enabled:        true,
				Category:       graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:          graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:          "Hello World Check",
				FileExtensions: []string{"sbt", "py"},
				FileContentsPredicate: PredicateInput{
					Type:  PredicateTypeContains,
					Value: "postgres",
				},
			})
		},
	},
	"UpdateRepositorySearch": {
		fixture: "check/update_repo_search",
		body: func(c *Client) (*Check, error) {
			return c.UpdateCheckRepositorySearch(CheckRepositorySearchUpdateInput{
				Id:             graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:           "Hello World",
				Enabled:        true,
				Category:       graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:          graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:          "Hello World Check",
				FileExtensions: []string{"sbt", "py"},
				FileContentsPredicate: PredicateInput{
					Type:  PredicateTypeContains,
					Value: "postgres",
				},
			})
		},
	},
	"CreateServiceConfiguration": {
		fixture: "check/create_service_configuration",
		body: func(c *Client) (*Check, error) {
			return c.CreateCheckServiceConfiguration(CheckServiceConfigurationCreateInput{
				Name:     "Hello World",
				Enabled:  true,
				Category: graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
			})
		},
	},
	"UpdateServiceConfiguration": {
		fixture: "check/update_service_configuration",
		body: func(c *Client) (*Check, error) {
			return c.UpdateCheckServiceConfiguration(CheckServiceConfigurationUpdateInput{
				Id:       graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:     "Hello World",
				Enabled:  true,
				Category: graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
			})
		},
	},
	"CreateServiceOwnership": {
		fixture: "check/create_service_ownership",
		body: func(c *Client) (*Check, error) {
			return c.CreateCheckServiceOwnership(CheckServiceOwnershipCreateInput{
				Name:     "Hello World",
				Enabled:  true,
				Category: graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
			})
		},
	},
	"UpdateServiceOwnership": {
		fixture: "check/update_service_ownership",
		body: func(c *Client) (*Check, error) {
			return c.UpdateCheckServiceOwnership(CheckServiceOwnershipUpdateInput{
				Id:       graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:     "Hello World",
				Enabled:  true,
				Category: graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
			})
		},
	},
	"CreateServiceProperty": {
		fixture: "check/create_service_property",
		body: func(c *Client) (*Check, error) {
			return c.CreateCheckServiceProperty(CheckServicePropertyCreateInput{
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
		},
	},
	"UpdateServiceProperty": {
		fixture: "check/update_service_property",
		body: func(c *Client) (*Check, error) {
			return c.UpdateCheckServiceProperty(CheckServicePropertyUpdateInput{
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
		},
	},
	"CreateTagDefined": {
		fixture: "check/create_tag_defined",
		body: func(c *Client) (*Check, error) {
			return c.CreateCheckTagDefined(CheckTagDefinedCreateInput{
				Name:     "Hello World",
				Enabled:  true,
				Category: graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
				TagKey:   "app",
				TagPredicate: PredicateInput{
					Type:  PredicateTypeEquals,
					Value: "postgres",
				},
			})
		},
	},
	"UpdateTagDefined": {
		fixture: "check/update_tag_defined",
		body: func(c *Client) (*Check, error) {
			return c.UpdateCheckTagDefined(CheckTagDefinedUpdateInput{
				Id:       graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:     "Hello World",
				Enabled:  true,
				Category: graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:    graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:    "Hello World Check",
				TagKey:   "app",
				TagPredicate: PredicateInput{
					Type:  PredicateTypeEquals,
					Value: "postgres",
				},
			})
		},
	},
	"CreateToolUsage": {
		fixture: "check/create_tool_usage",
		body: func(c *Client) (*Check, error) {
			return c.CreateCheckToolUsage(CheckToolUsageCreateInput{
				Name:         "Hello World",
				Enabled:      true,
				Category:     graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:        graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:        "Hello World Check",
				ToolCategory: ToolCategoryMetrics,
				ToolNamePredicate: PredicateInput{
					Type:  PredicateTypeEquals,
					Value: "datadog",
				},
				EnvironmentPredicate: PredicateInput{
					Type:  PredicateTypeEquals,
					Value: "production",
				},
			})
		},
	},
	"UpdateToolUsage": {
		fixture: "check/update_tool_usage",
		body: func(c *Client) (*Check, error) {
			return c.UpdateCheckToolUsage(CheckToolUsageUpdateInput{
				Id:           graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
				Name:         "Hello World",
				Enabled:      true,
				Category:     graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
				Level:        graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
				Notes:        "Hello World Check",
				ToolCategory: ToolCategoryMetrics,
				ToolNamePredicate: PredicateInput{
					Type:  PredicateTypeEquals,
					Value: "datadog",
				},
				EnvironmentPredicate: PredicateInput{
					Type:  PredicateTypeEquals,
					Value: "production",
				},
			})
		},
	},
	"GetCheck": {
		fixture: "check/get",
		body: func(c *Client) (*Check, error) {
			return c.GetCheck("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4")
		},
	},
}

func TestChecks(t *testing.T) {
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			// Arrange
			client := ANewClient(t, tc.fixture)
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
