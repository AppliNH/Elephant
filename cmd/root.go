package cmd

import (
	"fmt"
	"os"

	"github.com/applinh/elephant/kvdb"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var db *bolt.DB

var rootCmd = &cobra.Command{
	Use:   "elephant",
	Short: "Big and solid docker orchestrator 🐘",
	Long: `Elephant is a big and solid docker orchestrator.
	Or at least, he tries to be. Don't be too hard with him.`,
}

// Execute executes the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	fmt.Println("🐘")
	fmt.Println()
	var err error
	db, err = kvdb.InitDB()
	if err != nil {
		panic(err)
	}

}
