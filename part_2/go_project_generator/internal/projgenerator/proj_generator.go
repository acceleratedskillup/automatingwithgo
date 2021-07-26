package projgenerator

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	webappTemplateFileLocation         = "./templates/webapp.tmpl"
	progressAfterPrepOfOutputDirectory = 0.3
	progressAfterCreationOfMainFile    = 0.3
)

type (
	ProjectCreationInProgressMsg float64
	ProjectCreatedMsg            struct{}
)
type AppConfig struct {
	ProjectName       string
	UseRouter         bool
	RouterType        string
	RouterImportPath  string
	RouterConstructor string
	ConfigSource      string // "flags" or "env"
}

func GenerateProject(appConfig AppConfig, progressChannel chan tea.Msg) {
	defer close(progressChannel)
	outputPath := "../" + appConfig.ProjectName
	// Assume config is defined as shown previously
	prepareOutputDirectory(outputPath)
	progressChannel <- ProjectCreationInProgressMsg(progressAfterPrepOfOutputDirectory)
	generateMainFile(appConfig, outputPath)
	log.Println("after generateMainFile()")
	progressChannel <- ProjectCreationInProgressMsg(progressAfterCreationOfMainFile)
	log.Println("after progressChannel <- ProjectCreationInProgressMsg(0.3)")
	initGoModuleAndDependencies(appConfig.ProjectName, outputPath, appConfig)
	log.Println("after initGoModuleAndDependencies")
	progressChannel <- ProjectCreatedMsg{}
	log.Println("after progressChannel <- ProjectCreatedMsg{}")

}

func generateMainFile(config AppConfig, outputPath string) {

	// Specify the path for the main.go file within the output directory
	mainFilePath := filepath.Join(outputPath, "main.go")

	// Open the template file
	tmpl, err := template.ParseFiles(webappTemplateFileLocation)
	if err != nil {
		log.Fatalf("Error opening template file: %v", err)
	}

	// Create the main.go file in the specified output path
	f, err := os.Create(mainFilePath)
	if err != nil {
		log.Fatalf("Failed to create main.go at %s: %v", mainFilePath, err)
	}
	defer f.Close()

	// Execute the template with the AppConfig struct
	err = tmpl.Execute(f, config)
	if err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}
}

func initGoModuleAndDependencies(projectName, outputPath string, config AppConfig) {

	// Initialize go.mod file
	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = outputPath
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to initialize go.mod: %v", err)
	}

	if config.UseRouter {
		cmd = exec.Command("go", "get", config.RouterImportPath)
		cmd.Dir = outputPath
		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed to download router package: %v", err)
		}
	}

	// Add any other dependencies here
}

func prepareOutputDirectory(outputPath string) {
	// Check if the output directory exists
	if _, err := os.Stat(outputPath); !os.IsNotExist(err) {
		// Directory exists, attempt to remove it
		err := os.RemoveAll(outputPath)
		if err != nil {
			log.Fatalf("Failed to remove existing output directory: %v", err)
		}
	}

	// Create the output directory
	err := os.MkdirAll(outputPath, 0755) // Use appropriate permissions
	if err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}
}
