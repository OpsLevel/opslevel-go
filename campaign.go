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
		variables["status"] = *v.Status
	}
	return &variables
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
			Campaigns CampaignConnection `graphql:"campaigns(first: $first, after: $after, sortBy: $sortBy, status: $status)"`
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
