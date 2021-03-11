# go-hackday

## Getting started
* I didn't have go. So went and installed it.
* Opened InteliJ and started a new project.
* Followed the getting started and that made me `go mod init main` not sure really. Sounded like a package manager.
* Printed something to the console.

## Data
* Lets try and load some JSON.
* Loaded some with `os.Open()`
* Chris helped my decode that with `json.NewDecoder().Decode()`
* Tried to make a reusable function. I thought about generics, but they do not exist. Dave says I need to not be so generic.

## Web server
* Create a get request
* Refactor the file names to be more goey `main.go` etc.
* Lets test this endpoint. `main_test.go` - Went to [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server) and read what was there. This was way to fancy for me.
