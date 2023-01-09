CREATE TABLE `goods_stock` (
                               `id` int(11) NOT NULL AUTO_INCREMENT,
                               `goods_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品ID',
                               `goods_num` int(11) NOT NULL DEFAULT '0',
                               `goods_alert_num` int(11) NOT NULL DEFAULT '0' COMMENT '商品报警数据量',
                               `update_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
                               `create_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
                               PRIMARY KEY (`id`),
                               KEY `goods_id` (`goods_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品库存表';