package opslevel_test

import (
	"encoding/json"
	"testing"

	ol "github.com/opslevel/opslevel-go/v2024"
	"github.com/rocktavious/autopilot/v2023"
)

func TestJsonMarshalPredicateUpdateInputNull(t *testing.T) {
	// Arrange
	predicateNull := "null"
	outputNull := &ol.PredicateUpdateInput{}
	// Act
	marshalledNullPredicate, err := json.Marshal(outputNull)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, predicateNull, string(marshalledNullPredicate))
}

func TestJsonMarshalPredicateUpdateInputNoValue(t *testing.T) {
	// Arrange
	existsEnum := ol.PredicateTypeEnumExists
	predicateNoValue := `{"type":"exists"}`
	outputNoValue := &ol.PredicateUpdateInput{Type: &existsEnum}
	// Act
	marshalledNullPredicate, err := json.Marshal(outputNoValue)
	autopilot.Ok(t, err)
	autopilot.Equals(t, predicateNoValue, string(marshalledNullPredicate))
}

func TestJsonMarshalPredicateUpdateInputWithValue(t *testing.T) {
	// Arrange
	containsEnum := ol.PredicateTypeEnumContains
	predicateWithValue := `{"type":"contains","value":"go"}`
	outputWithValue := &ol.PredicateUpdateInput{
		Type:  &containsEnum,
		Value: ol.NewNullableFrom("go"),
	}
	// Act
	marshalledNullPredicate, err := json.Marshal(outputWithValue)
	autopilot.Ok(t, err)
	autopilot.Equals(t, predicateWithValue, string(marshalledNullPredicate))
}
