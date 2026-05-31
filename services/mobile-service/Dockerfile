FROM golang:1.23-alpine AS build
WORKDIR /src
COPY shared ./shared
COPY services/mobile-service ./services/mobile-service
WORKDIR /src/services/mobile-service
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o /out/service ./cmd

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/service /service
EXPOSE 8080
ENTRYPOINT ["/service"]
