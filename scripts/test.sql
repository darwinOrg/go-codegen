create table group_info
(
    id           bigint(20) not null AUTO_INCREMENT primary key,
    name         varchar(200)   not null comment '名称',
    main_data    json           not null,
    content      text           not null,
    bin_data     blob           not null,
    total_amount decimal(10, 2) not null,
    remark       varchar(128),
    created_by   bigint(20) not null,
    created_at   datetime       not null,
    modified_by  bigint(20) not null,
    modified_at  datetime       not null
) ENGINE=innodb CHARACTER SET utf8mb4 comment '组';

create table user_info
(
    id          bigint(20) not null AUTO_INCREMENT primary key,
    name        varchar(200) not null comment '名称',
    mobile      varchar(200) not null comment '手机号',
    age         int          not null comment '年龄',
    status      int          not null comment '状态(1:正常,2:注销)',
    sex         varchar(2)   not null comment '性别(Man:男,Woman:女)',
    sign_out_at datetime     not null comment '注销时间',
    created_by  bigint(20) not null,
    created_at  datetime     not null,
    modified_by bigint(20) not null,
    modified_at datetime     not null
) ENGINE=innodb CHARACTER SET utf8mb4 comment '用户';
