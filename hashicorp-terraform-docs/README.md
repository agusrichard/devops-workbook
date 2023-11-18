# Hashicorp Terraform Docs

<br />

## List of Contents:

### 1. [What is Infrastructure as Code with Terraform?](#content-1)
### 2. [Install Terraform](#content-2)
### 3. [Build Infrastructure](#content-3)
### 4. [Change Infrastructure](#content-4)
### 5. [Destroy Infrastructure](#content-5)
### 6. [Define Input Variables](#content-6)

<br />

---

## Contents:
## [What is Infrastructure as Code with Terraform?](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/infrastructure-as-code) <span id="content-1"></span>

### Standardize your deployment workflow
- To deploy infrastructure with Terraform:
  - Scope - Identify the infrastructure for your project.
  - Author - Write the configuration for your infrastructure.
  - Initialize - Install the plugins Terraform needs to manage the infrastructure.
  - Plan - Preview the changes Terraform will make to match your configuration.
  - Apply - Make the planned changes.
  Commands

## [Install Terraform](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli) <span id="content-2"></span>

### Install
- Using Homebrew
  - brew tap hashicorp/tap
  - brew install hashicorp/tap/terraform
  - brew update
  - brew upgrade hashicorp/tap/terraform

### Verify Installation
- terraform -help
- terraform -help plan

## [Build Infrastructure](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/aws-build) <span id="content-3"></span>

### Write Configuration
- Inside `main.tf`:
  ```text
  terraform {
    required_providers {
      aws = {
        source  = "hashicorp/aws"
        version = "~> 4.16"
      }
    }

    required_version = ">= 1.2.0"
  }

  provider "aws" {
    region  = "us-west-2"
  }

  resource "aws_instance" "app_server" {
    ami           = "ami-830c94e3"
    instance_type = "t2.micro"

    tags = {
      Name = "ExampleAppServerInstance"
    }
  }
  ```

### Terraform Block
- The terraform {} block contains Terraform settings, including the required providers Terraform will use to provision your infrastructure.
- For each provider, the source attribute defines an optional hostname, a namespace, and the provider type.

### Providers
- The provider block configures the specified provider, in this case aws. A provider is a plugin that Terraform uses to create and manage your resources.
- You can use multiple provider blocks in your Terraform configuration to manage resources from different providers. You can even use different providers together. For example, you could pass the IP address of your AWS EC2 instance to a monitoring resource from DataDog.

### Resources
- Use resource blocks to define components of your infrastructure. A resource might be a physical or virtual component such as an EC2 instance, or it can be a logical resource such as a Heroku application.
- Resource blocks have two strings before the block: the resource type and the resource name. In this example, the resource type is aws_instance and the name is app_server.
- The prefix of the type maps to the name of the provider. In the example configuration, Terraform manages the aws_instance resource with the aws provider. Together, the resource type and resource name form a unique ID for the resource. For example, the ID for your EC2 instance is aws_instance.app_server.
- Resource blocks contain arguments which you use to configure the resource. Arguments can include things like machine sizes, disk image names, or VPC IDs. Our providers reference lists the required and optional arguments for each resource. For your EC2 instance, the example configuration sets the AMI ID to an Ubuntu image, and the instance type to t2.micro, which qualifies for AWS' free tier. It also sets a tag to give the instance a name.

### Initialize the directory
- When you create a new configuration — or check out an existing configuration from version control — you need to initialize the directory with terraform init.
- Initializing a configuration directory downloads and installs the providers defined in the configuration, which in this case is the aws provider.
- Terraform downloads the aws provider and installs it in a hidden subdirectory of your current working directory, named .terraform.
- The terraform init command prints out which version of the provider was installed. Terraform also creates a lock file named .terraform.lock.hcl which specifies the exact provider versions used, so that you can control when you want to update the providers used for your project.

### Format and validate the configuration
- The terraform fmt command automatically updates configurations in the current directory for readability and consistency.
- You can also make sure your configuration is syntactically valid and internally consistent by using the terraform validate command.

### Create infrastructure
- Apply the configuration now with the terraform apply command.
- Before it applies any changes, Terraform prints out the execution plan which describes the actions Terraform will take in order to change your infrastructure to match the configuration.

### Inspect state
- When you applied your configuration, Terraform wrote data into a file called terraform.tfstate.
- Terraform stores the IDs and properties of the resources it manages in this file, so that it can update or destroy those resources going forward.
- The Terraform state file is the only way Terraform can track which resources it manages, and often contains sensitive information, so you must store your state file securely and restrict access to only trusted team members who need to manage your infrastructure.
- In production, we recommend storing your state remotely with Terraform Cloud or Terraform Enterprise. Terraform also supports several other remote backends you can use to store and manage your state.
- Inspect the current state using `terraform show`


## [Change Infrastructure](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/aws-change) <span id="content-4"></span>

### Configuration
- Now update the ami of your instance. Change the aws_instance.app_server resource under the provider block in main.tf by replacing the current AMI ID with a new one.

### Apply Changes
- After changing the configuration, run terraform apply again to see how Terraform will apply this change to the existing resources.
- The prefix -/+ means that Terraform will destroy and recreate the resource, rather than updating it in-place.
- Terraform can update some attributes in-place (indicated with the ~ prefix), but changing the AMI for an EC2 instance requires recreating it. Terraform handles these details for you, and the execution plan displays what Terraform will do.
- Additionally, the execution plan shows that the AMI change is what forces Terraform to replace the instance. Using this information, you can adjust your changes to avoid destructive updates if necessary.


## [Destroy Infrastructure](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/aws-destroy) <span id="content-5"></span>

### Destroy
- The terraform destroy command terminates resources managed by your Terraform project.
- This command is the inverse of terraform apply in that it terminates all the resources specified in your Terraform state. It does not destroy resources running elsewhere that are not managed by the current Terraform project.
- Run `terraform destroy`
- The - prefix indicates that the instance will be destroyed. As with apply, Terraform shows its execution plan and waits for approval before making any changes.
- Just like with apply, Terraform determines the order to destroy your resources. In this case, Terraform identified a single instance with no other dependencies, so it destroyed the instance. In more complicated cases with multiple resources, Terraform will destroy them in a suitable order to respect dependencies.

## [Define Input Variables](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/aws-variables) <span id="content-6"></span>

### Set the instance name with a variable
- The current configuration includes a number of hard-coded values. Terraform variables allow you to write configuration that is flexible and easier to re-use.
- Create a new file called variables.tf with a block defining a new instance_name variable.
  ```text
  variable "instance_name" {
    description = "Value of the Name tag for the EC2 instance"
    type        = string
    default     = "ExampleAppServerInstance"
  }
  ```
- In main.tf, update the aws_instance resource block to use the new variable. The instance_name variable block will default to its default value ("ExampleAppServerInstance") unless you declare a different value.
  ```text
  resource "aws_instance" "app_server" {
     ami           = "ami-08d70e59c07c61a3a"
     instance_type = "t2.micro"

     tags = {
  -    Name = "ExampleAppServerInstance"
  +    Name = var.instance_name
     }
   }
  ```
  
### Apply Your Configuration
- Run `terraform apply`
- Now apply the configuration again, this time overriding the default instance name by passing in a variable using the -var flag. Terraform will update the instance's Name tag with the new name. Respond to the confirmation prompt with yes.
  ```shell
  terraform apply -var "instance_name=YetAnotherName"
  ```


**[⬆ back to top](#list-of-contents)**

<br />

---

## References: