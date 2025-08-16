package main

import (
	"fmt"
	"net/http"
)

// deleteUser deletes a user by ID
func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id

	// Create a DELETE request
	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}

	// Set the API key header
	req.Header.Set("X-API-Key", apiKey)

	// Make the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Check the status code
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("failed to delete user, status code: %d", res.StatusCode)
	}

	return nil
}
