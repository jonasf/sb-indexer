FROM centurylink/ca-certs

WORKDIR /app

COPY sb-indexer /app/

EXPOSE 8080

ENTRYPOINT ["./sb-indexer"]