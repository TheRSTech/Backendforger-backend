package generator

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/TheRSTech/Backendforger-backend/cmd/utils"
	"github.com/fatih/color"
)

// GenerateGoProject creates a Go project structure.
func GenerateGoProject(projectName, framework, database, orm string) error {
	startTime := time.Now()

	// Create project root directory
	err := os.MkdirAll(projectName, 0755)
	if err != nil {
		fmt.Println("Error creating project directory:", err)
		return err
	}

	if err := utils.GoMod(projectName); err != nil {
		return err
	}

	// Copy template files based on the framework
	switch framework {
	case "gin":
		generateGinProject(projectName, database, orm)
	case "fiber":
		generateFiberProject(projectName, database, orm)
	case "echo":
		generateEchoProject(projectName, database, orm)
	case "http":
		generateHttpProject(projectName, database, orm)
	case "mux":
		generateMuxProject(projectName, database, orm)
	default:
		fmt.Println("Unsupported framework:", framework)
	}

	if err := utils.GoTidy(projectName); err != nil {
		return err
	}

	fmt.Printf("Go project '%s' generated in %v %s\n", color.BlueString(projectName), time.Since(startTime).Round(time.Millisecond), "🚀🚀\n")
	fmt.Printf("Navigate to the project directory using:\n\tcd %s\n\n", color.BlueString(projectName))
	fmt.Printf("Run your project using:\n\t%s\n", color.MagentaString(fmt.Sprintf("go run %s.go\n", "main")))
	fmt.Println(color.HiGreenString("Happy coding! 🎉🎉🎉"))

	return nil
}
func generateGinProject(projectName, database, orm string) {
	var wg sync.WaitGroup
	os.Mkdir(fmt.Sprintf("%s/api", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/middleware", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/models", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/config", projectName), 0755)

	// Copy template files concurrently
	wg.Add(3)
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/go/gin/main.txt", fmt.Sprintf("%s/main.go", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/go/gin/api/hello.txt", fmt.Sprintf("%s/api/api.go", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		if orm == "gorm" {
			if database == "mysql" {
				utils.CopyTemplate("backendforger", "templates/go/databases/gorm/init_mysql.txt", fmt.Sprintf("%s/config/init_db.go", projectName), map[string]string{"yourapp": projectName})
			} else if database == "postgres" {
				utils.CopyTemplate("backendforger", "templates/go/databases/gorm/init_pg.txt", fmt.Sprintf("%s/config/init_db.go", projectName), map[string]string{"yourapp": projectName})
			} else {
				utils.CopyTemplate("backendforger", "templates/go/databases/gorm/init_sqlite.txt", fmt.Sprintf("%s/config/init_db.go", projectName), map[string]string{"yourapp": projectName})
			}
		}
	}()

	wg.Wait()
}
func generateFiberProject(projectName, database, orm string) {
	var wg sync.WaitGroup
	os.Mkdir(fmt.Sprintf("%s/controllers", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/models", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/config", projectName), 0755)

	// Copy template files concurrently
	wg.Add(4)
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/go/fiber/main.txt", fmt.Sprintf("%s/main.go", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/go/fiber/controllers/user_controller.txt", fmt.Sprintf("%s/controllers/user_controller.go", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/go/fiber/models/user.txt", fmt.Sprintf("%s/models/user.go", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		if orm == "gorm" {
			if database == "mysql" {
				utils.CopyTemplate("backendforger", "templates/go/databases/gorm/init_mysql.txt", fmt.Sprintf("%s/config/init_db.go", projectName), map[string]string{"yourapp": projectName})
			} else if database == "postgres" {
				utils.CopyTemplate("backendforger", "templates/go/databases/gorm/init_pg.txt", fmt.Sprintf("%s/config/init_db.go", projectName), map[string]string{"yourapp": projectName})
			} else {
				utils.CopyTemplate("backendforger", "templates/go/databases/gorm/init_sqlite.txt", fmt.Sprintf("%s/config/init_db.go", projectName), map[string]string{"yourapp": projectName})
			}
		}
	}()

	wg.Wait()
}

func generateEchoProject(projectName, database, orm string) {
	var wg sync.WaitGroup
	os.Mkdir(fmt.Sprintf("%s/controllers", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/models", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/config", projectName), 0755)

	// Copy template files concurrently
	wg.Add(4)
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/go/echo/main.txt", fmt.Sprintf("%s/main.go", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/go/echo/controllers/user_controller.txt", fmt.Sprintf("%s/controllers/user_controller.go", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/go/echo/models/user.txt", fmt.Sprintf("%s/models/user.go", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		if orm == "gorm" {
			if database == "mysql" {
				utils.CopyTemplate("backendforger", "templates/go/databases/gorm/init_mysql.txt", fmt.Sprintf("%s/config/init_db.go", projectName), map[string]string{"yourapp": projectName})
			} else if database == "postgres" {
				utils.CopyTemplate("backendforger", "templates/go/databases/gorm/init_pg.txt", fmt.Sprintf("%s/config/init_db.go", projectName), map[string]string{"yourapp": projectName})
			} else {
				utils.CopyTemplate("backendforger", "templates/go/databases/gorm/init_sqlite.txt", fmt.Sprintf("%s/config/init_db.go", projectName), map[string]string{"yourapp": projectName})
			}
		}
	}()

	wg.Wait()
}

