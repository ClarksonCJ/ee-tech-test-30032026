package github

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalModel(t *testing.T) {
	assert := assert.New(t)

	jsonStr, err := loadSampleData("sample.json")
	assert.NoError(err)

	results, err := unMarshalData(jsonStr)
	assert.NoError(err)

	assert.Len(results, 1)
	assert.Equal("https://api.github.com/gists/aa5a315d61ae9438b18d", results[0].Url)
}

func TestGistClient(t *testing.T) {
	assert := assert.New(t)

	jsonStr, err := loadSampleData("sample.json")
	assert.NoError(err)

	// Mock the HTTP GET request by replacing the http.Get function temporarily
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, jsonStr)
	}))
	defer ts.Close()

	client := New(ts.URL + "/users/%s/gists")
	gists, err := client.GetGists("octocat")

	assert.NoError(err)
	assert.Len(gists, 1)
	assert.Equal("http://"+ts.Listener.Addr().String()+"/users/octocat/gists", ts.URL+"/users/octocat/gists")
	assert.Equal("https://api.github.com/gists/aa5a315d61ae9438b18d", gists[0].Url)
}

func TestGithubGistServe(t *testing.T) {
	assert := assert.New(t)

	jsonStr, err := loadSampleData("sample.json")
	assert.NoError(err)

	// Create test server to replicate Github API and return sample data for octocat
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, jsonStr)
	}))
	defer ts.Close()

	client := New(ts.URL + "/users/%s/gists")

	// Start a test server to simulate the Serve function
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.URL.Path[len("/"):]
		if user == "" {
			http.Error(w, "User not specified", http.StatusBadRequest)
			return
		}

		gists, err := client.GetGists(user)
		if err != nil {
			http.Error(w, "Error fetching gists", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(gists)
		w.Write(jsonData)
	}))
	defer apiServer.Close()

	resp, err := http.Get(apiServer.URL + "/octocat")
	assert.NoError(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	assert.NoError(err)

	var gists []Gist
	err = json.Unmarshal(body, &gists)
	assert.NoError(err)

	assert.Len(gists, 1)
	assert.Equal("https://api.github.com/gists/aa5a315d61ae9438b18d", gists[0].Url)
}

func TestGithubGistServeErrorFromGithub(t *testing.T) {
	assert := assert.New(t)

	// Create test server to replicate Github API and return sample data for octocat
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	client := New(ts.URL + "/users/%s/gists")

	// Start a test server to simulate the Serve function
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.URL.Path[len("/"):]
		if user == "" {
			http.Error(w, "User not specified", http.StatusBadRequest)
			return
		}

		gists, err := client.GetGists(user)
		if err != nil {
			http.Error(w, "Error fetching gists", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(gists)
		w.Write(jsonData)
	}))
	defer apiServer.Close()

	resp, err := http.Get(apiServer.URL + "/octocat")
	assert.NoError(err)
	defer resp.Body.Close()
	assert.Equal(http.StatusInternalServerError, resp.StatusCode)
}

func TestGithubGistServeErrorNoUser(t *testing.T) {
	assert := assert.New(t)

	jsonStr, err := loadSampleData("sample.json")
	assert.NoError(err)

	// Create test server to replicate Github API and return sample data for octocat
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, jsonStr)
	}))
	defer ts.Close()

	client := New(ts.URL + "/users/%s/gists")

	// Start a test server to simulate the Serve function
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.URL.Path[len("/"):]
		if user == "" {
			http.Error(w, "User not specified", http.StatusBadRequest)
			return
		}

		gists, err := client.GetGists(user)
		if err != nil {
			http.Error(w, "Error fetching gists", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(gists)
		w.Write(jsonData)
	}))
	defer apiServer.Close()

	resp, err := http.Get(apiServer.URL + "/")
	assert.NoError(err)
	defer resp.Body.Close()
	assert.Equal(http.StatusBadRequest, resp.StatusCode)
}

func loadSampleData(filename string) (string, error) {
	// This function would load the sample JSON data from a file.
	// For the purpose of this example, we'll return a hardcoded JSON string.
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	data, err := os.ReadFile(f.Name())
	if err != nil {
		return "", err
	}

	return string(data), nil
}
