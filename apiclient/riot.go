package apiclient

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const ApiKeyPath = "./.api_key"

type Riot interface {
	performRequest(url string, data Parser) (int, error)
	GetSummoner(summonerName string) (*SummonerDTO, error)
	GetTopWinRatio(queue string, tier string) (*LeagueEntryDTO, error)
	GetLeagueEntries(queue string, tier string, division string) (LeagueEntriesDTO, error)
	GetMatch(matchID int64) (*MatchDTO, error)
	GetMatchlist(accountID string) (*MatchlistDTO, error)
	GetCurrentGame(summonerName string) (*CurrentGameInfo, error)
}

type riot struct {
	client *http.Client
	apiKey string
}

func New() (Riot, error) {
	client := &http.Client{}

	keyFile, err := os.Open(ApiKeyPath)
	if err != nil {
		return nil, err
	}

	apiKey := make([]byte, 50)
	n, err := keyFile.Read(apiKey)
	if err != nil {
		return nil, err
	}

	apiKey = apiKey[:n-1]

	return &riot{client, string(apiKey)}, nil
}

func (r *riot) performRequest(url string, data Parser) (int, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 1, err
	}
	req.Header.Add("X-Riot-Token", r.apiKey)

	resp, err := r.client.Do(req)
	if err != nil {
		log.Printf("Couldn't perform the request: %v\n", err)
		log.Println(url)
		return resp.StatusCode, err
	}
	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(b))
		return resp.StatusCode, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Couldn't read response body: %v\n", err)
		return resp.StatusCode, err
	}

	if err := data.parse(respBody); err != nil {
		log.Printf("Couldn't unmarshall json response: %v\n", err)
		return resp.StatusCode, err
	}

	return resp.StatusCode, err
}
