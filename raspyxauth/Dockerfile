FROM golang:1.25.3-alpine3.21 as builder

WORKDIR /app

ADD go.mod .

COPY ./ ./

RUN go build -mod vendor -o jwtauth ./cmd/main.go

FROM alpine:3.22.0

RUN apk update && \
    apk add --no-cache curl

RUN rm -rf /var/cache/apk/*

WORKDIR /jwtauth-app

COPY --from=builder /app/jwtauth /jwtauth-app/jwtauth
COPY --from=builder /app/.env /jwtauth-app/.env
COPY --from=builder /app/docs /jwtauth-app/docs

CMD [ "./jwtauth" ]
