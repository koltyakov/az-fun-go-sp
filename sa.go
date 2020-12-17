package main

import (
	"os"

	"github.com/Azure/azure-sdk-for-go/storage"
)

// Gets Azure Storage Account client
func getSA() (storage.Client, error) {
	sa, err := storage.NewClientFromConnectionString(os.Getenv("AzureWebJobsStorage"))
	return sa, err
}
