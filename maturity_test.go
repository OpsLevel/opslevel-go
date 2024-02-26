package opslevel_test

import (
	"testing"

	"github.com/rocktavious/autopilot/v2023"
)

func TestGetServiceMaturityWithAlias(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query ($service:String!){account{service(alias:$service){name,maturityReport{categoryBreakdown{category{description,id,name},level{alias,description,id,index,name}},overallLevel{alias,description,id,index,name}}}}}`,
		`{"service": "cert-manager"}`,
		`{
  "data": {
    "account": {
      "service": {
        "name": "cert-manager",
        "maturityReport": {
          "categoryBreakdown": [
            {
              "category": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNzA1",
                "name": "Service Ownership"
              },
              "level": {
                "alias": "beginner",
                "description": "Services in this level are below the minimum standard to ship to production. You should address your failing checks as soon as possible.",
                "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDQ0",
                "index": 0,
                "name": "Beginner"
              }
            },
            {
              "category": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNzA2",
                "name": "Security"
              },
              "level": {
                "alias": "beginner",
                "description": "Services in this level are below the minimum standard to ship to production. You should address your failing checks as soon as possible.",
                "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDQ0",
                "index": 0,
                "name": "Beginner"
              }
            },
            {
              "category": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNzA3",
                "name": "Reliability"
              },
              "level": null
            },
            {
              "category": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNzA4",
                "name": "Performance"
              },
              "level": null
            },
            {
              "category": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNzA5",
                "name": "Scalability"
              },
              "level": null
            },
            {
              "category": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNzEw",
                "name": "Observability"
              },
              "level": null
            },
            {
              "category": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNzEx",
                "name": "Infrastructure"
              },
              "level": null
            },
            {
              "category": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNzEy",
                "name": "Quality"
              },
              "level": null
            },
            {
              "category": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNzEz",
                "name": "Service Ownership"
              },
              "level": null
            },
            {
              "category": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNzU2",
                "name": "Test"
              },
              "level": null
            },
            {
              "category": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNzU3",
                "name": "Test"
              },
              "level": null
            }
          ],
          "overallLevel": {
            "alias": "beginner",
            "description": "Services in this level are below the minimum standard to ship to production. You should address your failing checks as soon as possible.",
            "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDQ0",
            "index": 0,
            "name": "Beginner"
          }
        }
      }
    }
  }
}`,
	)
	client := BestTestClient(t, "maturity/get_service_maturity_with_alias", testRequest)
	// Act
	result, err := client.GetServiceMaturityWithAlias("cert-manager")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "cert-manager", result.Name)
	autopilot.Equals(t, "beginner", result.MaturityReport.OverallLevel.Alias)
}

func TestListServicesMaturity(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query ($after:String!$first:Int!){account{services(after: $after, first: $first){nodes{name,maturityReport{categoryBreakdown{category{description,id,name},level{alias,description,id,index,name}},overallLevel{alias,description,id,index,name}}},pageInfo{endCursor,hasNextPage,hasPreviousPage,startCursor}}}}`,
		`{"after":"", "first":100}`,
		`{
  "data": {
    "account": {
      "services": {
        "nodes": [
          {
            "name": "Example",
            "maturityReport": {
              "categoryBreakdown": [
                {
                  "category": {
                    "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvMTU",
                    "name": "Security"
                  },
                  "level": {
                    "alias": "gold",
                    "description": "Services in this level satisfy critical, important and useful checks. This is the requirement for your highest tier services but all services should aspire to be in this level.",
                    "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMjI2",
                    "index": 3,
                    "name": "Gold"
                  }
                },
                {
                  "category": {
                    "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvMTY",
                    "name": "Reliability"
                  },
                  "level": null
                },
                {
                  "category": {
                    "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvMTc",
                    "name": "Performance"
                  },
                  "level": null
                },
                {
                  "category": {
                    "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvMTg",
                    "name": "Scalability"
                  },
                  "level": null
                },
                {
                  "category": {
                    "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvMTk",
                    "name": "Observability"
                  },
                  "level": null
                },
                {
                  "category": {
                    "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvMjA",
                    "name": "Infrastructure"
                  },
                  "level": null
                },
                {
                  "category": {
                    "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvMjE",
                    "name": "Quality"
                  },
                  "level": {
                    "alias": "gold",
                    "description": "Services in this level satisfy critical, important and useful checks. This is the requirement for your highest tier services but all services should aspire to be in this level.",
                    "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMjI2",
                    "index": 3,
                    "name": "Gold"
                  }
                }
              ],
              "overallLevel": {
                "alias": "gold",
                "description": "Services in this level satisfy critical, important and useful checks. This is the requirement for your highest tier services but all services should aspire to be in this level.",
                "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMjI2",
                "index": 3,
                "name": "Gold"
              }
            }
          }
        ],
        "pageInfo": {
          "hasNextPage": false,
          "hasPreviousPage": false,
          "startCursor": "MQ",
          "endCursor": "MTA"
        }
      }
    }
  }
}`,
	)
	client := BestTestClient(t, "maturity/services", testRequest)
	// Act
	result, err := client.ListServicesMaturity()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
	autopilot.Equals(t, "Example", result[0].Name)
	autopilot.Equals(t, "Gold", result[0].MaturityReport.Get("Quality").Name)
}
