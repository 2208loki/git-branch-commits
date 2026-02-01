# Git Branch Commits

This deployment will be helpful for viewing latest commit details of different branches. esp. for viewing commit links (commited to others repo through new branchs). This might not be helpful to check details of own repo with commits done to same branch/master.

As a prerequestie just have docker in your environment. 
Then get image as follows and run easily :)
```bash
curl -O https://raw.githubusercontent.com/2208loki/git-branch-commits/main/git-commits-app.tar
docker load -i git-commits-app.tar
docker run -p 8080:8080 git-commits-app
```

Now deployment should be now accessible through this simple website with link  `http://localhost:8080/` in your local machine.
