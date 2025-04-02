package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var (
	inputFile string
	rootCmd = &cobra.Command{
		Use:   "namemash",
		Short: "Creating a user name list for brute force attacks",
		Run: func(cmd *cobra.Command, args []string) { 
			fmt.Println("Reading data from:", inputFile)
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&inputFile, "input", "", "File to read initial names from")
	rootCmd.MarkPersistentFlagRequired("input")
}


