create table `terminal_location`
(
    `id`          int unsigned not null auto_increment,
    `name`        varchar(100)          default '' comment '位置名称',
    `summary`     varchar(100)          default '' comment '简介',

    # 公共字段
    `created_at`  int unsigned not null default '0' comment '创建时间',
    `updated_at` int unsigned not null default '0' comment '修改时间',
    primary key (`id`)
) engine = InnoDB
  default charset = utf8mb4 comment '终端放置位置'