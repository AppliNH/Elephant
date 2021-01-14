package cmd

import (
	"github.com/applinh/elephant/commands"
	"github.com/spf13/cobra"
)

var stompCmd = &cobra.Command{
	Use:   "stomp",
	Short: "Stop a stack",
	Long:  `Stop a stack by providing an elephant name`,
	Run: func(cmd *cobra.Command, args []string) {

		commands.Stomp(db, args[0])
	},
}

func init() {
	rootCmd.AddCommand(stompCmd)

	//walkCmd.Flags().StringP("file", "f", "", "compose file")
}
