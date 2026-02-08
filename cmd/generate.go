/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate [name]",
	Short: "Scaffold a complete module",
	Args:  cobra.ExactArgs(1), // Ensures the user provides exactly one name
	Run: func(cmd *cobra.Command, args []string) {
		entity := args[0] // e.g., "User"

		// Define the suffixes you want
		suffixes := []string{
			"Repository.go",
			"Service.go",
			"Handler.go",
			"Routes.go",
			"Presenter.go",
		}

		fmt.Printf("Generating boilerplate for: %s...\n", entity)

		for _, suffix := range suffixes {
			fileName := fmt.Sprintf("%s%s", entity, suffix)
			content := fmt.Sprintf("package main\n\n// %s logic goes here\ntype %s%s struct {}\n", entity, entity, suffix[:len(suffix)-3])

			err := os.WriteFile(fileName, []byte(content), 0644)
			if err != nil {
				fmt.Printf("Failed to create %s: %v\n", fileName, err)
				continue
			}
			fmt.Printf("  - Created %s\n", fileName)
		}
		
		fmt.Println("Done!")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
