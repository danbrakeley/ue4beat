FROM golang:1.11.0-alpine3.8

ARG APK_PACKAGES="git"
RUN apk update && apk add --no-cache $APK_PACKAGES

COPY . /app
WORKDIR /app

RUN go test .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ue4beat .
RUN gzip ue4beat
