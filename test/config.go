package main

import (
	"fmt"
	"net/http"
	"os"

	// For reading configuration entries
    "github.com/ilyakaznacheev/cleanenv"
)

// Config is an application configuration structure
type Config struct {
	TerraformAzureModules struct {
		TerraformAzureVNETDir      string `yaml:"vnet_module_dir" env:"TERRAFORM_AZURE_VNET_MODULE_DIR" env-description:"Terraform Azure Module VNET Directory location in the file system"`
	} `yaml:"terraform_azure_modules"`
	TerraformAwsModules struct {
	    TerraformAwsVPCDir         string `yaml:"vpc_module_dir" env:"TERRAFORM_AWS_VPC_MODULE_DIR" env-description:"Terraform AWS Module VPC Directory location in the file system"`
		// Host string `yaml:"host" env:"SRV_HOST,HOST" env-description:"Server host" env-default:"localhost"`
		// Port string `yaml:"port" env:"SRV_PORT,PORT" env-description:"Server port" env-default:"8080"`
	} `yaml:"terraform_aws_modules"`
	Greeting string `env:"GREETING" env-description:"Greeting phrase" env-default:"Hello!"`
}

// Args command-line parameters
type Args struct {
	ConfigPath string
}