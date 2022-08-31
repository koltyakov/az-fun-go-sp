// https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/service_plan
resource "azurerm_service_plan" "service_plan" {
  name                = var.function_app
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  os_type             = "Linux"
  sku_name            = "Y1"
  tags                = var.tags
}

// https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/linux_function_app
resource "azurerm_linux_function_app" "functions" {
  name                       = var.function_app
  resource_group_name        = azurerm_resource_group.rg.name
  service_plan_id            = azurerm_service_plan.service_plan.id
  location                   = azurerm_resource_group.rg.location
  storage_account_name       = azurerm_storage_account.storage.name
  storage_account_access_key = azurerm_storage_account.storage.primary_access_key
  https_only                 = true

  app_settings = {
    FUNCTIONS_WORKER_RUNTIME = "custom"
    SPAUTH_SITEURL           = var.sharepoint_siteurl
    SPAUTH_CLIENTID          = var.sharepoint_clientid
    SPAUTH_CLIENTSECRET      = var.sharepoint_clientsecret

    FUNCTION_APP_EDIT_MODE   = "readonly"
    HASH                     = base64encode(filesha256(var.package))
    WEBSITE_RUN_FROM_PACKAGE = "https://${azurerm_storage_account.storage.name}.blob.core.windows.net/${azurerm_storage_container.deployments.name}/${azurerm_storage_blob.appcode.name}${data.azurerm_storage_account_sas.sas.sas}"

    APPINSIGHTS_INSTRUMENTATIONKEY = azurerm_application_insights.appinsights.instrumentation_key
  }

  site_config {}

  tags = var.tags
}

resource "azurerm_application_insights" "appinsights" {
  name                = var.function_app
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  application_type    = "other"
}

data "azurerm_function_app_host_keys" "keys" {
  name                = azurerm_linux_function_app.functions.name
  resource_group_name = azurerm_linux_function_app.functions.resource_group_name

  depends_on = [azurerm_linux_function_app.functions]
}

output "function_app_hostname" {
  value = azurerm_linux_function_app.functions.default_hostname
}

output "function_app_key" {
  value     = data.azurerm_function_app_host_keys.keys.default_function_key
  sensitive = true
}
