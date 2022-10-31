FROM alpine:3.16.2

COPY ./sqlModel /sqlModel
COPY ./dist/chat /usr/bin
COPY ./secrets/rsa /

CMD ["/usr/bin/chat" ]