CREATE TABLE `user_wallet` (
    `user_id` BIGINT UNSIGNED NOT NULL,
    `gold` INT DEFAULT 0 COMMENT '金币',
    `silver` INT DEFAULT 0 COMMENT '银币',
    `token` INT DEFAULT 0 COMMENT '点券',
    
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户货币钱包';

CREATE TABLE `wallet_log` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `type` TINYINT NOT NULL COMMENT '操作类型：1广告 2充值 3购买商品 4工坊 5转赠',
    `change_gold` INT DEFAULT 0 COMMENT '金币变化',
    `change_silver` INT DEFAULT 0 COMMENT '银币变化',
    `change_token` INT DEFAULT 0 COMMENT '点券变化',
    `reason` VARCHAR(255) COMMENT '说明',
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    KEY `idx_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='货币流水记录';

CREATE TABLE `gold_transfer` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `from_user` BIGINT UNSIGNED COMMENT '发送方',
    `to_user` BIGINT UNSIGNED COMMENT '接收方',
    `amount` INT NOT NULL COMMENT '转赠数量',
    `remark` VARCHAR(255) DEFAULT NULL COMMENT '备注',
    `status` TINYINT DEFAULT 1 COMMENT '1已发放',
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='金币转赠记录';
