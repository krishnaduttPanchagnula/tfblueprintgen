package vpc

import "os"

func CreateVPCVariablesFile(filePath string) error {
	content := `variable "vpc_name" {
  description = "Name of the VPC"
  type        = string
}

variable "vpc_cidr" {
  description = "CIDR block for the VPC"
  type        = string
}

variable "avaliability_zones" {
  description = "List of availability zones"
  type        = list(string)
}

variable "private_subnets" {
  description = "List of private subnet CIDR blocks"
  type        = list(string)
}

variable "public_subnets" {
  description = "List of public subnet CIDR blocks"
  type        = list(string)
}

variable "enable_nat_gateway" {
  description = "Flag to enable NAT gateway"
  type        = bool
}

variable "enable_vpn_gateway" {
  description = "Flag to enable VPN gateway"
  type        = bool
}

variable "tags" {
  description = "Tags for the VPC resources"
  type        = map(string)
}`
	return os.WriteFile(filePath, []byte(content), os.ModePerm)
}

func CreateVPCModuleFile(filePath string) error {
	content := `
	module "vpc" {
	  source = "terraform-aws-modules/vpc/aws"
	
	  name = var.vpc_name
	  cidr = var.vpc_cidr
	
	  azs             = var.avaliability_zones
	  private_subnets = var.private_subnets
	  public_subnets  = var.public_subnets
	
	  enable_nat_gateway = var.enable_nat_gateway
	  enable_vpn_gateway = var.enable_vpn_gateway
	
	  tags = var.tags
	  }
	}`
	return os.WriteFile(filePath, []byte(content), os.ModePerm)
}
