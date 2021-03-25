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
* Testing a package from `main` turned out to cause problems you have to run `go test *.go` to run the tests.

## Docker
* Create a Dockerfile
* Had issues with docker no recreated my image. Make sure the `CMD` is `./` to my executable.
* Create `build.sh`

## AWS
* Create an account! That was pretty easy but needed a credit card.
* Created an alert for spending as I once spend £700 for 4 days on Azure _(they refunded don't panic)_.
* I installed the AWS CLI again.
* Login to ECR to deploy my image
* Create a Cluster
* Create a Task
* Run Task
* Update VPC

## Access
Go Hackday [Visit](http://ec2-54-155-236-221.eu-west-1.compute.amazonaws.com/api/devs)

## CI/CD
* Creating a VPC and mack sure you have private and public subnets.
* You need an IGW for public and a NAT for private. The nat must point at the public.
* You can then create a RT for public and private traffic.
* You can now create a CodePipeline and CodeBuild.
