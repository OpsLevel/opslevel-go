package opslevel_test

import (
	"encoding/json"
	"testing"

	"github.com/hasura/go-graphql-client"
	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
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
	result3 := ol.NewJSONInput(`{"foo":"bar"}`)

	result4 := ol.NewJSONInput(true)
	result4a := ol.NewJSONInput(false)
	result5 := ol.NewJSONInput(1.32)
	result5a := ol.NewJSONInput(0)
	result6 := ol.NewJSONInput("hello world")
	result7 := ol.NewJSONInput([]any{"foo", "bar"})
	result8 := ol.NewJSONInput(map[string]any{"foo": "bar"})
	// Assert
	autopilot.Ok(t, err1)
	autopilot.Ok(t, err2)
	autopilot.Equals(t, data1, data2)
	autopilot.Assert(t, &data1 != &data2, "The JSON objects have the same memory address")
	autopilot.Equals(t, result1, result2)
	autopilot.Equals(t, string(result1), string(result3))
	autopilot.Equals(t, string(result2), string(result3))

	autopilot.Equals(t, `true`, string(result4))
	autopilot.Equals(t, `false`, string(result4a))
	autopilot.Equals(t, `1.32`, string(result5))
	autopilot.Equals(t, `0`, string(result5a))
	autopilot.Equals(t, `"hello world"`, string(result6))
	autopilot.Equals(t, `["foo","bar"]`, string(result7))
	autopilot.Equals(t, `{"foo":"bar"}`, string(result8))
}

func TestMarshalJSON(t *testing.T) {
	// Arrange
	id1 := ol.JSON{}
	id2 := ol.JSON{
		"foo": "bar",
		"nested": map[string]any{
			"one": 1,
			"two": "two",
		},
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
	autopilot.Equals(t, `{"key1":"{\"foo\":\"bar\",\"nested\":{\"one\":1,\"two\":\"two\"}}","key2":"{\"foo\":\"bar\",\"nested\":{\"one\":1,\"two\":\"two\"}}","key3":"{\"foo\":\"bar\",\"nested\":{\"one\":1,\"two\":\"two\"}}","key4":"{\"foo\":\"bar\",\"nested\":{\"one\":1,\"two\":\"two\"}}"}`, string(buf3))
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
			} `graphql:"myQuery(id1: $id1 id2: $id2 id3: $id3)"`
		}
	}
	v := ol.PayloadVariables{
		"id1": data,
		"id2": &data,
		"id3": ol.NewJSONInput(map[string]any{
			"foo": "bar",
		}),
	}
	// Act
	query, err := graphql.ConstructQuery(q, v, ol.WithName("MyQuery"))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, `query MyQuery($id1:JSON!$id2:JSON$id3:JsonString!){account{myQuery(id1: $id1 id2: $id2 id3: $id3){data}}}`, query)
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
			} `graphql:"myMutation(id1: $id1 id2: $id2 id3: $id3)"`
		}
	}
	v := ol.PayloadVariables{
		"id1": data,
		"id2": &data,
		"id3": ol.NewJSONInput(map[string]any{
			"foo": "bar",
		}),
	}
	// Act
	query, err := graphql.ConstructMutation(q, v, ol.WithName("MyMutation"))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, `mutation MyMutation($id1:JSON!$id2:JSON$id3:JsonString!){account{myMutation(id1: $id1 id2: $id2 id3: $id3){data}}}`, query)
}

func TestUnmarshalJSONString(t *testing.T) {
	// Arrange
	data1 := true
	data1a := false
	data2 := 1.32
	data2a := 0
	data3 := "hello world"
	data4 := []any{"foo", "bar"}
	data5 := map[string]any{"foo": "bar"}
	// Act
	result1 := ol.NewJSONInput(data1).AsBool()
	result1a := ol.NewJSONInput(data1a).AsBool()
	result2 := ol.NewJSONInput(data2).AsFloat64()
	result2a := ol.NewJSONInput(data2a).AsInt()
	result3 := ol.NewJSONInput(data3).AsString()
	result4 := ol.NewJSONInput(data4).AsArray()
	result5 := ol.NewJSONInput(data5).AsMap()
	_, err := ol.JsonStringAs[float32](ol.NewJSONInput(data1))
	// Assert
	autopilot.Equals(t, data1, result1)
	autopilot.Equals(t, data1a, result1a)
	autopilot.Equals(t, data2, result2)
	autopilot.Equals(t, data2a, result2a)
	autopilot.Equals(t, data3, result3)
	autopilot.Equals(t, data4, result4)
	autopilot.Equals(t, data5, result5)
	autopilot.Assert(t, err != nil, "The JSON string of type bool should be unable to unmarshalled into a float32")
}
