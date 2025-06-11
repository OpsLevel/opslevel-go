package opslevel

func (client *Client) ListCampaigns(variables *PayloadVariables) (*CampaignConnection, error) {
	var q struct {
		Account struct {
			Campaigns CampaignConnection `graphql:"campaigns(first: $first, after: $after, filter: $filter)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}

	if err := client.Query(&q, *variables, WithName("CampaignsList")); err != nil {
		return nil, err
	}

	if q.Account.Campaigns.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Campaigns.PageInfo.End
		resp, err := client.ListCampaigns(variables)
		if err != nil {
			return nil, err
		}
		q.Account.Campaigns.Nodes = append(q.Account.Campaigns.Nodes, resp.Nodes...)
		q.Account.Campaigns.PageInfo = resp.PageInfo

	}
	return &q.Account.Campaigns, nil
}
