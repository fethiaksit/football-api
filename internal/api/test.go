package api

import (
	"encoding/json"
	"net/http"
)

type testWriter struct {
	Wrt string `json:"write"`
}

func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := testWriter{
		Wrt: "test",
	}
	response, _ := json.Marshal(resp)
	w.Write(response)
}
