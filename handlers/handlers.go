package handlers

import "github.com/koltyakov/gosip/api"

// Handlers base struct
type Handlers struct {
	sp *api.SP
}

// NewHandlers handlers constructor
func NewHandlers(sp *api.SP) *Handlers {
	return &Handlers{
		sp: sp,
	}
}
