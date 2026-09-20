FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/web ./cmd/web

FROM gcr.io/distroless/static-debian12 AS build-release-stage
WORKDIR /app
COPY --from=build /out/web /app/web
EXPOSE 4000
USER nonroot:nonroot
ENTRYPOINT ["/app/web"]

