# Base image for building the go project
FROM golang:1.16.3-alpine AS build

# Updates the repository and installs git
RUN apk update && apk upgrade && \
    apk add --no-cache git


# Switches to /tmp/app as the working directory, similar to 'cd'
WORKDIR /tmp/app


## If you have a go.mod and go.sum file in your project, uncomment lines 13, 14, 15

# COPY go.mod .
# COPY go.sum .
# RUN go mod download

COPY .. .

# Builds the current project to a binary file called api
# The location of the binary file is /tmp/app/out/api
RUN go mod init biges



