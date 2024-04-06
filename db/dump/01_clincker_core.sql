CREATE DATABASE `clincker_core` CHARACTER SET utf8 COLLATE utf8_bin;
USE `clincker_core`;

CREATE TABLE users (
   id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
   email VARCHAR(80) NOT NULL,
   name VARCHAR(100) NULL,
   is_admin SET('0', '1') NULL DEFAULT '0',
   password VARCHAR(80) NOT NULL,
   created_at DATETIME DEFAULT CURRENT_TIMESTAMP NULL
) CHARACTER SET utf8 COLLATE utf8_bin;

CREATE UNIQUE INDEX users_email_uindex
    ON users (email);

INSERT INTO users (email, name, is_admin, password)
VALUES
    ("contact@contact.com.br", "Contato par Contato", "1", "swiheiwheweuwheuwhe"),
    ("extra@extra.com.br", "Extra Extra", "0", "hsuhuhuhsuhdushds");
