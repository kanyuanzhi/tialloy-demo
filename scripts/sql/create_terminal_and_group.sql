create table `terminal_and_group`
(
    `id`          int unsigned not null auto_increment,
    `terminal_id` int                   default '0' comment '终端id',
    `group_id`    int                   default '0' comment '终端组id',

    # 公共字段
    `created_at`  int unsigned not null default '0' comment '创建时间',
    `updated_at` int unsigned not null default '0' comment '修改时间',
    primary key (`id`)
) engine = InnoDB
  default charset = utf8mb4 comment '终端和终端组关联'