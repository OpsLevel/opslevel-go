package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2022"

	"github.com/rocktavious/autopilot"
)

func TestCreateFilter(t *testing.T) {
	// Arrange
	client := ATestClient(t, "filter/create")
	// Act
	result, err := client.CreateFilter(ol.FilterCreateInput{
		Name:       "Kubernetes",
		Connective: ol.ConnectiveEnumAnd,
		Predicates: []ol.FilterPredicate{{
			Key:   ol.PredicateKeyEnumTierIndex,
			Type:  ol.PredicateTypeEnumEquals,
			Value: "1",
		}},
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Kubernetes", result.Name)
	autopilot.Equals(t, ol.PredicateKeyEnumTierIndex, result.Predicates[0].Key)
	autopilot.Equals(t, ol.PredicateTypeEnumEquals, result.Predicates[0].Type)
}

func TestGetFilter(t *testing.T) {
	// Arrange
	client := ATestClient(t, "filter/get")
	// Act
	result, err := client.GetFilter("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg")
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Test", result.Name)
	autopilot.Equals(t, ol.PredicateKeyEnumTierIndex, result.Predicates[0].Key)
}

func TestGetMissingFilter(t *testing.T) {
	// Arrange
	client := ATestClient(t, "filter/get_missing")
	// Act
	_, err := client.GetFilter("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMf")
	// Assert
	autopilot.Assert(t, err != nil, "This test should throw an error.")
}

func TestListFilters(t *testing.T) {
	// Arrange
	client := ATestClient(t, "filter/list")
	// Act
	result, err := client.ListFilters()
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, 4, len(result))
	autopilot.Equals(t, "Test", result[1].Name)
	autopilot.Equals(t, ol.PredicateKeyEnumTierIndex, result[3].Predicates[0].Key)
}

func TestUpdateFilter(t *testing.T) {
	// Arrange
	client := ATestClient(t, "filter/update")
	// Act
	result, err := client.UpdateFilter(ol.FilterUpdateInput{
		Id:   ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg"),
		Name: "Test Updated",
		Predicates: []ol.FilterPredicate{{
			Key:   ol.PredicateKeyEnumTierIndex,
			Type:  ol.PredicateTypeEnumEquals,
			Value: "1",
		}},
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Test Updated", result.Name)
	autopilot.Equals(t, ol.PredicateKeyEnumTierIndex, result.Predicates[0].Key)
}

func TestDeleteFilter(t *testing.T) {
	// Arrange
	client := ATestClient(t, "filter/delete")
	// Act
	err := client.DeleteFilter("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz")
	// Assert
	autopilot.Equals(t, nil, err)
}
