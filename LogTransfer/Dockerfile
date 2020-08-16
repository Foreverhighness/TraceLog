FROM golang:1.14-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./

RUN GO111MODULE=on GOPROXY=https://goproxy.cn go mod download

COPY . .
RUN go build

FROM alpine:latest as develop

VOLUME [ "/conf" ]

COPY --from=builder /app/logtransfer /usr/bin

CMD [ "logtransfer" ]
