package riot

const (
	/* API ENDPOINTS */
	// LEAGUE-V4
	EndpointEntries = "https://eun1.api.riotgames.com/lol/league/v4/entries"

	// MATCH-V4
	EndpointMatches    = "https://eun1.api.riotgames.com/lol/match/v4/matches"
	EndpointMatchlists = "https://eun1.api.riotgames.com/lol/match/v4/matchlists/by-account"

	// SPECTATOR-V4
	EndpointSpectator = "https://eun1.api.riotgames.com/lol/spectator/v4/active-games/by-summoner"

	// SUMMONER-V4
	EndpointSummoner = "https://eun1.api.riotgames.com/lol/summoner/v4/summoners/by-name"

	/* PATH PARAMETERS */
	// DIVISIONS
	Division1 = "I"
	Division2 = "II"
	Division3 = "III"
	Division4 = "IV"

	// TIERS
	TierDiamond  = "DIAMOND"
	TierPlatinum = "PLATINUM"
	TierGold     = "GOLD"
	TierSilver   = "SILVER"
	TierBronze   = "BRONZE"
	TierIron     = "IRON"

	// QUEUES
	QueueRankedSolo5  = "RANKED_SOLO_5x5"
	QueueRankedFlexSR = "RANKED_FLEX_SR"
	QueueRankedFlexTT = "RANKED_FLEX_TT"
)
