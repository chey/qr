FROM golang:alpine

ARG VERSION=dev

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go test -v ./... &&  go build -v -ldflags "-s -w -X $(go list -m)/code.version=${VERSION/v/}" -o /usr/local/bin/qr .

FROM scratch

COPY --from=0 /usr/local/bin/qr /usr/local/bin/

ENTRYPOINT ["qr"]
CMD ["--help"]
