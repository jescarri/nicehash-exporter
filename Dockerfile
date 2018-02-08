FROM alpine:3.5
RUN apk update && \
    apk add dumb-init && \
    rm -rf /var/cache/apk/*

ADD server /bin/server
ENTRYPOINT ["dumb-init","/bin/server"]
