package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func createUser(url, apiKey string, data User) (User, error) {
	// Marshal the user struct into JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	// Create a POST request with the JSON body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}

	// Set required headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	// Execute the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	// Decode the JSON response body into a User
	var createdUser User
	if err := json.NewDecoder(res.Body).Decode(&createdUser); err != nil {
		return User{}, err
	}

	return createdUser, nil
}
