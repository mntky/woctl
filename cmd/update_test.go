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
	RootCmd.AddCommand(newupdateCmd())
}

func newupdateCmd() *cobra.Command {
	updateCmd := &cobra.Command {
		Use:		"update",
		Short:	"update lxc",
		Run:		func(cmd *cobra.Command, args []string) {
			endpointurl := viper.GetString("url")

			containername, err := cmd.Flags().GetString("name")
			if err != nil {
				fmt.Println(err)
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
				"update",
				spec,
			}
			err = pkg.Send(pdata)
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	//TODO updateする際の名前指定はyamlにmetadataフィールド追加してその名前を基準にアップデートする。
	updateCmd.PersistentFlags().StringP("name", "n", "", "container name")
	updateCmd.PersistentFlags().StringP("file", "f", "", "select yaml path")
	return updateCmd
}

