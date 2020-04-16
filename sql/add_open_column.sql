ALTER TABLE `open-auth`.`application`
ADD COLUMN `open` TINYINT NOT NULL DEFAULT 0 COMMENT '是否公开该应用到用户的可用应用列表' AFTER `mode`;
