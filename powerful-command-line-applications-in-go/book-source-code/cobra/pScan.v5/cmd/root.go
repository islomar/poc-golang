/*
Copyright © 2020 The Pragmatic Programmers, LLC
Copyrights apply to this source code.
Check LICENSE for details.

*/
package cmd

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"

  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "pScan",
  Short: "Fast TCP port scanner",
  Long: `pScan - short for Port Scanner - executes TCP port scan
on a list of hosts.

pScan allows you to add, list, and delete hosts from the list.

pScan executes a port scan on specified TCP ports. You can customize the
target ports using a command line flag.`,
  Version: "0.1",
  // Uncomment the following line if your bare application
  // has an action associated with it:
  //  Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and
// sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  // Here you will define your flags and configuration settings.
  // Cobra supports persistent flags, which, if defined here,
  // will be global for your application.

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
    "config file (default is $HOME/.pScan.yaml)")

  rootCmd.PersistentFlags().StringP("hosts-file", "f", "pScan.hosts",
    "pScan hosts file")

  versionTemplate := `{{printf "%s: %s - version %s\n" .Name .Short .Version}}`
  rootCmd.SetVersionTemplate(versionTemplate)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
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

    // Search config in home directory with name ".pScan" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".pScan")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}
