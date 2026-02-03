# Git Branch Commits

This deployment will be helpful for getting latest commit links of different branches. esp. for viewing commit links (commited to others repos through new branchs). This might not be helpful to get links of own repo or commits done to same branch/master.

## Quick Demo
![Recording 2026-02-03 145715](https://github.com/user-attachments/assets/a7c6e73b-a24c-41ef-8332-d6f7e7a303c1)


As a prerequestie just have docker in your environment. 
Then get image as follows and run easily :)
```bash
curl -O https://raw.githubusercontent.com/2208loki/git-branch-commits/main/git-commits-app.tar
docker load -i git-commits-app.tar
docker run -p 8080:8080 git-commits-app
```

Now deployment should be accessible through this simple website with link  `http://localhost:8080/` in your local machine.
