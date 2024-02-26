package opslevel

type ManualCheckFragment struct {
	UpdateFrequency       *ManualCheckFrequency `graphql:"updateFrequency"`
	UpdateRequiresComment bool                  `graphql:"updateRequiresComment"`
}

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
