package apicalls

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Commit struct {
	SHA    string `json:"sha"`
	Commit struct {
		Message string `json:"message"`
		Author  struct {
			Name  string `json:"name"`
			Date  string `json:"date"`
			Login string `json:"login"`
		} `json:"author"`
	} `json:"commit"`
}

type Branch struct {
	Name   string `json:"name"`
	Commit struct {
		SHA string `json:"sha"`
	} `json:"commit"`
}

func FetchBranches(owner, repoName string) ([]Branch, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/branches", owner, repoName)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return nil, fmt.Errorf("\nfailed to fetch branches: %s \n %s", resp.Status, resp.Body)
	}

	var branches []Branch
	if err := json.NewDecoder(resp.Body).Decode(&branches); err != nil {
		return nil, err
	}

	return branches, nil
}

func FetchCommits(username, repoName string) ([]Commit, error) {
	var fullCommits []Commit
	page := 1

	for {
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits?page=%d&per_page=100", username, repoName, page)

		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {

			return nil, fmt.Errorf("\nfailed to fetch commits: %s \n %s", resp.Status, resp.Body)
		}

		var pageCommits []Commit
		if err := json.NewDecoder(resp.Body).Decode(&pageCommits); err != nil {
			return nil, err
		}

		// Break if no more commits are returned
		if len(pageCommits) == 0 {
			break
		}

		// Append filtered commits for the given username
		for _, commit := range pageCommits {
			if commit.Commit.Author.Login == username {
				fullCommits = append(fullCommits, commit)
			}
		}

		page++
	}

	return fullCommits, nil
}
