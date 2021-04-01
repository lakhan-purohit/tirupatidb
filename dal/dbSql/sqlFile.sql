CREATE TABLE `cloth_photo` (
 `id` bigint(11) NOT NULL AUTO_INCREMENT,
 `photo` varchar(500) NOT NULL,
 `photo_thumb` varchar(500) NOT NULL,
 `clothTable_id` bigint(20) NOT NULL,
 PRIMARY KEY (`id`),
 KEY `cloth_id_f` (`clothTable_id`),
 CONSTRAINT `cloth_id_f` FOREIGN KEY (`clothTable_id`) REFERENCES `t_shirt_table` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `order_list` (
 `id` bigint(20) NOT NULL AUTO_INCREMENT,
 `user_id` bigint(20) NOT NULL,
 `product_id` bigint(20) NOT NULL,
 `qty` int(11) NOT NULL,
 `status` tinyint(1) NOT NULL,
 `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
 `updateAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
 `total` bigint(20) NOT NULL,
 `paid` bigint(20) NOT NULL,
 `price` bigint(20) NOT NULL,
 PRIMARY KEY (`id`),
 KEY `user_id_o` (`user_id`),
 CONSTRAINT `user_id_o` FOREIGN KEY (`user_id`) REFERENCES `userinfo` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `save_product` (
 `id` bigint(20) NOT NULL AUTO_INCREMENT,
 `user_id` bigint(20) NOT NULL,
 `product_id` bigint(20) NOT NULL,
 PRIMARY KEY (`id`),
 KEY `user_id_s` (`user_id`),
 CONSTRAINT `user_id_s` FOREIGN KEY (`user_id`) REFERENCES `userinfo` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `t_shirt_table` (
 `id` bigint(11) NOT NULL AUTO_INCREMENT,
 `color` varchar(50) NOT NULL,
 `nack` varchar(50) NOT NULL,
 `cloth_type` varchar(50) NOT NULL,
 `size` varchar(100) NOT NULL,
 `brand_name` varchar(50) DEFAULT NULL,
 `hand_slive` varchar(50) NOT NULL,
 `pocket` tinyint(1) NOT NULL,
 `role` int(11) NOT NULL,
 `price` bigint(20) NOT NULL DEFAULT 0,
 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

	CREATE TABLE `userinfo` (
 `id` bigint(20) NOT NULL AUTO_INCREMENT,
 `user_mobile_number` varchar(20) NOT NULL,
 `user_name` varchar(80) NOT NULL,
 `email_id` varchar(80) NOT NULL,
 `firm_name` varchar(150) NOT NULL,
 `password` varchar(200) NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
 `blocked` tinyint(1) NOT NULL,
 `role` smallint(6) NOT NULL,
 PRIMARY KEY (`id`),
 UNIQUE KEY `unique_email` (`email_id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4;

