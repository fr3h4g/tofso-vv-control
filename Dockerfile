# Build source
FROM golang:1.24.0 AS build-stage

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app .

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...


# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /app

COPY --from=build-stage /usr/share/zoneinfo/Europe/Stockholm /etc/localtime

COPY --from=build-stage /app /app

USER root:root

ENTRYPOINT ["/app/app"]
