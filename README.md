# Systembolaget article indexer

[![Build Status](https://travis-ci.org/jonasf/systembolaget-article-indexer.svg?branch=master)](https://travis-ci.org/jonasf/systembolaget-article-indexer)

This application will download the article XML file from [Systembolaget API](https://www.systembolaget.se/api) and index that data into Elasticsearch.

## Getting started

Simplest way is use [Docker](https://www.docker.com/) and [Docker-compose](https://github.com/docker/compose)

1. Download the docker-compose.yml
2. Run `docker-compose up`

## Building

1. Make sure [Go](https://golang.org/) is installed and set up properly
2. Install [dep](https://github.com/golang/dep)
3. Clone the repo
4. Run `make all`