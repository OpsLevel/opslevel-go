package opslevel

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

const (
	ExampleAlias       = "alias123"
	ExampleEmail       = "john.doe@example.com"
	ExampleEnvironment = "production"
	ExampleId          = "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"
	ExampleName        = "John Doe"
	ExampleUrl         = "https://example.com"
)

// Generate JSON formatted string of an OpsLevel resource -
// use T.Example() as the argument
func GenJsonFrom[T any](opslevelResource T) string {
	out, err := json.Marshal(opslevelResource)
	if err != nil {
		panic(err)
	}
	return string(out)
}

// Generate yaml formatted string of an OpsLevel resource -
// use T.Example() as the argument
func GenYamlFrom[T any](opslevelResource T) string {
	out, err := yaml.Marshal(opslevelResource)
	if err != nil {
		panic(err)
	}
	return string(out)
}
