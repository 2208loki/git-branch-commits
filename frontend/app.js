const usernameInput = document.getElementById('username');
const searchButton = document.getElementById('searchButton');
const resultsDiv = document.getElementById('results');

// Function to fetch and display repositories
async function fetchRepositories() {
    const username = usernameInput.value.trim();

    if (!username) {
        resultsDiv.innerHTML = 'Please enter a GitHub username.';
        return;
    }

    resultsDiv.innerHTML = 'Loading...';

    try {
        const response = await fetch(`http://localhost:8080/repos?username=${username}`);
        if (!response.ok) {
            throw new Error(`Error: ${response.statusText}`);
        }

        const repos = await response.json();

        if (repos.length === 0) {
            resultsDiv.innerHTML = `<p>No repositories found for user: ${username}</p>`;
        } else {
            resultsDiv.innerHTML = `<h2>Repositories for ${username}:</h2>`;
            const repoList = document.createElement('ul');

            repos.forEach(repo => {
                const listItem = document.createElement('li');
                const forkStatus = repo.is_fork ? "(Fork)" : "(Original)";
                listItem.innerHTML = `
                    <strong>${repo.name}</strong>: ${repo.description || 'No description'} ${forkStatus}
                    (<a href="${repo.html_url}" target="_blank">View</a>) 
                    <button onclick="fetchCommits('${username}', '${repo.name}')">View Commits</button>
                    <div class="commit-list" id="commits-${repo.name}"></div>
                `;
                repoList.appendChild(listItem);
            });

            resultsDiv.appendChild(repoList);
        }
    } catch (error) {
        resultsDiv.innerHTML = `Error: ${error.message}`;
    }
}

// Function to fetch and display commits for a repository
async function fetchCommits(username, repoName) {
    const commitsDiv = document.getElementById(`commits-${repoName}`);
    commitsDiv.innerHTML = 'Loading commits...';

    try {
        const response = await fetch(`http://localhost:8080/commits?username=${username}&repo=${repoName}`);
        if (!response.ok) {
            throw new Error(`Error: ${response.statusText}`);
        }

        const commits = await response.json();       
        if (commits.length === 0) {
            commitsDiv.innerHTML = '<p>No commits found for this repository.</p>';
        } else {
            const commitList = document.createElement('ul');
            commits.forEach(commit => {
                const commitItem = document.createElement('li');
                commitItem.innerHTML = `
                <a href="https://github.com/${username}/${repoName}/commit/${commit.sha}" target="_blank">
                    <strong>${commit.sha}</strong>
                </a>: ${commit.commit.message} 
                by ${commit.commit.author.name} on ${commit.commit.author.date}
            `;
                commitList.appendChild(commitItem);
            });

            commitsDiv.innerHTML = '<h3>Commits:</h3>';
            commitsDiv.appendChild(commitList);
        }
    } catch (error) {
        commitsDiv.innerHTML = `Error: ${error.message}`;
    }
}

// Add event listener for the search button
searchButton.addEventListener('click', fetchRepositories);

// Add event listener for the Enter key
usernameInput.addEventListener('keypress', function (event) {
    if (event.key === 'Enter') {
        fetchRepositories();
    }
});