FROM alpine:latest AS webapp
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=gobuilder /go/src/github.com/alexellis/href-counter/app .
CMD ["./app"]

FROM golang:1.7.3 AS gobuilder
WORKDIR /go/src/github.com/alexellis/href-counter/
RUN go get -d -v golang.org/x/net/html  
COPY app.go    .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

