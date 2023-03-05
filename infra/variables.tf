variable "location" {
  type        = string
  description = "Azure DC Location"
}

variable "subscription_id" {
  type        = string
  description = "Azure Subscription ID"
}

variable "function_app" {
  type        = string
  description = "Azure Function App Name"
}

variable "tags" {
  type = map(any)

  default = {
    Environment = "Dev"
    Stack       = "Go"
  }
}

# SharePoint Bindings

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

# Custom handlers package

variable "package" {
  type    = string
  default = "./package/functions.zip"
}
