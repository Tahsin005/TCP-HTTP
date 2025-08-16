package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// updateUser updates a user by ID and returns the updated user
func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id

	// Encode the user data as JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	// Create a PUT request with the JSON data as body
	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	// Make the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	// Decode the response JSON into a User struct
	var updatedUser User
	if err := json.NewDecoder(res.Body).Decode(&updatedUser); err != nil {
		return User{}, err
	}

	return updatedUser, nil
}

// getUserById fetches a user by ID
func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	// Create a GET request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return User{}, err
	}

	// Set the API key header
	req.Header.Set("X-API-Key", apiKey)

	// Make the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	// Decode the response JSON into a User struct
	var user User
	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		return User{}, err
	}

	return user, nil
}
