# Azure Functions Golang & SharePoint Sample

## Prerequisites

- [Azure Functions Core Tools, v3](https://www.npmjs.com/package/azure-functions-core-tools)
- [Go SDK](https://golang.org/dl/)
- [Visual Studio Code](https://code.visualstudio.com) & [Azure Functions extension](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-azurefunctions)

## Create Azure Function Resource

> This flow is dev friendly when it comes for quick experiments. For production deployments please concider infrustructure as code and automation via DevOps.

In VS Code with Azure Functions entension installed:

- `CMD+Shift+P` type `Azure Functions: Create Function App in Azure... (Advanced)` and folow the wizard (further is the example of the steps)
- Select dev subscription
- Enter a globally unique name for the new function app
- Select a runtime stack: `Custom Handler`
- Select an OS: `Linux`
- Select a hosting plan: `Consumption`
- Create a new resource group
- Enter the name of the new resource group
- Create new storage account
- Enter the name of the new storage account
- Create new Application Insights resource
- Enter the name of the new Application Insights resource
- Select a location for new resources (closer to your location)

> Don't forget to remove the resource group after experiments to avoid any unpredictable charges

## Local debug

- Copy/rename `local.settings.sample.json` to `local.settings.json`.
- Provide authentication parameters:
    - SPAUTH_SITEURL
    - SPAUTH_CLIENTID
    - SPAUTH_CLIENTSECRET

[Add-in Only auth](https://go.spflow.com/auth/strategies/addin) is used as a sample.

- Start local server with ```make start```
- Navigate to one of the URL endpoints printed in the console

## Build and deploy

1\. Run build command:

```bash
make build-linux
```

The assumption is that a Linux plan was selected for the function app.

The build creates `bin/server` binary with Go custom handler.

2\. In VSCode,

- `CMD+Shift+P` type `Azure Functions: Deploy to Function App...`
- Select subscription
- Select Function App in Azure
- Confirm deployment

## Configure environment variables

In Azure Function app, create and provide the following environment variables:
- SPAUTH_SITEURL
- SPAUTH_CLIENTID
- SPAUTH_CLIENTSECRET

Depending on SharePoint environment and use case, [auth strategy](https://go.spflow.com/auth/strategies) can be different. For a production installation [Azure Certificate Auth](https://go.spflow.com/auth/custom-auth/azure-certificate-auth) might be prefered.

`spauth.go/init` handler should be aligned with authentication parameters.

## Reference

- [Azure Functions custom handlers](https://docs.microsoft.com/en-us/azure/azure-functions/functions-custom-handlers)
- [Create a Go function in Azure using VSCode](https://docs.microsoft.com/en-us/azure/azure-functions/create-first-function-vs-code-other)