package sportmonks

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Highlights struct
type Highlights struct {
	FixtureId int                    `json:"fixture_id"`
	Location  string                 `json:"location"`
	CreatedAt map[string]interface{} `json:"created_at"`
}

//GetHighlight get highlight by fixture Id
func GetHighLight(fixtureId int) ([]Highlights, error) {
	uri := "highlights/fixture/" + strconv.Itoa(fixtureId)

	response, err := getAnythingArray(uri)
	fmt.Println(uri)

	if err != nil {
		return []Highlights{}, fmt.Errorf("Error getting the data: %v", err)
	}

	var highlight []Highlights

	jsonEncodedHighlight, err := json.Marshal(response.Data)

	if err != nil {
		return []Highlights{}, fmt.Errorf("Error marshalling the highlight: %v", err)
	}

	if err = json.Unmarshal(jsonEncodedHighlight, &highlight); err != nil {
		return []Highlights{}, fmt.Errorf("Error unmarshalling the response data: %v", err)
	}

	return highlight, nil
}
