FROM alpine:edge

RUN apk add --update curl && \
    rm -rf /var/cache/apk/*

COPY realtime /opt/realtime

EXPOSE 8080
CMD ["/opt/realtime"]
