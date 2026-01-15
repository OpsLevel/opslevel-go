package opslevel_test

import (
	"encoding/json"
	"testing"

	"github.com/hasura/go-graphql-client"

	ol "github.com/opslevel/opslevel-go/v2026"
	"github.com/rocktavious/autopilot/v2023"
)

var (
	wrappedObject    = `{"access":{"aws":"admin","okta":"admin"},"isIntern":false,"name":"Thomas","tags":["org:engineering","team:platform"]}`
	wrappedObjectMap = map[string]any{
		"name":     "Thomas",
		"isIntern": false,
		"access": map[string]any{
			"aws":  "admin",
			"okta": "admin",
		},
		"tags": []any{
			"org:engineering",
			"team:platform",
		},
	}
	wrappedArray      = `["Thomas",false,{"aws":"admin","okta":"admin"},["org:engineering","team:platform"]]`
	wrappedArraySlice = []any{
		"Thomas",
		false,
		map[string]any{
			"aws":  "admin",
			"okta": "admin",
		},
		[]string{
			"org:engineering",
			"team:platform",
		},
	}
)

type JSONTester struct {
	Key1 ol.JSON  `json:"key1"`
	Key2 ol.JSON  `json:"key2,omitempty"`
	Key3 *ol.JSON `json:"key3"`
	Key4 *ol.JSON `json:"key4,omitempty"`
}

func TestNewJsonString(t *testing.T) {
	type TestCase struct {
		Data   any
		Output ol.JsonString
	}
	testCases := map[string]TestCase{
		// object
		"wrappedObject":         {wrappedObject, ol.JsonString(wrappedObject)},
		"wrappedObjectMap":      {wrappedObjectMap, ol.JsonString(wrappedObject)},
		"wrappedObjectEmpty":    {"{}", "{}"},
		"wrappedObjectMapEmpty": {map[string]any{}, "{}"},
		// array
		"wrappedArray":           {wrappedArray, ol.JsonString(wrappedArray)},
		"wrappedArraySlice":      {wrappedArraySlice, ol.JsonString(wrappedArray)},
		"wrappedArrayEmpty":      {"[]", "[]"},
		"wrappedArraySliceEmpty": {[]any{}, "[]"},
		// string
		"string":      {"hello world", `"hello world"`},
		"stringEmpty": {"", `""`},
		// bool
		"wrappedBoolTrue": {"true", `"true"`},
		"boolTrue":        {true, "true"},
		"boolFalse":       {false, "false"},
		// number
		"wrappedDecimal": {"1.32", `"1.32"`},
		"numberDecimal":  {1.32, "1.32"},
		"numberIntZero":  {0, "0"},
	}
	for k, v := range testCases {
		t.Run(k, func(t *testing.T) {
			res, err := ol.NewJSONInput(v.Data)
			autopilot.Ok(t, err)
			autopilot.Equals(t, v.Output, *res)
		})
	}
}

func TestNewJSON(t *testing.T) {
	type TestCase struct {
		Data   string
		Output ol.JSON
	}
	testCases := map[string]TestCase{
		"wrappedObject":      {wrappedObject, ol.JSON(wrappedObjectMap)},
		"wrappedObjectEmpty": {"{}", ol.JSON{}},
	}
	for k, v := range testCases {
		t.Run(k, func(t *testing.T) {
			res, err := ol.NewJSON(v.Data)
			autopilot.Ok(t, err)
			autopilot.Equals(t, v.Output, *res)
		})
	}
}

func TestNewJSONSchema(t *testing.T) {
	type TestCase struct {
		Data   string
		Output ol.JSONSchema
	}
	testCases := map[string]TestCase{
		"wrappedObject":      {wrappedObject, ol.JSONSchema(wrappedObjectMap)},
		"wrappedObjectEmpty": {"{}", ol.JSONSchema{}},
	}
	for k, v := range testCases {
		t.Run(k, func(t *testing.T) {
			res, err := ol.NewJSONSchema(v.Data)
			autopilot.Ok(t, err)
			autopilot.Equals(t, v.Output, *res)
		})
	}
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
	var3, err := ol.NewJSONInput(map[string]any{
		"foo": "bar",
	})
	autopilot.Assert(t, err == nil, "unexpected error")
	v := ol.PayloadVariables{
		"id1": data,
		"id2": &data,
		"id3": *var3,
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
	var3, err := ol.NewJSONInput(map[string]any{
		"foo": "bar",
	})
	autopilot.Assert(t, err == nil, "unexpected error")
	v := ol.PayloadVariables{
		"id1": data,
		"id2": &data,
		"id3": *var3,
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
	result1, result1Err := ol.NewJSONInput(data1)
	result1a, result1aErr := ol.NewJSONInput(data1a)
	result2, result2Err := ol.NewJSONInput(data2)
	result2a, result2aErr := ol.NewJSONInput(data2a)
	result3, result3Err := ol.NewJSONInput(data3)
	result4, result4Err := ol.NewJSONInput(data4)
	result5, result5Err := ol.NewJSONInput(data5)
	_, err := ol.JsonStringAs[float64](*result1)
	// Assert
	autopilot.Ok(t, result1Err)
	autopilot.Ok(t, result1aErr)
	autopilot.Ok(t, result2Err)
	autopilot.Ok(t, result2aErr)
	autopilot.Ok(t, result3Err)
	autopilot.Ok(t, result4Err)
	autopilot.Ok(t, result5Err)
	autopilot.Equals(t, data1, result1.AsBool())
	autopilot.Equals(t, data1a, result1a.AsBool())
	autopilot.Equals(t, data2, result2.AsFloat64())
	autopilot.Equals(t, data2a, result2a.AsInt())
	autopilot.Equals(t, data3, result3.AsString())
	autopilot.Equals(t, data4, result4.AsArray())
	autopilot.Equals(t, data5, result5.AsMap())
	autopilot.Assert(t, err != nil, "The JSON string of type bool should be unable to unmarshalled into a float64")
}
