package main

import (
	"fmt"
	"log"
	"riot-API-client/riot"
)

func main() {
	client, err := riot.New()
	if err != nil {
		log.Println(err)
	}

	leagueEntries, err := client.GetLeagueEntries(riot.QueueRankedSolo5, riot.TierDiamond, riot.Division1)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(leagueEntries[0].SummonerName)
}
