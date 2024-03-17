# syntax=docker/dockerfile:1

FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# Build
RUN cd cmd/ && CGO_ENABLED=0 GOOS=linux go build -o cli

#EXPOSE 8080

# Run
CMD ["cmd/cli"]