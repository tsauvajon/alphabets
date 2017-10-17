package sportmonks

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Response : SportMonks soccer API response
type Response struct {
	// Data : string
	Data map[string]interface{} `json:"data"`
	// Number
	Meta map[string]interface{} `json:"meta"`
}

type ArrayResponse struct {
	Data []map[string]interface{} `json:"data"`
	// Number
	Meta map[string]interface{} `json:"meta"`
}

var (
	apiURI      = "https://soccer.sportmonks.com/api/v2.0/"
	suffix      = "?api_token="
	token       string
	initialized bool
)

// Initialize : reads the token from the env var or the config file
func Initialize() error {
	tk := os.Getenv("API_TOKEN")

	// If the API_TOKEN environment variable is set
	if tk != "" {
		token = tk
		initialized = true
		fmt.Println("Read the SportMonks token from the environment")
		return nil
	}

	config, err := getConfig()

	if err != nil {
		return err
	}

	if config.Token == "" {
		return errors.New("Can't find the token in secret.json")
	}

	token = config.Token
	initialized = true
	fmt.Println("Read the SportMonks token from secret.json")

	return nil
}

func getAnything(endpointWithParameters string) (Response, error) {
	if !initialized {
		return Response{}, errors.New("Error: the SportMonks token isn't initialized")
	}

	client := &http.Client{}

	url := apiURI + endpointWithParameters + suffix + token

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

func getAnythingArray(endpointWithParameters string) (ArrayResponse, error) {
	if !initialized {
		return ArrayResponse{}, errors.New("Error: the SportMonks token isn't initialized")
	}

	client := &http.Client{}

	url := apiURI + endpointWithParameters + suffix + token

	// Prepare the request
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return ArrayResponse{}, fmt.Errorf("Error creating the request: %v", err)
	}

	// Execute the request
	res, err := client.Do(req)

	if err != nil {
		return ArrayResponse{}, fmt.Errorf("Error executing the request: %v", err)
	}

	defer res.Body.Close()

	var response ArrayResponse

	// Read the response
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return ArrayResponse{}, fmt.Errorf("Error reading the response body: %v", err)
	}

	// Unmarshal the response : { data: { ... }, meta: { ... }}
	if err = json.Unmarshal(body, &response); err != nil {
		return ArrayResponse{}, fmt.Errorf("Error unmarshalling the response: %v", err)
	}

	return response, nil
}
