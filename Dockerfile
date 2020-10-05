FROM golang:1.15.2-alpine3.12 as builder

LABEL maintainer="jearzamendia@gmail.com"

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/jarzamendia/dockerreports

RUN  apk add bash git ca-certificates

COPY get.sh .

RUN bash get.sh

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download dependencies
RUN go get -d -v ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/dockerreports .

######## Start a new stage from scratch #######
FROM alpine:3.9  

ENV DOCKER_API_VERSION='1.40'

WORKDIR /root

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/bin/dockerreports .

ENTRYPOINT ["./dockerreports"]

ARG BUILD_DATE

# Labels.
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.build-date=$BUILD_DATE
LABEL org.label-schema.name="dockerreports"
LABEL org.label-schema.description="Dockerreports, a Docker Swarm CPU, RAM, ENV reporter."
LABEL org.label-schema.url="https://golang.org/"
LABEL org.label-schema.vendor="Jarza"
LABEL org.label-schema.version="1"
LABEL org.label-schema.docker.cmd="docker run -it -v /var/run/docker.sock:/var/run/docker.sock jarzamendia/dockerreports:1.0.0"