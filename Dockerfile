FROM golang:alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go test -v ./... &&  go build -v -ldflags "-s -w" -o /usr/local/bin/qr .

FROM alpine

COPY --from=0 /usr/local/bin/qr /usr/local/bin/

ENTRYPOINT ["qr"]
CMD ["--help"]
