package opslevel_example

import (
	"fmt"

	"github.com/opslevel/opslevel-go/v2022"
)

func init() {
	var query struct {
		Account struct {
			Tiers []opslevel.Tier
		}
	}

	client := opslevel.NewClient("xxx")
	if err := client.Query(&query, nil); err != nil {
		panic(err)
	}
	for _, tier := range query.Account.Tiers {
		fmt.Println(tier.Name)
	}
}
