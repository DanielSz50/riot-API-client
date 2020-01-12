package apiclient

import "fmt"

type SummonerDTO struct {
	ProfileIconId int    `json:"profileIconId"`
	Name          string `json:"name"`
	Puuid         string `json:"puuid"`
	SummonerLevel int64  `json:"summonerLevel"`
	RevisionDate  int64  `json:"revisionDate"`
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
}

func (r *riot) GetSummoner(summonerName string) (*SummonerDTO, error) {
	requestUrl := fmt.Sprintf("%s/%s", SummonerURL, summonerName)

	var summoner SummonerDTO
	if _, err := r.performRequest(requestUrl, &summoner); err != nil {
		return nil, err
	}

	return &summoner, nil
}
