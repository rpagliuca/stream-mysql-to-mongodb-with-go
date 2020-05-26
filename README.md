# About

Using the library go-mysql, I've setup a proof of concept environment where changes to a MySQL table are reflected into a MongoDB collection

# Usage

1. Start servers:
`docker-compose up`

2. List records from MongoDB collection
`docker-compose run mongodb-fetcher`

3. Insert records into MySQL table
`docker-compose run mysql-cli 'INSERT INTO mytable(value) VALUES ("Record1"), ("Record2"), ("Another record")`
`docker-compose run mysql-cli 'UPDATE mytable SET value="RecordTwo" WHERE value="Record2"'
`docker-compose run mysql-cli 'DELETE FROM mytable WHERE value="Record1"'

4. List records from MongoDB collection
`docker-compose run mongodb-fetcher`
