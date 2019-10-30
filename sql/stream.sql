/*
 Navicat Premium Data Transfer

 Source Server         : Docker-Local
 Source Server Type    : MySQL
 Source Server Version : 80017
 Source Host           : 192.168.56.108:3306
 Source Schema         : stream

 Target Server Type    : MySQL
 Target Server Version : 80017
 File Encoding         : 65001

 Date: 30/10/2019 21:53:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for pp_set_serial
-- ----------------------------
DROP TABLE IF EXISTS `pp_set_serial`;
CREATE TABLE `pp_set_serial`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `hardware_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '硬件编码',
  `serial_code` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '序列号',
  `valid_time` bigint(20) NULL DEFAULT 0 COMMENT '有效期',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `create_id` int(11) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `update_id` int(11) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `create_time` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_create_id`(`create_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pp_set_serial
-- ----------------------------
INSERT INTO `pp_set_serial` VALUES (1, '', '', 4102415999, '请输入序列号注册', 0, 0, 0, 0);

-- ----------------------------
-- Table structure for pp_uc_admin
-- ----------------------------
DROP TABLE IF EXISTS `pp_uc_admin`;
CREATE TABLE `pp_uc_admin`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `login_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `real_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '真实姓名',
  `password` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `role_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '角色id字符串，如：2,3,4',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '手机号码',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `salt` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码盐',
  `last_login` int(11) NOT NULL DEFAULT 0 COMMENT '最后登录时间',
  `last_ip` char(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态，1-正常 0禁用',
  `create_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者ID',
  `update_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改者ID',
  `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_user_name`(`login_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pp_uc_admin
-- ----------------------------
INSERT INTO `pp_uc_admin` VALUES (1, 'admin', '超级管理员', 'eadee0dc2dc2cde2f8d93c2056d577b0', '0,2,1', '13888888889', 'jingjie520@gmail.com', '5gfG', 1572442706, '[::1]', 1, 0, 1, 0, 1562597561);

-- ----------------------------
-- Table structure for pp_uc_auth
-- ----------------------------
DROP TABLE IF EXISTS `pp_uc_auth`;
CREATE TABLE `pp_uc_auth`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `pid` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级ID，0为顶级',
  `auth_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '权限名称',
  `auth_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT 'URL地址',
  `sort` int(1) UNSIGNED NOT NULL DEFAULT 999 COMMENT '排序，越小越前',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `is_show` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否显示，0-隐藏，1-显示',
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作者ID',
  `create_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者ID',
  `update_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改者ID',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态，1-正常，0-删除',
  `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 59 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限因子' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pp_uc_auth
-- ----------------------------
INSERT INTO `pp_uc_auth` VALUES (1, 0, '所有权限', '/', 1, '', 0, 1, 1, 1, 1, 1505620970, 1505620970);
INSERT INTO `pp_uc_auth` VALUES (2, 1, '权限管理', '/', 999, 'fa-id-card', 1, 1, 0, 1, 1, 0, 1505622360);
INSERT INTO `pp_uc_auth` VALUES (3, 2, '管理员', '/admin/list', 1, 'fa-user-o', 1, 1, 1, 1, 1, 1505621186, 1505621186);
INSERT INTO `pp_uc_auth` VALUES (4, 2, '角色管理', '/role/list', 2, 'fa-user-circle-o', 1, 1, 0, 1, 1, 0, 1505621852);
INSERT INTO `pp_uc_auth` VALUES (5, 3, '新增', '/admin/add', 1, '', 0, 1, 0, 1, 1, 0, 1505621685);
INSERT INTO `pp_uc_auth` VALUES (6, 3, '修改', '/admin/edit', 2, '', 0, 1, 0, 1, 1, 0, 1505621697);
INSERT INTO `pp_uc_auth` VALUES (7, 3, '删除', '/admin/ajaxdel', 3, '', 0, 1, 1, 1, 1, 1505621756, 1505621756);
INSERT INTO `pp_uc_auth` VALUES (8, 4, '新增', '/role/add', 1, '', 1, 1, 0, 1, 1, 0, 1505698716);
INSERT INTO `pp_uc_auth` VALUES (9, 4, '修改', '/role/edit', 2, '', 0, 1, 1, 1, 1, 1505621912, 1505621912);
INSERT INTO `pp_uc_auth` VALUES (10, 4, '删除', '/role/ajaxdel', 3, '', 0, 1, 1, 1, 1, 1505621951, 1505621951);
INSERT INTO `pp_uc_auth` VALUES (11, 2, '权限因子', '/auth/list', 3, 'fa-list', 1, 1, 1, 1, 1, 1505621986, 1505621986);
INSERT INTO `pp_uc_auth` VALUES (12, 11, '新增', '/auth/add', 1, '', 0, 1, 1, 1, 1, 1505622009, 1505622009);
INSERT INTO `pp_uc_auth` VALUES (13, 11, '修改', '/auth/edit', 2, '', 0, 1, 1, 1, 1, 1505622047, 1505622047);
INSERT INTO `pp_uc_auth` VALUES (14, 11, '删除', '/auth/ajaxdel', 3, '', 0, 1, 1, 1, 1, 1505622111, 1505622111);
INSERT INTO `pp_uc_auth` VALUES (15, 1, '个人中心', 'profile/edit', 1001, 'fa-user-circle-o', 1, 1, 0, 1, 1, 0, 1506001114);
INSERT INTO `pp_uc_auth` VALUES (20, 1, '基础设置', '/', 2, 'fa-cogs', 1, 1, 1, 1, 1, 1505622601, 1505622601);
INSERT INTO `pp_uc_auth` VALUES (22, 20, '序列号设置', '/serial/detail', 2, 'fa-tree', 1, 1, 0, 1, 1, 0, 1572364878);
INSERT INTO `pp_uc_auth` VALUES (24, 15, '资料修改', '/user/edit', 1, 'fa-edit', 1, 1, 0, 1, 1, 0, 1506057468);
INSERT INTO `pp_uc_auth` VALUES (51, 1, '频道管理', '/', 1, 'fa-podcast', 1, 1, 1, 1, 1, 1561915094, 1561915094);
INSERT INTO `pp_uc_auth` VALUES (52, 51, '所有频道', '/channel/list', 1, 'fa-list', 1, 1, 1, 1, 1, 1561915228, 1561915228);
INSERT INTO `pp_uc_auth` VALUES (53, 51, '新增频道', '/channel/add', 2, 'fa-plus', 1, 1, 0, 1, 1, 0, 1561993844);
INSERT INTO `pp_uc_auth` VALUES (54, 52, '修改频道', '/channel/edit', 1, 'fa-link', 0, 1, 0, 1, 1, 0, 1566918787);
INSERT INTO `pp_uc_auth` VALUES (55, 52, '删除频道', '/channel/ajaxdel', 9, 'fa-link', 0, 1, 0, 1, 1, 0, 1568817057);
INSERT INTO `pp_uc_auth` VALUES (56, 52, '启动频道', '/channel/actionstart', 2, 'fa-list', 0, 1, 0, 1, 1, 0, 1566918773);
INSERT INTO `pp_uc_auth` VALUES (57, 52, '停止频道', '/channel/actionstop', 4, 'fa-list', 0, 1, 0, 1, 1, 0, 1568817035);
INSERT INTO `pp_uc_auth` VALUES (58, 52, '启动频道保存', '/channel/ajaxstartsave', 3, 'fa-list', 0, 1, 1, 1, 1, 1568817026, 1568817026);
INSERT INTO `pp_uc_auth` VALUES (59, 52, '停止频道保存', '/channel/ajaxstopsave', 5, 'fa-list', 0, 1, 0, 1, 1, 0, 1569335890);

-- ----------------------------
-- Table structure for pp_uc_role
-- ----------------------------
DROP TABLE IF EXISTS `pp_uc_role`;
CREATE TABLE `pp_uc_role`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '角色名称',
  `detail` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '备注',
  `create_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者ID',
  `update_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改这ID',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态1-正常，0-删除',
  `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pp_uc_role
-- ----------------------------
INSERT INTO `pp_uc_role` VALUES (1, '频道管理员', '拥有频道所有权限', 0, 1, 1, 1572443556, 1572443556);
INSERT INTO `pp_uc_role` VALUES (2, '系统管理员', '系统管理员', 0, 1, 1, 1568817080, 1568817080);

-- ----------------------------
-- Table structure for pp_uc_role_auth
-- ----------------------------
DROP TABLE IF EXISTS `pp_uc_role_auth`;
CREATE TABLE `pp_uc_role_auth`  (
  `role_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  `auth_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限ID',
  PRIMARY KEY (`role_id`, `auth_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限和角色关系表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pp_uc_role_auth
-- ----------------------------
INSERT INTO `pp_uc_role_auth` VALUES (1, 0);
INSERT INTO `pp_uc_role_auth` VALUES (1, 1);
INSERT INTO `pp_uc_role_auth` VALUES (1, 15);
INSERT INTO `pp_uc_role_auth` VALUES (1, 24);
INSERT INTO `pp_uc_role_auth` VALUES (1, 51);
INSERT INTO `pp_uc_role_auth` VALUES (1, 52);
INSERT INTO `pp_uc_role_auth` VALUES (1, 53);
INSERT INTO `pp_uc_role_auth` VALUES (1, 54);
INSERT INTO `pp_uc_role_auth` VALUES (1, 55);
INSERT INTO `pp_uc_role_auth` VALUES (1, 56);
INSERT INTO `pp_uc_role_auth` VALUES (1, 57);
INSERT INTO `pp_uc_role_auth` VALUES (1, 58);
INSERT INTO `pp_uc_role_auth` VALUES (1, 59);
INSERT INTO `pp_uc_role_auth` VALUES (2, 1);
INSERT INTO `pp_uc_role_auth` VALUES (2, 2);
INSERT INTO `pp_uc_role_auth` VALUES (2, 3);
INSERT INTO `pp_uc_role_auth` VALUES (2, 4);
INSERT INTO `pp_uc_role_auth` VALUES (2, 5);
INSERT INTO `pp_uc_role_auth` VALUES (2, 6);
INSERT INTO `pp_uc_role_auth` VALUES (2, 7);
INSERT INTO `pp_uc_role_auth` VALUES (2, 8);
INSERT INTO `pp_uc_role_auth` VALUES (2, 9);
INSERT INTO `pp_uc_role_auth` VALUES (2, 10);
INSERT INTO `pp_uc_role_auth` VALUES (2, 11);
INSERT INTO `pp_uc_role_auth` VALUES (2, 12);
INSERT INTO `pp_uc_role_auth` VALUES (2, 13);
INSERT INTO `pp_uc_role_auth` VALUES (2, 14);
INSERT INTO `pp_uc_role_auth` VALUES (2, 15);
INSERT INTO `pp_uc_role_auth` VALUES (2, 20);
INSERT INTO `pp_uc_role_auth` VALUES (2, 22);
INSERT INTO `pp_uc_role_auth` VALUES (2, 24);
INSERT INTO `pp_uc_role_auth` VALUES (2, 51);
INSERT INTO `pp_uc_role_auth` VALUES (2, 52);
INSERT INTO `pp_uc_role_auth` VALUES (2, 53);
INSERT INTO `pp_uc_role_auth` VALUES (2, 54);
INSERT INTO `pp_uc_role_auth` VALUES (2, 55);
INSERT INTO `pp_uc_role_auth` VALUES (2, 56);
INSERT INTO `pp_uc_role_auth` VALUES (2, 57);
INSERT INTO `pp_uc_role_auth` VALUES (2, 58);
INSERT INTO `pp_uc_role_auth` VALUES (2, 59);

SET FOREIGN_KEY_CHECKS = 1;
