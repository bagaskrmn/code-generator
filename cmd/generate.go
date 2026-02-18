/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/bagaskrmn/code-generator/internal/generator"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate [name]",
	Short: "Generate ur master code",
	Args:  cobra.ExactArgs(1), // Ensures the user provides exactly one name
	Run: func(cmd *cobra.Command, args []string) {
		entity := args[0]

		var repoName string
		fmt.Print("Enter ABA-Developer Repository Name (e.g., my-project-api): ")
		fmt.Scanln(&repoName)

		if repoName == "" {
			log.Fatal("Repository name cannot be empty")
		}
		fmt.Printf("Generating code for: %s\n", entity)

		// 1. Generate Repository
		if err := generator.GenerateRepository(entity); err != nil {
			log.Fatalf("Error generating Repository: %v", err)
		}

		// 2. Generate Service
		if err := generator.GenerateService(entity); err != nil {
			log.Fatalf("Error generating Service: %v", err)
		}

		// 3. Generate Handler
		if err := generator.GenerateHandler(entity); err != nil {
			log.Fatalf("Error generating Handler: %v", err)
		}

		// 4. Generate Presenter
		if err := generator.GeneratePresenter(entity); err != nil {
			log.Fatalf("Error generating Presenter: %v", err)
		}

		// 5. Generate Routes
		if err := generator.GenerateRoutes(entity, repoName); err != nil {
			log.Fatalf("Error generating Routes: %v", err)
		}

		fmt.Printf("\nSuccessfully generated all modules for %s!\n", entity)
		fmt.Println("Check your project folders to see the new files.")

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
