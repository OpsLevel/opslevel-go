package opslevel

// YAML represents a YAML document serialized as a string for use with the OpsLevel API.
type YAML string

func (y YAML) GetGraphQLType() string { return "YAML" }
