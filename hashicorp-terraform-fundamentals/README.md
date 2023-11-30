# Hashicorp Terraform Fundamentals

<br />

## List of Contents:

### 1. [Initialize Terraform configuration](#content-1)

<br />

---

## Contents:
## [Initialize Terraform configuration](https://developer.hashicorp.com/terraform/tutorials/cli/init) <span id="content-1"></span>

### Introduction
- The core Terraform workflow consists of three main steps after you have written your Terraform configuration:
  - Initialize prepares your workspace so Terraform can apply your configuration.
  - Plan allows you to preview the changes Terraform will make before you apply them.
  - Apply makes the changes defined by your plan to create, update, or destroy resources.
- When you initialize a Terraform workspace, Terraform configures the backend, installs all providers and modules referred to in your configuration, and creates a version lock file if one doesn't already exist.
- In addition, you can use the terraform init command to change your workspace's backend and upgrade your workspace's providers and modules.

### Review configuration
- The directory contains Terraform configuration that uses multiple providers, a local module, and a remote module.
- The example repository includes the following:
  - LICENSE includes the text of the Mozilla Public License under which HashiCorp distributes the example configuration.
  - README.md describes the example configuration.
  - main.tf includes the resources and data sources used by the example configuration. 
  - the modules/aws-ec2-instance directory includes a Terraform module to provision an EC2 instance on AWS.
  - terraform.tf defines the terraform block, which defines the providers, remote backend, and the Terraform version(s) to be used with this configuration.
  - variables.tf defines the variables used in this configuration.

### Initialize your workspace
- Initialize the workspace with `terraform init`
- Terraform recognizes that the module "ec2-instance" block uses the local modules/aws-ec2-instance module. Next, Terraform determines that the module "hello" block references a remote module, so it downloads it from the public Terraform Registry.
- Since the configuration does not yet have a lock file, Terraform downloaded the aws and random providers specified in the required_providers block found in terraform.tf.
- When you initialize a workspace, Terraform will attempt to download the provider versions specified by the workspace's lock file.
- If the lock file does not exist, Terraform will use the required_providers block to determine the provider version and create a new lock file. If neither exists, Terraform will search for a matching provider and download the latest version.
- Next, Terraform creates the lock file if it does not already exist, or updates it if necessary.
- Terraform's lock file, .terraform.lock.hcl, records the versions and hashes of the providers used in this run. This ensures consistent Terraform runs in different environments, since Terraform will download the versions recorded in the lock file for future runs by default.
- When you manage Terraform configuration in a source control repository, commit the .terraform.lock.hcl file along with your configuration files.
- Finally, Terraform prints out a success message and reminds you how to plan your configuration, and to re-run terraform init if you change your modules or backend configuration.

### Review initialization artifacts
- When you initialize a new Terraform workspace, it creates a lock file named .terraform.lock.hcl and the .terraform directory.
- The lock file ensures that Terraform uses the same provider versions across your team and in remote execution environments. During initialization, Terraform will download the provider versions specified by this file rather than the latest versions.
- If the versions defined in the lock file's provider block do not match the versions defined in your configuration's required_providers block, Terraform will prompt you to re-initialize your configuration using the -upgrade flag. You will do this in the next section.
- Terraform uses the .terraform directory to store the project's providers and modules. Terraform will refer to these components when you run validate, plan, and apply,
- Terraform automatically manages the .terraform directory. Do not check it into version control, and do not directly modify this directory's contents.
- The aws-ec2-instance module refers to a local module, so Terraform refers directly to the module's configuration found within the modules/aws-ec2-instance directory in the example repository. This means that if you make changes to a local module, Terraform will recognize them immediately.
- Since the hello module is remote, Terraform downloaded the module from its source and saved a local copy in the .terraform/modules/hello directory when you initialized your workspace. Open the files in .terraform/modules/hello to view the module's configuration. These files are intended to be read-only, like the other contents in .terraform. Do not modify them. Terraform will only update a remote module when you run terraform init -upgrade or terraform get.
- The .terraform/providers directory stores cached versions of all of the configuration's providers.
- View the .terraform/providers directory. When you ran terraform init earlier, Terraform downloaded the providers defined in your configuration from the provider's source (defined by the required_providers block) and saved them in their respective directories, defined as [hostname]/[namespace]/[name]/[version]/[os_arch].

### Reinitialize configuration
- Since you updated the provider and module versions, you must re-initialize the configuration for Terraform to install the updated versions.
- If you attempt to validate, plan, or apply your configuration before doing so, Terraform will prompt you to re-initialize.
- Re-initialize your configuration to have Terraform upgrade the module to match the new version you configured in the previous step. Terraform will report an error for the provider, however.
- Re-initialize your configuration with the -upgrade flag. This tells Terraform to upgrade the provider to the most recent version that matches the version attribute in that provider's required_version block.
- Open the lock file. Notice that the random provider now uses version 3.5.1. Even though there are two versions of the random provider in .terraform/providers, Terraform will always use the version recorded in the lock file.
- Since you have updated your provider and module version, check whether your configuration is still valid: `terraform validate`
- Initialize your Terraform workspace with terraform init when:
  - You create new Terraform configuration and are ready to use it to create a workspace and provision infrastructure.
  - You clone a version control repository containing Terraform configuration, and are ready to use it to create a workspace and provision infrastructure.
  - You add, remove, or change the version of a module or provider in an existing workspace.
  - You add, remove, or change the backend or cloud blocks within the terraform block of an existing workspace.



**[⬆ back to top](#list-of-contents)**

<br />

---

## References: