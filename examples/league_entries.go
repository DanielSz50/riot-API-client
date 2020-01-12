package main

import (
	"fmt"
	riot "riot-API-client/apiclient"
)

func main() {
	client, err := riot.New()
	if err != nil {
		fmt.Println(err)
	}

	leagueEntries, err := client.GetLeagueEntries("RANKED_SOLO_5x5", "DIAMOND", "I")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(leagueEntries[0].SummonerName)

}
