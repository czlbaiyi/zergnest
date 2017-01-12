DROP DATABASE IF EXISTS by_zjz;
CREATE DATABASE by_zjz;
USE by_zjz;
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account` (  
  `id` BIGINT NOT NULL ,
	`account_type` TINYINT NOT NULL ,
	`account` VARCHAR(12) NOT NULL ,
  `password` VARCHAR(12) NOT NULL,
  `nickname` VARCHAR(20) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8;