CREATE DATABASE some_domain CHARACTER SET utf8mb4;
GRANT ALL ON *.* TO 'admin'@'%' WITH GRANT OPTION;
CREATE USER 'admin'@'localshot' IDENTIFIED BY 'admin';
CREATE USER 'admin'@'127.0.0.1' IDENTIFIED BY 'admin';
FLUSH PRIVILEGES;
