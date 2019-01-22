FROM alpine:3.7
WORKDIR /server
ADD app /server
ADD static /server/static
ADD config.yaml /server

EXPOSE 8000

CMD ["./app"]