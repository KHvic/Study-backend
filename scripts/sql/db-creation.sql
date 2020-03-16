DROP TABLE IF EXISTS `question`;
CREATE TABLE `question` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `category` varchar(100) DEFAULT '',
    `sub_category` varchar(100) DEFAULT '',
    `type` varchar(100) DEFAULT '',
    `description` varchar(1020) NOT NULL,
    `options` varchar(255) NOT NULL,
    `answers` varchar(255) NOT NULL,
    `image` varchar(255),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
