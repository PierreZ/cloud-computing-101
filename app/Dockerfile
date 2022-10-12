##
## Build
##
FROM golang:1.19-buster AS build

WORKDIR /app

COPY . ./

RUN go build -o /my-app

##
## Deploy
##
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /my-app /my-app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/my-app"]