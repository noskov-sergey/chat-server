FROM golang:1.20.3-alpine AS builder

COPY . /github.com/noskov-sergey/chat-server/source/
WORKDIR /github.com/noskov-sergey/chat-server/source/

RUN go mod download
RUN go build -o ./bin/chat-server cmd/main.go


FROM alpine:latest

WORKDIR /root/

ARG grpc_host=grpc_host
ARG grpc_port=grpc_port
ARG pg_dsn=pg_dsn
ENV GRPC_HOST=$grpc_host
ENV GRPC_PORT=$grpc_port
ENV PG_DSN=$pg_dsn

RUN touch .env
RUN echo GRPC_HOST=$GRPC_HOST >> .env
RUN echo GRPC_PORT=$GRPC_PORT >> .env
RUN echo PG_DSN=$PG_DSN >> .env

COPY --from=builder /github.com/noskov-sergey/chat-server/source/bin/chat-server .

CMD ["./chat-server"]