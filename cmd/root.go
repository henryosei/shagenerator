/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shagenerator [text to hash]",
	Short: "Generates a SHA256 HASH for any string value or set of string values provided.",
	Long: `Generates a SHA256 HASH for any string value or set of string values provided. 
This program accepts at least one string argument. Single string values with spaces should be placed in quotes. 
Multiple strings to be hashed must be passed as arguments to the command separated by a space.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			hasher := sha256.New()
			hasher.Write([]byte(arg))
			sha := hasher.Sum(nil)
			fmt.Printf("Original Text: %s \nHashed text: %x \n----------------\n", arg, sha)
		}
	},
	Args: cobra.MinimumNArgs(1),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.shagenerator.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
