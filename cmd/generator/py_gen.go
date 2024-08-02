package generator

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/TheRSTech/Backendforger-backend/cmd/utils"
	"github.com/fatih/color"
)

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
		fmt.Println("Before you run your project you will have to migrate it using the following steps")
		fmt.Println("Enter the following lines of codes")
		fmt.Println(color.MagentaString(fmt.Sprintf(".venv/Scripts/Activate")))
		fmt.Printf("Run your project using:\n\t%s\n", color.MagentaString(fmt.Sprintf("uvicorn app.main:app --reload")))
		fmt.Println(color.HiGreenString("Happy coding! 🎉🎉🎉"))

	case "flask":

		generateFlaskProject(projectName, database, orm)

		utils.FlaskInit(projectName)
		utils.FlaskMigrate(projectName)
		utils.FlaskUpgrade(projectName)

		fmt.Printf("Python project '%s' generated in %v %s\n", color.BlueString(projectName), time.Since(startTime).Round(time.Millisecond), "🚀🚀\n")
		fmt.Printf("Navigate to the project directory using:\n\tcd %s\n\n", color.BlueString(projectName))
		fmt.Println("Before you run your project you will have to migrate it using the following steps")
		fmt.Println("Enter the following lines of codes")
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

	os.Mkdir(fmt.Sprintf("%s/app", projectName), 0755)

	// Copy template files concurrently
	wg.Add(8)
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/fast_api/requirements.txt.txt", fmt.Sprintf("%s/requirements.txt", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/fast_api/.gitignore.txt", fmt.Sprintf("%s/.gitignore", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/fast_api/app/__init__.txt", fmt.Sprintf("%s/app/__init__.py", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/fast_api/app/crud.txt", fmt.Sprintf("%s/app/crud.py", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/fast_api/app/database.txt", fmt.Sprintf("%s/app/database.py", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/fast_api/app/main.txt", fmt.Sprintf("%s/app/main.py", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/fast_api/app/models.txt", fmt.Sprintf("%s/app/models.py", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/fast_api/app/schemas.txt", fmt.Sprintf("%s/app/schemas.py", projectName), map[string]string{"yourapp": projectName})
	}()

	wg.Wait()
}

func generateFlaskProject(projectName, database, orm string) {

	var wg sync.WaitGroup

	os.Mkdir(fmt.Sprintf("%s/app", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/static", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/static/images", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/static/js", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/static/css", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/templates", projectName), 0755)
	os.Mkdir(fmt.Sprintf("%s/migrations", projectName), 0755)

	// Copy template files concurrently
	wg.Add(12)
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/flask/run.txt", fmt.Sprintf("%s/run.py", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/flask/.env.txt", fmt.Sprintf("%s/.env", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/flask/requirements.txt.txt", fmt.Sprintf("%s/requirements.txt", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/flask/.gitignore.txt", fmt.Sprintf("%s/.gitignore", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/flask/app/extensions.txt", fmt.Sprintf("%s/app/extensions.py", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/flask/app/__init__.txt", fmt.Sprintf("%s/app/__init__.py", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/flask/app/models.txt", fmt.Sprintf("%s/app/models.py", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/flask/app/routes.txt", fmt.Sprintf("%s/app/routes.py", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/flask/app/utils.txt", fmt.Sprintf("%s/app/utils.py", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/flask/templates/index.txt", fmt.Sprintf("%s/templates/index.html", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		utils.CopyTemplate("backendforger", "templates/python/flask/templates/layout.txt", fmt.Sprintf("%s/templates/layout.html", projectName), map[string]string{"yourapp": projectName})
	}()
	go func() {
		defer wg.Done()
		if database == "postgres" {
			utils.CopyTemplate("backendforger", "templates/python/flask/app/database/config_postgres.txt", fmt.Sprintf("%s/app/config.py", projectName), map[string]string{"yourapp": projectName})
		}
		if database == "sqlite" {
			utils.CopyTemplate("backendforger", "templates/python/flask/app/database/config_sqlite.txt", fmt.Sprintf("%s/app/config.py", projectName), map[string]string{"yourapp": projectName})
		}
		if database == "mysql" {
			utils.CopyTemplate("backendforger", "templates/python/flask/app/database/config_mysql.txt", fmt.Sprintf("%s/app/config.py", projectName), map[string]string{"yourapp": projectName})
		} else {
			utils.CopyTemplate("backendforger", "templates/python/flask/app/database/config_sqlite.txt", fmt.Sprintf("%s/app/config.py", projectName), map[string]string{"yourapp": projectName})
		}
	}()
	// Wait for all goroutines to complete
	wg.Wait()
}
