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

