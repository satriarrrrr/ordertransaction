-- Database migration
DROP DATABASE IF EXISTS store;
CREATE DATABASE store character set=utf8;

USE store;

DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
  `id` int(12) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(256) NOT NULL,
  `quantity` int(8) NOT NULL DEFAULT 0,
  `price` int(12) NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_price` (`price`),
  KEY `ix_quantity` (`quantity`)
);

DROP TABLE IF EXISTS `coupons`;
CREATE TABLE `coupons` (
  `id` int(12) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(64) NOT NULL,
  `discount_type` tinyint(2) NOT NULL,
  `discount_nominal` int(12) NOT NULL,
  `max_used` int(12) NOT NULL,
  `current_used` int(12) NOT NULL DEFAULT 0,
  `valid_date_start` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `valid_date_end` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_discount_type` (`discount_type`),
  KEY `ix_discount_nominal` (`discount_nominal`)
);

DROP TABLE IF EXISTS `discount_type`;
CREATE TABLE `discount_type` (
  `id` tinyint(2) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(256) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `used_coupons`;
CREATE TABLE `used_coupons` (
  `id` int(12) unsigned NOT NULL AUTO_INCREMENT,
  `coupon_id` int(12) NOT NULL,
  `order_id` int(12) NOT NULL,
  `total_discount` int(12) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_order_id` (`order_id`),
  KEY `ix_coupon_id` (`coupon_id`)
);

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `id` int(12) unsigned NOT NULL AUTO_INCREMENT,
  `customer_name` varchar(256) NOT NULL,
  `customer_email` varchar(256) NOT NULL,
  `customer_phonenumber` varchar(256) NOT NULL,
  `customer_address` text NOT NULL,
  `order_status` tinyint(2) NOT NULL,
  `total_price` int(12) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_customer_phonenumber` (`customer_phonenumber`),
  KEY `ix_customer_email` (`customer_email`),
  KEY `ix_order_status` (`order_status`)
);

DROP TABLE IF EXISTS `order_status`;
CREATE TABLE `order_status` (
  `id` tinyint(2) unsigned NOT NULL,
  `name` varchar(256) NOT NULL,
  `description` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `order_items`;
CREATE TABLE `order_items` (
  `id` int(12) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` int(12) unsigned NOT NULL,
  `product_id` int(12) unsigned NOT NULL,
  `quantity` int(12) unsigned NOT NULL DEFAULT 1,
  `price` int(12) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_order_id` (`order_id`),
  KEY `ix_product_id` (`product_id`)
);

DROP TABLE IF EXISTS `payments`;
CREATE TABLE `payments` (
  `id` int(12) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` int(12) unsigned NOT NULL,
  `payment_type` tinyint(2) NOT NULL,
  `payment_status` tinyint(2) NOT NULL,
  `proof` varchar(256),
  `total_paid` int(12) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_order_id` (`order_id`),
  KEY `ix_payment_type` (`payment_type`),
  KEY `ix_payment_status` (`payment_status`)
);

DROP TABLE IF EXISTS `payment_type`;
CREATE TABLE `payment_type` (
  `id` tinyint(2) unsigned NOT NULL,
  `name` varchar(256) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `payment_status`;
CREATE TABLE `payment_status` (
  `id` tinyint(2) unsigned NOT NULL,
  `name` varchar(256) NOT NULL,
  `description` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `shippings`;
CREATE TABLE `shippings` (
  `id` int(12) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` int(12) unsigned NOT NULL,
  `shipping_operator_id` int(5) NOT NULL,
  `shipping_status_id` tinyint(2) NOT NULL,
  `no_resi` varchar(32),
  `shipping_cost` int(12) NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_order_id` (`order_id`),
  KEY `ix_shipping_operator_id` (`shipping_operator_id`),
  KEY `ix_shipping_status_id` (`shipping_status_id`)
);

DROP TABLE IF EXISTS `shipping_operators`;
CREATE TABLE `shipping_operators` (
  `id` int(5) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(256) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `shipping_status`;
CREATE TABLE `shipping_status` (
  `id` tinyint(2) unsigned NOT NULL,
  `name` varchar(256) NOT NULL,
  `description` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
);