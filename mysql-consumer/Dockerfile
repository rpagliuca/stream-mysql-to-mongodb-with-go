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
ENTRYPOINT sh -c "echo godoc at: http://$(show-ip|tr -d ' '):6060; (godoc -http 0.0.0.0:6060 > /dev/null &); echo application at: http://$(show-ip|tr -d ' '):3000; gin --appPort 8080 --immediate --bin build/gin-bin run main.go"
