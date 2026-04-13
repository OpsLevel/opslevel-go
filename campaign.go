package opslevel

type ListCampaignsVariables struct {
	After  *string
	First  *int
	SortBy *CampaignSortEnum
	Status *CampaignStatusEnum
}

func (v *ListCampaignsVariables) AsPayloadVariables() *PayloadVariables {
	variables := PayloadVariables{}
	if v.After != nil {
		variables["after"] = *v.After
	}
	if v.First != nil {
		variables["first"] = *v.First
	}
	if v.SortBy != nil {
		variables["sortBy"] = *v.SortBy
	}
	if v.Status != nil {
		// cast status to match filter argument type
		variables["status"] = string(*v.Status)
	}
	return &variables
}

func (client *Client) CreateCampaign(input CampaignCreateInput) (*Campaign, error) {
	var m struct {
		Payload CampaignCreatePayload `graphql:"campaignCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CampaignCreate"))
	return &m.Payload.Campaign, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) GetCampaign(id ID) (*Campaign, error) {
	var q struct {
		Account struct {
			Campaign Campaign `graphql:"campaign(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id": id,
	}
	err := client.Query(&q, v, WithName("CampaignGet"))
	return &q.Account.Campaign, HandleErrors(err, nil)
}

func (client *Client) UpdateCampaign(input CampaignUpdateInput) (*Campaign, error) {
	var m struct {
		Payload CampaignUpdatePayload `graphql:"campaignUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CampaignUpdate"))
	return &m.Payload.Campaign, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteCampaign(id ID) error {
	var m struct {
		Payload CampaignDeletePayload `graphql:"campaignDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": CampaignDeleteInput{Id: id},
	}
	err := client.Mutate(&m, v, WithName("CampaignDelete"))
	return HandleErrors(err, m.Payload.Errors)
}

func (client *Client) ScheduleCampaign(input CampaignScheduleUpdateInput) (*Campaign, error) {
	var m struct {
		Payload CampaignSchedulePayload `graphql:"campaignSchedule(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CampaignSchedule"))
	return &m.Payload.Campaign, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UnscheduleCampaign(id ID) (*Campaign, error) {
	var m struct {
		Payload CampaignUnschedulePayload `graphql:"campaignUnschedule(input: $input)"`
	}
	v := PayloadVariables{
		"input": struct {
			Id ID `json:"id"`
		}{Id: id},
	}
	err := client.Mutate(&m, v, WithName("CampaignUnschedule"))
	return &m.Payload.Campaign, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) ListCampaigns(campaignVariables *ListCampaignsVariables) (*CampaignConnection, error) {
	if campaignVariables == nil {
		campaignVariables = &ListCampaignsVariables{}
	}

	defaultPages := client.InitialPageVariablesPointer()

	if campaignVariables.First == nil {
		defaultFirst := (*defaultPages)["first"].(int)
		campaignVariables.First = &defaultFirst
	}

	if campaignVariables.After == nil {
		defaultAfter := (*defaultPages)["after"].(string)
		campaignVariables.After = &defaultAfter
	}

	if campaignVariables.SortBy == nil {
		campaignVariables.SortBy = &CampaignSortEnumStartDateDesc
	}

	if campaignVariables.Status == nil {
		campaignVariables.Status = &CampaignStatusEnumInProgress
	}

	variables := campaignVariables.AsPayloadVariables()

	var q struct {
		Account struct {
			Campaigns CampaignConnection `graphql:"campaigns(first: $first, after: $after, sortBy: $sortBy, filter: [{key: status, arg: $status}])"`
		}
	}

	if err := client.Query(&q, *variables, WithName("CampaignsList")); err != nil {
		return nil, err
	}

	if q.Account.Campaigns.PageInfo.HasNextPage {
		campaignVariables.After = &q.Account.Campaigns.PageInfo.End
		resp, err := client.ListCampaigns(campaignVariables)
		if err != nil {
			return nil, err
		}
		q.Account.Campaigns.Nodes = append(q.Account.Campaigns.Nodes, resp.Nodes...)
		q.Account.Campaigns.PageInfo = resp.PageInfo

	}
	q.Account.Campaigns.TotalCount = len(q.Account.Campaigns.Nodes)
	return &q.Account.Campaigns, nil
}
