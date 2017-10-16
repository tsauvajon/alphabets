# AlphaBets
[![Build Status](https://travis-ci.org/tsauvajon/ws-blockchain.svg?branch=dev)](https://travis-ci.org/tsauvajon/ws-blockchain)

![logo](https://github.com/tsauvajon/ws-blockchain/blob/dev/logo-mini.png?raw=true)

Chainblock applied to betting.

## SportMonks

Special thanks to [SportMonks](https://sportmonks.com) for giving us access to the Ligue 1
![SportMonks logo](https://www.sportmonks.com/images/logos/logo_black_top.png)

Alpha for innovation - Bets for Betting => AlphaBets

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

Node
- lint avec ESLint
- tests unitaires avec Karma / Mocha / Sinon
- build

Go
- TODO : lint
- TODO : tests unitaires
- build

## Deployment

TODO
