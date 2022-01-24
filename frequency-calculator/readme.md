# Frequency Calculator

## Description
Server that accepts input as text and provides Json Output as Top ten most used words and times of occurrence in the text.

## Installation
If you dont have a docker, docker-compose setup - 

Install docker and docker-compose

Clone this project 

```bash
cd frequency-calculator/
```

## Usage
Create a .env file in app frequency calculator folder and set the PORT environment variable.

```bash
cp .env.example .env
```

Start the server.

```bash
go run main.go or sudo docker-compose up
```

## REST API ENDPOINT

Get occurrences of top 10 words in text.

```
Request :

POST http://localhost:5000/api/v1/wordsCount
Body (Content-Type application/json):
{
	"input": "Lorem ipsum dolor sit amet consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum"
}
```
```
Response :
[{"Word":"in","Frequency":3},{"Word":"dolor","Frequency":2},{"Word":"dolore","Frequency":2},{"Word":"ut","Frequency":2},{"Word":"non","Frequency":1},{"Word":"reprehenderit","Frequency":1},{"Word":"Duis","Frequency":1},{"Word":"nisi","Frequency":1},{"Word":"irure","Frequency":1},{"Word":"commodo","Frequency":1}]
```

## Licence
[![MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)