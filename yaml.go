package opslevel

import (
	"gopkg.in/yaml.v3"
)

// Generate example yaml files for OpsLevel resources
func GenYamlFromStruct[T any](opslevelResource T) (string, error) {
	out, err := yaml.Marshal(opslevelResource)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
