############################
# STEP 1 build executable binary
############################
#
# Start from the latest golang base image
FROM golang:alpine as builder

# Add Maintainer Info
LABEL maintainer="Waldir Borba Junior <wborbajr@gmail.com>"

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

# Set the Current Working Directory inside the container
WORKDIR /mobapi

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. 
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo -v \
    . 

############################
# STEP 2 build a small image
############################
#
FROM alpine:latest

# Add non root user and certs
#RUN apk --no-cache add ca-certificates  

# Import from builder. 
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

WORKDIR /mobapi

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /main mobapiserver

# Use an unprivileged user.
USER appuser:appuser

# Setting timezone
ENV TZ=America/Sao_Paulo

# Command to run the executable
CMD ["./mobapiserver"]
