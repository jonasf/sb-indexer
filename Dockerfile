# Build stage
FROM golang:1.10 AS build-env
ADD . $GOPATH/src/github.com/jonasf/systembolaget-article-indexer
WORKDIR $GOPATH/src/github.com/jonasf/systembolaget-article-indexer

## Install dependencies
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -vendor-only

RUN cd cmd/systembolaget-article-indexer && CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o ../../build/systembolaget-article-indexer

# Package stage
FROM centurylink/ca-certs

WORKDIR /app

# NOTE: hard coded $GOPATH
COPY --from=build-env /go/src/github.com/jonasf/systembolaget-article-indexer/build/systembolaget-article-indexer /app/

ENTRYPOINT ["./systembolaget-article-indexer"]