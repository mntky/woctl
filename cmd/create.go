package cmd

import (
	//"fmt"
	"io/ioutil"
	//"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"woctl/pkg"
)

func init() {
	RootCmd.AddCommand(newcreateCmd())
}

func newcreateCmd() *cobra.Command {
	createCmd := &cobra.Command {
		Use:		"create",
		Short:	"create lxc",
		RunE:		func(cmd *cobra.Command, args []string) error {
			endpointurl := viper.GetString("url")
			specname, err := cmd.Flags().GetString("file")
			if err != nil {
				return err
			}
			//TODO yamlの構文チェック機能をつけたい。
			spec, err := ioutil.ReadFile(specname)
			if err != nil {
				return err
			}

			//body
			pdata := pkg.Postdata{
				endpointurl,
				"create",
				spec,
			}
			err = pkg.Send(pdata)
			if err != nil {
				return err
			}
			return nil
		},
	}

	createCmd.PersistentFlags().StringP("file", "f", "", "select yaml path")
	return createCmd
}
