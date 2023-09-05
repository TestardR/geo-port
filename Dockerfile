FROM golang:1.21-alpine3.18

WORKDIR /app

COPY ./bin/geo-port .

# Image should be run as a non-root user for improved security
# Principle of least priviledge applies to ensure limiting access to system resources
# https://snyk.io/blog/containerizing-go-applications-with-docker/ & https://www.baeldung.com/linux/docker-alpine-add-user
# To refine
RUN adduser -D gp-user
USER gp-user

ENTRYPOINT ["./geo-port"]