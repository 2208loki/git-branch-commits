package apicalls

type GitHubError struct {
	Message          string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
}
