package ./main

type Team struct {
	// Id : id of team
	Id int `json:"id"`
	// Legacy_id : id of legacy
	LegacyId int `json:"legacy_id"`
	// twitter
	Twitter string `json:"twitter"`
	// Name : name of team
	Name string `json:"name"`
	// country_id : id of team country
	CountryId int `json:"country_id"`
	// national_team : yes or no it's an national team
	NationalTeam bool `json:"national_team"`
	// founded : date of team
	Founded int `json:"founded"`
	// logo_path : logo of team
	LogoPath string `json:"logo_path"`
	//
	VenueId int `json:"venue_id"`
}