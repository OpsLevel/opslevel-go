package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
	"github.com/shurcooL/graphql"
)

func TestCreateFilter(t *testing.T) {
	// Arrange
	client := ANewClient(t, "filter/create")
	// Act
	result, err := client.CreateFilter(FilterCreateInput{
		Name:       "Kubernetes",
		Connective: ConnectiveTypeAnd,
		Predicates: []FilterPredicate{FilterPredicate{
			Key:   "tier_index",
			Type:  PredicateTypeEquals,
			Value: "1",
		}},
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Kubernetes", result.Name)
	autopilot.Equals(t, "tier_index", result.Predicates[0].Key)
	autopilot.Equals(t, PredicateTypeEquals, result.Predicates[0].Type)
}

func TestListFilters(t *testing.T) {
	// Arrange
	client := ANewClient(t, "filter/list")
	// Act
	result, err := client.ListFilters()
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, 4, len(result))
	autopilot.Equals(t, "Test", result[1].Name)
	autopilot.Equals(t, "tier_index", result[3].Predicates[0].Key)
}

func TestUpdateFilter(t *testing.T) {
	// Arrange
	client := ANewClient(t, "filter/update")
	// Act
	result, err := client.UpdateFilter(FilterUpdateInput{
		Id:   graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg"),
		Name: "Test Updated",
		Predicates: []FilterPredicate{FilterPredicate{
			Key:   "tier_index",
			Type:  PredicateTypeEquals,
			Value: "1",
		}},
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Test Updated", result.Name)
	autopilot.Equals(t, "tier_index", result.Predicates[0].Key)
}

func TestDeleteFilter(t *testing.T) {
	// Arrange
	client := ANewClient(t, "filter/delete")
	// Act
	err := client.DeleteFilter("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz")
	// Assert
	autopilot.Equals(t, nil, err)
}
