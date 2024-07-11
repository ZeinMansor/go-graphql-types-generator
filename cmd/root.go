package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Root CMD",
	Short: "Shorte Description",
	Long:  `Root command`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}
}
