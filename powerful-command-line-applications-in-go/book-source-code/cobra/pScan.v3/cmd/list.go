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
  "pragprog.com/rggo/cobra/pScan/scan"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
  Use:     "list",
  Aliases: []string{"l"},
  Short:   "List hosts in hosts list",
  RunE: func(cmd *cobra.Command, args []string) error {
    hostsFile, err := cmd.Flags().GetString("hosts-file")
    if err != nil {
      return err
    }

    return listAction(os.Stdout, hostsFile, args)
  },
}

func listAction(out io.Writer, hostsFile string, args []string) error {
  hl := &scan.HostsList{}

  if err := hl.Load(hostsFile); err != nil {
    return err
  }

  for _, h := range hl.Hosts {
    if _, err := fmt.Fprintln(out, h); err != nil {
      return err
    }
  }

  return nil
}

func init() {
  hostsCmd.AddCommand(listCmd)

  // Here you will define your flags and configuration settings.

  // Cobra supports Persistent Flags which will work for this command
  // and all subcommands, e.g.:
  // listCmd.PersistentFlags().String("foo", "", "A help for foo")

  // Cobra supports local flags which will only run when this command
  // is called directly, e.g.:
  // listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
