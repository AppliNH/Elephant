package cmd

import (
	"github.com/applinh/elephant/commands"
	"github.com/spf13/cobra"
)

var walkCmd = &cobra.Command{
	Use:   "walk",
	Short: "Run a stack",
	Long:  `Run a stack from a compose file.`,
	Run: func(cmd *cobra.Command, args []string) {
		composeFile, _ := cmd.Flags().GetString("file")
		elephantName, _ := cmd.Flags().GetString("elephant-name")

		commands.Walk(db, composeFile, elephantName)
	},
}

func init() {
	rootCmd.AddCommand(walkCmd)

	walkCmd.Flags().StringP("file", "f", "", "compose file")
	walkCmd.Flags().StringP("elephant-name", "e", "", "elephant name that will hold your stack")
}
