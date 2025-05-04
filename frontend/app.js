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
            const errorText = await response.text(); // Read raw error message from server
            throw new Error(errorText);
        }

        const repos = await response.json();

        if (repos.length === 0) {
            resultsDiv.innerHTML = `<p>No repositories found for user: ${username}</p>`;
        } else {
            resultsDiv.innerHTML = `<h2>Repositories for ${username}:</h2>`;
            const repoList = document.createElement('ul');

            repos.forEach(repo => {
                const listItem = document.createElement('li');
                //const forkStatus = repo.is_fork ? "(Fork)" : "(Original)";
                listItem.innerHTML = `
                    <strong>${repo.name}</strong>: ${repo.description || 'No description'}
                    (<a href="${repo.html_url}" target="_blank">View Repo</a>) 
                        <a href="commits.html?username=${username}&repo=${repo.name}" target="_blank" rel="noopener noreferrer">
                            <button>View Branches</button>
                        </a>
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


// Add event listener for the search button
searchButton.addEventListener('click', fetchRepositories);

// Add event listener for the Enter key
usernameInput.addEventListener('keypress', function (event) {
    if (event.key === 'Enter') {
        fetchRepositories();
    }
});