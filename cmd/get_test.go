package cmd

import (
	//"fmt"
	"errors"
	"woctl/pkg"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(newgetCmd())
}

func newgetCmd() *cobra.Command {
	getCmd := &cobra.Command {
		Use:		"get",
		Short:	"get woyendetsa config and status",
		RunE:		func(cmd *cobra.Command, args []string) error {
			endpointurl := viper.GetString("url")
			containername, err := cmd.Flags().GetString("container")
			if err != nil {
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
				"get",
				spec,
			}
			err = pkg.Send(pdata)
			if err != nil {
				return err
			}
			return nil
		},
	}

	//getCmd.PersistentFlags().StringP("node", "", "", "node name")
	getCmd.PersistentFlags().StringP("name", "", "", "spec name")
	return getCmd
}
