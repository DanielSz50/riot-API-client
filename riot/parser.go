package riot

import (
	"encoding/json"
	"fmt"
)

type Parser interface {
	parse(response []byte) error
}

func (s *SummonerDTO) parse(response []byte) error {
	if err := json.Unmarshal(response, s); err != nil {
		return fmt.Errorf("Couldn't unmarshall SummonerDTO: %w\n", err)
	}
	return nil
}

func (g *CurrentGameInfo) parse(response []byte) error {
	if err := json.Unmarshal(response, g); err != nil {
		return fmt.Errorf("Couldn't unmarshall CurrentGameInfo: %w\n", err)
	}
	return nil
}

func (l *LeagueEntriesDTO) parse(response []byte) error {
	if err := json.Unmarshal(response, l); err != nil {
		return fmt.Errorf("Couldn't unmarshall LeagueEntriesDTO: %w\n", err)
	}
	return nil
}

func (m *MatchDTO) parse(response []byte) error {
	if err := json.Unmarshal(response, m); err != nil {
		return fmt.Errorf("Couldn't unmarshall MatchDTO: %w\n", err)
	}
	return nil
}

func (m *MatchlistDTO) parse(response []byte) error {
	if err := json.Unmarshal(response, m); err != nil {
		return fmt.Errorf("Couldn't unmarshall MatchlistDTO: %w\n", err)
	}
	return nil
}
