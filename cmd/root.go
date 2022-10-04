package cmd

import (
	"better-dev-container/cmd/docker"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:                   "better-dev-container [shellCmd ...]",
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	Aliases: []string{
		"bdev",
	},
	Args: cobra.MinimumNArgs(1),
	Run:  docker.RunCommandInContainer,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find working directory.
		workingDir, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in working directory with name ".better-dev-container" (without extension).
		viper.AddConfigPath(workingDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".better-dev-container")

		viper.SetDefault("name", filepath.Base(workingDir))
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
	}
}
