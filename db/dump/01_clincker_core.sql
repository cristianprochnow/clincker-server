CREATE DATABASE `clincker_core` CHARACTER SET utf8 COLLATE utf8_bin;
USE `clincker_core`;

CREATE TABLE users (
   id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
   email VARCHAR(80) NULL DEFAULT NULL,
   is_admin SET('0', '1') NULL DEFAULT '0'
) CHARACTER SET utf8 COLLATE utf8_bin;

INSERT INTO users (email, is_admin)
VALUES
    ("contact@contact.com.br", "1"),
    ("extra@extra.com.br", "0");