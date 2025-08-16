package main

import "fmt"

func fetchTasks(baseURL, availability string) []Issue {
	limit := 0
	switch availability {
	case "Low":
		limit = 1
	case "Medium":
		limit = 3
	case "High":
		limit = 5
	}

	fullURL := baseURL + "?sort=estimate&limit=" + fmt.Sprintf("%d", limit)
	return getIssues(fullURL)
}
