FROM golang:1.18

# Bad: No multi-stage build, keeping the entire Go toolchain
WORKDIR /go/src/app
COPY . .

# Bad: Fetching dependencies in a separate layer without caching optimization
RUN go get -d -v ./...
RUN go install -v ./...

# Bad: Final image is unnecessarily large
CMD ["app"]
