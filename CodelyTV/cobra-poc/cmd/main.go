package main

import (
	"github.com/islomar/poc-golang/codely.tv/cobra-poc/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd())
	rootCmd.AddCommand(cli.InitBeerStoresCmd())
	rootCmd.Execute()
}