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
