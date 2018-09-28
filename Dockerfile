FROM alpine:3.8

# Use 'make build' to build bin/app
COPY bin/app /bin/app
EXPOSE 8080

ENTRYPOINT ["/bin/app"]
