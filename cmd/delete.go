package cmd

import (
	"errors"
	"woctl/pkg"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(newdeleteCmd())
}

func newdeleteCmd() *cobra.Command {
	deleteCmd := &cobra.Command {
		Use:		"delete",
		Short:	"delete lxc and spec",
		RunE:		func(cmd *cobra.Command, args []string) error{
			endpointurl := viper.GetString("url")
			containername, err := cmd.Flags().GetString("name")
			if containername == "" {
				return errors.New("missing argment")
			}
			m := map[string]interface{}{
				"name": containername,
			}
			spec, err := json.Marshal(m)
			if err != nil {
				return err
			}
			pdata := pkg.Postdata{
				endpointurl,
				containername,
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

	deleteCmd.PersistentFlags().StringP("name", "n", "", "container name")
	return deleteCmd
}
