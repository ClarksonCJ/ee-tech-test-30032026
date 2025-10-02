FROM golang:1.24.4-alpine@sha256:25f0ae8a9452540a6ffa309395ca983e199b28dae84e9611c99a7587cff38e73 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o tech-test .

FROM alpine:3.22.1@sha256:4562b419adf48c5f3c763995d6014c123b3ce1d2e0ef2613b189779caa787192

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/tech-test .

USER appuser

ENTRYPOINT ["/app/tech-test"]
EXPOSE 8080
CMD ["serve"]

