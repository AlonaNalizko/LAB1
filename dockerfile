FROM golang:1.24.1-alpine3.21 AS builder

ENV GOCACHE=/root/.cache/go-build

WORKDIR /app

COPY ./go.mod ./
COPY ./cmd ./cmd


RUN go mod tidy


RUN --mount=type=cache,target="/root/.cache/go-build" \
   go build -o app ./cmd/main.go


FROM alpine:3.21

WORKDIR /myapp

COPY --from=builder /app ./

EXPOSE 6080

CMD ["/myapp/app"]