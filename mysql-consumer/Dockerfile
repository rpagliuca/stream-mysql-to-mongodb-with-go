FROM golang
WORKDIR /var/
RUN apt-get update
RUN apt-get install mysql-client -y
COPY app/ app/
WORKDIR /var/app/
RUN go build 
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["sh", "/entrypoint.sh"]
