# About

Using the library go-mysql, I've setup a proof of concept environment where changes to a MySQL table are reflected into a MongoDB collection

# Usage

1. Start servers:
`docker-compose up`

2. List records from MongoDB collection (should return 0 records)
`docker-compose run mongodb-fetcher`

3. Create MySQL table and insert 3 records
`docker-compose run mysql-cli init`

4. List records from MongoDB collection (should return 3 records)
`docker-compose run mongodb-fetcher`

5. Insert, update and delete records from MySQL table
`docker-compose run mysql-cli 'INSERT INTO mytable(value) VALUES ("Record1"), ("Record2"), ("Another record")`
`docker-compose run mysql-cli 'UPDATE mytable SET value="RecordTwo" WHERE value="Record2"'
`docker-compose run mysql-cli 'DELETE FROM mytable WHERE value="Record1"'

6. List records from MongoDB collection (should new set of records)
`docker-compose run mongodb-fetcher`
