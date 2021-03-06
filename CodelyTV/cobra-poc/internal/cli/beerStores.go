package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var stores = map[string]string{
	"01DC9ZAPGKEQJS4P4A48EG3P43": "Mercadona",
	"01DC9ZB23EW0J0ARAER09SJDKC": "Carrefour",
	"01DC9ZB89V1PQD977ZE6QXSQHH": "Alcampo",
}

func InitBeerStoresCmd() *cobra.Command {
	beerStoresCmd := &cobra.Command{
		Use:   "stores",
		Short: "Print stores where beers are sold",
		Run:   runBeerStoresFn(),
	}

	beerStoresCmd.Flags().StringP(idFlag, "i", "", "id of the beer")

	return beerStoresCmd
}

func runBeerStoresFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			fmt.Println(stores[id])
		} else {
			fmt.Println(stores)
		}
	}
}
