FROM alpine:3.16.2

COPY ./sqlModel /sqlModel
COPY ./dist/chat /usr/bin

CMD ["/usr/bin/chat" ]