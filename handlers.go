package main

import (
	"encoding/json"
	"net/http"
)

func getListsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := sp.Web().Lists().Select("Id,Title").Top(500).Get()
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
	json.NewEncoder(w).Encode(lists)

	// fmt.Fprint(w, fmt.Sprintf("%s", data.Normalized()))
}

func getFieldsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := sp.Web().Fields().Select("InternalName,Title").Top(500).Get()
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
	json.NewEncoder(w).Encode(fields)
}
