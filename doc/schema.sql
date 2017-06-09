CREATE DATABASE IF NOT EXISTS `scheduler` DEFAULT CHARACTER SET utf8;

USE `scheduler`;

DROP TABLE IF EXISTS `one_time_jobs`;
CREATE TABLE `one_time_jobs` (
  `id` VARCHAR(36) NOT NULL,
  `name` VARCHAR(50) NOT NULL,
  `description` VARCHAR(256) NOT NULL,
  `retry` TINYINT(1) NOT NULL DEFAULT '0',
  `method` VARCHAR(50) NOT NULL DEFAULT 'GET',
  `addr` VARCHAR(50) NOT NULL,
  `path` VARCHAR(50) NOT NULL,
  `params` VARCHAR(256) NOT NULL DEFAULT '',
  `start_at` BIGINT(20) NOT NULL,  -- 单位：秒
  `created_at` BIGINT(20) NOT NULL,-- 单位：秒
  `ret_err` VARCHAR(128) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `one_time_jobs_start_at` (`start_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `every_jobs`;
CREATE TABLE `every_jobs` (
  `id` VARCHAR(36) NOT NULL,
  `name` VARCHAR(50) NOT NULL,
  `description` VARCHAR(256) NOT NULL,
  `retry` TINYINT(1) NOT NULL DEFAULT '0', -- 失败后重试次数
  `method` VARCHAR(50) NOT NULL DEFAULT 'GET',
  `addr` VARCHAR(50) NOT NULL,
  `path` VARCHAR(50) NOT NULL,
  `params` VARCHAR(256) NOT NULL DEFAULT '',
  `Spec` VARCHAR(32) NOT NULL,
  `created_at` BIGINT(20) NOT NULL,-- 单位：秒
  `ret_err` VARCHAR(128) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
