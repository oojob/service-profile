package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string

	// RootCmd root coommand
	RootCmd = &cobra.Command{
		Use:   "company [OPTIONS] [COMMANDS]",
		Short: "Company Service Backend Application",
		Long:  `Company Service is part of job board microservice and is created as a CLI application using Go.This application is responsible for company entity operation.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
)

// Execute executes the root command.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("unable to read config: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is config.yaml)")
	RootCmd.PersistentFlags().StringP("author", "a", "nirajgeorgian", "author name for copyright attribution")
	RootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
	viper.SetDefault("author", "nirajgeorgian <nirajgeorgian@gmail.com>")
}
