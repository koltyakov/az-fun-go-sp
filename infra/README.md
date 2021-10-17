# Function App Provisioning

The sample shown automated reproducable Azure Functions App provisioning and deployment using Terraform.

> Terraform is infrastructure as code software tool that provides a consistent CLI workflow to manage hundreds of cloud services. Terraform codifies cloud APIs into declarative configuration files `.tf`.

## Prerequisites

- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli?in=terraform/azure-get-started), [Azure Provider](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs)
- [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli)

## Initiate infra (Terraform)

Read `State Backend` before initiating terraform.

```bash
make init
```

### State Backend

The project uses [Azure Backend](https://www.terraform.io/docs/language/settings/backends/azurerm.html) for storing Terraform state.

Before starting, copy `.env.sample` to `.env` and provide a valid `TERRAFORM_BACKEND_STORAGE_ACCOUNT`. The account should be created in your infrastructure first:

- Create resource group: `Terraform`
- Add storage account (which name goes to `TERRAFORM_BACKEND_STORAGE_ACCOUNT`)
- Create `tfstate` container with private access

Then run init command.

### Local State

To use local state, comment `backend "azurerm" {}` in `01_main.tf` model.

## Project structure

Terraform `.tf` files' content composition mostly doesn't matter much until you feel comfortable with what is where and how structured for reusability or maintenance robustness. It can be all in one file on domain areas split between multiple `.tf`s.

- `main.tf` - main Terraform file, sort of an entry point. I used to place the provider, resource group, and network stanzas here.
  - `storage.tf` - storage account is a pillar of an Azure Functions app. Here storage is combined with deployment via storage container blobs artifacts.
  - `functions.tf` - functions app and service plan definitions, data automation outputs (such as provisioned functions hostname and function key).
- `variables.tf` - dynamic input variables definition.
  - `terraform.sample.tfvars` - sample variables file, ignored. For simpler copy & paste into `terraform.tfvars`.
  - `terraform.tfvars` - variables values files. Should not be included in the repo, might contain private data and secrets. The variables also can be provided with command line arguments.
- `Makefile` - automation commands aliases.
- Miscellaneous:
  - `.terraform.lock.hcl` - dependencies lock file used with `terraform init`.
  - `.terraform/` folder - Terraform provider's binaries. Aka "node_modules/" of terraformation. Might have tangible size.
  - `terraform.tfstate` - remote state sync file.
  - `package/functions.zip` - custom handler deployment package produced with `make pack` command.

## Initiate project

Aka "npm install" for terraform providers. Initiates Terraform project and downloads missed provisioning providers.

```bash
make init
```

## Variables file

Terraform variables are private dynamic values/inputs which are used in the templates allowing targeting the same tested definitions agains different environments or tenants. 

- Copy & rename `terraform.sample.tfvars` to `terraform.tfvars`
- Provide your variables

```hcl
subscription_id         = "529f730b-4ac8-43b6-9851-221e735a71cf"
location                = "westus2"
function_app            = "az-fun-go-042"
sharepoint_siteurl      = "https://contoso.sharepoint.com/sites/site"
sharepoint_clientid     = "428b492b-575d-4d4b-991e-16195a3c496e"
sharepoint_clientsecret = "CgnihMbRphqR....7XLlZ/0QCgw="
```

> Avoid variables commit to Git repository

## Plan before apply

Plans provisioning, apply stuff in a dry mode. Can be useful for trubleshooting and verification of what a deployment will do. E.g. in some cases a deployment of a resource can't be done without destroying and recreating, which however might be unwanted behavior. With planning, you'd see what property causes such effect.

```bash
make plan
```

## Functions app code packaging

Before applying Terraform templates, the Azure Functions package should be created with:

```bash
make pack
```

This will create `./package/functions.zip`, which is used as an app code source deployed to Blobs and bounded with the Azure Functions app via `WEBSITE_RUN_FROM_PACKAGE` environment variable and corresponding automation.

## Applying infrastructure as code

```bash
make apply
```

The command is an alias of `terraform apply -auto-approve`, won't prompt for confirmation.

Azure Functions and Custom Handlers in Go are now published and ready to use.

## Destroying when done testing

```bash
make destroy
```

> Don't run this for a prod as all the resources provisioned with Terraform templates will be removed.

### Functions App choices

#### Dynamic Linux plan

I tend to believe and prefer serverless options first approach. Also, I'd chose Linux on a server in all cases when it works.

```hcl
## functions.tf

# Service Plan
resource "azurerm_app_service_plan" "service_plan" {
  name = "${var.function_app}_Plan"
  # ...

  kind = "FunctionApp" # Azure Functions

  sku {
    tier = "Dynamic" # Dynamic App Service Plan
    size = "Y1"      # Linux
  }
}

# Function App
resource "azurerm_function_app" "functions" {
  name = var.function_app
  # ...

  os_type = "linux" # Currently if omitted might lead to Functions App redeployment on each apply
  version = "~3"    # Latest Azure Functions version, required for Custom Handler support

  app_settings = {
    FUNCTIONS_WORKER_RUNTIME = "custom" # Custom Handler runtime
    # ...
  }
}
```

#### SharePoint bindings

```hcl
## functions.tf

resource "azurerm_function_app" "functions" {
  name = var.function_app
  # ...

  app_settings = {
    # SharePoint auth binding configuration, see more in a root project
    SPAUTH_SITEURL           = var.sharepoint_siteurl
    SPAUTH_CLIENTID          = var.sharepoint_clientid
    SPAUTH_CLIENTSECRET      = var.sharepoint_clientsecret

    # ...
  }
}

## variables.tf

variable "sharepoint_siteurl" {
  type        = string
  description = "SharePoint SiteURL"
}

variable "sharepoint_clientid" {
  type        = string
  description = "SharePoint ClientID"
}

variable "sharepoint_clientsecret" {
  type        = string
  description = "SharePoint CLient Secret"
}

## terraform.tfvars

sharepoint_siteurl      = "https://contoso.sharepoint.com/sites/site"
sharepoint_clientid     = "428b492b-575d-4d4b-991e-16195a3c496e"
sharepoint_clientsecret = "CgnihMbRphqR....7XLlZ/0QCgw="
```

#### Package deployment

```hcl
## functions.tf

resource "azurerm_function_app" "functions" {
  name = var.function_app
  # ...

  app_settings = {
    # Deployment from storage container blob settings
    FUNCTION_APP_EDIT_MODE   = "readonly"
    HASH                     = base64encode(filesha256(var.package))
    WEBSITE_RUN_FROM_PACKAGE = "https://${azurerm_storage_account.storage.name}.blob.core.windows.net/${azurerm_storage_container.deployments.name}/${azurerm_storage_blob.appcode.name}${data.azurerm_storage_account_sas.sas.sas}"

    # ...
  }
}

## storage.tf

resource "azurerm_storage_container" "deployments" {
  # ...
}

resource "azurerm_storage_blob" "appcode" {
  name = "functions.zip"
  type = "Block"
  source = var.package # Package archive upload

  # ...
}

# Deployment blob access connection string
data "azurerm_storage_account_sas" "sas" {
   # ...
}

## variables.tf

variable "package" {
  type    = string
  default = "./package/functions.zip" # Package default location
}
```

## Nice to read article(s) around the topic

- [Deploying an Azure Function App with Terraform](https://adrianhall.github.io/typescript/2019/10/23/terraform-functions/)
