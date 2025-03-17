create table `user_info` (
    id int(11) not null primary key auto_increment,
    username varchar(64) not null comment '用户名',
    age int(4) not null default -1 comment '年龄',
    created_time timestamp not null default current_timestamp comment '创建时间',
    updated_time timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
    unique key `uk_username`(`username`)
) engine=InnoDB charset=utf8mb4 comment '用户信息表';

create table `announcement_info` (
                             id int(11) not null primary key auto_increment,
                             content text null comment '公告内容',
                             status int(4) not null default 0 comment '状态 0 未上线 1 已上线',
                             created_time timestamp not null default current_timestamp comment '创建时间',
                             updated_time timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
                             unique key `uk_username`(`username`)
) engine=InnoDB charset=utf8mb4 comment '公告信息表';