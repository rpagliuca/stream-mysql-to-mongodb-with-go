FROM golang
WORKDIR /var/
COPY app/ app/
WORKDIR /var/app/
RUN go build .
ENTRYPOINT ["/var/app/app"]
