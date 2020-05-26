# About

Using the library go-mysql, I've setup a proof of concept environment where changes to a MySQL table are reflected into a MongoDB collection

# Usage

* Start servers:
```
docker-compose up
```

* List records from MongoDB collection (should return 0 records)
```
docker-compose run mongodb-fetcher
```

* Create MySQL table and insert 3 records
```
docker-compose run mysql-cli init
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

* List records from MongoDB collection (should new set of records)
```
docker-compose run mongodb-fetcher
```
