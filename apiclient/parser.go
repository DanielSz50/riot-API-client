package apiclient

import (
	"encoding/json"
	"log"
)

type Parser interface {
	parse(response []byte) error
}

func (s *SummonerDTO) parse(response []byte) error {
	if err := json.Unmarshal(response, s); err != nil {
		log.Printf("Couldn't unmarshall SummonerDTO JSON: %v\n", err)
		return err
	}
	return nil
}

func (g *CurrentGameInfo) parse(response []byte) error {
	if err := json.Unmarshal(response, g); err != nil {
		log.Printf("Couldn't unmarshall CurrentGameInfo JSON: %v\n", err)
		return err
	}

	return nil
}

func (l *LeagueEntriesDTO) parse(response []byte) error {
	if err := json.Unmarshal(response, l); err != nil {
		log.Printf("Couldn't unmarshall LeagueEntriesDTO: %v\n", err)
		return err
	}
	return nil
}

func (m *MatchDTO) parse(response []byte) error {
	if err := json.Unmarshal(response, m); err != nil {
		log.Printf("Couldn't unmarshall MatchDTO: %v\n", err)
		return err
	}
	return nil
}

func (m *MatchlistDTO) parse(response []byte) error {
	if err := json.Unmarshal(response, m); err != nil {
		log.Printf("Couldn't unmarshall MatchlistDTO: %v\n", err)
		return err
	}
	return nil
}
