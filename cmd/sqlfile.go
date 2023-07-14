package cmd

import (
	"github.com/spf13/cobra"
)

var sqlfile string

var sqlFileToStructCmd = &cobra.Command{
	Use:   "file",
	Short: "sqlfile to struct",
	Long:  "sqlfile to struct",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	sqlFileToStructCmd.Flags().StringVarP(&sqlfile, "file", "f", "", "SQL文件地址")
}
