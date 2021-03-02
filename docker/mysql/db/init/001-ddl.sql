CREATE DATABASE IF NOT EXISTS `ddd_sample`;
USE `ddd_sample`;

DROP TABLE IF EXISTS `users`;

CREATE TABLE IF NOT EXISTS `users` (
  `id`              CHAR(26)     NOT NULL COMMENT 'ユーザID',
  `name`            VARCHAR(64)  NOT NULL COMMENT 'ユーザ名',
  `mail_address`    VARCHAR(255) NOT NULL COMMENT 'メールアドレス',
  `hashed_password` VARCHAR(255) NOT NULL COMMENT 'パスワード',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_mail_address` (`mail_address`))
ENGINE=InnoDB
COMMENT='ユーザ'
DEFAULT CHARSET=utf8mb4;
