package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/storage"
)

// To make it run locally, don't forget providing `AzureWebJobsStorage`
// in functions/local.settings.json

const (
	tableName    = "state"
	partitionKey = "settings"
	rowKey       = "workflow"
)

// Storage shows Azure Storage Account API basics
func (h *Handlers) Storage(w http.ResponseWriter, r *http.Request) {
	tableCli := h.sa.GetTableService()
	table := tableCli.GetTableReference(tableName)
	entity := table.GetEntityReference(partitionKey, rowKey)

	// Ensure table
	if err := entity.Table.Get(30, storage.FullMetadata); err != nil {
		if err := entity.Table.Create(20, storage.FullMetadata, nil); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	// Get existing state info
	_ = entity.Get(30, storage.FullMetadata, nil)
	result := entity.Properties
	if result == nil {
		// If the row key value is nil, return some defaults
		result = map[string]interface{}{
			"LastModifiedOn": time.Now().Add(-5 * 365 * 24 * time.Hour),
		}
	}

	// Update state with a new value
	entity.Properties = map[string]interface{}{
		"LastModifiedOn": time.Now(),
	}
	if err := entity.InsertOrReplace(nil); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Return previous state data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
