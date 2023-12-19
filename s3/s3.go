package s3

import (
	"os"
	"path/filepath"
)

func CreateS3FilePathNames(basepath string) []string {
	return []string{filepath.Join(basepath, "modules", "s3"),
		filepath.Join(basepath, "modules", "s3", "main.tf"),
		filepath.Join(basepath, "modules", "s3", "variables.tf"),
		filepath.Join(basepath, "modules", "s3", "outputs.tf"),
	}
}

func CreateS3MainFile(filePath string) error {
	content := `resource "aws_s3_bucket" "products" {
		bucket = var.bucketname
		tags = var.tags
	  }
	  
	  resource "aws_s3_bucket_public_access_block" "products" {
		bucket = aws_s3_bucket.products.id
	  
		block_public_acls       = var.block_public_acls
		block_public_policy     = var.block_public_policy
		ignore_public_acls      = var.ignore_public_acls
		restrict_public_buckets = var.restrict_public_buckets
	  }
	  
	  resource "aws_s3_bucket_versioning" "versioning_example" {
		bucket = aws_s3_bucket.products.id
		versioning_configuration {
		  status = var.versioning_status
		}
	  }`
	return os.WriteFile(filePath, []byte(content), os.ModePerm)
}

func CreateS3VariablesFile(filePath string) error {
	content := `variable "bucketname" {
		description = "Name of the bucket to be created"
		type = string
	  }
	  
	  variable "tags"{
		  description = "THe list of tags to be associated to the bucket"
		  type        = map(string)
	  }
	  variable "versioning_status" {
		  description = "Describes what should be the versioning status"
		  type =string
	  }
	  
	  variable "block_public_acls" {
		  description = "Whether Amazon S3 should block public ACLs for this bucket"
		  type = bool
		
	  }
	  variable "block_public_policy" {
		 description = "Whether Amazon S3 should block public bucket policies for this bucket"
		  type = bool
	  }
	  variable "ignore_public_acls" {
		 description = "Whether Amazon S3 should ignore public ACLs for this bucket."
		  type = bool
	  }
	  variable "restrict_public_buckets" {
		 description = "Whether Amazon S3 should restrict public bucket policies for this bucket"
		  type = bool
	  }`
	return os.WriteFile(filePath, []byte(content), os.ModePerm)
}
