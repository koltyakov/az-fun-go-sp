resource "azurerm_app_service_plan" "service_plan" {
  name                = "${var.function_app}_Plan"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  kind                = "FunctionApp"
  reserved            = true

  sku {
    tier = "Dynamic"
    size = "Y1"
  }

  tags = var.tags
}

resource "azurerm_function_app" "functions" {
  name                       = var.function_app
  resource_group_name        = azurerm_resource_group.rg.name
  location                   = azurerm_resource_group.rg.location
  app_service_plan_id        = azurerm_app_service_plan.service_plan.id
  storage_account_name       = azurerm_storage_account.storage.name
  storage_account_access_key = azurerm_storage_account.storage.primary_access_key
  https_only                 = true
  os_type                    = "linux"
  version                    = "~3"

  app_settings = {
    FUNCTIONS_WORKER_RUNTIME = "custom"
    SPAUTH_SITEURL           = var.sharepoint_siteurl
    SPAUTH_CLIENTID          = var.sharepoint_clientid
    SPAUTH_CLIENTSECRET      = var.sharepoint_clientsecret
    FUNCTION_APP_EDIT_MODE   = "readonly"
    HASH                     = base64encode(filesha256(var.package))
    WEBSITE_RUN_FROM_PACKAGE = "https://${azurerm_storage_account.storage.name}.blob.core.windows.net/${azurerm_storage_container.deployments.name}/${azurerm_storage_blob.appcode.name}${data.azurerm_storage_account_sas.sas.sas}"
  }

  lifecycle {
    ignore_changes = [
      app_settings["WEBSITE_RUN_FROM_PACKAGE"], # prevent TF reporting configuration drift after app code is deployed
    ]
  }

  tags = var.tags
}

data "azurerm_function_app_host_keys" "keys" {
  name                = azurerm_function_app.functions.name
  resource_group_name = azurerm_function_app.functions.resource_group_name

  depends_on = [azurerm_function_app.functions]
}

output "function_app_default_hostname" {
  value = azurerm_function_app.functions.default_hostname
}

output "azurerm_function_app_host_keys" {
  value = data.azurerm_function_app_host_keys.keys.default_function_key
}