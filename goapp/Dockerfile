#first stage - builder

FROM golang:1.13.7-alpine3.11 as builder
COPY . /opt/
WORKDIR /opt/
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o challengesite

#second stage
FROM alpine:latest
WORKDIR /opt/
COPY --from=builder /MultiStage .CMD ["./challengesite"]
