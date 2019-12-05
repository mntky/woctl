package cmd

import (
	"fmt"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"woctl/pkg"
	//"github.com/lxc/lxd/shared/api"
)

func init() {
	RootCmd.AddCommand(newcreateCmd())
}

func newcreateCmd() *cobra.Command {
	createCmd := &cobra.Command {
		Use:		"create",
		Short:	"create lxc",
		Run:		func(cmd *cobra.Command, args []string) {
			//urlname, err := cmd.Flags().GetString("url")
			//if err != nil {
			//	fmt.Println(err)
			//}
			endpointurl := viper.GetString("url")

			containername, err := cmd.Flags().GetString("name")
			if err != nil {
				fmt.Println(err)
			}else if containername == "" {
				//ランダムなコンテナ名つける
				containername = pkg.Naming()
			}

			m := map[string]interface{}{
				"name":	containername,
				"replica": 2,
				"test": "test",
			}
			spec, err := json.Marshal(m)
			if err != nil {
				fmt.Println(err)
			}

			//body
			pdata := pkg.Postdata{
				endpointurl,
				containername,
				"create",
				spec,
			}
			err = pkg.Send(pdata)
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Println(containername)
			//fmt.Println(endpointurl)
		},
	}

	createCmd.PersistentFlags().StringP("name", "n", "", "container name")
	return createCmd
}
