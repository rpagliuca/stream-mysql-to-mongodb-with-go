# About

Using the library go-mysql, I've setup a proof of concept environment where changes to a MySQL table are reflected into a MongoDB collection

# Usage

* Start servers:
```
docker-compose up mysql-server mongodb-server
```

* Initialize MySQL table with 3 records
```
docker-compose run mysql-cli init
```

* List records from MongoDB collection (should return no records)
```
docker-compose run mongodb-fetcher
```

* Initialize mysql-consumer, which will read changes from MySQL server
```
docker-compose up mysql-consumer
```

* List records from MongoDB collection (should return 3 records)
```
docker-compose run mongodb-fetcher
```

* Insert, update and delete records from MySQL table
```
docker-compose run mysql-cli 'INSERT INTO mytable(value) VALUES ("Record1"), ("Record2"), ("Another record")'

docker-compose run mysql-cli 'UPDATE mytable SET value="RecordTwo" WHERE value="Record2"'

docker-compose run mysql-cli 'DELETE FROM mytable WHERE value="Record1"'
```

* List records from MongoDB collection (should return the new set of records)
```
docker-compose run mongodb-fetcher
```
