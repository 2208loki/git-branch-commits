package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	apicalls "gitcommits/packages"
)

/*func reposHandler(w http.ResponseWriter, r *http.Request) {
	// Get username from query parameters
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "username query parameter is required", http.StatusBadRequest)
		return
	}

	// Fetch repositories
	repositories, err := apicalls.FetchRepositories(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repositories)
}

func handleFetchRepositories(username string) error {
	repositories, err := apicalls.FetchRepositories(username)
	if err != nil {
		return fmt.Errorf("error fetching repositories: %v", err)
	}

	fmt.Printf("Repositories for user %s:\n", username)
	for _, repo := range repositories {
		fmt.Printf("- %s: %s (%s)\n", repo.Name, repo.Description, repo.HTMLURL)
	}

	return nil
}

func handleFetchBranches(username string, repo string) error {
	branches, err := apicalls.FetchBranches(username, repo)
	if err != nil {
		return fmt.Errorf("error fetching branches: %v", err)
	}

	fmt.Printf("Branches for repositories %s:\n", repo)
	for _, repo_branch := range branches {
		fmt.Printf("- %s: %s\n", repo_branch.Name, repo_branch.Sha)
	}

	return nil
}
*/

func handleRepositories(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "username query parameter is required", http.StatusBadRequest)
		return
	}

	repositories, err := apicalls.FetchRepositories(username)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching repositories: %v", err), http.StatusInternalServerError)
		return
	}

	for i, repo := range repositories {
		isFork, err := apicalls.IsFork(username, repo.Name)
		if err != nil {
			log.Printf("error checking fork status for %s/%s: %v", username, repo.Name, err)
			repositories[i].IsFork = false // Default to false if there's an error
		} else {
			repositories[i].IsFork = isFork
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repositories)
}

func handleCommits(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	repoName := r.URL.Query().Get("repo")
	if username == "" || repoName == "" {
		http.Error(w, "username and repo query parameters are required", http.StatusBadRequest)
		return
	}

	commitList, err := apicalls.FetchCommits(username, repoName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching commits: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commitList)
}

func main() {
	// Serve static files (e.g., index.html) from the current directory
	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs)

	// API endpoints
	http.HandleFunc("/repos", handleRepositories)
	http.HandleFunc("/commits", handleCommits)

	// Start the server on port 8080
	port := "8080"
	fmt.Printf("Server running at http://localhost:%s/\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}
