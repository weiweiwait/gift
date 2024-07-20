create database gift;
use gift;
create table inventory (
                           id int(11) NOT NULL AUTO_INCREMENT COMMENT '奖品id,自增',
                           name varchar(20) NOT NULL COMMENT '奖品名称',
                           description varchar(100) not null default '' comment '奖品描述',
                           picture varchar(200) not null default '0' comment '奖品图片',
                           price int(11) not null default '0' comment '价值',
                           count int(11) not null default '0' comment '库存量',
                           primary key(id)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8 COMMENT='奖品库存表'
create table orders (
                        id int(11) NOT NULL AUTO_INCREMENT COMMENT '订单id,自增',
                        gift_id int(11) not null comment '商品id',
                        user_id int(11) not null comment '用户id',
                        count int(11) not null default '1' comment '购买数量',
                        create_time datetime default current_timestamp comment '订单创建时间',
                        primary key(id),
                        key id_user (user_id)
) ENGINE=InnoDB AUTO_INCREMENT=189549 DEFAULT CHARSET=utf8mb4 COMMENT='订单表'