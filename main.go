package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	xstrings "github.com/charmbracelet/x/exp/strings"
	"github.com/krishnaduttPanchagnula/Tfblueprintgen/lambda"
	"github.com/krishnaduttPanchagnula/Tfblueprintgen/readme"
	"github.com/krishnaduttPanchagnula/Tfblueprintgen/s3"
	progressbar "github.com/krishnaduttPanchagnula/Tfblueprintgen/utils"
	"github.com/krishnaduttPanchagnula/Tfblueprintgen/vpc"
)

var BASE_FOLDER_NAME = "terraform-aws"

var Files_list = []string{"main.tf", "variables.tf", "outputs.tf"}

var ProdFiles = []string{
	filepath.Join(BASE_FOLDER_NAME, "Environments", "Production"),
	filepath.Join(BASE_FOLDER_NAME, "Environments", "Production", "main.tf"),
	filepath.Join(BASE_FOLDER_NAME, "Environments", "Production", "variables.tf"),
	filepath.Join(BASE_FOLDER_NAME, "Environments", "Production", "outputs.tf"),
}

var StagingFiles = []string{
	filepath.Join(BASE_FOLDER_NAME, "Environments", "Staging"),
	filepath.Join(BASE_FOLDER_NAME, "Environments", "Staging", "main.tf"),
	filepath.Join(BASE_FOLDER_NAME, "Environments", "Staging", "variables.tf"),
	filepath.Join(BASE_FOLDER_NAME, "Environments", "Staging", "outputs.tf"),
}
var UATFiles = []string{
	filepath.Join(BASE_FOLDER_NAME, "Environments", "UAT"),
	filepath.Join(BASE_FOLDER_NAME, "Environments", "UAT", "main.tf"),
	filepath.Join(BASE_FOLDER_NAME, "Environments", "UAT", "variables.tf"),
	filepath.Join(BASE_FOLDER_NAME, "Environments", "UAT", "outputs.tf"),
}
var DevFiles = []string{
	filepath.Join(BASE_FOLDER_NAME, "Environments", "dev"),
	filepath.Join(BASE_FOLDER_NAME, "Environments", "dev", "main.tf"),
	filepath.Join(BASE_FOLDER_NAME, "Environments", "dev", "variables.tf"),
	filepath.Join(BASE_FOLDER_NAME, "Environments", "dev", "outputs.tf"),
}

var ReadmeFile = []string{
	filepath.Join("terraform-aws", "README.md"),
}

type Package struct {
	Environments []string
	Resources    []string
}

func main() {

	var Package Package

	form := huh.NewForm(
		huh.NewGroup(huh.NewNote().
			Title("Tfblueprintgen").
			Description("Welcome to _Tfblueprintgen™_.\n\n Press enter to start creating terraform folder structure.")),

		// Choose a burger.
		// We'll need to know what topping to add too.
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Choose List of Environments ").
				Description("Select the list of environments used in your organization .").
				Options(
					huh.NewOption("Production", "production"),
					huh.NewOption("Development", "development"),
					huh.NewOption("UAT", "uat"),
					huh.NewOption("TEST", "test"),
				).
				Validate(func(t []string) error {
					if len(t) <= 0 {
						return fmt.Errorf("at least one ENV is required")
					}
					return nil
				}).
				Value(&Package.Environments),

			huh.NewMultiSelect[string]().
				Title("Resources").
				Description("Choose all the resources that you want modules for.").
				Options(
					huh.NewOption("RDS", "rds"),
					huh.NewOption("VPC", "vpc").Selected(true),
					huh.NewOption("Lambda", "lambda"),
				).
				Validate(func(t []string) error {
					if len(t) <= 0 {
						return fmt.Errorf("at least one topping is required")
					}
					return nil
				}).
				Value(&Package.Resources).
				Filterable(true).
				Limit(4),
		),
	)

	err := form.Run()

	combined := []string{}
	for _, env := range Package.Environments {
		for _, resources := range Package.Resources {
			fp := filepath.Join(BASE_FOLDER_NAME + "/" + "environments" + "/" + env + "/" + resources + ".tf")
			combined = append(combined, fp)
		}

	}
	fmt.Println(combined)

	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	// Define the directory structure

	lambdafilenames := lambda.CreateLambdaFilePathNames(BASE_FOLDER_NAME)
	s3filenames := s3.CreateS3FilePathNames(BASE_FOLDER_NAME)
	vpcfilenames := vpc.CreateVpcFilePathNames(BASE_FOLDER_NAME)

	fileNames := concatMultipleSlices([][]string{lambdafilenames, s3filenames, vpcfilenames, combined})

	directoryStructure := append([]string{BASE_FOLDER_NAME}, fileNames...)

	// Create directories and files
	for _, path := range directoryStructure {
		switch {
		case filepath.Ext(path) == ".tf" && path == filepath.Join(BASE_FOLDER_NAME, "modules", "lambda", "variables.tf"):

			err := lambda.CreateLambdaVariablesFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".tf" && path == filepath.Join(BASE_FOLDER_NAME, "modules", "lambda", "main.tf"):

			err := lambda.CreateLambdamoduleFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".tf" && path == filepath.Join(BASE_FOLDER_NAME, "modules", "vpc", "variables.tf"):

			err := vpc.CreateVPCVariablesFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".tf" && path == filepath.Join(BASE_FOLDER_NAME, "modules", "vpc", "main.tf"):

			err := vpc.CreateVPCModuleFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".tf" && path == filepath.Join(BASE_FOLDER_NAME, "modules", "s3", "variables.tf"):

			err := s3.CreateS3VariablesFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".tf" && path == filepath.Join(BASE_FOLDER_NAME, "modules", "s3", "main.tf"):
			// Create variables.tf file with dynamic content for lambda module
			err := s3.CreateS3MainFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".md" && path == filepath.Join(BASE_FOLDER_NAME, "README.md"):
			// Create variables.tf file with dynamic content for lambda module
			err := readme.CreateReadmeFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == "":
			// Create directory
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				fmt.Printf("Error creating directory %s: %v\n", path, err)
				os.Exit(1)
			}
		default:
			// Create file
			err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
			if err != nil {
				fmt.Printf("Error creating file %s: %v\n", path, err)
				os.Exit(1)
			}
			file, err := os.Create(path)
			if err != nil {
				log.Fatalf("Error creating file: %v", err)
			}
			defer file.Close()
		}

		// // Print status message for each file/directory created
		// fmt.Printf("[%d/%d] Created: %s\n", i+1, len(directoryStructure), path)
	}

	// fmt.Printf("File structure for %s created successfully.", BASE_FOLDER_NAME)

	m := progressbar.Model{
		Progress: progress.New(progress.WithDefaultGradient()),
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}

	// Print order summary.
	{
		var sb strings.Builder
		keyword := func(s string) string {
			return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
		}
		fmt.Fprintf(&sb,
			"%s\n\n The Resources of %s, are created in  %s Folders.",
			lipgloss.NewStyle().Bold(true).Render("Tfblueprintgen Report"),
			keyword(xstrings.EnglishJoin(Package.Resources, true)),
			keyword(xstrings.EnglishJoin(Package.Environments, true)),
		)

		fmt.Println(
			lipgloss.NewStyle().
				Width(40).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render(sb.String()),
		)
	}
}

func concatMultipleSlices[T any](slices [][]T) []T {
	var totalLen int

	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]T, totalLen)

	var i int

	for _, s := range slices {
		i += copy(result[i:], s)
	}

	return result
}
