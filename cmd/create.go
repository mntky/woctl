package cmd

import (
	"fmt"
	"io/ioutil"
	//"encoding/json"

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
			endpointurl := viper.GetString("url")

			containername, err := cmd.Flags().GetString("name")
			if err != nil {
				fmt.Println(err)
			}else if containername == "" {
				containername = pkg.Naming()
			}

			yamlfilename, err := cmd.Flags().GetString("file")
			if err != nil {
				fmt.Println(err)
			}
			//TODO yamlの構文チェック機能をつけたい。
			spec, err := ioutil.ReadFile(yamlfilename)
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
		},
	}

	createCmd.PersistentFlags().StringP("name", "n", "", "container name")
	createCmd.PersistentFlags().StringP("file", "f", "", "select yaml path")
	return createCmd
}
