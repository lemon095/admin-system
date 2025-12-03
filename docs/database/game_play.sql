CREATE TABLE `game_play_log` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `type` TINYINT NOT NULL COMMENT '玩法类型',
    `materials` JSON DEFAULT NULL COMMENT '产出材料',
    `silver` INT DEFAULT 0 COMMENT '产出银币',
    `duration` INT DEFAULT 0 COMMENT '时长（秒）',
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩法产出记录';
