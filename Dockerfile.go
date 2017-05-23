FROM golang:1.7.3
WORKDIR /go/src/github.com/pchico/demo
RUN go get github.com/Sirupsen/logrus  
COPY app.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
WORKDIR /root/
COPY --from=0 /go/src/github.com/pchico/demo/app .
CMD ["./app"]
