package cmd

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/spf13/cobra"
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
			containername, err := cmd.Flags().GetString("name")
			if err != nil {
				fmt.Println(err)
			}else if containername == "" {
				containername = strconv.Itoa(int(rand.Int63()))
			}
			fmt.Println(containername)
		},
	}

	createCmd.PersistentFlags().StringP("name", "n", "", "container name")
	return createCmd
}
