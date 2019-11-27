package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(newdeleteCmd())
}

func newdeleteCmd() *cobra.Command {
	deleteCmd := &cobra.Command {
		Use:		"delete",
		Short:	"delete lxc",
		Run:		func(cmd *cobra.Command, args []string) {
			containername, err := cmd.Flags().GetString("name")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(containername)
		},
	}

	deleteCmd.PersistentFlags().StringP("name", "n", "", "container name")
	return deleteCmd
}
