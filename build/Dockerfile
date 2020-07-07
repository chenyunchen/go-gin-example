FROM golang:1.14 as builder

WORKDIR /mockserver
COPY . /mockserver

ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN go mod tidy && go mod download
RUN go build -mod=readonly -v -o mockserver *.go
# RUN go build -mod=vendor -v -o order *.go

FROM alpine:3

RUN apk add --no-cache ca-certificates

RUN mkdir -p /mockserver
WORKDIR /mockserver

COPY --from=builder /mockserver/mockserver /mockserver/mockserver
COPY --from=builder /mockserver/conf /mockserver/conf

# Create appuser.
ENV USER=appuser
ENV UID=54666

RUN adduser \
--disabled-password \
--gecos "application user" \
--no-create-home \
--uid "${UID}" \
"${USER}"

RUN chown appuser:appuser /mockserver
RUN chown appuser:appuser /mockserver/*
USER appuser:appuser

ENTRYPOINT ./mockserver
