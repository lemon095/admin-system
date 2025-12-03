-- 用户表设计
CREATE TABLE IF NOT EXISTS `users` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` VARCHAR(50) NOT NULL COMMENT '用户名',
  `avatar_id` VARCHAR(255) DEFAULT NULL COMMENT '头像ID/URL',
  `role` TINYINT NOT NULL DEFAULT 3 COMMENT '用户身份：1-超级管理员，2-管理员，3-普通用户',
  `password` VARCHAR(32) NOT NULL COMMENT '密码（MD5值）',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1-启用，0-禁用',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_by` BIGINT UNSIGNED DEFAULT NULL COMMENT '创建者ID',
  `last_login_at` DATETIME DEFAULT NULL COMMENT '最后登录时间',
  `last_login_ip` VARCHAR(50) DEFAULT NULL COMMENT '最后登录IP',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`),
  KEY `idx_role` (`role`),
  KEY `idx_status` (`status`),
  KEY `idx_created_by` (`created_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 插入默认超级管理员账号
-- 用户名: chuchangkeji
-- 密码: chuchangkeji666 (MD5: 需要计算)
-- 注意：实际插入时需要使用MD5值
INSERT INTO `users` (`username`, `avatar_id`, `role`, `password`, `status`, `created_by`)
VALUES ('chuchangkeji', 'default-avatar-001', 1, MD5('chuchangkeji666'), 1, 1)
ON DUPLICATE KEY UPDATE `username`=`username`;

CREATE TABLE `admin_role` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT '角色ID',
    `name` VARCHAR(50) NOT NULL COMMENT '角色名称：超管/运营/客服',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台管理员角色表';

CREATE TABLE `admin_user` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
    `username` VARCHAR(50) NOT NULL UNIQUE COMMENT '登录用户名',
    `password` VARCHAR(255) NOT NULL COMMENT '加密后的密码',
    `role_id` INT NOT NULL COMMENT '角色ID（外键）',
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台管理员账户表';

CREATE TABLE `invite_relation` (
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID（下级）',
    `parent_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '一级邀请人（上级）',
    `grand_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '二级邀请人（上上级）',
    `bind_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '绑定时间',
    
    PRIMARY KEY (`user_id`),
    KEY `idx_parent` (`parent_id`),
    KEY `idx_grand` (`grand_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户邀请绑定关系表';

