package cmd

import (
	"fmt"

	"github.com/TheRSTech/Backendforger-backend/cmd/generator"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "backendforger",
	Short: "Backendforger is a backend project starter generation tool",
	Long:  "Backendforger is a CLI tool to generate backend projects with different languages and frameworks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'backendforger --help' to see available commands")
	},
}

var createGoAppCmd = &cobra.Command{
	Use:   "create-go-app [name]",
	Short: "Create a new backend project in Go",
	Args:  cobra.ExactArgs(1), // Expect exactly one argument (the app name)
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]
		framework, _ := cmd.Flags().GetString("framework")
		database, _ := cmd.Flags().GetString("database")
		orm, _ := cmd.Flags().GetString("orm")

		fmt.Printf("Creating golang app '%s' with framework: %s, database: %s, orm: %s\n",
			appName, framework, database, orm)

		// Call your generateGoProject function with appName, framework, database, orm
		generator.GenerateGoProject(appName, framework, database, orm)
	},
}

var createPythonAppCmd = &cobra.Command{
	Use:   "create-python-app [name]",
	Short: "Create a new backend project in Python",
	Args:  cobra.ExactArgs(1), // Expect exactly one argument (the app name)
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]
		framework, _ := cmd.Flags().GetString("framework")
		database, _ := cmd.Flags().GetString("database")
		orm, _ := cmd.Flags().GetString("orm")

		fmt.Printf("Creating python app '%s' with framework: %s, database: %s, orm: %s\n",
			appName, framework, database, orm)

		// Call your generatePythonProject function with appName, framework, database, orm
		generator.GeneratePythonProject(appName, framework, database, orm)
	},
}

var createNodeAppCmd = &cobra.Command{
	Use:   "create-node-app [name]",
	Short: "Create a new Node.js project",
	Args:  cobra.ExactArgs(1), // Expect exactly one argument (the app name)
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]
		ts, _ := cmd.Flags().GetBool("typescript")
		framework, _ := cmd.Flags().GetString("framework")
		database, _ := cmd.Flags().GetString("database")
		orm, _ := cmd.Flags().GetString("orm")

		if ts {
			fmt.Printf("Creating Node.js app '%s' with TypeScript, framework: %s, database: %s, orm: %s\n",
				appName, framework, database, orm)
		} else {
			fmt.Printf("Creating Node.js app '%s' with JavaScript, framework: %s, database: %s, orm: %s\n",
				appName, framework, database, orm)
		}

		generator.GenerateNodeProject(appName, framework, database, orm, ts)
	},
}

func init() {
	// Define flags for createGoAppCmd
	createGoAppCmd.Flags().StringP("framework", "f", "", "Framework (e.g. gin, echo, flask, express, fastapi, fiber, mux)")
	createGoAppCmd.Flags().StringP("database", "d", "", "Database (e.g. sqlite, postgres, mysql, mongodb)")
	createGoAppCmd.Flags().StringP("orm", "o", "", "ORM (optional)")
	createGoAppCmd.MarkFlagRequired("framework")

	// Define flags for createPythonAppCmd
	createPythonAppCmd.Flags().StringP("framework", "f", "", "Framework (e.g. flask, fast api)")
	createPythonAppCmd.Flags().StringP("database", "d", "", "Database (e.g. sqlite, postgres, mysql")
	createPythonAppCmd.Flags().StringP("orm", "o", "", "ORM (optional)")
	createPythonAppCmd.MarkFlagRequired("framework")

	// Define flags for createNodeAppCmd
	createNodeAppCmd.Flags().BoolP("typescript", "t", false, "Use TypeScript for Node.js")
	createNodeAppCmd.Flags().StringP("framework", "f", "", "Framework (e.g. express)")
	createNodeAppCmd.Flags().StringP("database", "d", "", "Database (e.g. mongodb)")
	createNodeAppCmd.Flags().StringP("orm", "o", "", "ORM (optional)")
	createNodeAppCmd.MarkFlagRequired("framework")

	// Add createGoAppCmd and createNodeAppCmd to rootCmd

	RootCmd.AddCommand(createGoAppCmd)
	RootCmd.AddCommand(createPythonAppCmd)
	RootCmd.AddCommand(createNodeAppCmd)

	// Set a custom error handler
	RootCmd.CompletionOptions.DisableDefaultCmd = true

	RootCmd.SuggestionsMinimumDistance = 1
}
