package handlers

import (
	"encoding/json"
	"net/http"
)

// GetFields get SharePoint fields sample handler
func (h *Handlers) GetFields(w http.ResponseWriter, r *http.Request) {
	data, err := h.sp.Web().Fields().Select("InternalName,Title").Top(500).Get()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	type fieldInfo struct {
		InternalName string
		Title        string
	}

	var fields []*fieldInfo

	if err := json.Unmarshal(data.Normalized(), &fields); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(fields)
}
