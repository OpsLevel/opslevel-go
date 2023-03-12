package opslevel

type Lifecycle struct {
	Alias       string `json:"alias"`
	Description string `json:"description"`
	Id          ID     `json:"id"`
	Index       int    `json:"index"`
	Name        string `json:"name"`
}

func (client *Client) ListLifecycles() ([]Lifecycle, error) {
	var q struct {
		Account struct {
			Lifecycles []Lifecycle
		}
	}
	err := client.Query(&q, nil, WithName("LifecycleList"))
	return q.Account.Lifecycles, HandleErrors(err, nil)
}
