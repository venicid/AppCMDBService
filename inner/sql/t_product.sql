/*
Navicat MySQL Data Transfer

Source Server         : mysql
Source Server Version : 50736
Source Host           : 150.158.171.205:3306
Source Database       : t_cmdb

Target Server Type    : MYSQL
Target Server Version : 50736
File Encoding         : 65001

Date: 2022-05-27 00:02:30
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `t_product`
-- ----------------------------
DROP TABLE IF EXISTS `t_product`;
CREATE TABLE `t_product` (
  `id` int(10) NOT NULL COMMENT '主键',
  `product_name` varchar(100) NOT NULL COMMENT '名称',
  `code` varchar(50) NOT NULL COMMENT '代码',
  `product_type` varchar(50) NOT NULL COMMENT '类型',
  `is_delete` smallint(6) NOT NULL COMMENT '是否删除',
  `create_user` varchar(50) NOT NULL COMMENT '创建人',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `parent_id` int(11) DEFAULT NULL COMMENT '父节点ID，t_product.id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_product
-- ----------------------------
