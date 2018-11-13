CREATE DATABASE IF NOT EXISTS cdnupload DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE cdnupload;

CREATE TABLE IF NOT EXISTS `t_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除: 0:正常/1:已删除',
  `username` varchar(30) NOT NULL DEFAULT '' COMMENT '用户ID',
  `password` varchar(100) NOT NULL DEFAULT '' COMMENT '用户密码',
  `is_admin` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否管理员: 0:管理员/1:非管理员',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `status` varchar(50) NOT NULL DEFAULT '未使用' COMMENT '状态: 使用中/未使用/已删除',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '用户信息表';


CREATE TABLE IF NOT EXISTS `t_project` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '项目ID',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除: 0:正常/1:已删除',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '项目名称',
  `platform_id` int(11) NOT NULL DEFAULT '0' COMMENT 'CDN平台ID',
  `path` varchar(100) NOT NULL DEFAULT '' COMMENT '存储路径',
  `domain` varchar(100) NOT NULL DEFAULT '' COMMENT 'CDN域名',
  `cdn_id` int(11) COMMENT 'CDN厂商ID',
  `status` varchar(50) NOT NULL DEFAULT '未使用' COMMENT '状态: 使用中/未使用/已删除',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `project_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '项目信息表';


CREATE TABLE IF NOT EXISTS `t_cdn_platform` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'CDN平台ID',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除: 0:正常/1:已删除',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT 'CDN名称',
  `secret_id` varchar(100) NOT NULL DEFAULT '' COMMENT '加密id',
  `secret_key` varchar(100) NOT NULL DEFAULT '' COMMENT '加密key',
  `api_url` varchar(100) NOT NULL DEFAULT '' COMMENT '',
  `action` varchar(100) COMMENT 'CDN调用方法',
  `status` varchar(50) NOT NULL DEFAULT '未使用' COMMENT '状态: 使用中/未使用/已删除',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `platform_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT 'CDN平台信息表';

CREATE TABLE IF NOT EXISTS `t_file_suffix` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除: 0:正常/1:已删除',
  `file_suffix` varchar(50) NOT NULL DEFAULT '' COMMENT '文件后缀',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `file_suffix` (`file_suffix`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '允许的文件后缀信息表';

CREATE TABLE IF NOT EXISTS `t_cdn_file` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'CDN文件ID',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除: 0:正常/1:已删除',
  `project_id` int(11) NOT NULL DEFAULT '0' COMMENT '项目ID',
  `file_name` varchar(100) BINARY NOT NULL DEFAULT '' COMMENT '文件名',
  `file_md5` varchar(50) NOT NULL DEFAULT '' COMMENT '文件md5',
  `file_size` varchar(50) NOT NULL DEFAULT '' COMMENT '文件大小',
  `comment` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT 'CDN FILE信息表';

ALTER TABLE `t_project` ADD CONSTRAINT `fk_project_platform` FOREIGN KEY (`platform_id`) REFERENCES `t_cdn_platform` (`id`);
ALTER TABLE `t_cdn_file` ADD CONSTRAINT `fk_cdnfile_project` FOREIGN KEY (`project_id`) REFERENCES `t_project` (`id`);