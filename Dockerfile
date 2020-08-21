FROM alpine:3.7
WORKDIR /server
ADD app /server
ADD template /server/template
ADD web/build /server/web/build
ADD config.yaml /server

EXPOSE 8000

CMD ["./app"]