FROM alpine:latest AS webapp
RUN apk --no-cache add ca-certificates
WORKDIR /root/
CMD ["./app"]
