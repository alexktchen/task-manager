FROM golang:latest


WORKDIR $GOPATH/src/github.com/alexktchen/task-manager
COPY . $GOPATH/src/github.com/alexktchen/task-manager
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./task-manager"]
