drop table  if exists `user`;
create TABLE `user`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) NOT NULL UNIQUE,
    `real_name` CHAR(10) DEFAULT '未认证' COMMENT '用于团队人员认证',
    `role` TINYINT(10) DEFAULT 0,
    `phone` CHAR(20) NOT NULL UNIQUE,
    `e_mail` VARCHAR(40) NOT NULL UNIQUE,
    `college` CHAR(20) NOT NULL,
    primary key (`id`)
)engine=innodb DEFAULT CHARSET=UTF8MB4;