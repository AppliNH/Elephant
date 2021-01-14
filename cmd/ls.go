package cmd

import (
	"github.com/applinh/elephant/commands"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all running stacks",
	Long:  `List all running stacks`,
	Run: func(cmd *cobra.Command, args []string) {

		commands.List(db)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	//walkCmd.Flags().StringP("file", "f", "", "compose file")
}
