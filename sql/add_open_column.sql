ALTER TABLE `open-auth`.`application`
ADD COLUMN `open` TINYINT NOT NULL DEFAULT 0 COMMENT '是否公开该应用到用户的可用应用列表' AFTER `mode`;
ADD COLUMN `icon` VARCHAR(1000) NOT NULL DEFAULT '' COMMENT '应用图标' AFTER `open`,
ADD COLUMN `home_page` VARCHAR(1000) NOT NULL DEFAULT '' COMMENT '应用首页' AFTER `icon`;
