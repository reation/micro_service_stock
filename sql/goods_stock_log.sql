CREATE TABLE `goods_stock_log` (
                                   `id` int(11) NOT NULL AUTO_INCREMENT,
                                   `goods_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品ID',
                                   `operation_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '操作类型 0:减少 1:增加',
                                   `operation_num` int(11) NOT NULL DEFAULT '1' COMMENT '操作数量',
                                   `type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '操作类型 1:订单减少 2:退订单增加 3:供应商供货 4:供应商回收',
                                   `operation_id` int(11) NOT NULL DEFAULT '1' COMMENT '操作ID type为1、2时为订单ID，为3、4时为供应商ID',
                                   `update_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
                                   `create_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
                                   PRIMARY KEY (`id`),
                                   KEY `goods_id` (`goods_id`,`operation_id`,`type`) USING BTREE,
                                   KEY `operation_id` (`operation_id`,`type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='库存操作日志表';