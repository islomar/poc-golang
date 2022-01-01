/*
Copyright © 2020 The Pragmatic Programmers, LLC
Copyrights apply to this source code.
Check LICENSE for details.

*/
package cmd

import (
  "fmt"
  "io"
  "os"

  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  "pragprog.com/rggo/cobra/pScan/scan"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
  Use:          "add <host1>...<hostn>",
  Aliases:      []string{"a"},
  Short:        "Add new host(s) to list",
  SilenceUsage: true,
  Args:         cobra.MinimumNArgs(1),
  RunE: func(cmd *cobra.Command, args []string) error {
    hostsFile := viper.GetString("hosts-file")

    return addAction(os.Stdout, hostsFile, args)
  },
}

func addAction(out io.Writer, hostsFile string, args []string) error {
  hl := &scan.HostsList{}

  if err := hl.Load(hostsFile); err != nil {
    return err
  }

  for _, h := range args {
    if err := hl.Add(h); err != nil {
      return err
    }

    fmt.Fprintln(out, "Added host:", h)
  }

  return hl.Save(hostsFile)
}

func init() {
  hostsCmd.AddCommand(addCmd)

  // Here you will define your flags and configuration settings.

  // Cobra supports Persistent Flags which will work for this command
  // and all subcommands, e.g.:
  // addCmd.PersistentFlags().String("foo", "", "A help for foo")

  // Cobra supports local flags which will only run when this command
  // is called directly, e.g.:
  // addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
