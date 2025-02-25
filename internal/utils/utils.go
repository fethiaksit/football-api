package utils

import (
	"encoding/json"
	"fmt"
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
func FormatScore(homeScore interface{}, awayScore interface{}) string {
	return fmt.Sprintf("%v - %v", homeScore, awayScore)
}
