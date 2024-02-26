package opslevel

func (client *Client) ListTiers() ([]Tier, error) {
	var q struct {
		Account struct {
			Tiers []Tier
		}
	}
	err := client.Query(&q, nil, WithName("TierList"))
	return q.Account.Tiers, HandleErrors(err, nil)
}
