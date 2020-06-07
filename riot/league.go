package riot

import (
	"fmt"
)

type LeagueEntryDTO struct {
	QueueType    string `json:"queueType"`
	LeagueID     string `json:"leagueID"`
	SummonerID   string `json:"summonerID"`
	SummonerName string `json:"summonerName"`
	LeaguePoints int    `json:"leaguePoints"`
	Rank         string `json:"rank"`
	Tier         string `json:"tier"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
	Veteran      bool   `json:"veteran"`
	Inactive     bool   `json:"inactive"`
	FreshBlood   bool   `json:"freshBlood"`
	HotStreak    bool   `json:"hotStreak"`
}

type LeagueEntriesDTO []LeagueEntryDTO

type LeagueListDTO struct {
	Tier     string          `json:"tier"`
	LeagueID string          `json:"leagueID"`
	Queue    string          `json:"queue"`
	Name     string          `json:"name"`
	Entries  []LeagueItemDTO `json:"entries"`
}

type LeagueItemDTO struct {
	SummonerID   string `json:"summonerID"`
	SummonerName string `json:"summonerName"`
	LeaguePoints int    `json:"leaguePoints"`
	Rank         string `json:"rank"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
	Veteran      bool   `json:"veteran"`
	Inactive     bool   `json:"inactive"`
	FreshBlood   bool   `json:"freshBlood"`
	HotStreak    bool   `json:"hotStreak"`
}

func (r *riot) GetLeagueEntries(queue string, tier string, division string) (LeagueEntriesDTO, error) {
	requestUrl := fmt.Sprintf("%s/%s/%s/%s?page=1", EndpointEntries, queue, tier, division)

	var leagueEntries LeagueEntriesDTO
	if _, err := r.sendRequest(requestUrl, &leagueEntries); err != nil {
		return nil, err
	}

	return leagueEntries, nil
}
