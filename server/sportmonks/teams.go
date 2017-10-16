package sportmonks

// Team : équipe
type Team struct {
	// ID of team
	ID int `json:"id"`
	// LegacyID : unused
	LegacyID int `json:"legacy_id"`
	// Twitter : twitter account
	Twitter string `json:"twitter"`
	// Name : team's name
	Name string `json:"name"`
	// CountryID :  ID of the team's country
	CountryID int `json:"country_id"`
	// NationalTeam : is it a national team
	NationalTeam bool `json:"national_team"`
	// Founded : date of team creation
	Founded int `json:"founded"`
	// LogoPath : url of the team's logo
	LogoPath string `json:"logo_path"`
	// VenueID : match ID
	VenueID int `json:"venue_id"`
}
