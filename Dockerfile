################
# BUILD BINARY #
################

FROM golang:1.21.5-alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.Version=v1.0.0'" .

#####################
# MAKE SMALL BINARY #
#####################
FROM scratch

# Copy the executable.
WORKDIR /app

COPY --from=builder /app/golang_agency_clean_architecture /app
COPY --from=builder /app/config.prod.yml /app