FROM golang:latest
EXPOSE 8080
COPY main.go /go/src/main.go
RUN go build -o ~/start /go/src/main.go
CMD ~/start
