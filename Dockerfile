### Client
# Latest LTS
FROM node as client

# Copy the local package files to the container's workspace.
COPY . /usr/src/app
WORKDIR /usr/src/app/client

# Build the client
RUN yarn
RUN yarn build

### Server
# Debian image with the latest version of Go installed
FROM golang as server

# Copy the local package files to the container's workspace.
WORKDIR /go/src/github.com/tsauvajon/ws-blockchain/server
COPY --from=client /usr/src/app/server .

# Build the server for alpine, executable name: alphabets
RUN go get
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o alphabets

### Small container to execute the server
# Small container to run the app
# FROM alpine:3.6
FROM octoblu/alpine-ca-certificates

# Add root ca certificates
# RUN echo "http://dl-cdn.alpinelinux.org/alpine/v3.6/main/" >> /etc/apk/repositories \
#     && echo "http://dl-1.alpinelinux.org/alpine/v3.6/main/" >> /etc/apk/repositories \
#     && echo "http://dl-2.alpinelinux.org/alpine/v3.6/main/" >> /etc/apk/repositories \
#     && echo "http://dl-3.alpinelinux.org/alpine/v3.6/main/" >> /etc/apk/repositories \
#     && echo "http://dl-4.alpinelinux.org/alpine/v3.6/main/" >> /etc/apk/repositories \
#     && echo "http://dl-5.alpinelinux.org/alpine/v3.6/main/" >> /etc/apk/repositories \
#     && apk add --update --no-cache ca-certificates

WORKDIR /root/

COPY --from=server /go/src/github.com/tsauvajon/ws-blockchain/server/alphabets .
COPY --from=server /go/src/github.com/tsauvajon/ws-blockchain/server/dist ./dist/
COPY --from=server /go/src/github.com/tsauvajon/ws-blockchain/server/sportmonks/secret.json ./sportmonks/

# Run the server when the container starts
CMD ./alphabets

# Listen to 3333
EXPOSE 3333