# 

## Description
Script that reads text from a file, calls frequency-calculator service and gets top 10 word occurrences in text.

## Installation
If you dont have a Golang setup, install go and setup the environment: 

Install [Go](https://go.dev/doc/install) as given in the official documentation.

Set GOPATH and update Path environment variable.

```bash
export GOPATH=path/to/go/working/directory
export PATH=$PATH:$GOPATH/bin
```

Clone this project and install all project dependencies.

```bash
cd golang-task
cd project2
go mod download
```

## Usage
Make sure the frequency-calculator server is running. Please visit: https://github.com/prithvichouhan/golang-task/tree/main/frequency-calculator.

Create a .env file in project2 folder and set the URL in env file, Add text in the input.txt

```bash
cp .env.example .env
```

Run Script

```bash
go run main.go
```

JSON output will print onto the CLI

```
Top Ten element -  [{"Word":"in","Frequency":3},{"Word":"dolor","Frequency":2},{"Word":"dolore","Frequency":2},{"Word":"ut","Frequency":2},{"Word":"non","Frequency":1},{"Word":"reprehenderit","Frequency":1},{"Word":"Duis","Frequency":1},{"Word":"nisi","Frequency":1},{"Word":"irure","Frequency":1},{"Word":"commodo","Frequency":1}]
```

## Licence
[![MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)