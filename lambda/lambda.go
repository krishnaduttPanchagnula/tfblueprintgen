package lambda

import "os"

func CreateLambdaVariablesFile(filePath string) error {
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

func CreateLambdamoduleFile(filePath string) error {
	content := `module "lambda-function" {
    source  = "mineiros-io/lambda-function/aws"
    version = "~> 0.5.0"
  
    function_name = var.function_name
    description   = var.description
    filename      = data.archive_file.lambda.output_path
    runtime       = var.runtime
    handler       = var.handler
    timeout       = 30
    memory_size   = 128
  
    role_arn = module.iam_role.role.arn
  
    module_tags = {
      Environment = var.environment
    }
  }
  
  # ----------------------------------------------------------------------------------------------------------------------
  # CREATE AN IAM LAMBDA EXECUTION ROLE WHICH WILL BE ATTACHED TO THE FUNCTION
  # ----------------------------------------------------------------------------------------------------------------------
  
  module "iam_role" {
    source  = "mineiros-io/iam-role/aws"
    version = "~> 0.6.0"
  
    name = var.iam_role_name
  
    assume_role_principals = [
      {
        type        = "Service"
        identifiers = ["lambda.amazonaws.com"]
      }
    ]
  
    tags = {
      Environment = var.environment
    }
  }`
	return os.WriteFile(filePath, []byte(content), os.ModePerm)
}
