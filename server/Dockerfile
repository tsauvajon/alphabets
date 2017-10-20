# Debian image with the latest version of Go installed
FROM golang

# Copy the local package files to the container's workspace.
COPY . /go/src/github.com/tsauvajon/ws-blockchain/server
WORKDIR /go/src/github.com/tsauvajon/ws-blockchain/server

# Build the server
RUN go get
RUN go build

# Run the "./server" command when the container starts
ENTRYPOINT /go/src/github.com/tsauvajon/ws-blockchain/server/server

# Listen to 3333
EXPOSE 3333