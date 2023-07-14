package cmd

import "github.com/spf13/cobra"

func init() {}

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(sqlToStructCmd)
	rootCmd.AddCommand(sqlFileToStructCmd)
}
