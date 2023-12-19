package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	// Define the directory structure
	directoryStructure := []string{
		"terraform-aws",
		filepath.Join("terraform-aws", "modules", "lambda"),
		filepath.Join("terraform-aws", "modules", "lambda", "main.tf"),
		filepath.Join("terraform-aws", "modules", "lambda", "variables.tf"),
		filepath.Join("terraform-aws", "modules", "lambda", "outputs.tf"),
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
			err := createLambdaVariablesFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".tf" && path == filepath.Join("terraform-aws", "modules", "vpc", "variables.tf"):
			// Create variables.tf file with dynamic content for vpc module
			err := vpc.createVPCVariablesFile(path)
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

func createLambdaVariablesFile(filePath string) error {
	content := `variable "function_name" {
  description = "Name of the Lambda function"
  type        = string
}

variable "description" {
  description = "Description of the Lambda function"
  type        = string
}

variable "runtime" {
  description = "Runtime for the Lambda function"
  type        = string
}

variable "handler" {
  description = "Handler for the Lambda function"
  type        = string
}

variable "iam_role_name" {
  description = "Name of the IAM role for Lambda execution"
  type        = string
}

variable "environment" {
  description = "Environment tag for Lambda resources"
  type        = string
}`
	return os.WriteFile(filePath, []byte(content), os.ModePerm)
}
