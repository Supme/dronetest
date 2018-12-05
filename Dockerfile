FROM golang:latest
COPY main.go /go/src/main.go
RUN go build -o ~/start /go/src/main.go
CMD ~/start