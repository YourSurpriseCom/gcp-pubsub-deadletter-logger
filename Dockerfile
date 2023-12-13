FROM alpine:latest AS build

RUN apk add --update --no-cache ca-certificates


FROM scratch

COPY --from=build  /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY  ./gcp-pubsub-deadletter-logger /gcp-pubsub-deadletter-logger

ENTRYPOINT [ "/gcp-pubsub-deadletter-logger" ]