package opslevel_test

import (
	"encoding/json"
	"testing"

	"github.com/hasura/go-graphql-client"
	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
)

type IDTester struct {
	Key1 ol.ID  `json:"key1"`
	Key2 ol.ID  `json:"key2,omitempty"`
	Key3 *ol.ID `json:"key3"`
	Key4 *ol.ID `json:"key4,omitempty"`
}

func TestMarshalID(t *testing.T) {
	// Arrange
	id1 := ol.NewID()
	id2 := ol.NewID("Z2lkOi8vMTIzNDU2Nzg5")
	case1 := IDTester{}
	case2 := IDTester{
		Key1: *id1,
		Key2: *id1,
		Key3: id1,
		Key4: id1,
	}
	case3 := IDTester{
		Key1: *id2,
		Key2: *id2,
		Key3: id2,
		Key4: id2,
	}
	// Act
	buf1, err1 := json.Marshal(case1)
	buf2, err2 := json.Marshal(case2)
	buf3, err3 := json.Marshal(case3)
	// Assert
	autopilot.Ok(t, err1)
	autopilot.Equals(t, `{"key1":"","key3":null}`, string(buf1))
	autopilot.Ok(t, err2)
	autopilot.Equals(t, `{"key1":"","key3":null,"key4":null}`, string(buf2))
	autopilot.Ok(t, err3)
	autopilot.Equals(t, `{"key1":"Z2lkOi8vMTIzNDU2Nzg5","key2":"Z2lkOi8vMTIzNDU2Nzg5","key3":"Z2lkOi8vMTIzNDU2Nzg5","key4":"Z2lkOi8vMTIzNDU2Nzg5"}`, string(buf3))
}

func TestConstructQueryID(t *testing.T) {
	// Arrange
	id := ol.ID("1234")
	var q struct {
		Account struct {
			Output struct {
				Id ol.ID `graphql:"id"`
			} `graphql:"myQuery(id1: $id1 id2: $id2)"`
		}
	}
	v := ol.PayloadVariables{
		"id1": id,
		"id2": &id,
	}
	// Act
	query, err := graphql.ConstructQuery(q, v, ol.WithName("MyQuery"))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, `query MyQuery($id1:ID!$id2:ID){account{myQuery(id1: $id1 id2: $id2){id}}}`, query)
}

func TestConstructMutationID(t *testing.T) {
	// Arrange
	id := ol.ID("1234")
	var q struct {
		Account struct {
			Output struct {
				Id ol.ID `graphql:"id"`
			} `graphql:"myMutation(id1: $id1 id2: $id2)"`
		}
	}
	v := ol.PayloadVariables{
		"id1": id,
		"id2": &id,
	}
	// Act
	query, err := graphql.ConstructMutation(q, v, ol.WithName("MyMutation"))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, `mutation MyMutation($id1:ID!$id2:ID){account{myMutation(id1: $id1 id2: $id2){id}}}`, query)
}

type IdentifierTester struct {
	Key1 ol.IdentifierInput  `json:"key1"`
	Key2 ol.IdentifierInput  `json:"key2,omitempty"`
	Key3 *ol.IdentifierInput `json:"key3"`
	Key4 *ol.IdentifierInput `json:"key4,omitempty"`
}

func TestMarshalIdentifier(t *testing.T) {
	// Arrange
	id1 := ol.NewIdentifier("")
	id2 := ol.NewIdentifier("my-service")
	id3 := ol.NewIdentifier("Z2lkOi8vMTIzNDU2Nzg5")
	case1 := IdentifierTester{}
	case2 := IdentifierTester{
		Key1: *id1,
		Key2: *id1,
		Key3: id1,
		Key4: id1,
	}
	case3 := IdentifierTester{
		Key1: *id2,
		Key2: *id2,
		Key3: id2,
		Key4: id2,
	}
	case4 := IdentifierTester{
		Key1: *id3,
		Key2: *id3,
		Key3: id3,
		Key4: id3,
	}
	// Act
	buf1, err1 := json.Marshal(case1)
	buf2, err2 := json.Marshal(case2)
	buf3, err3 := json.Marshal(case3)
	buf4, err4 := json.Marshal(case4)
	// Assert
	autopilot.Ok(t, err1)
	autopilot.Equals(t, `{"key1":{},"key2":{},"key3":null}`, string(buf1))
	autopilot.Ok(t, err2)
	autopilot.Equals(t, `{"key1":{},"key2":{},"key3":{},"key4":{}}`, string(buf2))
	autopilot.Ok(t, err3)
	autopilot.Equals(t, `{"key1":{"alias":"my-service"},"key2":{"alias":"my-service"},"key3":{"alias":"my-service"},"key4":{"alias":"my-service"}}`, string(buf3))
	autopilot.Ok(t, err4)
	autopilot.Equals(t, `{"key1":{"id":"Z2lkOi8vMTIzNDU2Nzg5"},"key2":{"id":"Z2lkOi8vMTIzNDU2Nzg5"},"key3":{"id":"Z2lkOi8vMTIzNDU2Nzg5"},"key4":{"id":"Z2lkOi8vMTIzNDU2Nzg5"}}`, string(buf4))
}

func TestConstructQueryIdentifier(t *testing.T) {
	// Arrange
	id1 := ol.NewIdentifier("my-service")
	id2 := ol.NewIdentifier("Z2lkOi8vMTIzNDU2Nzg5")
	var q struct {
		Account struct {
			Output ol.Identifier `graphql:"myQuery(id1: $id1 id2: $id2 id3: $id3 id4: $id4)"`
		}
	}
	v := ol.PayloadVariables{
		"id1": *id1,
		"id2": *id2,
		"id3": id1,
		"id4": id2,
	}
	// Act
	query, err := graphql.ConstructQuery(q, v, ol.WithName("MyQuery"))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, `query MyQuery($id1:IdentifierInput!$id2:IdentifierInput!$id3:IdentifierInput$id4:IdentifierInput){account{myQuery(id1: $id1 id2: $id2 id3: $id3 id4: $id4){id,aliases}}}`, query)
}

func TestConstructMutationIdentifier(t *testing.T) {
	// Arrange
	id1 := ol.NewIdentifier("my-service")
	id2 := ol.NewIdentifier("Z2lkOi8vMTIzNDU2Nzg5")
	var q struct {
		Account struct {
			Output ol.Identifier `graphql:"myMutation(id1: $id1 id2: $id2 id3: $id3 id4: $id4)"`
		}
	}
	v := ol.PayloadVariables{
		"id1": *id1,
		"id2": *id2,
		"id3": id1,
		"id4": id2,
	}
	// Act
	query, err := graphql.ConstructMutation(q, v, ol.WithName("MyMutation"))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, `mutation MyMutation($id1:IdentifierInput!$id2:IdentifierInput!$id3:IdentifierInput$id4:IdentifierInput){account{myMutation(id1: $id1 id2: $id2 id3: $id3 id4: $id4){id,aliases}}}`, query)
}

func TestNewIdentifierArray(t *testing.T) {
	// Arrange
	s := []string{"my-service", "Z2lkOi8vMTIzNDU2Nzg5"}
	result := ol.NewIdentifierArray(s)
	// Assert
	autopilot.Equals(t, "my-service", result[0].Alias)
	autopilot.Equals(t, ol.ID("Z2lkOi8vMTIzNDU2Nzg5"), result[1].Id)
}
