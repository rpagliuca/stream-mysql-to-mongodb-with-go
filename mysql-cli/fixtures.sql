DROP DATABASE IF EXISTS mydb;
CREATE DATABASE mydb;
USE mydb;
CREATE TABLE mytable (
    id INT AUTO_INCREMENT PRIMARY KEY,
    value TEXT NULL
);
INSERT INTO mytable(value) VALUES ("value1"), ("value2"), ("value3");
