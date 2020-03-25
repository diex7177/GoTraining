FROM golang:latest AS build
WORKDIR /go/src/app
COPY . .
# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

FROM scratch
COPY --from=build /go/bin/app/HelloWorld /go/bin/app/HelloWorld
ENTRYPOINT ["/go/bin/app/HelloWorld"]