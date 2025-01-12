package utils

import (
	"encoding/json"
	"net/http"
)

type testWriter struct {
	Wrt string `json:"write"`
}

func ResponseWriter(w http.ResponseWriter, wrt string) interface{} {
	resp := testWriter{
		Wrt: wrt,
	}
	response, _ := json.Marshal(resp)
	w.Write(response)
	return response
}
