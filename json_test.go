package opslevel_test

import (
	"encoding/json"
	"testing"

	"github.com/hasura/go-graphql-client"
	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
)

type JSONTester struct {
	Key1 ol.JSON  `json:"key1"`
	Key2 ol.JSON  `json:"key2,omitempty"`
	Key3 *ol.JSON `json:"key3"`
	Key4 *ol.JSON `json:"key4,omitempty"`
}

func TestNewJSON(t *testing.T) {
	// Arrange
	data1 := ol.JSON{"foo": "bar"}
	data2 := ol.NewJSON(`{"foo":"bar"}`)
	// Act
	result1, err1 := json.Marshal(data1)
	result2, err2 := json.Marshal(data2)
	// Assert
	autopilot.Ok(t, err1)
	autopilot.Ok(t, err2)
	autopilot.Equals(t, data1, data2)
	autopilot.Assert(t, &data1 != &data2, "The JSON objects have the same memory address")
	autopilot.Equals(t, result1, result2)
}

func TestMarshalJSON(t *testing.T) {
	// Arrange
	id1 := ol.JSON{}
	id2 := ol.JSON{
		"foo": "bar",
	}
	case1 := JSONTester{}
	case2 := JSONTester{
		Key1: id1,
		Key2: id1,
		Key3: &id1,
		Key4: &id1,
	}
	case3 := JSONTester{
		Key1: id2,
		Key2: id2,
		Key3: &id2,
		Key4: &id2,
	}
	// Act
	buf1, err1 := json.Marshal(case1)
	buf2, err2 := json.Marshal(case2)
	buf3, err3 := json.Marshal(case3)
	// Assert
	autopilot.Ok(t, err1)
	autopilot.Equals(t, `{"key1":"{}","key3":null}`, string(buf1))
	autopilot.Ok(t, err2)
	autopilot.Equals(t, `{"key1":"{}","key3":"{}","key4":"{}"}`, string(buf2))
	autopilot.Ok(t, err3)
	autopilot.Equals(t, `{"key1":"{\"foo\":\"bar\"}","key2":"{\"foo\":\"bar\"}","key3":"{\"foo\":\"bar\"}","key4":"{\"foo\":\"bar\"}"}`, string(buf3))
}

func TestConstructQueryJSON(t *testing.T) {
	// Arrange
	data := ol.JSON{
		"foo": "bar",
	}
	var q struct {
		Account struct {
			Output struct {
				Data ol.JSON `json:"data" scalar:"true"`
			} `graphql:"myQuery(id1: $id1 id2: $id2)"`
		}
	}
	v := ol.PayloadVariables{
		"id1": data,
		"id2": &data,
	}
	// Act
	query, err := graphql.ConstructQuery(q, v, ol.WithName("MyQuery"))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, `query MyQuery($id1:JSON!$id2:JSON){account{myQuery(id1: $id1 id2: $id2){data}}}`, query)
}

func TestConstructMutationJSON(t *testing.T) {
	// Arrange
	data := ol.JSON{
		"foo": "bar",
	}
	var q struct {
		Account struct {
			Output struct {
				Data ol.JSON `json:"data" scalar:"true"`
			} `graphql:"myMutation(id1: $id1 id2: $id2)"`
		}
	}
	v := ol.PayloadVariables{
		"id1": data,
		"id2": &data,
	}
	// Act
	query, err := graphql.ConstructMutation(q, v, ol.WithName("MyMutation"))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, `mutation MyMutation($id1:JSON!$id2:JSON){account{myMutation(id1: $id1 id2: $id2){data}}}`, query)
}
