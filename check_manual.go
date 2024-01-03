package opslevel

import "github.com/relvacode/iso8601"

type ManualCheckFragment struct {
	UpdateFrequency       *ManualCheckFrequency `graphql:"updateFrequency"`
	UpdateRequiresComment bool                  `graphql:"updateRequiresComment"`
}

type ManualCheckFrequency struct {
	StartingDate       iso8601.Time       `graphql:"startingDate"`
	FrequencyTimeScale FrequencyTimeScale `graphql:"frequencyTimeScale"`
	FrequencyValue     int                `graphql:"frequencyValue"`
}

// type ManualCheckFrequencyInput struct {
// 	StartingDate       iso8601.Time       `json:"startingDate" yaml:"startingDate" default:"2023-11-05T20:22:44.427Z"`
// 	FrequencyTimeScale FrequencyTimeScale `json:"frequencyTimeScale" yaml:"frequencyTimeScale" default:"day"`
// 	FrequencyValue     int                `json:"frequencyValue" yaml:"frequencyValue" default:"1"`
// }

func NewManualCheckFrequencyInput(startingDate string, timeScale FrequencyTimeScale, value int) *ManualCheckFrequencyInput {
	return &ManualCheckFrequencyInput{
		StartingDate:       NewISO8601Date(startingDate),
		FrequencyTimeScale: timeScale,
		FrequencyValue:     value,
	}
}

func NewManualCheckFrequencyUpdateInput(startingDate string, timeScale FrequencyTimeScale, value int) *ManualCheckFrequencyUpdateInput {
	startingDateIso := NewISO8601Date(startingDate)
	return &ManualCheckFrequencyUpdateInput{
		StartingDate:       &startingDateIso,
		FrequencyTimeScale: &timeScale,
		FrequencyValue:     &value,
	}
}

// type CheckManualCreateInput struct {
// 	CheckCreateInput

// 	UpdateFrequency       *ManualCheckFrequencyInput `json:"updateFrequency,omitempty" yaml:"updateFrequency,omitempty"`
// 	UpdateRequiresComment bool                       `json:"updateRequiresComment" yaml:"updateRequiresComment" default:"false"`
// }

// type CheckManualUpdateInput struct {
// 	CheckUpdateInput

// 	UpdateFrequency       *ManualCheckFrequencyInput `json:"updateFrequency,omitempty"`
// 	UpdateRequiresComment bool                       `json:"updateRequiresComment,omitempty"`
// }

func (client *Client) CreateCheckManual(input CheckManualCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkManualCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckManualCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckManual(input CheckManualUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkManualUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckManualUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}
