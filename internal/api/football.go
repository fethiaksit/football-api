package api

import (
	"io/ioutil"
	"net/http"
)

type Match struct {
	HomeTeam string `json:"homeTeam"`
	AwayTeam string `json:"awayTeam"`
	Score    string `json:"score"`
	Status   string `json:"status"`
}

func NewRedisClient() *RedisClient {
	return &RedisClient{}
}

func (api *RedisClient) FetchLiveMatches(w http.ResponseWriter, r *http.Request) {
	url := "https://api.collectapi.com/sport/leaguesList"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "apikey 5lh280zbxpz265bVaRpqEu:50SxWSCbug7vJaQswNQQbc")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to fetch matches", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func (api *RedisClient) FetchSuperLigResults(w http.ResponseWriter, r *http.Request) {
	url := "https://api.collectapi.com/football/league?data.league=super-lig"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "apikey 5lh280zbxpz265bVaRpqEu:50SxWSCbug7vJaQswNQQbc")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to fetch results", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
