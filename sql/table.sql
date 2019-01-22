CREATE SCHEMA IF NOT EXISTS `oauth` DEFAULT CHARACTER SET utf8 ;
USE `oauth` ;

CREATE TABLE IF NOT EXISTS `oauth`.`user` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` VARCHAR(100) NOT NULL DEFAULT 'default' COMMENT '账户名',
  `password` VARCHAR(255) NOT NULL DEFAULT 'default' COMMENT '密码',
  `nickname` VARCHAR(255) NOT NULL DEFAULT 'default' COMMENT '昵称',
  `email` VARCHAR(245) NOT NULL DEFAULT 'default' COMMENT '邮箱',
  `avater` VARCHAR(255) NOT NULL DEFAULT 'default' COMMENT '头像',
  `created_at` BIGINT(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uniq_id` (`id` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COMMENT = '用户表';


CREATE TABLE IF NOT EXISTS `oauth`.`application` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` VARCHAR(255) NOT NULL DEFAULT 'default' COMMENT '用户id',
  `name` VARCHAR(255) NOT NULL DEFAULT 'default' COMMENT '应用名称',
  `client_id` VARCHAR(255) NOT NULL DEFAULT 'default' COMMENT '客户端ID',
  `private_key` VARCHAR(255) NOT NULL DEFAULT 'default' COMMENT '客户端私钥',
  `callback` VARCHAR(255) NOT NULL DEFAULT 'default' COMMENT '客户端回调地址',
  `mode` VARCHAR(255) NOT NULL DEFAULT 'default' COMMENT '客户端运行模式',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uniq_id` (`id` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COMMENT = '应用表';


CREATE TABLE IF NOT EXISTS `oauth`.`app_user_list` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` BIGINT(20) NOT NULL DEFAULT 0 COMMENT '用户id',
  `app_id` BIGINT(20) NOT NULL DEFAULT 0 COMMENT '应用id',
  `category` VARCHAR(255) NOT NULL DEFAULT 'black' COMMENT '黑白名单',
  `role_id` BIGINT(20) NOT NULL DEFAULT 0 COMMENT '角色id',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uniq_id` (`id` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COMMENT = '应用的用户列表';


CREATE TABLE IF NOT EXISTS `oauth`.`app_role` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `app_id` BIGINT(20) NOT NULL DEFAULT 0 COMMENT '应用id',
  `name` VARCHAR(255) NOT NULL DEFAULT 'default' COMMENT '角色名称',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uniq_id` (`id` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COMMENT = '应用的角色信息表';

CREATE TABLE IF NOT EXISTS `oauth`.`app_role_permission` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_id` BIGINT(20) NOT NULL DEFAULT 0 COMMENT '角色id',
  `name` VARCHAR(255) NOT NULL DEFAULT 'default' COMMENT '权限别称',
  `method` VARCHAR(45) NOT NULL DEFAULT 'GET' COMMENT 'HTTP方法',
  `pattern` VARCHAR(255) NOT NULL DEFAULT '*' COMMENT '路径匹配正则',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uniq_id` (`id` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COMMENT = '应用角色的权限列表';