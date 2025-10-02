package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func New(url string) GithubService {
	return &GithubClient{url: url}
}

func (c *GithubClient) GetGists(username string) ([]Gist, error) {
	// Construct the URL to fetch gists for the given username
	addr := fmt.Sprintf(c.url, username)

	// Make the HTTP GET request and handle errors
	resp, err := http.Get(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch gists: %s", resp.Status)
	}

	// Ensure the response body is closed after reading
	defer resp.Body.Close()

	// Check if the response status is OK (200)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch gists: %s", resp.Status)
	}

	// Read the response body
	jsondata, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a slice of Gist structs and handle errors
	data, err := unMarshalData(fmt.Sprintf("%s", jsondata))
	if err != nil {
		return nil, err
	}

	// Return the slice of Gist structs
	return data, nil
}

func (c *GithubClient) Serve(port int) {
	http.HandleFunc("/{user}", func(w http.ResponseWriter, r *http.Request) {
		user := r.URL.Path[len("/"):]
		if user == "" {
			http.Error(w, "User not specified", http.StatusBadRequest)
			return
		}

		gists, err := c.GetGists(user)
		if err != nil {
			http.Error(w, "Error fetching gists", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gists)
	})
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func unMarshalData(jsondata string) ([]Gist, error) {
	// This is a placeholder function to avoid "no non-test Go files" error.
	var data []Gist

	if err := json.Unmarshal([]byte(jsondata), &data); err != nil {
		return nil, err
	}
	return data, nil
}
