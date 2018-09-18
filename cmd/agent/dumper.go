package main

import (
	"fmt"
	"github.com/SantoDE/datahamster/configuration"
	"github.com/SantoDE/datahamster/dumper"
	"github.com/SantoDE/datahamster/log"
	"github.com/Sirupsen/logrus"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var cfgFile string
var logLevel string
var server string
var interval int32
var config configuration.DumperConfiguration

// RootCmd represents the base command when called without any subcommands
var dumperCommand = &cobra.Command{
	Use:   "datahamster-agent",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.Unmarshal(&config)

		if err != nil {
			fmt.Printf("unable to decode into struct, %v", err)
		}

		level, err := logrus.ParseLevel(strings.ToLower(config.LogLevel))
		if err != nil {
			log.Error("Error getting level", err)
		}
		log.SetLevel(level)

		dumper.StartWorker(&config)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := dumperCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initDumperConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	dumperCommand.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config.toml)")
	dumperCommand.PersistentFlags().StringVar(&logLevel, "logLevel", "error", "Dumper Application Log Level")
	dumperCommand.PersistentFlags().StringVar(&server, "server", "error", "Server to connect to")
	dumperCommand.PersistentFlags().Int32Var(&interval, "interval", 60, "Interval to Poll Jobs")

	viper.BindPFlag("config", dumperCommand.PersistentFlags().Lookup("config"))
	viper.BindPFlag("logLevel", dumperCommand.PersistentFlags().Lookup("logLevel"))
	viper.BindPFlag("server", dumperCommand.PersistentFlags().Lookup("server"))
	viper.BindPFlag("interval", dumperCommand.PersistentFlags().Lookup("interval"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	dumperCommand.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initDumperConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".dodo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("Error reading config file %s", err.Error())
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
