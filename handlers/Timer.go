package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// InvokeResponse ...
type InvokeResponse struct {
	Outputs     map[string]interface{}
	ReturnValue interface{}
	Logs        []string
}

// TimerInfo ...
type TimerInfo struct {
	Data     map[string]interface{}
	Metadata map[string]interface{}
}

// Timer executes renewal sample process
func (h *Handlers) Timer(w http.ResponseWriter, r *http.Request) {
	defer func() { _ = r.Body.Close() }()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// fmt.Printf("%s\n", data) // TimerInfo

	response := &InvokeResponse{
		ReturnValue: fmt.Sprintf("%s", data),
		Outputs: map[string]interface{}{
			"output1": "",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
