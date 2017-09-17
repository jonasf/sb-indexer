# Build stage
FROM golang:1.9 AS build-env
ADD . /src

# TODO: Replace with dependency manager
RUN go get golang.org/x/net/context
RUN go get gopkg.in/olivere/elastic.v5

RUN cd /src && CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o sb-indexer

# Package stage
FROM centurylink/ca-certs

WORKDIR /app

COPY --from=build-env /src/sb-indexer /app/

ENTRYPOINT ["./sb-indexer"]