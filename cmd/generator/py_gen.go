package generator

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/TheRSTech/Backendforger-backend/cmd/utils"
	"github.com/fatih/color"
)

// GeneratePythonProject generates a Python project based on the specified framework.
func GeneratePythonProject(projectName, framework, database, orm string) error {
	startTime := time.Now()

	// Create project root directory
	err := os.MkdirAll(projectName, 0755)
	if err != nil {
		fmt.Println("Error creating project directory:", err)
		return err
	}

	// Copy template files based on the framework
	switch framework {
	case "fastapi":
		generateFastProject(projectName, database, orm)

		utils.VenvSetup(projectName)
		utils.VenvActivate(projectName)
		utils.Python_Install(projectName)

		fmt.Printf("Python project '%s' generated in %v %s\n", color.BlueString(projectName), time.Since(startTime).Round(time.Millisecond), "🚀🚀\n")
		fmt.Printf("Navigate to the project directory using:\n\tcd %s\n\n", color.BlueString(projectName))
		fmt.Println("Before you run your project, you will have to migrate it using the following steps:")
		fmt.Println("Enter the following lines of code:")
		fmt.Println(color.MagentaString(".venv/Scripts/Activate"))
		fmt.Printf("Run your project using:\n\t%s\n", color.MagentaString("uvicorn app.main:app --reload"))
		fmt.Println(color.HiGreenString("Happy coding! 🎉🎉🎉"))

	case "flask":
		generateFlaskProject(projectName, database, orm)

		utils.FlaskInit(projectName)
		utils.FlaskMigrate(projectName)
		utils.FlaskUpgrade(projectName)

		fmt.Printf("Python project '%s' generated in %v %s\n", color.BlueString(projectName), time.Since(startTime).Round(time.Millisecond), "🚀🚀\n")
		fmt.Printf("Navigate to the project directory using:\n\tcd %s\n\n", color.BlueString(projectName))
		fmt.Println("Before you run your project, you will have to migrate it using the following steps:")
		fmt.Println("Enter the following lines of code:")
		fmt.Println(color.MagentaString("flask db init\n\n"))
		fmt.Println(color.MagentaString("flask db migrate -m \"Your migration message\"\n\n"))
		fmt.Println(color.MagentaString("flask db upgrade\n\n"))
		fmt.Printf("Run your project using:\n\t%s\n", color.MagentaString(fmt.Sprintf("python %s.py\n", projectName)))
		fmt.Println(color.HiGreenString("Happy coding! 🎉🎉🎉"))

	default:
		fmt.Println("Unsupported framework:", framework)
	}

	return nil
}

func generateFastProject(projectName, database, orm string) {
	var wg sync.WaitGroup

	// Create necessary directories
	os.Mkdir(fmt.Sprintf("%s/app", projectName), 0755)

	// Define the source and destination map for FastAPI templates
	templates := map[string]string{
		"templates/python/fast_api/requirements.txt.txt": fmt.Sprintf("%s/requirements.txt", projectName),
		"templates/python/fast_api/.gitignore.txt":       fmt.Sprintf("%s/.gitignore", projectName),
		"templates/python/fast_api/app/__init__.txt":     fmt.Sprintf("%s/app/__init__.py", projectName),
		"templates/python/fast_api/app/crud.txt":         fmt.Sprintf("%s/app/crud.py", projectName),
		"templates/python/fast_api/app/database.txt":     fmt.Sprintf("%s/app/database.py", projectName),
		"templates/python/fast_api/app/main.txt":         fmt.Sprintf("%s/app/main.py", projectName),
		"templates/python/fast_api/app/models.txt":       fmt.Sprintf("%s/app/models.py", projectName),
		"templates/python/fast_api/app/schemas.txt":      fmt.Sprintf("%s/app/schemas.py", projectName),
	}

	// Copy template files concurrently
	wg.Add(len(templates))
	for src, dest := range templates {
		go func(src, dest string) {
			defer wg.Done()
			utils.CopyTemplate("backendforger", src, dest, map[string]string{"yourapp": projectName})
		}(src, dest)
	}

	wg.Wait()
}

func generateFlaskProject(projectName, database, orm string) {
	var wg sync.WaitGroup

	// Create necessary directories
	os.Mkdir(fmt.Sprintf("%s/app", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/static", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/static/images", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/static/js", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/static/css", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/templates", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/migrations", projectName), 0755)

	// Define the source and destination map for Flask templates
	templates := map[string]string{
		"templates/python/flask/run.txt":                 fmt.Sprintf("%s/run.py", projectName),
		"templates/python/flask/.env.txt":                fmt.Sprintf("%s/.env", projectName),
		"templates/python/flask/requirements.txt.txt":    fmt.Sprintf("%s/requirements.txt", projectName),
		"templates/python/flask/.gitignore.txt":          fmt.Sprintf("%s/.gitignore", projectName),
		"templates/python/flask/app/extensions.txt":      fmt.Sprintf("%s/app/extensions.py", projectName),
		"templates/python/flask/app/__init__.txt":        fmt.Sprintf("%s/app/__init__.py", projectName),
		"templates/python/flask/app/models.txt":          fmt.Sprintf("%s/app/models.py", projectName),
		"templates/python/flask/app/routes.txt":          fmt.Sprintf("%s/app/routes.py", projectName),
		"templates/python/flask/app/utils.txt":           fmt.Sprintf("%s/app/utils.py", projectName),
		"templates/python/flask/templates/index.txt":     fmt.Sprintf("%s/templates/index.html", projectName),
		"templates/python/flask/templates/layout.txt":    fmt.Sprintf("%s/templates/layout.html", projectName),
		"templates/python/flask/app/database/config.txt": fmt.Sprintf("%s/app/config.py", projectName),
	}

	// Determine the correct database config template
	switch database {
	case "postgres":
		templates["templates/python/flask/app/database/config_postgres.txt"] = fmt.Sprintf("%s/app/config.py", projectName)
	case "sqlite":
		templates["templates/python/flask/app/database/config_sqlite.txt"] = fmt.Sprintf("%s/app/config.py", projectName)
	case "mysql":
		templates["templates/python/flask/app/database/config_mysql.txt"] = fmt.Sprintf("%s/app/config.py", projectName)
	default:
		templates["templates/python/flask/app/database/config_sqlite.txt"] = fmt.Sprintf("%s/app/config.py", projectName)
	}

	// Copy template files concurrently
	wg.Add(len(templates))
	for src, dest := range templates {
		go func(src, dest string) {
			defer wg.Done()
			utils.CopyTemplate("backendforger", src, dest, map[string]string{"yourapp": projectName})
		}(src, dest)
	}

	// Wait for all goroutines to complete
	wg.Wait()
}
