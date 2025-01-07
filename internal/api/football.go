package api

import (
	"fmt"
	"net/http"
)

type Match struct {
	HomeTeam string `json:"home_team"`
	AwayTeam string `json:"Away_team"`
	Score    string `json:"Score"`
}

func (api *Api) Api_FootballSite() {
	url := "https://api.football-data.org/v4/matches"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("url call failed ", err)
		return
	}

	defer resp.Body.Close()

}
