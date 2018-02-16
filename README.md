# AlphaBets
[![Build Status](https://travis-ci.org/tsauvajon/alphabets.svg?branch=dev)](https://travis-ci.org/tsauvajon/alphabets)

![logo](https://github.com/tsauvajon/alphabets/blob/dev/logo-mini.png?raw=true)

Chainblock applied to betting.

Alpha for innovation - Bets for football betting => AlphaBets

AlphaBets was created during the [IBM Blockchain Hackaton](https://www.openmainframeproject.org/blog/2018/01/23/blockchain-hackathon-unleashing-creativity-linuxone-scale) and got the 1st place out of 80 teams.
The challenge was to create an application using the Hyperledger Fabric technology in 4 days.

It is a decentralized football betting application, using the blockchain technology.

We built a [Hyperledger Composer](https://github.com/hyperledger/composer) REST API, and used a football API (see below).
We used [Vue.js](https://github.com/vuejs/vue) for the front-end and [Golang](https://github.com/golang/go) for the back-end.

[Here](https://docs.google.com/presentation/d/e/2PACX-1vQ0tF8cIwfmjP1pICG4sKH8bDuxNF0NEXaei5wdC0zo7vcgcj8yG_dVImsQ1mCMBn_obYaRCbtdE7di/pub?start=true&loop=false&delayms=60000)'s an in-depth Google Slides presentation of the project, architecture, technological choices and UI.

![Screenshot of the UI: Fixtures](https://i.imgur.com/oUQrhQM.png)

## SportMonks

Special thanks to [SportMonks](https://sportmonks.com) for giving us access to the **Ligue 1**.
![SportMonks logo](https://www.sportmonks.com/images/logos/logo_black_top.png)

## Getting started

``` bash
# Front-end
cd client

# Get the dependencies
yarn

# Build the client
yarn build

# Run the Client with Hot-Reloading
yarn dev

# Back-end
cd ../server

# Get the dependencies
dep ensure

# Build the server
go build

# If your port 3333 is already used, set another
# Unix / MacOS
PORT=...
# Windows
SET PORT=...
```

## Setup the SportMonks API Token
Get an API Token on https://www.sportmonks.com/sports/soccer/documentation

Either set the environment variable `API_TOKEN` before running the server,
or create a `secret.json` file under the `server/sportmonks` directory:

``` json
{
    "token": "your_token"
}
```

## Run the server
```
# Unix / MacOS
./server
# Windows
server.exe
```

Browse http://localhost:{PORT}

## CI

**Node**
- lint avec ESLint

- TODO : unit testing
- tests unitaires avec Karma / Mocha / Sinon
- build

**Go**
- lint : go lint
- auto format : go fmt
- build

**Travis**

Travis automatically runs tests and builds evrey push and pull request to every branch.

## Deployment

**Local deployment with Docker**

Requires Docker to be installed and running.

``` bash
docker build -t alphabets .
docker run --publish 3333:3333 --name alphabets --rm alphabets
```
