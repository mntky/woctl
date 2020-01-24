package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"fmt"
)

var (
	cfgFile string

	//default cobra command
	RootCmd = &cobra.Command{
		Use:		"woyectl",
		Short:	"woyectl control lxc",
		Long:		`woyectl control lxc container.
woyectl get node,pod status and spec`,
		Run: runHelp,
	}

)

// create new cobra command 
func NewWoyectl() *cobra.Command {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringP("url", "", "", "api-server url")
	viper.BindPFlag("url", RootCmd.PersistentFlags().Lookup("url"))

	return RootCmd
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.w8aclient/")
	viper.SetConfigName("config")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}


func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
