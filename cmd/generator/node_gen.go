package generator

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/TheRSTech/Backendforger-backend/cmd/utils"
	"github.com/fatih/color"
)

// GenerateNodeProject generates a Node.js project with the specified options.
func GenerateNodeProject(projectName, framework, database, orm string, ts bool) error {
	startTime := time.Now()
	fmt.Println("Generating Node.js project...")

	// Create the project directory
	if err := os.Mkdir(projectName, 0755); err != nil {
		fmt.Println("Error creating project directory:", err)
		return err
	}

	var wg sync.WaitGroup

	// Initialize a new Node.js project
	cmd := exec.Command("npm", "init", "-y")
	cmd.Dir = projectName
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error initializing Node.js project:", err)
		return err
	}

	if ts {
		// Handle TypeScript project setup
		// Implement TypeScript project setup if needed
		// Create src and nested directories
		var dirs []string

		if orm == "drizzle" {
			dirs = []string{"src", "src/routes", "src/db", "src/controllers", "src/db/schema", "src/constants", "src/middlewares", "src/utils", "src/types"}

			for _, dir := range dirs {
				if err := os.MkdirAll(fmt.Sprintf("%s/%s", projectName, dir), 0755); err != nil {
					fmt.Println("Error creating directory:", err)
					return err
				}
			}

			if database == "mysql" {
				files := map[string]string{
					fmt.Sprintf("%s/tsconfig.json", projectName):                      "templates/node/ts/tsconfig.txt",
					fmt.Sprintf("%s/package.json", projectName):                       "templates/node/ts/drizzle/db/ms/package.txt",
					fmt.Sprintf("%s/src/routes/user-routes.ts", projectName):          "templates/node/ts/drizzle/user-routes.txt",
					fmt.Sprintf("%s/src/db/schema/user.ts", projectName):              "templates/node/ts/drizzle/db/ms/user-ms.txt",
					fmt.Sprintf("%s/src/db/setup.ts", projectName):                    "templates/node/ts/drizzle/db/ms/msql-setup.txt",
					fmt.Sprintf("%s/src/controllers/user-controller.ts", projectName): "templates/node/ts/drizzle/db/ms/user-controller.txt",
					fmt.Sprintf("%s/src/index.ts", projectName):                       "templates/node/ts/drizzle/index.txt",
					fmt.Sprintf("%s/drizzle.config.ts", projectName):                  "templates/node/ts/drizzle/drizzle.config.txt",
				}
				// Use a buffered channel to limit concurrent template copying
				copyCh := make(chan struct{}, 5) // Adjust the buffer size as needed

				for dest, src := range files {
					wg.Add(1)
					copyCh <- struct{}{} // Add a token to limit concurrency
					go func(dest, src string) {
						defer func() {
							<-copyCh // Release token after copying is done
							wg.Done()
						}()
						if err := utils.CopyTemplate("backendforger", src, dest); err != nil {
							fmt.Println("Error copying template:", err)
						}
					}(dest, src)
				}

				// Close the channel after all copying tasks are added
				close(copyCh)
			} else if database == "postgres" {
				files := map[string]string{
					fmt.Sprintf("%s/tsconfig.json", projectName):                      "templates/node/ts/tsconfig.txt",
					fmt.Sprintf("%s/package.json", projectName):                       "templates/node/ts/drizzle/db/pg/package.txt",
					fmt.Sprintf("%s/src/routes/user-routes.ts", projectName):          "templates/node/ts/drizzle/user-routes.txt",
					fmt.Sprintf("%s/src/db/schema/user.ts", projectName):              "templates/node/ts/drizzle/db/pg/user-pg.txt",
					fmt.Sprintf("%s/src/db/setup.ts", projectName):                    "templates/node/ts/drizzle/db/pg/pg-setup.txt",
					fmt.Sprintf("%s/src/controllers/user-controller.ts", projectName): "templates/node/ts/drizzle/db/pg/user-controller.txt",
					fmt.Sprintf("%s/src/index.ts", projectName):                       "templates/node/ts/drizzle/index.txt",
					fmt.Sprintf("%s/drizzle.config.ts", projectName):                  "templates/node/ts/drizzle/drizzle.config.txt",
				}
				// Use a buffered channel to limit concurrent template copying
				copyCh := make(chan struct{}, 5) // Adjust the buffer size as needed

				for dest, src := range files {
					wg.Add(1)
					copyCh <- struct{}{} // Add a token to limit concurrency
					go func(dest, src string) {
						defer func() {
							<-copyCh // Release token after copying is done
							wg.Done()
						}()
						if err := utils.CopyTemplate("backendforger", src, dest); err != nil {
							fmt.Println("Error copying template:", err)
						}
					}(dest, src)
				}

				// Close the channel after all copying tasks are added
				close(copyCh)
			}

			// Install dependencies concurrently
			wg.Add(1)
			go func() {
				defer wg.Done()
				cmd := exec.Command("npm", "install", "--prefer-offline", "--frozen-lockfile")
				cmd.Dir = projectName
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					fmt.Println("Error installing dependencies:", err)
					return
				}
			}()

		} else {
			// for mongodb
			dirs = []string{"src", "src/routes", "src/models", "src/controllers"}
			for _, dir := range dirs {
				if err := os.MkdirAll(fmt.Sprintf("%s/%s", projectName, dir), 0755); err != nil {
					fmt.Println("Error creating directory:", err)
					return err
				}
			}

			// Copy templates to create example files
			files := map[string]string{
				fmt.Sprintf("%s/tsconfig.json", projectName):                     "templates/node/ts/tsconfig.txt",
				fmt.Sprintf("%s/package.json", projectName):                      "templates/node/ts/package.txt",
				fmt.Sprintf("%s/src/routes/user-routes.ts", projectName):         "templates/node/ts/src/routes/user-routes.txt",
				fmt.Sprintf("%s/src/models/user.ts", projectName):                "templates/node/ts/src/models/user.txt",
				fmt.Sprintf("%s/src/controllers/userController.ts", projectName): "templates/node/ts/src/controllers/user-controller.txt",
				fmt.Sprintf("%s/src/index.ts", projectName):                      "templates/node/ts/src/index.txt",
			}

			// Use a buffered channel to limit concurrent template copying
			copyCh := make(chan struct{}, 5) // Adjust the buffer size as needed

			for dest, src := range files {
				wg.Add(1)
				copyCh <- struct{}{} // Add a token to limit concurrency
				go func(dest, src string) {
					defer func() {
						<-copyCh // Release token after copying is done
						wg.Done()
					}()
					if err := utils.CopyTemplate("backendforger", src, dest); err != nil {
						fmt.Println("Error copying template:", err)
					}
				}(dest, src)
			}

			// Close the channel after all copying tasks are added
			close(copyCh)
		}

	} else {

		// Create src and nested directories
		dirs := []string{"src", "src/routes", "src/models", "src/controllers"}
		for _, dir := range dirs {
			if err := os.MkdirAll(fmt.Sprintf("%s/%s", projectName, dir), 0755); err != nil {
				fmt.Println("Error creating directory:", err)
				return err
			}
		}

		// Copy templates to create example files
		files := map[string]string{
			fmt.Sprintf("%s/package.json", projectName):                      "templates/node/js/package.txt",
			fmt.Sprintf("%s/src/routes/user.js", projectName):                "templates/node/js/src/routes/user-routes.txt",
			fmt.Sprintf("%s/src/models/user.js", projectName):                "templates/node/js/src/models/user.txt",
			fmt.Sprintf("%s/src/controllers/userController.js", projectName): "templates/node/js/src/controllers/user-controller.txt",
			fmt.Sprintf("%s/src/index.js", projectName):                      "templates/node/js/src/index.txt",
		}

		// Use a buffered channel to limit concurrent template copying
		copyCh := make(chan struct{}, 5) // Adjust the buffer size as needed

		for dest, src := range files {
			wg.Add(1)
			copyCh <- struct{}{} // Add a token to limit concurrency
			go func(dest, src string) {
				defer func() {
					<-copyCh // Release token after copying is done
					wg.Done()
				}()
				if err := utils.CopyTemplate("backendforger", src, dest); err != nil {
					fmt.Println("Error copying template:", err)
				}
			}(dest, src)
		}

		// Close the channel after all copying tasks are added
		close(copyCh)

		// Install dependencies concurrently
		wg.Add(1)
		go func() {
			defer wg.Done()
			cmd := exec.Command("npm", "install", "--prefer-offline", "--frozen-lockfile")
			cmd.Dir = projectName
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Println("Error installing dependencies:", err)
				return
			}
		}()
	}

	// Wait for all tasks to finish
	wg.Wait()

	fmt.Printf("Node.js project '%s' generated in %v %s\n", color.BlueString(projectName), time.Since(startTime).Round(time.Millisecond), "🚀🚀\n")
	fmt.Printf("Navigate to the project directory using:\n\tcd %s\n\n", color.BlueString(projectName))
	fmt.Printf("Run your project using:\n\t%s\n", color.MagentaString(fmt.Sprintln("npm run dev")))
	fmt.Println(color.HiGreenString("Happy coding! 🎉🎉🎉"))

	return nil
}
