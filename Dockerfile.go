FROM 694770204095.dkr.ecr.us-east-1.amazonaws.com/webapp-service:server-develop358 AS builder
FROM alpine:latest AS webapp-server
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/alexellis/href-counter/app .
CMD ["./app"]  
