######################
# dependencies layer #
######################
FROM golang:1.26.2-alpine AS deps

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download


#############
# dev layer #
#############
FROM deps AS dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
COPY . .

CMD ["air"]

###############
# build layer #
###############
FROM golang:1.26.2-alpine AS builder

WORKDIR /app

COPY --from=deps /go/pkg /go/pkg
COPY . .

# choose binary: api or grpc
ARG APP=api

RUN CGO_ENABLED=0 GOOS=linux go build -o $APP ./cmd/$APP


#################
# runtime layer #
#################
FROM alpine:3.22

WORKDIR /app

RUN addgroup -S app && adduser -S -G app app

ARG APP=api

COPY --from=builder --chown=app:app /app/$APP ./server

USER app

CMD ["./server"]
