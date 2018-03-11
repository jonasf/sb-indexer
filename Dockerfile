# Build stage
FROM golang:1.10 AS build-env
ADD . $GOPATH/src/build
WORKDIR $GOPATH/src/build

## Install dependencies
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -vendor-only

RUN CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o sb-indexer

# Package stage
FROM centurylink/ca-certs

WORKDIR /app

# NOTE: hard coded $GOPATH
COPY --from=build-env /go/src/build/sb-indexer /app/

ENTRYPOINT ["./sb-indexer"]