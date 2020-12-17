package handlers

import (
	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/koltyakov/gosip/api"
)

// Handlers base struct
type Handlers struct {
	sp *api.SP        // SharePoint API root object/client
	sa storage.Client // Storage Account API client
}

// NewHandlers handlers constructor
func NewHandlers(sp *api.SP, sa storage.Client) *Handlers {
	return &Handlers{
		sp: sp,
		sa: sa,
	}
}
