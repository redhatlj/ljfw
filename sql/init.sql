# 创建数据库
create database  if not exists cmdb default charset utf8mb4 ;

use cmdb;

# 创建表
create table  if not exists  user (
                                      id bigint primary key auto_increment comment '用户ID',
                                      name varchar(64) not null default '' comment '用户名',
                                      password varchar(1024) not null default '' comment '用户密码'
) engine=innodb default charset utf8mb4;

#插入数据
insert into user( name, password) VALUES ('test',md5("123abc"));

select id,name,password from user where name='test';
# select * from user;
# 清空表
# truncate table  user;
# delete from user where id = 1;
