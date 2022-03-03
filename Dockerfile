# Building binary.
FROM golang:1.17 AS binary

RUN apt-get update && apt-get upgrade -yq

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download -x

COPY . .

RUN make build_all BIN_DIR=/build

# Building container to run binary.
FROM debian:bullseye

RUN apt-get update && apt-get install -y --no-install-recommends

WORKDIR /app

COPY --from=binary /build/ /app/

EXPOSE 9000
