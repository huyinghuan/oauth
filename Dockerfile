FROM golang:latest
WORKDIR /go/src/oauth
COPY . /go/src/oauth
RUN go build
EXPOSE 8000
CMD [ "./oauth" ]