package cmd

import (
	"fmt"
	"io/ioutil"
	"woctl/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(newdeleteCmd())
}

func newdeleteCmd() *cobra.Command {
	deleteCmd := &cobra.Command {
		Use:		"delete",
		Short:	"delete spec",
		RunE:		func(cmd *cobra.Command, args []string) error{
			endpointurl := viper.GetString("url")
			specyaml, err := cmd.Flags().GetString("file")
			if err != nil {
				return err
			}
			spec, err := ioutil.ReadFile(specyaml)
			if err != nil {
				return err
			}

			//debug
			fmt.Printf("%T\n",spec)

			pdata := pkg.Postdata{
				endpointurl,
				"delete",
				spec,
			}
			err = pkg.Send(pdata)
			if err != nil {
				return err
			}
			return nil
		},
	}

	deleteCmd.PersistentFlags().StringP("file", "f", "", "spec file name")
	return deleteCmd
}
