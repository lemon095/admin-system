CREATE TABLE `shop_item` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100),
    `icon` VARCHAR(255),
    `description` VARCHAR(255),
    `price` INT NOT NULL COMMENT '点券价格',
    `stock` INT DEFAULT NULL COMMENT '库存（可为空无限）',
    `daily_limit` INT DEFAULT NULL COMMENT '每日限购',
    `weekly_limit` INT DEFAULT NULL COMMENT '每周限购',
    `start_time` DATETIME DEFAULT NULL,
    `end_time` DATETIME DEFAULT NULL,
    `status` TINYINT DEFAULT 1 COMMENT '1上架 0下架',

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城商品';


CREATE TABLE `order` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `item_id` BIGINT UNSIGNED NOT NULL,
    `price` INT NOT NULL COMMENT '金额',
    `token_paid` INT NOT NULL COMMENT '支付点券',
    `status` TINYINT NOT NULL COMMENT '0未支付 1成功 2失败 3超时',
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    KEY `idx_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城订单';
