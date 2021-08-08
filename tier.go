package opslevel

import "github.com/shurcooL/graphql"

type Tier struct {
	Alias       string
	Description string
	Id          graphql.ID
	Index       int
	Name        string
}

func (client *Client) ListTiers() ([]Tier, error) {
	var q struct {
		Account struct {
			Tiers []Tier
		}
	}
	if err := client.Query(&q, nil); err != nil {
		return []Tier{}, err
	}
	return q.Account.Tiers, nil
}
