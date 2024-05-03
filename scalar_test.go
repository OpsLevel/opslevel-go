package opslevel_test

import (
	"encoding/json"
	"testing"

	"github.com/hasura/go-graphql-client"

	ol "github.com/opslevel/opslevel-go/v2024"
	"github.com/rocktavious/autopilot/v2023"
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

func TestMarshalIdentifiers(t *testing.T) {
	type TestCase struct {
		Name         string
		Identifier   *ol.IdentifierInput
		OutputBuffer string
	}
	testCases := []TestCase{
		{
			Name:         "empty identifier",
			Identifier:   ol.NewIdentifier(),
			OutputBuffer: `null`,
		},
		{
			Name:         "identifier with empty arg",
			Identifier:   ol.NewIdentifier(""),
			OutputBuffer: `{"alias":""}`,
		},
		{
			Name:         "identifier with valid ID",
			Identifier:   ol.NewIdentifier("Z2lkOi8vb3BzbGV2ZWwvSGVsbG9Xb3JsZC8xMDEw"),
			OutputBuffer: `{"id":"Z2lkOi8vb3BzbGV2ZWwvSGVsbG9Xb3JsZC8xMDEw"}`,
		},
		{
			Name:         "identifier with valid alias",
			Identifier:   ol.NewIdentifier("hello_world"),
			OutputBuffer: `{"alias":"hello_world"}`,
		},
	}

	for _, testCase := range testCases {
		buf, err := json.Marshal(testCase.Identifier)
		autopilot.Ok(t, err)
		autopilot.Equals(t, testCase.OutputBuffer, string(buf))
	}
}

func TestMarshalIdentifiersOmitBehavior(t *testing.T) {
	type TestCase struct {
		Name         string
		Owner        *ol.IdentifierInput
		Maintainer   *ol.IdentifierInput
		OutputBuffer string
	}
	testCases := []TestCase{
		{
			Name:         "pass nil, owner should omitempty",
			Owner:        nil,
			Maintainer:   nil,
			OutputBuffer: `{"maintainer":null}`,
		},
		{
			Name:         "pass empty identifier, owner should null",
			Owner:        ol.NewIdentifier(),
			Maintainer:   ol.NewIdentifier(),
			OutputBuffer: `{"owner":null,"maintainer":null}`,
		},
		{
			Name:         "pass normal identifiers",
			Owner:        ol.NewIdentifier("Z2lkOi8vb3BzbGV2ZWwvSGVsbG9Xb3JsZC8xMDEw"),
			Maintainer:   ol.NewIdentifier("team2"),
			OutputBuffer: `{"owner":{"id":"Z2lkOi8vb3BzbGV2ZWwvSGVsbG9Xb3JsZC8xMDEw"},"maintainer":{"alias":"team2"}}`,
		},
	}

	for _, testCase := range testCases {
		type SomethingUpdateInput struct {
			Owner      *ol.IdentifierInput `json:"owner,omitempty"`
			Maintainer *ol.IdentifierInput `json:"maintainer"`
		}
		input := SomethingUpdateInput{
			Owner:      testCase.Owner,
			Maintainer: testCase.Maintainer,
		}

		buf, err := json.Marshal(input)
		autopilot.Ok(t, err)
		autopilot.Equals(t, testCase.OutputBuffer, string(buf))
	}
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
	autopilot.Equals(t, "my-service", *result[0].Alias)
	autopilot.Equals(t, ol.ID("Z2lkOi8vMTIzNDU2Nzg5"), *result[1].Id)
}

func TestOptionalString(t *testing.T) {
	type TestCase struct {
		Name         string
		Input        string
		OutputBuffer string
	}
	testCases := []TestCase{
		{
			Name:         "empty input",
			Input:        "",
			OutputBuffer: `null`,
		},
		{
			Name:         "spaces",
			Input:        "              ",
			OutputBuffer: `"              "`,
		},
		{
			Name:         "the string null",
			Input:        "null",
			OutputBuffer: `"null"`,
		},
		{
			Name:         "simple hello world",
			Input:        "hello world",
			OutputBuffer: `"hello world"`,
		},
		{
			Name:         "quoted hello world",
			Input:        `"hello world"`,
			OutputBuffer: `"\"hello world\""`,
		},
	}

	for _, testCase := range testCases {
		buf, err := json.Marshal(ol.NewOptionalString(testCase.Input))
		autopilot.Ok(t, err)
		autopilot.Equals(t, testCase.OutputBuffer, string(buf))
	}
}
