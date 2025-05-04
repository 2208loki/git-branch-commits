package apicalls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Branches represents a GitHub Branches.
type Branches struct {
	Name string `json:"name"`
	Sha  string `json:"sha"`
}

// FetchBranches fetches repositories for a given GitHub username.
func FetchBranches_temp(username string, user_repo string) ([]Branches, error) {
	//https://api.github.com/repos/2208loki/jsoncons/branches
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/branches", username, user_repo)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return nil, fmt.Errorf("\nfailed to fetch branches %s \n %s", resp.Status, resp.Body)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var repo_branch []Branches
	if err := json.Unmarshal(body, &repo_branch); err != nil {
		return nil, err
	}

	return repo_branch, nil
}
