CREATE TABLE `ad_watch_log` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `ad_platform` VARCHAR(50) COMMENT '广告平台',
    `status` TINYINT DEFAULT 1 COMMENT '0失败 1成功',
    `reward_gold` INT DEFAULT 0,
    `reward_silver` INT DEFAULT 0,
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='广告观看记录';


CREATE TABLE `ad_config` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `type` TINYINT NOT NULL COMMENT '广告类型',
    `base_reward_silver` INT DEFAULT 0 COMMENT '基础银币奖励',
    `base_reward_material` JSON DEFAULT NULL COMMENT '基础材料奖励',
    `double_period` JSON DEFAULT NULL COMMENT '双倍时段'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='广告奖励配置';
