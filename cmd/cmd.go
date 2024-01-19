package cmd

import (
	"CoinDirectCli/cdclient"
	"fmt"

	"github.com/spf13/cobra"
)

func Init() {
	var client = &cobra.Command{
		Use:   "cmd",
		Short: "a cli for retrieving data from the .coindirect.com api",
		Long:  "A cli for retrieving data from the .coindirect.com api\n\n-a ascending : defaults to false\n-c currencymap : return a map of which currencies are used by which lands",
		Run: func(cmd *cobra.Command, args []string) {

			sortKey, _ := cmd.Flags().GetString("sortkey")
			isDescending, _ := cmd.Flags().GetBool("descending")
			isCurrencyMap, _ := cmd.Flags().GetBool("currencymap")

			sortBy, err := cdclient.ParseSortBy(sortKey)
			if err != nil {
				fmt.Printf("Error parsing sortvalue, should be one of id, name or currency")
				return
			}

			cdclient.FetchCountries(isDescending, sortBy, isCurrencyMap)
			fmt.Println("-h  for help")
		},
	}
	client.Flags().StringP("sortkey", "s", "id", "id|name|currency")
	client.Flags().BoolP("descending", "d", false, "descending sort by sortkey, default is ascending order")
	client.Flags().BoolP("currencymap", "c", false, "return a map of which currencies are used by which lands")
	client.Execute()
}
