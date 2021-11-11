FROM 694770204095.dkr.ecr.us-east-1.amazonaws.com/webapp-service/cache:latest AS gobuilder
FROM alpine:latest AS webapp
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=gobuilder /go/src/github.com/alexellis/href-counter/app .
CMD ["./app"]
