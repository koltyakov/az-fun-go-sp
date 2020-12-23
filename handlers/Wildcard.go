package handlers

import (
	"fmt"
	"net/http"
)

// Wildcard route traps all /api/* requests and response with URL Path
func (h *Handlers) Wildcard(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", r.URL.Path)
}
