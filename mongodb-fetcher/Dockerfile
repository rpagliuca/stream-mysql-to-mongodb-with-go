FROM golang
WORKDIR /var/
COPY package_cache/ package_cache/
RUN ls -ahltr
RUN ls -ahltr
RUN cd package_cache && go mod download
RUN go get github.com/codegangsta/gin
RUN apt-get update
RUN apt-get install mysql-client -y
COPY app/ app/
WORKDIR /var/app/
RUN go build .
ENTRYPOINT ["/var/app/app"]
