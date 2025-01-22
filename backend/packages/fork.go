package apicalls

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RepoMetadata struct {
	Fork bool `json:"fork"`
}

func IsFork(username, repoName string) (bool, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", username, repoName)

	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch repository metadata: %s", resp.Status)
	}

	var metadata RepoMetadata
	if err := json.NewDecoder(resp.Body).Decode(&metadata); err != nil {
		return false, err
	}

	return metadata.Fork, nil
}
