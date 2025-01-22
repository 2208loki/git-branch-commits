package apicalls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Repository represents a GitHub repository.
type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	HTMLURL     string `json:"html_url"`
	IsFork      bool   `json:"is_fork"` // get this from fork.go package
}

// FetchRepositories fetches repositories for a given GitHub username.
func FetchRepositories(username string) ([]Repository, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch repositories: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var repos []Repository
	if err := json.Unmarshal(body, &repos); err != nil {
		return nil, err
	}

	return repos, nil
}
