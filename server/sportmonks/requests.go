package sportmonks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Response : SportMonks soccer API response
type Response struct {
	// Data : string
	Data map[string]interface{} `json:"data"`
	// Number
	Meta map[string]interface{} `json:"meta"`
}

var (
	apiURI = "https://soccer.sportmonks.com/api/v2.0/"
	token  = "?api_token=QCaGhUsahBbiKEQPGINe7G843Jr5qSsdKmPpUpvR4MJee7xGYw4t63z4YgkO"
)

func getAnything(endpointWithParameters string) (Response, error) {
	client := &http.Client{}

	url := apiURI + endpointWithParameters + token

	// Prepare the request
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return Response{}, fmt.Errorf("Error creating the request: %v", err)
	}

	// Execute the request
	res, err := client.Do(req)

	if err != nil {
		return Response{}, fmt.Errorf("Error executing the request: %v", err)
	}

	defer res.Body.Close()

	var response Response

	// Read the response
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return Response{}, fmt.Errorf("Error reading the response body: %v", err)
	}

	// Unmarshal the response : { data: { ... }, meta: { ... }}
	if err = json.Unmarshal(body, &response); err != nil {
		return Response{}, fmt.Errorf("Error unmarshalling the response: %v", err)
	}

	return response, nil
}
