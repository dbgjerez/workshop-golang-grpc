# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.19-buster AS build

WORKDIR /app

COPY comment/go.mod ./
COPY comment/go.sum ./
RUN go mod download

COPY comment ./


RUN go build -o /application server/main.go 

##
## Deploy
##
FROM gcr.io/distroless/base-debian11
ARG version
ARG serviceName
ARG buildTime
ENV SERVICE_VERSION=${version}
ENV SERVICE_NAME=${serviceName}
ENV SERVICE_BUILD_TIME=${buildTime}
WORKDIR /

COPY --from=build /application /application
COPY data.json /data.json

EXPOSE 8080

USER nonroot:nonroot

ENV GIN_MODE=release

ENTRYPOINT ["/application"]
CMD ["-port", "8080"]