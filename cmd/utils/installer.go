package utils

import (
	"fmt"
	"os/exec"
)

func GoMod(projectName string) error {
	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = projectName
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error initializing go module:", err)
		fmt.Println("Output:", string(output)) // Print command output for debugging
		return err
	}
	return nil
}

func GoTidy(projectName string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = projectName
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running go mod tidy:", err)
		fmt.Println("Output:", string(output)) // Print command output for debugging
		return err
	}
	return nil
}

func FlaskInit(projectName string) error {
	cmd := exec.Command("flask", "db", "init")
	cmd.Dir = projectName
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running flask db init:", err)
		fmt.Println("Output:", string(output)) // Print command output for debugging
		return err
	}
	return nil
}

func FlaskMigrate(projectName string) error {
	cmd := exec.Command("flask", "db", "migrate", "-m", "Your migration message")
	cmd.Dir = projectName
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("flask db migrate -m \"Your migration message\"", err)
		fmt.Println("Output:", string(output)) // Print command output for debugging
		return err
	}
	return nil
}

func FlaskUpgrade(projectName string) error {
	cmd := exec.Command("flask", "db", "upgrade")
	cmd.Dir = projectName
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("flask db upgrade", err)
		fmt.Println("Output:", string(output)) // Print command output for debugging
		return err
	}
	return nil
}

func Python_Install(projectName string) error {
	cmd := exec.Command("pip", "install", "requirements.txt")
	cmd.Dir = projectName
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("pip install requirements.txt", err)
		fmt.Println("Output:", string(output)) // Print command output for debugging
		return err
	}
	return nil
}

func VenvSetup(projectName string) error {
	cmd := exec.Command("python", "-m", "venv", "venv")
	cmd.Dir = projectName
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Creating virtual environment", err)
		fmt.Println("Output:", string(output)) // Print command output for debugging
		return err
	}
	return nil
}

func VenvActivate(projectName string) error {
	cmd := exec.Command(".venv/Scripts/Activate")
	cmd.Dir = projectName
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Activating virtual environment", err)
		fmt.Println("Output:", string(output)) // Print command output for debugging
		return err
	}
	return nil
}
