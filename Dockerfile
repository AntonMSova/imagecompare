FROM golang
WORKDIR /go/src/github.com/AntonMSova/imagecompare
COPY vendor vendor/
COPY cmd/ cmd/
COPY pkg/ pkg/

ARG VERSION
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server cmd/server/main.go
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o "/arch=amd64+os=linux" cmd/cli/main.go
RUN env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o "/arch=amd64+os=darwin" cmd/cli/main.go
RUN env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o "/arch=amd64+os=windows" cmd/cli/main.go

FROM alpine
RUN apk add --update --no-cache ca-certificates
COPY public/ public
COPY examples/ examples/
COPY --from=0 /server /server
COPY --from=0 /arch=amd64+os=linux /arch=amd64+os=linux
COPY --from=0 /arch=amd64+os=darwin /arch=amd64+os=darwin

ENTRYPOINT ["/server"]
