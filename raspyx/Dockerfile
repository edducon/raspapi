FROM golang:1.25.3-alpine3.21 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . .


RUN go build -mod vendor -o raspyx2 ./cmd/main.go

FROM alpine:3.22.0

RUN apk update && \
    apk add --no-cache curl

RUN rm -rf /var/cache/apk/*


WORKDIR /raspyx2

COPY --from=builder /app/raspyx2 /raspyx2/raspyx2
COPY --from=builder /app/.env /raspyx2/.env
COPY --from=builder /app/docs /raspyx2/docs

CMD [ "./raspyx2" ]
