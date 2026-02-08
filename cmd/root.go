package cmd

import (
	"fmt"
	"os"

	"github.com/bagaskrmn/code-generator/constant"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   constant.AppName,
	Short: constant.AppDesc,
	Version: constant.Version,
	Long: constant.AppDescLong,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var verbose bool
func init() {
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	versionTemplate := fmt.Sprintf(`
	{{.Name}} 
	Version: {{.Version}}
	Author:  %s
	`, constant.Author)
	rootCmd.SetVersionTemplate(versionTemplate)
}
