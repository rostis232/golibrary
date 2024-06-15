#base go image

FROM golang:1.22-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
    go install github.com/a-h/templ/cmd/templ@latest && \
    ln -s /go/bin/linux_amd64/migrate /usr/local/bin/migrate

RUN CGO_ENABLED=0 go build -o library ./cmd/web

RUN chmod +x ./library

#build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app /app
#add go tools binares
COPY --from=builder /go/bin/migrate /app/migrate
COPY --from=builder /go/bin/templ /app/templ

RUN chmod +x /app/entrypoint.sh

WORKDIR /app

CMD ["/app/library"]