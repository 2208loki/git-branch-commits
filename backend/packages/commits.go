package apicalls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Commit struct {
	SHA    string `json:"sha"`
	Commit struct {
		Message string `json:"message"`
		Author  struct {
			Name string `json:"name"`
			Date string `json:"date"`
		} `json:"author"`
	} `json:"commit"`
}

func FetchCommits(username, repoName string) ([]Commit, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", username, repoName)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch commits: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var commits []Commit
	if err := json.Unmarshal(body, &commits); err != nil {
		return nil, err
	}

	return commits, nil
}
