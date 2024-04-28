CREATE DATABASE `clincker_core` CHARACTER SET utf8 COLLATE utf8_bin;
USE `clincker_core`;

CREATE TABLE `users` (
     `id` int(11) NOT NULL AUTO_INCREMENT,
     `email` varchar(80) COLLATE utf8_bin NOT NULL,
     `name` varchar(100) COLLATE utf8_bin DEFAULT NULL,
     `is_admin` set('0','1') COLLATE utf8_bin DEFAULT '0',
     `password` varchar(80) COLLATE utf8_bin NOT NULL,
     `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
     `hash` varchar(16) COLLATE utf8_bin NOT NULL,
     PRIMARY KEY (`id`),
     UNIQUE KEY `users_email_uindex` (`email`),
     UNIQUE KEY `users_hash_uindex` (`hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `links` (
     `id` int(11) NOT NULL AUTO_INCREMENT,
     `hash` varchar(64) COLLATE utf8_bin NOT NULL,
     `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
     `edited_at` datetime DEFAULT NULL,
     `original_url` varchar(100) COLLATE utf8_bin DEFAULT NULL,
     `domain` varchar(80) COLLATE utf8_bin DEFAULT NULL,
     `resources` varchar(80) COLLATE utf8_bin DEFAULT NULL,
     `protocol` varchar(10) COLLATE utf8_bin DEFAULT NULL,
     `user` int(11) NOT NULL,
     PRIMARY KEY (`id`),
     UNIQUE KEY `links_id_uindex` (`id`),
     UNIQUE KEY `links_hash_uindex` (`hash`),
     KEY `links_fk_users_idx` (`user`),
     CONSTRAINT `links_fk_users` FOREIGN KEY (`user`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
