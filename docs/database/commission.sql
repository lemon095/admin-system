CREATE TABLE `commission_log` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '获得佣金的用户ID',
    `from_user_id` BIGINT UNSIGNED NOT NULL COMMENT '行为来源用户ID（下级或下下级）',
    `type` TINYINT NOT NULL COMMENT '佣金类型：1广告 2付费 3工坊产出 等',
    `level` TINYINT NOT NULL COMMENT '层级：1一级佣金 2二级佣金',
    `reward_gold` INT DEFAULT 0 COMMENT '奖励金币数量',
    `reward_silver` INT DEFAULT 0 COMMENT '奖励银币数量',
    `reward_material` JSON DEFAULT NULL COMMENT '奖励材料，JSON结构',
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '发放时间',

    PRIMARY KEY (`id`),
    KEY `idx_user` (`user_id`),
    KEY `idx_from_user` (`from_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='佣金发放记录表';

CREATE TABLE `commission_config` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT '主键',
    `type` TINYINT NOT NULL COMMENT '佣金类型：1广告 2付费 3工坊 等',
    `level` TINYINT NOT NULL COMMENT '层级：1一级 2二级',
    `percent` INT NOT NULL COMMENT '佣金比例（例如8%存8）',
    `daily_limit` INT DEFAULT NULL COMMENT '用户每日佣金上限',
    `monthly_limit` INT DEFAULT NULL COMMENT '用户月度佣金上限',

    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_type_level` (`type`,`level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='佣金比例与规则配置表';
