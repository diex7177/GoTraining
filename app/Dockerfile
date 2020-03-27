FROM golang:latest AS build
WORKDIR /go/src/app
COPY . .
# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...