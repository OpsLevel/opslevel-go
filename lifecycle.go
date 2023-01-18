package opslevel

type Lifecycle struct {
	Alias       string
	Description string
	Id          ID
	Index       int
	Name        string
}

func (client *Client) ListLifecycles() ([]Lifecycle, error) {
	var q struct {
		Account struct {
			Lifecycles []Lifecycle
		}
	}
	if err := client.Query(&q, nil); err != nil {
		return []Lifecycle{}, err
	}
	return q.Account.Lifecycles, nil
}
