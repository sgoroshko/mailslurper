FROM golang:alpine AS builder
LABEL maintainer="erguotou525@gmail.compute"

RUN apk --no-cache add git libc-dev gcc sqlite-dev

WORKDIR /mailslurper

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go generate cmd/mailslurper/mailslurper.go
RUN CGO_CFLAGS="-D_LARGEFILE64_SOURCE" CGO_ENABLED=1 go build -tags="sqlite_omit_load_extension" -o mailslurper cmd/main/main.go


FROM alpine:latest

WORKDIR /app
RUN apk add --no-cache ca-certificates \
 && echo -e '{\n\
  "dbEngine": "SQLite",\n\
  "dbHost": "",\n\
  "dbPort": 0,\n\
  "dbDatabase": "/app/data/mailslurper.db",\n\
  "dbUserName": "",\n\
  "dbPassword": "",\n\
  "keyFile": "",\n\
  "certFile": "",\n\
  "adminKeyFile": "",\n\
  "adminCertFile": ""\n\
}' >> config.json

COPY --from=builder /mailslurper/mailslurper /app/mailslurper

ENTRYPOINT [ "/app/mailslurper" ]
CMD server --debug
