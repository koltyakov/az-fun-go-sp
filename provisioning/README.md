# Function App Provisioning

## Prerequisites

- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli?in=terraform/azure-get-started)
- [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli)

## Initiate provisioning (Terraform)

```bash
make init
```

## Variables

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

```bash
make plan
```

Plans provisioning, apply stuff in a dry mode.

## Functions app code packaging

Before applying Terraform templates, Azure Functions package should be created with:

```bash
make pack
```

This will create `./package/functions.zip`, which is used as app code source deployed to Blobs and bounded with Azure Functions app via `WEBSITE_RUN_FROM_PACKAGE` environment variable and corresponding automation.

## Applying infrastructure as code

```bash
make apply
```

An alias of `terraform apply -auto-approve`, won't prompt for confirmation.

Azure Functions and Custom Handlers in Go are now published and ready to use.

## Destroying when done testing

```bash
make destroy
```

> Don't run this for a prod as all the resources provisioned with Terraform templates will be removed.

## Nice articles around topic

- [Deploying an Azure Function App with Terraform](https://adrianhall.github.io/typescript/2019/10/23/terraform-functions/)