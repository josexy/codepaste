CREATE DATABASE IF NOT EXISTS db_codepaste;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL,
  `nickname` varchar(20) NOT NULL,
  `password_digest` longtext NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_username` (`username`)
) ENGINE = InnoDB AUTO_INCREMENT = 7 DEFAULT CHARSET = utf8mb4;
CREATE TABLE `pastes` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(20) DEFAULT NULL,
  `title` varchar(30) DEFAULT NULL,
  `content` longtext NOT NULL,
  `lang` varchar(10) DEFAULT NULL,
  `password` varchar(20) DEFAULT NULL,
  `private` tinyint(1) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_pastes_key` (`key`),
  KEY `fk_pastes_user` (`user_id`),
  CONSTRAINT `fk_pastes_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 14 DEFAULT CHARSET = utf8mb4;