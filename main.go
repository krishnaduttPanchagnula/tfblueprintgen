package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/krishnaduttPanchagnula/Tfblueprintgen/lambda"
	"github.com/krishnaduttPanchagnula/Tfblueprintgen/readme"
	"github.com/krishnaduttPanchagnula/Tfblueprintgen/s3"
	"github.com/krishnaduttPanchagnula/Tfblueprintgen/vpc"
)

func main() {

	// Define the directory structure
	directoryStructure := []string{
		"terraform-aws",
		filepath.Join("terraform-aws", "modules", "lambda"),
		filepath.Join("terraform-aws", "modules", "lambda", "main.tf"),
		filepath.Join("terraform-aws", "modules", "lambda", "variables.tf"),
		filepath.Join("terraform-aws", "modules", "lambda", "outputs.tf"),
		filepath.Join("terraform-aws", "modules", "lambda"),
		filepath.Join("terraform-aws", "modules", "s3"),
		filepath.Join("terraform-aws", "modules", "s3", "main.tf"),
		filepath.Join("terraform-aws", "modules", "s3", "variables.tf"),
		filepath.Join("terraform-aws", "modules", "s3", "outputs.tf"),
		filepath.Join("terraform-aws", "modules", "vpc"),
		filepath.Join("terraform-aws", "modules", "vpc", "main.tf"),
		filepath.Join("terraform-aws", "modules", "vpc", "variables.tf"),
		filepath.Join("terraform-aws", "modules", "vpc", "outputs.tf"),
		filepath.Join("terraform-aws", "Environments", "Production"),
		filepath.Join("terraform-aws", "Environments", "Production", "main.tf"),
		filepath.Join("terraform-aws", "Environments", "Production", "variables.tf"),
		filepath.Join("terraform-aws", "Environments", "Production", "outputs.tf"),
		filepath.Join("terraform-aws", "Environments", "Staging"),
		filepath.Join("terraform-aws", "Environments", "Staging", "main.tf"),
		filepath.Join("terraform-aws", "Environments", "Staging", "variables.tf"),
		filepath.Join("terraform-aws", "Environments", "Staging", "outputs.tf"),
		filepath.Join("terraform-aws", "Environments", "UAT"),
		filepath.Join("terraform-aws", "Environments", "UAT", "main.tf"),
		filepath.Join("terraform-aws", "Environments", "UAT", "variables.tf"),
		filepath.Join("terraform-aws", "Environments", "UAT", "outputs.tf"),
		filepath.Join("terraform-aws", "Environments", "dev"),
		filepath.Join("terraform-aws", "Environments", "dev", "main.tf"),
		filepath.Join("terraform-aws", "Environments", "dev", "variables.tf"),
		filepath.Join("terraform-aws", "Environments", "dev", "outputs.tf"),
		filepath.Join("terraform-aws", "README.md"),
	}

	// Create directories and files
	for i, path := range directoryStructure {
		switch {
		case filepath.Ext(path) == ".tf" && path == filepath.Join("terraform-aws", "modules", "lambda", "variables.tf"):
			// Create variables.tf file with dynamic content for lambda module
			err := lambda.CreateLambdaVariablesFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".tf" && path == filepath.Join("terraform-aws", "modules", "lambda", "main.tf"):
			// Create variables.tf file with dynamic content for lambda module
			err := lambda.CreateLambdamoduleFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".tf" && path == filepath.Join("terraform-aws", "modules", "vpc", "variables.tf"):
			// Create variables.tf file with dynamic content for vpc module
			err := vpc.CreateVPCVariablesFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".tf" && path == filepath.Join("terraform-aws", "modules", "vpc", "main.tf"):
			// Create variables.tf file with dynamic content for vpc module
			err := vpc.CreateVPCModuleFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".tf" && path == filepath.Join("terraform-aws", "modules", "s3", "variables.tf"):
			// Create variables.tf file with dynamic content for lambda module
			err := s3.CreateS3VariablesFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".tf" && path == filepath.Join("terraform-aws", "modules", "s3", "main.tf"):
			// Create variables.tf file with dynamic content for lambda module
			err := s3.CreateS3MainFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".md" && path == filepath.Join("terraform-aws", "README.md"):
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
			_, err := os.Create(path)
			if err != nil {
				fmt.Printf("Error creating file %s: %v\n", path, err)
				os.Exit(1)
			}
		}

		// Print status message for each file/directory created
		fmt.Printf("[%d/%d] Created: %s\n", i+1, len(directoryStructure), path)
	}

	fmt.Println("File structure for terraform-aws created successfully.")
}
