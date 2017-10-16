# AlphaBets
[![Build Status](https://travis-ci.org/tsauvajon/ws-blockchain.svg?branch=dev)](https://travis-ci.org/tsauvajon/ws-blockchain)

![logo](https://github.com/tsauvajon/ws-blockchain/blob/dev/logo-mini.png?raw=true)

Chainblock applied to betting.

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

# Set the environment variables
# Unix / MacOS
PORT=...
# Windows
SET PORT=...

# Run the server
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
