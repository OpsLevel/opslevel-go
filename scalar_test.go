package opslevel_test

import (
	"encoding/json"
	"github.com/hasura/go-graphql-client"
	"github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
	"testing"
)

type TestInput struct {
	Key1 opslevel.ID  `json:"key1"`
	Key2 opslevel.ID  `json:"key2,omitempty"`
	Key3 *opslevel.ID `json:"key3"`
	Key4 *opslevel.ID `json:"key4,omitempty"`
}

func TestMarshalID(t *testing.T) {
	// Arrange
	id := opslevel.ID("")
	id2 := opslevel.ID("1234")
	one := TestInput{}
	two := TestInput{
		Key1: id,
		Key2: id,
		Key3: &id,
		Key4: &id,
	}
	three := TestInput{
		Key1: id2,
		Key2: id2,
		Key3: &id2,
		Key4: &id2,
	}
	// Act
	buf1, err1 := json.Marshal(one)
	buf2, err2 := json.Marshal(two)
	buf3, err3 := json.Marshal(three)
	// Assert
	autopilot.Ok(t, err1)
	autopilot.Equals(t, `{"key1":"","key3":null}`, string(buf1))
	autopilot.Ok(t, err2)
	autopilot.Equals(t, `{"key1":"","key3":null,"key4":null}`, string(buf2))
	autopilot.Ok(t, err3)
	autopilot.Equals(t, `{"key1":"1234","key2":"1234","key3":"1234","key4":"1234"}`, string(buf3))
}

func TestConstructQueryID(t *testing.T) {
	// Arrange
	id := opslevel.ID("1234")
	var q struct {
		Account struct {
			Output TestInput `graphql:"myMutation(id1: $id1 id2: $id2)"`
		}
	}
	v := opslevel.PayloadVariables{
		"id1": id,
		"id2": &id,
	}
	// Act
	query, err := graphql.ConstructQuery(q, v)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, `query ($id1:ID!$id2:ID){account{myMutation(id1: $id1 id2: $id2){key1,key2,key3,key4}}}`, query)
}

func TestConstructMutationID(t *testing.T) {
	// Arrange
	id := opslevel.ID("1234")
	var q struct {
		Account struct {
			Output TestInput `graphql:"myMutation(id1: $id1 id2: $id2)"`
		}
	}
	v := opslevel.PayloadVariables{
		"id1": id,
		"id2": &id,
	}
	// Act
	query, err := graphql.ConstructMutation(q, v)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, `mutation ($id1:ID!$id2:ID){account{myMutation(id1: $id1 id2: $id2){key1,key2,key3,key4}}}`, query)
}
