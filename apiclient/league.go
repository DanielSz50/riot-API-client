package apiclient

import (
	"fmt"
	"strconv"
	"time"
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
	requestUrl := fmt.Sprintf("%s/%s/%s/%s?page=1", EntriesURL, queue, tier, division)

	var leagueEntries LeagueEntriesDTO
	if _, err := r.performRequest(requestUrl, &leagueEntries); err != nil {
		return nil, err
	}

	return leagueEntries, nil
}

func (r *riot) GetTopWinRatio(queue string, tier string) (*LeagueEntryDTO, error) {
	requestUrl := fmt.Sprintf("%s/%s/%s", EntriesURL, queue, tier)
	divisions := []string{"I", "II", "III", "IV"}

	maxWinRatio := 0.0
	var bestSummoner LeagueEntryDTO
	for i := 0; i < 4; i++ {
		tempURL := requestUrl + "/" + divisions[i] + "/"
		divURL := tempURL
		for p := 1; p < 1000; p++ {
			tempURL += "?page=" + strconv.Itoa(p)
			var leagueEntries LeagueEntriesDTO
			statusCode, err := r.performRequest(tempURL, &leagueEntries)
			if err != nil {
				return nil, err
			}
			if statusCode != 200 {
				fmt.Println(tempURL)
				time.Sleep(5 * time.Second)
				p--
				tempURL = divURL
				continue
			}
			if len(leagueEntries) == 0 {
				break
			}

			for _, e := range leagueEntries {
				winRatio := float64(e.Wins) / float64(e.Wins+e.Losses)
				if winRatio > maxWinRatio {
					maxWinRatio = winRatio
					bestSummoner = e
				}
			}
			tempURL = divURL
		}
	}

	return &bestSummoner, nil
}
