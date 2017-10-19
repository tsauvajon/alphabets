package blockchain

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

// ArrayResponse : Useless
type ArrayResponse struct {
	Data []map[string]interface{} `json:"data"`
	// Number
	Meta map[string]interface{} `json:"meta"`
}

var (
	apiURI = "http://alpha-bets.world:3000/api/"
)

func getBcAnything(endpointWithParameters string) (Response, error) {
	client := &http.Client{}

	url := apiURI + endpointWithParameters

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

func getBcAnythingArray(endpointWithParameters string) (ArrayResponse, error) {
	client := &http.Client{}

	url := apiURI + endpointWithParameters
	fmt.Println(url)

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

	fmt.Println(body)

	if err != nil {
		return ArrayResponse{}, fmt.Errorf("Error reading the response body: %v", err)
	}

	// Unmarshal the response : { data: { ... }, meta: { ... }}
	if err = json.Unmarshal(body, &response); err != nil {
		return ArrayResponse{}, fmt.Errorf("Error unmarshalling the response array : %v", err)
	}

	return response, nil
}

func postBcAnything(endpointWithParameters string, strings []string) (Response, error) {
	client := &http.Client{}

	url := apiURI + endpointWithParameters

	// Prepare the request
	req, err := http.NewRequest("POST", url, nil) // strings.NewReader(r.Encode())

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
