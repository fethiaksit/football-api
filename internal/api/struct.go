package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Match struct {
	HomeTeam string `json:"home_team"`
	AwayTeam string `json:"Away_team"`
	Score    string `json:"Score"`
}

func Api_FootballSite() {
	url := "https://api.football-data.org/v4/matches"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("url error")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var matches []Match
	json.Unmarshal(body, &matches)

	for _, match := range matches {
		fmt.Printf("%s vs %s Score: %s", match.HomeTeam, match.AwayTeam, match.Score)
	}
}
