# GolangCodeSnippets
GolangCodeSnippets is a collection of go codes which are used for learning purposes

# Check for test Coverage in  Golang

```bash

#install cover package using go get golang.org/x/tools/cmd/cover

# add coverprofile flag to check for test coverage and thus store in cover.out file

go test -coverprofile cover.out


#This would then open cover.out in a web browser

go tool cover -html=cover.out

```

