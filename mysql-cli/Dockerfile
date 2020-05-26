FROM debian
RUN apt-get update
RUN apt-get install mysql-client -y
COPY ./entrypoint.sh /entrypoint.sh
COPY ./fixtures.sql /fixtures.sql
ENTRYPOINT ["sh", "/entrypoint.sh"]
