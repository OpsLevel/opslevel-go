package opslevel_example

var query struct {
	Account struct {
		Tiers []Tier
	}
}
if err := client.Query(&query, nil); err != nil {
	panic(err)
}
for _, tier := range m.Account.Tiers {
	fmt.Println(tier.Name)
}