func generateHttpProject(projectName, database, orm string) {
	var wg sync.WaitGroup
	os.Mkdir(fmt.Sprintf("%s/controllers", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/models", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/config", projectName), 0755)

	// Copy template files concurrently
	wg.Add(4)
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/go/http/main.txt", fmt.Sprintf("%s/main.go", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/go/http/controllers/user_controller.txt", fmt.Sprintf("%s/controllers/user_controller.go", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/go/http/models/user.txt", fmt.Sprintf("%s/models/user.go", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		if orm == "gorm" {
			if database == "mysql" {
				utils.CopyTemplate("backendforger", "templates/go/databases/gorm/init_mysql.txt", fmt.Sprintf("%s/config/init_db.go", projectName), map[string]string{"yourapp": projectName})
			} else if database == "postgres" {
				utils.CopyTemplate("backendforger", "templates/go/databases/gorm/init_pg.txt", fmt.Sprintf("%s/config/init_db.go", projectName), map[string]string{"yourapp": projectName})
			} else {
				utils.CopyTemplate("backendforger", "templates/go/databases/gorm/init_sqlite.txt", fmt.Sprintf("%s/config/init_db.go", projectName), map[string]string{"yourapp": projectName})
			}
		}
	}()

	wg.Wait()
}

func generateMuxProject(projectName, database, orm string) {
	var wg sync.WaitGroup
	os.Mkdir(fmt.Sprintf("%s/controllers", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/models", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/config", projectName), 0755)

	// Copy template files concurrently
	wg.Add(4)
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/go/mux/main.txt", fmt.Sprintf("%s/main.go", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/go/mux/controllers/user_controller.txt", fmt.Sprintf("%s/controllers/user_controller.go", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/go/mux/models/user.txt", fmt.Sprintf("%s/models/user.go", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		if orm == "gorm" {
			if database == "mysql" {
				utils.CopyTemplate("backendforger", "templates/go/databases/gorm/init_mysql.txt", fmt.Sprintf("%s/config/init_db.go", projectName), map[string]string{"yourapp": projectName})
			} else if database == "postgres" {
				utils.CopyTemplate("backendforger", "templates/go/databases/gorm/init_pg.txt", fmt.Sprintf("%s/config/init_db.go", projectName), map[string]string{"yourapp": projectName})
			} else {
				utils.CopyTemplate("backendforger", "templates/go/databases/gorm/init_sqlite.txt", fmt.Sprintf("%s/config/init_db.go", projectName), map[string]string{"yourapp": projectName})
			}
		}
	}()

	wg.Wait()
}
