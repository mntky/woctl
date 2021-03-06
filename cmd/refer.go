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
			specname, err := cmd.Flags().GetString("name")
			if err != nil {
				return errors.New("missing argment")
			}
			m := map[string]interface{}{
				"name": specname,
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

	getCmd.PersistentFlags().StringP("name", "n", "", "spec name")
	return getCmd
}
