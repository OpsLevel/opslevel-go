package opslevel

import (
	"encoding/json"
	"fmt"
)

// TypeIdentifier represents a GraphQL __typename that can be used to identify the concrete type of an union response
type TypeIdentifier string

func (r *RelationshipResource) MarshalJSON() ([]byte, error) {
	switch r.Type {
	case "Domain":
		return json.Marshal(r.Domain)
	case "Infrastructure":
		return json.Marshal(r.InfrastructureResource)
	case "Service":
		return json.Marshal(r.Service)
	case "System":
		return json.Marshal(r.System)
	case "Team":
		return json.Marshal(r.Team)
	default:
		return nil, fmt.Errorf("unknown type: %s", r.Type)
	}
}

// UnmarshalJSON implements custom JSON unmarshaling for RelationshipResource to support the union type
func (r *RelationshipResource) UnmarshalJSON(data []byte) error {
	// First unmarshal to get the type
	var tmp struct {
		Type TypeIdentifier `json:"__typename"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return fmt.Errorf("failed to unmarshal type: %w", err)
	}
	r.Type = tmp.Type

	// Then unmarshal the specific type
	switch string(r.Type) {
	case "Domain":
		if err := json.Unmarshal(data, &r.Domain); err != nil {
			return fmt.Errorf("failed to unmarshal Domain: %w", err)
		}
	case "Infrastructure":
		if err := json.Unmarshal(data, &r.InfrastructureResource); err != nil {
			return fmt.Errorf("failed to unmarshal Infrastructure Resource: %w", err)
		}
	case "Service":
		if err := json.Unmarshal(data, &r.Service); err != nil {
			return fmt.Errorf("failed to unmarshal Service: %w", err)
		}
	case "System":
		if err := json.Unmarshal(data, &r.System); err != nil {
			return fmt.Errorf("failed to unmarshal System: %w", err)
		}
	case "Team":
		if err := json.Unmarshal(data, &r.Team); err != nil {
			return fmt.Errorf("failed to unmarshal Team: %w", err)
		}
	default:
		return fmt.Errorf("unknown resource type: %s", r.Type)
	}
	return nil
}
