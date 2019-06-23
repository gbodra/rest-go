FROM golang:1.8

WORKDIR /go/src/github.com/gbodra/rest-go
COPY . .

RUN go get -v github.com/gorilla/mux
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build github.com/gbodra/rest-go" -command="./rest-go"