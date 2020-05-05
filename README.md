# Company API
> Company Service is responsible for company management section of job portal microservice

*   Create a company
*   Update a company
*   Delete a company
*   Read a single company by ID
*   Read all company data

## Setup
> Setup to start working on this project

### Install GoLang
[version as of writing: go version go1.12.4 darwin/amd64](https://golang.org/)

### setup `$GOPATH`
```bash
# In your bash profile
export GOPATH="/Users/<user>/<folder>/go"
export PATH=$PATH:$GOPATH/bin
```
> ### Use `gitflow` for project management
* [gitflow docs](https://www.atlassian.com/git/tutorials/comparing-workflows/)
* [gitflow cheatsheet](https://danielkummer.github.io/git-flow-cheatsheet/)

gitflow-workflow
> ### IMPORTANT! Make sure this repository is located or clone the project
```bash
# clone the project
cd cd $GOPATH/src/github.com/oojob
git clone @repo
$GOPATH/src/github.com/oojob/service-company
```

### Install protobuf
Mac: `make setup-protobuf-mac`
Linux: `make setup-protobuf-linux`
>   See: [Error](http://google.github.io/proto-lens/installing-protoc.html) if there are any failures

### Setup Go environment

#### Install go dep tool (https://github.com/golang/dep)
```bash
make setup-dep
```

Install go dependencies*

```bash
make setup-go
```
> these need to be managed outside of the vendor/ directory because they are used in proto code generation

## Development
> run the api's locally

### Build Services
```bash
make build
```


## Running company service
> company service is build as an command line application.
> After running `make build` run `./bin/server --help` to view available commands
```bash
Usage:
  company [OPTIONS] [COMMANDS] [flags]
  company [command]

Available Commands:
  help        Help about any command
  serve       serves the gRPC server
  version     Print the version number

Flags:
  -a, --author string   author name for copyright attribution (
default "nirajgeorgian")
      --config string   config file (default is config.yaml)
  -h, --help            help for company
      --viper           use Viper for configuration (default tr
ue)

Use "company [command] --help" for more information about a com
mand.

# create a company
grpc_cli call localhost:3000 CreateCompany "name: 'lorem',description: 'ipsum', no_of_employees: {min: 3, max: 4}"

# update a company
grpc_cli call localhost:30000 UpdateCompany "name: 'updated name',id: '5e85b903dd1dca85d6e1f8bd',no_of_employees: {min: 3, max: 4}"

# read a company by id
grpc_cli call localhost:3000 ReadCompany "id: '5e84efd6fbe77f66ced36613'"

# read all companies as stream data
grpc_cli call localhost:3000 ReadCompanies ""

```
