package riot

import (
	"fmt"
)

type CurrentGameInfo struct {
	GameId            int64                    `json:"gameId"`
	GameStartTime     int64                    `json:"gameStartTime"`
	PlatformId        string                   `json:"platformId"`
	GameMode          string                   `json:"gameMode"`
	MapId             int64                    `json:"mapId"`
	GameType          string                   `json:"gameType"`
	BannedChampions   []BannedChampion         `json:"bannedChampions"`
	Observers         Observer                 `json:"observers"`
	Participants      []CurrentGameParticipant `json:"participants"`
	GameLength        int64                    `json:"gameLength"`
	GameQueueConfigId int64                    `json:"gameQueueConfigId"`
}

type BannedChampion struct {
	PickTurn   int64 `json:"pickTurn"`
	ChampionId int64 `json:"championId"`
	TeamId     int64 `json:"teamId"`
}

type Observer struct {
	EncryptionKey string `json:"encryptionKey"`
}

type CurrentGameParticipant struct {
	ProfileIconId            int64                     `json:"profileIconId"`
	ChampionId               int64                     `json:"championId"`
	SummonerName             string                    `json:"summonerName"`
	GameCustomizationObjects []GameCustomizationObject `json:"gameCustomizationObjects"`
	Bot                      bool                      `json:"bot"`
	Perks                    Perks                     `json:"perks"`
	Spell2Id                 int64                     `json:"spell2Id"`
	TeamId                   int64                     `json:"teamId"`
	Spell1Id                 int64                     `json:"spell1Id"`
	SummonerId               string                    `json:"summonerId"`
}

type GameCustomizationObject struct {
	Category int64 `json:"category"`
	Content  int64 `json:"content"`
}

type Perks struct {
	PerkStyle    int64   `json:"perkStyle"`
	PerkIds      []int64 `json:"perkIds"`
	PerkSubStyle int64   `json:"perkSubStyle"`
}

func (r *riot) GetCurrentGame(summonerName string) (*CurrentGameInfo, error) {
	summoner, err := r.GetSummoner(summonerName)
	if err != nil {
		return nil, err
	}

	requestUrl := fmt.Sprintf("%s/%s", EndpointSpectator, summoner.ID)
	var game CurrentGameInfo
	if _, err := r.sendRequest(requestUrl, &game); err != nil {
		return nil, err
	}

	return &game, nil
}
