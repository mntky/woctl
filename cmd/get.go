package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(newgetCmd())
}

func newgetCmd() *cobra.Command {
	getCmd := &cobra.Command {
		Use:		"get",
		Short:	"get woyendetsa config and status",
		Run:		func(cmd *cobra.Command, args []string) {
			containername, err := cmd.Flags().GetString("name")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(containername)
			fmt.Println(len(args))
		},
	}

	getCmd.PersistentFlags().StringP("node", "", "", "node name")
	getCmd.PersistentFlags().StringP("container", "", "", "container name")
	return getCmd
}
