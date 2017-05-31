FROM alpine:3.4

MAINTAINER Nick Landolfi <nclandolfi@google.com>

RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY gub /gub

ENTRYPOINT ["/gub"]
