FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN apk update && apk add curl git && \
    curl -sL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/local/bin && \
    chmod +x /usr/local/bin/air

CMD ["/golang-fiber_app_1"]
