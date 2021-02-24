CREATE DATABASE IF NOT EXISTS `ddd_sample`;
USE `ddd_sample`;


CREATE TABLE IF NOT EXISTS `users` (
  `id`          VARCHAR(255) NOT NULL COMMENT 'ユーザID',
  `name`        VARCHAR(255) NOT NULL COMMENT 'ユーザ名',
  `mail_adress` VARCHAR(255) NOT NULL COMMENT 'メールアドレス',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_mail_adress` (`mail_adress`))
ENGINE=InnoDB
COMMENT='ユーザ'
DEFAULT CHARSET=utf8mb4;
