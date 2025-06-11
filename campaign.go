package opslevel

func (client *Client) ListCampaigns(variables *PayloadVariables, sortBy *CampaignSortEnum) (*CampaignConnection, error) {
	var q struct {
		Account struct {
			Campaigns CampaignConnection `graphql:"campaigns(first: $first, after: $after, sortBy: $sortBy)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if sortBy == nil {
		sortBy = &CampaignSortEnumStartDateDesc
	}

	(*variables)["sortBy"] = CampaignSortEnumStartDateDesc


	if err := client.Query(&q, *variables, WithName("CampaignsList")); err != nil {
		return nil, err
	}

	if q.Account.Campaigns.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Campaigns.PageInfo.End
		resp, err := client.ListCampaigns(variables, sortBy) // not sure if I need this with the page info
		if err != nil {
			return nil, err
		}
		q.Account.Campaigns.Nodes = append(q.Account.Campaigns.Nodes, resp.Nodes...)
		q.Account.Campaigns.PageInfo = resp.PageInfo

	}
	q.Account.Campaigns.TotalCount = len(q.Account.Campaigns.Nodes)
	return &q.Account.Campaigns, nil
}
