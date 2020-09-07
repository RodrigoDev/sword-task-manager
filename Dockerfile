FROM golang:1.14-alpine as build

LABEL Maintainer="Rodrigo Carneiro <rodrigo_carneiro@ymail.com>"

ARG COMMIT='local'

WORKDIR /go/src/github.com/RodrigoDev/sword-task-manager/

COPY . .

RUN apk update && \
    apk --no-cache add ca-certificates && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build -mod vendor -o /app ./cmd/server

ENTRYPOINT ["/app"]

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=build /app /

ENTRYPOINT ["/app"]
