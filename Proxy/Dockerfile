# syntax=docker/dockerfile:1

FROM golang:1.20 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
WORKDIR /app/cmd/proxy
RUN CGO_ENABLED=0 GOOS=linux go build -o /proxy

FROM scratch as runner
COPY --from=builder /proxy /proxy
EXPOSE 8080

CMD ["/proxy"]
