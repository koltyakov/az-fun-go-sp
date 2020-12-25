package handlers

import (
	"encoding/json"
	"net/http"
)

// Lists get SharePoint lists sample handler
func (h *Handlers) Lists(w http.ResponseWriter, r *http.Request) {
	data, err := h.sp.Web().Lists().Select("Id,Title").Top(500).Get()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	type listInfo struct {
		ID    string `json:"Id"`
		Title string
	}

	var lists []*listInfo

	if err := json.Unmarshal(data.Normalized(), &lists); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(lists)
}
