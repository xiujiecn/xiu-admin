SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for gen_table
-- ----------------------------
DROP TABLE IF EXISTS `gen_table`;
CREATE TABLE `gen_table` (
  `table_id` bigint(20) NOT NULL COMMENT '编号',
  `data_name` varchar(200) DEFAULT '' COMMENT '数据源名称',
  `table_name` varchar(200) DEFAULT '' COMMENT '表名称',
  `table_comment` varchar(500) DEFAULT '' COMMENT '表描述',
  `sub_table_name` varchar(64) DEFAULT NULL COMMENT '关联子表的表名',
  `sub_table_fk_name` varchar(64) DEFAULT NULL COMMENT '子表关联的外键名',
  `class_name` varchar(100) DEFAULT '' COMMENT '实体类名称',
  `tpl_category` varchar(200) DEFAULT 'crud' COMMENT '使用的模板（crud单表操作 tree树表操作）',
  `package_name` varchar(100) DEFAULT NULL COMMENT '生成包路径',
  `module_name` varchar(30) DEFAULT NULL COMMENT '生成模块名',
  `business_name` varchar(30) DEFAULT NULL COMMENT '生成业务名',
  `function_name` varchar(50) DEFAULT NULL COMMENT '生成功能名',
  `function_author` varchar(50) DEFAULT NULL COMMENT '生成功能作者',
  `gen_type` char(1) DEFAULT '0' COMMENT '生成代码方式（0zip压缩包 1自定义路径）',
  `gen_path` varchar(200) DEFAULT '/' COMMENT '生成路径（不填默认项目路径）',
  `options` varchar(1000) DEFAULT NULL COMMENT '其它生成选项',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`table_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='代码生成业务表';

-- ----------------------------
-- Records of gen_table
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for gen_table_column
-- ----------------------------
DROP TABLE IF EXISTS `gen_table_column`;
CREATE TABLE `gen_table_column` (
  `column_id` bigint(20) NOT NULL COMMENT '编号',
  `table_id` bigint(20) DEFAULT NULL COMMENT '归属表编号',
  `column_name` varchar(200) DEFAULT NULL COMMENT '列名称',
  `column_comment` varchar(500) DEFAULT NULL COMMENT '列描述',
  `column_type` varchar(100) DEFAULT NULL COMMENT '列类型',
  `java_type` varchar(500) DEFAULT NULL COMMENT 'JAVA类型',
  `java_field` varchar(200) DEFAULT NULL COMMENT 'JAVA字段名',
  `is_pk` char(1) DEFAULT NULL COMMENT '是否主键（1是）',
  `is_increment` char(1) DEFAULT NULL COMMENT '是否自增（1是）',
  `is_required` char(1) DEFAULT NULL COMMENT '是否必填（1是）',
  `is_insert` char(1) DEFAULT NULL COMMENT '是否为插入字段（1是）',
  `is_edit` char(1) DEFAULT NULL COMMENT '是否编辑字段（1是）',
  `is_list` char(1) DEFAULT NULL COMMENT '是否列表字段（1是）',
  `is_query` char(1) DEFAULT NULL COMMENT '是否查询字段（1是）',
  `query_type` varchar(200) DEFAULT 'EQ' COMMENT '查询方式（等于、不等于、大于、小于、范围）',
  `html_type` varchar(200) DEFAULT NULL COMMENT '显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）',
  `dict_type` varchar(200) DEFAULT '' COMMENT '字典类型',
  `sort` int(11) DEFAULT NULL COMMENT '排序',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`column_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='代码生成业务表字段';

-- ----------------------------
-- Records of gen_table_column
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_client
-- ----------------------------
DROP TABLE IF EXISTS `sys_client`;
CREATE TABLE `sys_client` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `client_id` varchar(64) DEFAULT NULL COMMENT '客户端id',
  `client_key` varchar(32) DEFAULT NULL COMMENT '客户端key',
  `client_secret` varchar(255) DEFAULT NULL COMMENT '客户端秘钥',
  `grant_type` varchar(255) DEFAULT NULL COMMENT '授权类型',
  `device_type` varchar(32) DEFAULT NULL COMMENT '设备类型',
  `active_timeout` int(11) DEFAULT 1800 COMMENT 'token活跃超时时间',
  `timeout` int(11) DEFAULT 604800 COMMENT 'token固定超时',
  `status` char(1) DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_by` bigint(20) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统授权表';

-- ----------------------------
-- Records of sys_client
-- ----------------------------
BEGIN;
INSERT INTO `sys_client` (`id`, `client_id`, `client_key`, `client_secret`, `grant_type`, `device_type`, `active_timeout`, `timeout`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (1, 'e5cd7e4891bf95d1d19206ce24a7b32e', 'pc', 'pc123', 'password,social', 'pc', 1800, 604800, '0', 103, 1, '2025-02-13 11:56:36', 1, '2025-03-14 14:51:44', NULL, NULL);
INSERT INTO `sys_client` (`id`, `client_id`, `client_key`, `client_secret`, `grant_type`, `device_type`, `active_timeout`, `timeout`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (2, '428a8310cd442757ae699df5d894f051', 'app', 'app123', 'password,sms,social', 'android', 1800, 604800, '0', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-13 11:56:36', NULL, NULL);
INSERT INTO `sys_client` (`id`, `client_id`, `client_key`, `client_secret`, `grant_type`, `device_type`, `active_timeout`, `timeout`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (3, '', 'dd', 'dd', 'sms,xcx', 'android', 1800, 604800, '0', 103, 1, '2025-03-14 14:24:31', 1, '2025-03-14 14:33:54', 1, '2025-03-14 14:33:54');
INSERT INTO `sys_client` (`id`, `client_id`, `client_key`, `client_secret`, `grant_type`, `device_type`, `active_timeout`, `timeout`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (4, '1f6fxdu15xzd8fs32vh5ywwe004xwmqv', 'dd', 'dd', 'sms,email', 'ios', 1800, 604800, '0', 103, 1, '2025-03-14 14:34:09', NULL, '2025-03-14 14:52:59', 1, '2025-03-14 14:52:59');
COMMIT;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config` (
  `config_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `tenant_id` varchar(20) DEFAULT '000000' COMMENT '租户编号',
  `config_name` varchar(100) DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(1024) DEFAULT '' COMMENT '参数键值',
  `config_type` char(1) DEFAULT 'N' COMMENT '系统内置（Y是 N否）',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`config_id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='参数配置表';

-- ----------------------------
-- Records of sys_config
-- ----------------------------
BEGIN;
INSERT INTO `sys_config` (`config_id`, `tenant_id`, `config_name`, `config_key`, `config_value`, `config_type`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1, '000000', '主框架页-默认皮肤样式名称', 'sys.index.skinName', 'skin-blue', 'Y', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow');
INSERT INTO `sys_config` (`config_id`, `tenant_id`, `config_name`, `config_key`, `config_value`, `config_type`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (2, '000000', '用户管理-账号初始密码', 'sys.user.initPassword', '123456', 'Y', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '初始化密码 123456');
INSERT INTO `sys_config` (`config_id`, `tenant_id`, `config_name`, `config_key`, `config_value`, `config_type`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (3, '000000', '主框架页-侧边栏主题', 'sys.index.sideTheme', 'theme-dark', 'Y', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '深色主题theme-dark，浅色主题theme-light');
INSERT INTO `sys_config` (`config_id`, `tenant_id`, `config_name`, `config_key`, `config_value`, `config_type`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (5, '000000', '账号自助-是否开启用户注册功能', 'sys.account.registerUser', 'false', 'Y', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '是否开启注册用户功能（true开启，false关闭）');
INSERT INTO `sys_config` (`config_id`, `tenant_id`, `config_name`, `config_key`, `config_value`, `config_type`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (11, '000000', 'OSS预览列表资源开关', 'sys.oss.previewListResource', 'true', 'Y', 103, 1, '2025-02-13 11:56:36', NULL, NULL, 'true:开启, false:关闭');
INSERT INTO `sys_config` (`config_id`, `tenant_id`, `config_name`, `config_key`, `config_value`, `config_type`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (14, '000000', '上传文件类型', 'sys.oss.fileType', 'txt,md,doc,docx,pdf,xls,xlsx,ppt,pptx,txt,jpg,jpeg,png,gif,mp4,avi,mov,mp3,wav,m4a,csv', 'Y', 103, 1, '2025-03-14 11:28:54', 1, '2025-03-14 11:29:11', '');
INSERT INTO `sys_config` (`config_id`, `tenant_id`, `config_name`, `config_key`, `config_value`, `config_type`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (15, '000000', '上传图片类型', 'sys.oss.imgType', 'jpg,jpeg,png,gif,webp', 'Y', 103, 1, '2025-03-14 11:29:32', 1, '2025-03-14 11:29:32', '');
INSERT INTO `sys_config` (`config_id`, `tenant_id`, `config_name`, `config_key`, `config_value`, `config_type`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (16, '000000', '上传文件大小', 'sys.oss.fileSize', '20M', 'Y', 103, 1, '2025-03-14 11:29:48', 1, '2025-03-14 11:29:48', '');
INSERT INTO `sys_config` (`config_id`, `tenant_id`, `config_name`, `config_key`, `config_value`, `config_type`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (17, '000000', '上传图片大小', 'sys.oss.imgSize', '5M', 'Y', 103, 1, '2025-03-14 11:30:07', 1, '2025-03-14 11:30:07', '');
INSERT INTO `sys_config` (`config_id`, `tenant_id`, `config_name`, `config_key`, `config_value`, `config_type`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (18, '000000', '上传文件Url路径', 'sys.oss.urlPath', 'resource/upload|/upload', 'Y', 103, 1, '2025-03-14 12:48:44', 1, '2025-03-14 12:48:44', '\"http://ab.com/\" 或 \"/\" 或 \"resource/upload|http://ab.com/upload\"');
INSERT INTO `sys_config` (`config_id`, `tenant_id`, `config_name`, `config_key`, `config_value`, `config_type`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (19, '000000', '是否启动在线用户强制退出', 'sys.online.forceLogout', 'true', 'Y', 103, 1, '2025-03-26 20:07:24', 1, '2025-03-26 20:08:33', 'true开启,false关闭');
COMMIT;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `dept_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `tenant_id` varchar(20) DEFAULT '000000' COMMENT '租户编号',
  `parent_id` bigint(20) DEFAULT 0 COMMENT '父部门id',
  `ancestors` varchar(500) DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) DEFAULT '' COMMENT '部门名称',
  `dept_category` varchar(100) DEFAULT NULL COMMENT '部门类别编码',
  `order_num` int(4) DEFAULT 0 COMMENT '显示顺序',
  `leader` bigint(20) DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `status` char(1) DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_by` bigint(20) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dept_id`)
) ENGINE=InnoDB AUTO_INCREMENT=118 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='部门表';

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_dept` (`dept_id`, `tenant_id`, `parent_id`, `ancestors`, `dept_name`, `dept_category`, `order_num`, `leader`, `phone`, `email`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (100, '000000', 0, '0,', 'XXX科技', NULL, 0, NULL, '15888888888', 'xxx@qq.com', '0', 103, 1, '2025-02-13 11:56:36', NULL, '2025-02-27 22:21:39', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `tenant_id`, `parent_id`, `ancestors`, `dept_name`, `dept_category`, `order_num`, `leader`, `phone`, `email`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (101, '000000', 100, '0,100,', '济南总公司', NULL, 1, NULL, '15888888888', 'xxx@qq.com', '0', 103, 1, '2025-02-13 11:56:36', NULL, '2025-02-27 22:21:39', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `tenant_id`, `parent_id`, `ancestors`, `dept_name`, `dept_category`, `order_num`, `leader`, `phone`, `email`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (102, '000000', 100, '0,100,', '上海分公司', NULL, 2, NULL, '15888888888', 'xxx@qq.com', '0', 103, 1, '2025-02-13 11:56:36', NULL, '2025-02-27 22:21:40', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `tenant_id`, `parent_id`, `ancestors`, `dept_name`, `dept_category`, `order_num`, `leader`, `phone`, `email`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (103, '000000', 101, '0,100,101,', '研发部门', NULL, 1, 1, '15888888888', 'xxx@qq.com', '0', 103, 1, '2025-02-13 11:56:36', NULL, '2025-02-27 22:21:40', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `tenant_id`, `parent_id`, `ancestors`, `dept_name`, `dept_category`, `order_num`, `leader`, `phone`, `email`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (104, '000000', 101, '0,100,101,', '市场部门', NULL, 2, NULL, '15888888888', 'xxx@qq.com', '0', 103, 1, '2025-02-13 11:56:36', NULL, '2025-02-27 22:21:40', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `tenant_id`, `parent_id`, `ancestors`, `dept_name`, `dept_category`, `order_num`, `leader`, `phone`, `email`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (105, '000000', 101, '0,100,101,', '测试部门', NULL, 3, NULL, '15888888888', 'xxx@qq.com', '0', 103, 1, '2025-02-13 11:56:36', NULL, '2025-02-27 22:21:40', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `tenant_id`, `parent_id`, `ancestors`, `dept_name`, `dept_category`, `order_num`, `leader`, `phone`, `email`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (106, '000000', 101, '0,100,101,', '财务部门', NULL, 4, NULL, '15888888888', 'xxx@qq.com', '0', 103, 1, '2025-02-13 11:56:36', NULL, '2025-02-27 22:21:40', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `tenant_id`, `parent_id`, `ancestors`, `dept_name`, `dept_category`, `order_num`, `leader`, `phone`, `email`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (107, '000000', 101, '0,100,101,', '运维部门', NULL, 5, NULL, '15888888888', 'xxx@qq.com', '0', 103, 1, '2025-02-13 11:56:36', NULL, '2025-02-27 22:21:40', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `tenant_id`, `parent_id`, `ancestors`, `dept_name`, `dept_category`, `order_num`, `leader`, `phone`, `email`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (108, '000000', 102, '0,100,102,', '市场部门', NULL, 1, NULL, '15888888888', 'xxx@qq.com', '0', 103, 1, '2025-02-13 11:56:36', NULL, '2025-02-27 22:21:40', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `tenant_id`, `parent_id`, `ancestors`, `dept_name`, `dept_category`, `order_num`, `leader`, `phone`, `email`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (109, '000000', 102, '0,100,102,', '财务部门', NULL, 2, NULL, '15888888888', 'xxx@qq.com', '0', 103, 1, '2025-02-13 11:56:36', NULL, '2025-02-27 22:21:41', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `tenant_id`, `parent_id`, `ancestors`, `dept_name`, `dept_category`, `order_num`, `leader`, `phone`, `email`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (110, '000000', 100, '0,100,', '青岛分公司', '33', 3, 0, '15888888888', 'lxj521w@163.com', '0', 103, 1, '2025-02-27 22:25:56', 1, '2025-02-27 22:31:24', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `tenant_id`, `parent_id`, `ancestors`, `dept_name`, `dept_category`, `order_num`, `leader`, `phone`, `email`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (111, '000000', 110, '0,100,110,', '市场部门', '', 1, 0, '', '', '0', 103, 1, '2025-02-27 22:31:49', 1, '2025-02-27 22:31:50', NULL, '2025-02-27 22:31:52');
INSERT INTO `sys_dept` (`dept_id`, `tenant_id`, `parent_id`, `ancestors`, `dept_name`, `dept_category`, `order_num`, `leader`, `phone`, `email`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (112, '000000', 110, '0,100,110,', '打点', '', 3, 0, '', '', '0', 103, 1, '2025-02-27 22:32:43', 1, '2025-02-27 22:32:43', NULL, '2025-02-27 22:32:49');
INSERT INTO `sys_dept` (`dept_id`, `tenant_id`, `parent_id`, `ancestors`, `dept_name`, `dept_category`, `order_num`, `leader`, `phone`, `email`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (113, '000000', 110, '0,100,110,', '市场部门', '', 1, 0, '', '', '0', 103, 1, '2025-03-02 08:59:25', 1, '2025-03-02 08:59:26', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `tenant_id`, `parent_id`, `ancestors`, `dept_name`, `dept_category`, `order_num`, `leader`, `phone`, `email`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (117, '100006', 0, '', '秀杰智联', NULL, 0, 10, NULL, NULL, '0', NULL, 0, '2025-03-16 14:36:52', 0, '2025-03-16 14:36:52', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data` (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `tenant_id` varchar(20) DEFAULT '000000' COMMENT '租户编号',
  `dict_sort` int(4) DEFAULT 0 COMMENT '字典排序',
  `dict_label` varchar(100) DEFAULT '' COMMENT '字典标签',
  `dict_value` varchar(100) DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) DEFAULT '' COMMENT '字典类型',
  `css_class` varchar(100) DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
  `list_class` varchar(100) DEFAULT NULL COMMENT '表格回显样式',
  `is_default` char(1) DEFAULT 'N' COMMENT '是否默认（Y是 N否）',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`dict_code`)
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='字典数据表';

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1, '000000', 1, '男', '0', 'sys_user_sex', '', '', 'Y', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '性别男');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (2, '000000', 2, '女', '1', 'sys_user_sex', '', '', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '性别女');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (3, '000000', 3, '未知', '2', 'sys_user_sex', '', '', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '性别未知');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (4, '000000', 1, '显示', '0', 'sys_show_hide', '', 'primary', 'Y', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '显示菜单');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (5, '000000', 2, '隐藏', '1', 'sys_show_hide', '', 'danger', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '隐藏菜单');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (6, '000000', 1, '正常', '0', 'sys_normal_disable', '', 'primary', 'Y', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '正常状态');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (7, '000000', 2, '停用', '1', 'sys_normal_disable', '', 'danger', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '停用状态');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (12, '000000', 1, '是', 'Y', 'sys_yes_no', '', 'primary', 'Y', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '系统默认是');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (13, '000000', 2, '否', 'N', 'sys_yes_no', '', 'danger', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '系统默认否');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (14, '000000', 1, '通知', '1', 'sys_notice_type', '', 'warning', 'Y', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '通知');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (15, '000000', 2, '公告', '2', 'sys_notice_type', '', 'success', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '公告');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (16, '000000', 1, '正常', '0', 'sys_notice_status', '', 'primary', 'Y', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '正常状态');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (17, '000000', 2, '关闭', '1', 'sys_notice_status', '', 'danger', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '关闭状态');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (18, '000000', 1, '新增', '1', 'sys_oper_type', '', 'info', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '新增操作');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (19, '000000', 2, '修改', '2', 'sys_oper_type', '', 'info', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '修改操作');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (20, '000000', 3, '删除', '3', 'sys_oper_type', '', 'danger', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '删除操作');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (21, '000000', 4, '授权', '4', 'sys_oper_type', '', 'primary', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '授权操作');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (22, '000000', 5, '导出', '5', 'sys_oper_type', '', 'warning', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '导出操作');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (23, '000000', 6, '导入', '6', 'sys_oper_type', '', 'warning', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '导入操作');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (24, '000000', 7, '强退', '7', 'sys_oper_type', '', 'danger', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '强退操作');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (25, '000000', 8, '生成代码', '8', 'sys_oper_type', '', 'warning', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '生成操作');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (26, '000000', 9, '清空数据', '9', 'sys_oper_type', '', 'danger', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '清空操作');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (27, '000000', 1, '成功', '0', 'sys_common_status', '', 'primary', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '正常状态');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (28, '000000', 2, '失败', '1', 'sys_common_status', '', 'danger', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '停用状态');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (29, '000000', 99, '其他', '0', 'sys_oper_type', '', 'info', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '其他操作');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (30, '000000', 0, '密码认证', 'password', 'sys_grant_type', 'el-check-tag', 'default', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '密码认证');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (31, '000000', 0, '短信认证', 'sms', 'sys_grant_type', 'el-check-tag', 'default', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '短信认证');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (32, '000000', 0, '邮件认证', 'email', 'sys_grant_type', 'el-check-tag', 'default', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '邮件认证');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (33, '000000', 0, '小程序认证', 'xcx', 'sys_grant_type', 'el-check-tag', 'default', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '小程序认证');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (34, '000000', 0, '三方登录认证', 'social', 'sys_grant_type', 'el-check-tag', 'default', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '三方登录认证');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (35, '000000', 0, 'PC', 'pc', 'sys_device_type', '', 'default', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, 'PC');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (36, '000000', 0, '安卓', 'android', 'sys_device_type', '', 'default', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '安卓');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (37, '000000', 0, 'iOS', 'ios', 'sys_device_type', '', 'default', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, 'iOS');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (38, '000000', 0, '小程序', 'xcx', 'sys_device_type', '', 'default', 'N', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '小程序');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (40, '000000', 55, '22', '33', 'addd', '44', '22', '', 103, 1, '2025-03-05 12:59:35', NULL, '2025-03-05 12:59:35', '44');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (41, '000000', 1, '是', '0', 'sys_yes_no_num', '', 'primary', '', 103, 1, '2025-03-13 16:42:15', 1, '2025-03-13 16:42:50', '');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (42, '000000', 2, '否', '1', 'sys_yes_no_num', '', 'danger', '', 103, 1, '2025-03-13 16:42:28', 1, '2025-03-13 16:43:06', '');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (43, '000000', 1, '私有', '0', 'oss_access_policy', '', 'orange', '', 103, 1, '2025-03-13 16:55:30', NULL, '2025-03-13 16:55:30', '');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (44, '000000', 2, '公开', '1', 'oss_access_policy', '', 'primary', '', 103, 1, '2025-03-13 16:55:52', NULL, '2025-03-13 16:55:52', '');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (45, '000000', 3, '自定义', '2', 'oss_access_policy', '', '', '', 103, 1, '2025-03-13 16:56:17', NULL, '2025-03-13 16:56:17', '');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (46, '000000', 1, '生成成功', '0', 'sys_gen_status', '', 'success', '', 103, 1, '2025-03-18 14:27:50', NULL, '2025-03-18 14:27:50', '');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (47, '000000', 2, '未生成', '1', 'sys_gen_status', '', '', '', 103, 1, '2025-03-18 14:28:09', 1, '2025-03-20 11:21:25', '');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (48, '000000', 3, '生成失败', '2', 'sys_gen_status', '', 'danger', '', 103, 1, '2025-03-18 14:28:24', 1, '2025-03-18 14:28:43', '');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (49, '000000', 1, '默认', '0', 'sys_job_group', '', '', '', 101, 11, '2025-03-19 17:37:14', 11, '2025-03-19 17:37:41', '');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (50, '000000', 1, '系统', '1', 'sys_job_group', '', '', '', 101, 11, '2025-03-19 17:37:34', NULL, '2025-03-19 17:37:34', '');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (51, '000000', 2, '数据中心', '2', 'sys_job_group', '', '', '', 101, 11, '2025-03-19 17:37:55', NULL, '2025-03-19 17:37:55', '');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (52, '000000', 1, '多次执行', '1', 'sys_missfire_policy', '', '', '', 101, 11, '2025-03-19 17:38:55', NULL, '2025-03-19 17:38:55', '');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (53, '000000', 2, '执行一次', '2', 'sys_missfire_policy', '', '', '', 101, 11, '2025-03-19 17:39:04', NULL, '2025-03-19 17:39:04', '');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (54, '000000', 1, '允许', '0', 'sys_job_concurrent', '', '', '', 101, 11, '2025-03-20 11:25:29', NULL, '2025-03-20 11:25:29', '');
INSERT INTO `sys_dict_data` (`dict_code`, `tenant_id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (55, '000000', 2, '禁止', '1', 'sys_job_concurrent', '', '', '', 101, 11, '2025-03-20 11:25:41', NULL, '2025-03-20 11:25:41', '');
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type` (
  `dict_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `tenant_id` varchar(20) DEFAULT '000000' COMMENT '租户编号',
  `dict_name` varchar(100) DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) DEFAULT '' COMMENT '字典类型',
  `is_sys` char(1) DEFAULT '1' COMMENT '是否系统（0是 1否)',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`dict_id`),
  UNIQUE KEY `tenant_id` (`tenant_id`,`dict_type`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='字典类型表';

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict_type` (`dict_id`, `tenant_id`, `dict_name`, `dict_type`, `is_sys`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1, '000000', '用户性别', 'sys_user_sex', '1', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '用户性别列表');
INSERT INTO `sys_dict_type` (`dict_id`, `tenant_id`, `dict_name`, `dict_type`, `is_sys`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (2, '000000', '菜单状态', 'sys_show_hide', '1', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '菜单状态列表');
INSERT INTO `sys_dict_type` (`dict_id`, `tenant_id`, `dict_name`, `dict_type`, `is_sys`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (3, '000000', '系统开关', 'sys_normal_disable', '1', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '系统开关列表');
INSERT INTO `sys_dict_type` (`dict_id`, `tenant_id`, `dict_name`, `dict_type`, `is_sys`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (6, '000000', '系统是否', 'sys_yes_no', '1', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '系统是否列表');
INSERT INTO `sys_dict_type` (`dict_id`, `tenant_id`, `dict_name`, `dict_type`, `is_sys`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (7, '000000', '通知类型', 'sys_notice_type', '1', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '通知类型列表');
INSERT INTO `sys_dict_type` (`dict_id`, `tenant_id`, `dict_name`, `dict_type`, `is_sys`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (8, '000000', '通知状态', 'sys_notice_status', '1', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '通知状态列表');
INSERT INTO `sys_dict_type` (`dict_id`, `tenant_id`, `dict_name`, `dict_type`, `is_sys`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (9, '000000', '操作类型', 'sys_oper_type', '1', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '操作类型列表');
INSERT INTO `sys_dict_type` (`dict_id`, `tenant_id`, `dict_name`, `dict_type`, `is_sys`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (10, '000000', '系统状态', 'sys_common_status', '1', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '登录状态列表');
INSERT INTO `sys_dict_type` (`dict_id`, `tenant_id`, `dict_name`, `dict_type`, `is_sys`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (11, '000000', '授权类型', 'sys_grant_type', '1', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '认证授权类型');
INSERT INTO `sys_dict_type` (`dict_id`, `tenant_id`, `dict_name`, `dict_type`, `is_sys`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (12, '000000', '设备类型', 'sys_device_type', '1', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '客户端设备类型');
INSERT INTO `sys_dict_type` (`dict_id`, `tenant_id`, `dict_name`, `dict_type`, `is_sys`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (16, '000000', '是否默认', 'sys_yes_no_num', '1', 103, 1, '2025-03-05 12:32:12', 1, '2025-03-13 16:41:48', '是否默认 0是 1否');
INSERT INTO `sys_dict_type` (`dict_id`, `tenant_id`, `dict_name`, `dict_type`, `is_sys`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (17, '000000', 'oss权限桶类型', 'oss_access_policy', '1', 103, 1, '2025-03-13 16:54:51', NULL, '2025-03-13 16:54:51', '');
INSERT INTO `sys_dict_type` (`dict_id`, `tenant_id`, `dict_name`, `dict_type`, `is_sys`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (18, '000000', '代码生成状态', 'sys_gen_status', '1', 103, 1, '2025-03-18 14:27:03', NULL, '2025-03-18 14:27:03', '');
INSERT INTO `sys_dict_type` (`dict_id`, `tenant_id`, `dict_name`, `dict_type`, `is_sys`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (19, '000000', '任务分组', 'sys_job_group', '1', 101, 11, '2025-03-19 17:36:37', 1, '2025-03-26 09:47:20', '任务分组（0 默认 1系统 2数据中心）');
INSERT INTO `sys_dict_type` (`dict_id`, `tenant_id`, `dict_name`, `dict_type`, `is_sys`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (20, '000000', '任务计划执行策略', 'sys_missfire_policy', '1', 101, 11, '2025-03-19 17:38:30', 1, '2025-03-26 09:46:20', '计划执行策略（1多次执行 2执行一次）');
INSERT INTO `sys_dict_type` (`dict_id`, `tenant_id`, `dict_name`, `dict_type`, `is_sys`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (21, '000000', '任务是否并发执行', 'sys_job_concurrent', '1', 101, 11, '2025-03-20 11:24:58', 11, '2025-03-20 11:25:14', '是否并发执行（0允许 1禁止）');
COMMIT;

-- ----------------------------
-- Table structure for sys_gen_table
-- ----------------------------
DROP TABLE IF EXISTS `sys_gen_table`;
CREATE TABLE `sys_gen_table` (
  `table_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '生成ID',
  `gen_type` int(10) unsigned NOT NULL COMMENT '生成类型',
  `gen_template` int(11) DEFAULT 0 COMMENT '生成模板',
  `var_name` varchar(255) NOT NULL COMMENT '实体命名',
  `options` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '配置选项' CHECK (json_valid(`options`)),
  `db_name` varchar(128) DEFAULT NULL COMMENT '数据库名称',
  `table_name` varchar(255) NOT NULL COMMENT '主表名称',
  `table_comment` varchar(255) DEFAULT NULL COMMENT '主表注释',
  `dao_name` varchar(255) DEFAULT NULL COMMENT '主表dao模型',
  `master_columns` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '主表字段' CHECK (json_valid(`master_columns`)),
  `addon_name` varchar(128) DEFAULT NULL COMMENT '插件名称',
  `status` char(1) NOT NULL COMMENT '生成状态（0成功 1未开始）',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`table_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='代码生成业务表';

-- ----------------------------
-- Records of sys_gen_table
-- ----------------------------
BEGIN;
INSERT INTO `sys_gen_table` (`table_id`, `gen_type`, `gen_template`, `var_name`, `options`, `db_name`, `table_name`, `table_comment`, `dao_name`, `master_columns`, `addon_name`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (3, 0, 0, 'TestDemo', '{\"autoOps\":[\"genMenuPermissions\",\"runDao\",\"runService\",\"forcedCover\"],\"columnOps\":[\"edit\",\"del\",\"view\",\"status\",\"check\"],\"funcDict\":{\"labelColumn\":null,\"valueColumn\":null},\"headOps\":[\"add\",\"batchDel\",\"export\"],\"join\":[{\"alias\":\"dept\",\"columns\":[{\"align\":\"left\",\"dataType\":\"bigint\",\"dc\":\"部门id\",\"defaultValue\":0,\"dictType\":\"\",\"extra\":\"auto_increment\",\"formGridSpan\":1,\"formMode\":\"\",\"formRole\":\"\",\"goName\":\"DeptDeptId\",\"goType\":\"int64\",\"id\":1,\"index\":\"PRI\",\"isAllowNull\":\"NO\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"dept_id\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"bigint(20)\",\"tsName\":\"deptDeptId\",\"tsType\":\"number\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"varchar\",\"dc\":\"租户编号\",\"defaultValue\":\"\'000000\'\",\"dictType\":\"\",\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"\",\"formRole\":\"\",\"goName\":\"DeptTenantId\",\"goType\":\"string\",\"id\":2,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":20,\"name\":\"tenant_id\",\"queryWhere\":\"LIKE\",\"required\":false,\"sqlType\":\"varchar(20)\",\"tsName\":\"deptTenantId\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"bigint\",\"dc\":\"父部门id\",\"defaultValue\":0,\"dictType\":\"\",\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"\",\"formRole\":\"\",\"goName\":\"DeptParentId\",\"goType\":\"int64\",\"id\":3,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"parent_id\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"bigint(20)\",\"tsName\":\"deptParentId\",\"tsType\":\"number\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"varchar\",\"dc\":\"祖级列表\",\"defaultValue\":\"\'\'\",\"dictType\":\"\",\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"\",\"formRole\":\"\",\"goName\":\"DeptAncestors\",\"goType\":\"string\",\"id\":4,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":500,\"name\":\"ancestors\",\"queryWhere\":\"LIKE %...%\",\"required\":false,\"sqlType\":\"varchar(500)\",\"tsName\":\"deptAncestors\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"varchar\",\"dc\":\"部门名称\",\"defaultValue\":\"\'\'\",\"dictType\":\"\",\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"\",\"formRole\":\"\",\"goName\":\"DeptDeptName\",\"goType\":\"string\",\"id\":5,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":true,\"isList\":true,\"isQuery\":true,\"isSort\":false,\"length\":30,\"name\":\"dept_name\",\"queryWhere\":\"LIKE\",\"required\":false,\"sqlType\":\"varchar(30)\",\"tsName\":\"deptDeptName\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"varchar\",\"dc\":\"部门类别编码\",\"defaultValue\":\"NULL\",\"dictType\":\"\",\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"\",\"formRole\":\"\",\"goName\":\"DeptDeptCategory\",\"goType\":\"string\",\"id\":6,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":100,\"name\":\"dept_category\",\"queryWhere\":\"LIKE\",\"required\":false,\"sqlType\":\"varchar(100)\",\"tsName\":\"deptDeptCategory\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"int\",\"dc\":\"显示顺序\",\"defaultValue\":0,\"dictType\":\"\",\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"\",\"formRole\":\"\",\"goName\":\"DeptOrderNum\",\"goType\":\"int\",\"id\":7,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"order_num\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"int(4)\",\"tsName\":\"deptOrderNum\",\"tsType\":\"number\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"bigint\",\"dc\":\"负责人\",\"defaultValue\":0,\"dictType\":\"\",\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"\",\"formRole\":\"\",\"goName\":\"DeptLeader\",\"goType\":\"int64\",\"id\":8,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"leader\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"bigint(20)\",\"tsName\":\"deptLeader\",\"tsType\":\"number\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"varchar\",\"dc\":\"联系电话\",\"defaultValue\":\"NULL\",\"dictType\":\"\",\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"\",\"formRole\":\"\",\"goName\":\"DeptPhone\",\"goType\":\"string\",\"id\":9,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":11,\"name\":\"phone\",\"queryWhere\":\"LIKE\",\"required\":false,\"sqlType\":\"varchar(11)\",\"tsName\":\"deptPhone\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"varchar\",\"dc\":\"邮箱\",\"defaultValue\":\"NULL\",\"dictType\":\"\",\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"\",\"formRole\":\"\",\"goName\":\"DeptEmail\",\"goType\":\"string\",\"id\":10,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":50,\"name\":\"email\",\"queryWhere\":\"LIKE\",\"required\":false,\"sqlType\":\"varchar(50)\",\"tsName\":\"deptEmail\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"char\",\"dc\":\"部门状态（0正常 1停用）\",\"defaultValue\":\"\'0\'\",\"dictType\":\"\",\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"\",\"formRole\":\"\",\"goName\":\"DeptStatus\",\"goType\":\"string\",\"id\":11,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":1,\"name\":\"status\",\"queryWhere\":\"LIKE\",\"required\":false,\"sqlType\":\"char(1)\",\"tsName\":\"deptStatus\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"bigint\",\"dc\":\"创建部门\",\"defaultValue\":0,\"dictType\":\"\",\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"\",\"formRole\":\"\",\"goName\":\"DeptCreatedDept\",\"goType\":\"int64\",\"id\":12,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"created_dept\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"bigint(20)\",\"tsName\":\"deptCreatedDept\",\"tsType\":\"number\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"bigint\",\"dc\":\"创建者\",\"defaultValue\":0,\"dictType\":\"\",\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"\",\"formRole\":\"\",\"goName\":\"DeptCreatedBy\",\"goType\":\"int64\",\"id\":13,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"created_by\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"bigint(20)\",\"tsName\":\"deptCreatedBy\",\"tsType\":\"number\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"datetime\",\"dc\":\"创建时间\",\"defaultValue\":\"\",\"dictType\":\"\",\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"\",\"formRole\":\"\",\"goName\":\"DeptCreatedAt\",\"goType\":\"*gtime.Time\",\"id\":14,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"created_at\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"datetime\",\"tsName\":\"deptCreatedAt\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"bigint\",\"dc\":\"更新者\",\"defaultValue\":0,\"dictType\":\"\",\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"\",\"formRole\":\"\",\"goName\":\"DeptUpdatedBy\",\"goType\":\"int64\",\"id\":15,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"updated_by\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"bigint(20)\",\"tsName\":\"deptUpdatedBy\",\"tsType\":\"number\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"datetime\",\"dc\":\"更新时间\",\"defaultValue\":\"\",\"dictType\":\"\",\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"\",\"formRole\":\"\",\"goName\":\"DeptUpdatedAt\",\"goType\":\"*gtime.Time\",\"id\":16,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"updated_at\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"datetime\",\"tsName\":\"deptUpdatedAt\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"bigint\",\"dc\":\"删除人\",\"defaultValue\":0,\"dictType\":\"\",\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"\",\"formRole\":\"\",\"goName\":\"DeptDeletedBy\",\"goType\":\"int64\",\"id\":17,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"deleted_by\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"bigint(20)\",\"tsName\":\"deptDeletedBy\",\"tsType\":\"number\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"datetime\",\"dc\":\"删除时间\",\"defaultValue\":\"\",\"dictType\":\"\",\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"\",\"formRole\":\"\",\"goName\":\"DeptDeletedAt\",\"goType\":\"*gtime.Time\",\"id\":18,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"deleted_at\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"datetime\",\"tsName\":\"deptDeletedAt\",\"tsType\":\"string\",\"unique\":false,\"width\":null}],\"daoName\":\"SysDept\",\"dbName\":\"default\",\"field\":\"dept_id\",\"linkMode\":1,\"linkTable\":\"sys_dept\",\"masterField\":\"dept_id\",\"masterTableName\":\"test_demo\",\"uuid\":\"1742806299441001\"}],\"menu\":{\"icon\":\"ant-design:home-twotone\",\"pid\":0,\"sort\":1},\"presetStep\":{\"formGridCols\":1},\"tree\":{\"styleType\":1,\"titleColumn\":\"\"}}', 'default', 'test_demo', '测试单表', 'TestDemo', '[{\"align\":\"left\",\"dataType\":\"bigint\",\"dc\":\"主键\",\"defaultValue\":0,\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"InputNumber\",\"formRole\":\"none\",\"goName\":\"Id\",\"goType\":\"int64\",\"id\":1,\"index\":\"PRI\",\"isAllowNull\":\"NO\",\"isEdit\":false,\"isExport\":true,\"isList\":true,\"isQuery\":true,\"isSort\":true,\"length\":0,\"name\":\"id\",\"queryWhere\":\"=\",\"required\":true,\"sqlType\":\"bigint(20)\",\"tsName\":\"id\",\"tsType\":\"number\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"varchar\",\"dc\":\"租户编号\",\"defaultValue\":\"\'000000\'\",\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"Input\",\"formRole\":\"none\",\"goName\":\"TenantId\",\"goType\":\"string\",\"id\":2,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":20,\"name\":\"tenant_id\",\"queryWhere\":\"LIKE\",\"required\":false,\"sqlType\":\"varchar(20)\",\"tsName\":\"tenantId\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"bigint\",\"dc\":\"部门id\",\"defaultValue\":0,\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"InputNumber\",\"formRole\":\"num\",\"goName\":\"DeptId\",\"goType\":\"int64\",\"id\":3,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":true,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"dept_id\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"bigint(20)\",\"tsName\":\"deptId\",\"tsType\":\"number\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"bigint\",\"dc\":\"用户id\",\"defaultValue\":0,\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"InputNumber\",\"formRole\":\"none\",\"goName\":\"UserId\",\"goType\":\"int64\",\"id\":4,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":true,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"user_id\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"bigint(20)\",\"tsName\":\"userId\",\"tsType\":\"number\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"int\",\"dc\":\"排序号\",\"defaultValue\":0,\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"InputNumber\",\"formRole\":\"none\",\"goName\":\"OrderNum\",\"goType\":\"int\",\"id\":5,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":true,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"order_num\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"int(11)\",\"tsName\":\"orderNum\",\"tsType\":\"number\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"varchar\",\"dc\":\"key键\",\"defaultValue\":\"NULL\",\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"Input\",\"formRole\":\"none\",\"goName\":\"TestKey\",\"goType\":\"string\",\"id\":6,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":true,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":255,\"name\":\"test_key\",\"queryWhere\":\"LIKE\",\"required\":false,\"sqlType\":\"varchar(255)\",\"tsName\":\"testKey\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"varchar\",\"dc\":\"值\",\"defaultValue\":\"NULL\",\"dictType\":\"sys_yes_no_num\",\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"Select\",\"formRole\":\"none\",\"goName\":\"Value\",\"goType\":\"string\",\"id\":7,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":true,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":255,\"name\":\"value\",\"queryWhere\":\"LIKE\",\"required\":false,\"sqlType\":\"varchar(255)\",\"tsName\":\"value\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"int\",\"dc\":\"版本\",\"defaultValue\":0,\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"InputNumber\",\"formRole\":\"none\",\"goName\":\"Version\",\"goType\":\"int\",\"id\":8,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":true,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"version\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"int(11)\",\"tsName\":\"version\",\"tsType\":\"number\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"bigint\",\"dc\":\"创建部门\",\"defaultValue\":0,\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"InputNumber\",\"formRole\":\"none\",\"goName\":\"CreatedDept\",\"goType\":\"int64\",\"id\":9,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"created_dept\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"bigint(20)\",\"tsName\":\"createdDept\",\"tsType\":\"number\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"datetime\",\"dc\":\"创建时间\",\"defaultValue\":\"\",\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"DateRange\",\"formRole\":\"none\",\"goName\":\"CreatedAt\",\"goType\":\"*gtime.Time\",\"id\":10,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":true,\"isList\":true,\"isQuery\":true,\"isSort\":false,\"length\":0,\"name\":\"created_at\",\"queryWhere\":\"BETWEEN\",\"required\":false,\"sqlType\":\"datetime\",\"tsName\":\"createdAt\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"bigint\",\"dc\":\"创建者\",\"defaultValue\":0,\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"InputNumber\",\"formRole\":\"none\",\"goName\":\"CreatedBy\",\"goType\":\"int64\",\"id\":11,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"created_by\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"bigint(20)\",\"tsName\":\"createdBy\",\"tsType\":\"number\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"datetime\",\"dc\":\"更新时间\",\"defaultValue\":\"\",\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"Time\",\"formRole\":\"none\",\"goName\":\"UpdatedAt\",\"goType\":\"*gtime.Time\",\"id\":12,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"updated_at\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"datetime\",\"tsName\":\"updatedAt\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"bigint\",\"dc\":\"更新者\",\"defaultValue\":0,\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"InputNumber\",\"formRole\":\"none\",\"goName\":\"UpdatedBy\",\"goType\":\"int64\",\"id\":13,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"updated_by\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"bigint(20)\",\"tsName\":\"updatedBy\",\"tsType\":\"number\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"bigint\",\"dc\":\"删除人\",\"defaultValue\":0,\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"InputNumber\",\"formRole\":\"none\",\"goName\":\"DeletedBy\",\"goType\":\"int64\",\"id\":14,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"deleted_by\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"bigint(20)\",\"tsName\":\"deletedBy\",\"tsType\":\"number\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"datetime\",\"dc\":\"删除时间\",\"defaultValue\":\"\",\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"Time\",\"formRole\":\"none\",\"goName\":\"DeletedAt\",\"goType\":\"*gtime.Time\",\"id\":15,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"deleted_at\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"datetime\",\"tsName\":\"deletedAt\",\"tsType\":\"string\",\"unique\":false,\"width\":null}]', '', '0', 103, 1, '2025-03-21 16:20:01', 1, '2025-03-26 17:50:05');
INSERT INTO `sys_gen_table` (`table_id`, `gen_type`, `gen_template`, `var_name`, `options`, `db_name`, `table_name`, `table_comment`, `dao_name`, `master_columns`, `addon_name`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (4, 0, 0, 'SysUserOnline', '{\"autoOps\":[\"genMenuPermissions\",\"runDao\",\"runService\",\"forcedCover\"],\"columnOps\":[\"edit\",\"del\",\"view\",\"check\"],\"funcDict\":{\"labelColumn\":null,\"valueColumn\":null},\"headOps\":[\"add\",\"batchDel\",\"export\"],\"join\":[],\"menu\":{\"icon\":\"ant-design:home-twotone\",\"pid\":0,\"sort\":1},\"presetStep\":{\"formGridCols\":1},\"tree\":{\"styleType\":1,\"titleColumn\":\"\"}}', 'default', 'sys_user_online', '用户在线列表', 'SysUserOnline', '[{\"align\":\"left\",\"dataType\":\"bigint\",\"dc\":\"访问ID\",\"defaultValue\":0,\"dictType\":null,\"extra\":\"auto_increment\",\"formGridSpan\":1,\"formMode\":\"InputNumber\",\"formRole\":\"none\",\"goName\":\"OnlineId\",\"goType\":\"int64\",\"id\":1,\"index\":\"PRI\",\"isAllowNull\":\"NO\",\"isEdit\":false,\"isExport\":true,\"isList\":true,\"isQuery\":true,\"isSort\":false,\"length\":0,\"name\":\"online_id\",\"queryWhere\":\"=\",\"required\":true,\"sqlType\":\"bigint(20)\",\"tsName\":\"onlineId\",\"tsType\":\"number\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"varchar\",\"dc\":\"租户编号\",\"defaultValue\":\"\'000000\'\",\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"Input\",\"formRole\":\"none\",\"goName\":\"TenantId\",\"goType\":\"string\",\"id\":2,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":20,\"name\":\"tenant_id\",\"queryWhere\":\"LIKE\",\"required\":false,\"sqlType\":\"varchar(20)\",\"tsName\":\"tenantId\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"varchar\",\"dc\":\"UUID\",\"defaultValue\":\"\'\'\",\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"Input\",\"formRole\":\"email\",\"goName\":\"Uuid\",\"goType\":\"string\",\"id\":3,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":true,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":50,\"name\":\"uuid\",\"queryWhere\":\"LIKE\",\"required\":true,\"sqlType\":\"varchar(50)\",\"tsName\":\"uuid\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"varchar\",\"dc\":\"用户账号\",\"defaultValue\":\"\'\'\",\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"Input\",\"formRole\":\"none\",\"goName\":\"UserName\",\"goType\":\"string\",\"id\":4,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":true,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":50,\"name\":\"user_name\",\"queryWhere\":\"LIKE\",\"required\":true,\"sqlType\":\"varchar(50)\",\"tsName\":\"userName\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"varchar\",\"dc\":\"客户端\",\"defaultValue\":\"\'\'\",\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"Input\",\"formRole\":\"none\",\"goName\":\"ClientKey\",\"goType\":\"string\",\"id\":5,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":true,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":32,\"name\":\"client_key\",\"queryWhere\":\"LIKE\",\"required\":false,\"sqlType\":\"varchar(32)\",\"tsName\":\"clientKey\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"varchar\",\"dc\":\"设备类型\",\"defaultValue\":\"\'\'\",\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"Input\",\"formRole\":\"none\",\"goName\":\"DeviceType\",\"goType\":\"string\",\"id\":6,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":true,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":32,\"name\":\"device_type\",\"queryWhere\":\"LIKE\",\"required\":false,\"sqlType\":\"varchar(32)\",\"tsName\":\"deviceType\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"varchar\",\"dc\":\"登录IP地址\",\"defaultValue\":\"\'\'\",\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"Input\",\"formRole\":\"none\",\"goName\":\"Ipaddr\",\"goType\":\"string\",\"id\":7,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":true,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":128,\"name\":\"ipaddr\",\"queryWhere\":\"LIKE\",\"required\":false,\"sqlType\":\"varchar(128)\",\"tsName\":\"ipaddr\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"varchar\",\"dc\":\"登录地点\",\"defaultValue\":\"\'\'\",\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"Input\",\"formRole\":\"none\",\"goName\":\"LoginLocation\",\"goType\":\"string\",\"id\":8,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":true,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":255,\"name\":\"login_location\",\"queryWhere\":\"LIKE\",\"required\":false,\"sqlType\":\"varchar(255)\",\"tsName\":\"loginLocation\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"varchar\",\"dc\":\"浏览器类型\",\"defaultValue\":\"\'\'\",\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"Input\",\"formRole\":\"none\",\"goName\":\"Browser\",\"goType\":\"string\",\"id\":9,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":true,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":50,\"name\":\"browser\",\"queryWhere\":\"LIKE\",\"required\":false,\"sqlType\":\"varchar(50)\",\"tsName\":\"browser\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"varchar\",\"dc\":\"操作系统\",\"defaultValue\":\"\'\'\",\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"Input\",\"formRole\":\"none\",\"goName\":\"Os\",\"goType\":\"string\",\"id\":10,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":true,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":50,\"name\":\"os\",\"queryWhere\":\"LIKE\",\"required\":false,\"sqlType\":\"varchar(50)\",\"tsName\":\"os\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"text\",\"dc\":\"Token\",\"defaultValue\":\"\'\'\",\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"InputEditor\",\"formRole\":\"none\",\"goName\":\"Token\",\"goType\":\"string\",\"id\":11,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":true,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":65535,\"name\":\"token\",\"queryWhere\":\"LIKE %...%\",\"required\":false,\"sqlType\":\"text\",\"tsName\":\"token\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"datetime\",\"dc\":\"访问时间\",\"defaultValue\":\"\",\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"Time\",\"formRole\":\"none\",\"goName\":\"LoginTime\",\"goType\":\"*gtime.Time\",\"id\":12,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":true,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"login_time\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"datetime\",\"tsName\":\"loginTime\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"datetime\",\"dc\":\"过期时间\",\"defaultValue\":\"\",\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"Time\",\"formRole\":\"none\",\"goName\":\"ExpireTime\",\"goType\":\"*gtime.Time\",\"id\":13,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":true,\"isExport\":true,\"isList\":true,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"expire_time\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"datetime\",\"tsName\":\"expireTime\",\"tsType\":\"string\",\"unique\":false,\"width\":null},{\"align\":\"left\",\"dataType\":\"datetime\",\"dc\":\"删除时间\",\"defaultValue\":\"\",\"dictType\":null,\"extra\":\"\",\"formGridSpan\":1,\"formMode\":\"Time\",\"formRole\":\"none\",\"goName\":\"DeletedAt\",\"goType\":\"*gtime.Time\",\"id\":14,\"index\":\"\",\"isAllowNull\":\"YES\",\"isEdit\":false,\"isExport\":false,\"isList\":false,\"isQuery\":false,\"isSort\":false,\"length\":0,\"name\":\"deleted_at\",\"queryWhere\":\"=\",\"required\":false,\"sqlType\":\"datetime\",\"tsName\":\"deletedAt\",\"tsType\":\"string\",\"unique\":false,\"width\":null}]', '', '0', 103, 1, '2025-03-27 14:50:04', 1, '2025-03-28 08:37:11');
COMMIT;

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job` (
  `job_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `job_name` varchar(64) NOT NULL DEFAULT '' COMMENT '任务名称',
  `job_params` varchar(200) DEFAULT '' COMMENT '参数',
  `job_group` varchar(64) NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
  `invoke_target` varchar(500) NOT NULL COMMENT '调用目标字符串',
  `cron_expression` varchar(255) DEFAULT '' COMMENT 'cron执行表达式',
  `misfire_policy` tinyint(4) DEFAULT 1 COMMENT '计划执行策略（1多次执行 2执行一次）',
  `concurrent` tinyint(4) DEFAULT 1 COMMENT '是否并发执行（0允许 1禁止）',
  `status` char(1) DEFAULT '1' COMMENT '状态（0正常 1暂停）',
  `remark` varchar(500) DEFAULT '' COMMENT '备注信息',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_by` bigint(20) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`job_id`,`job_name`,`job_group`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='定时任务调度表';

-- ----------------------------
-- Records of sys_job
-- ----------------------------
BEGIN;
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `remark`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (1, 'test1', '', '1', 'test.abb', '0 * * * ', 1, 1, '1', 'Lorem elit', 101, 11, '2025-03-19 12:00:33', 11, '2025-03-20 11:28:07', 11, '2025-03-20 11:28:07');
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `remark`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (2, '韩成', 'Duis aute', '0', '0', ' sss ', 1, 0, '0', '0', 101, 11, '2025-03-19 13:13:29', 11, '2025-03-20 11:28:13', 11, '2025-03-20 11:28:13');
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `remark`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (3, 'test2', '收拾收拾', '0', 'test2', 'test2', 2, 0, '1', 'test2', 101, 11, '2025-03-20 10:29:48', 11, '2025-03-20 11:28:13', 11, '2025-03-20 11:28:13');
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `remark`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (4, 'test31', '1', '0', 'clear_data', '3/10 * * * * ?', 1, 0, '1', 'test31', 101, 11, '2025-03-20 10:33:23', 1, '2025-03-25 16:40:38', NULL, NULL);
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `remark`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (5, 'test11', '', '0', 'test', '3/5 * * * 3/5 ?', 1, 0, '1', 'test11', 101, 11, '2025-03-20 11:28:57', 1, '2025-03-25 19:07:56', NULL, NULL);
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `remark`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (6, 'test121', 'test121', '0', 'test121', '3/5 * 2,3 * * ? *', 1, 0, '1', 'test121', 101, 11, '2025-03-20 11:29:16', 11, '2025-03-24 13:01:50', NULL, NULL);
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `remark`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (7, 'test2', '1', '1', 'test', '3/5 * * * * ?', 1, 1, '1', 'test2', 103, 1, '2025-03-25 16:53:33', 1, '2025-03-25 16:54:56', 1, '2025-03-25 16:54:56');
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `remark`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (8, 'test1', '2', '2', 'test', '* * * ? * 2/1', 2, 1, '1', 'test2', 103, 1, '2025-03-26 09:54:08', 1, '2025-03-26 09:54:08', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_logininfor
-- ----------------------------
DROP TABLE IF EXISTS `sys_logininfor`;
CREATE TABLE `sys_logininfor` (
  `info_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `tenant_id` varchar(20) DEFAULT '000000' COMMENT '租户编号',
  `user_name` varchar(50) DEFAULT '' COMMENT '用户账号',
  `client_key` varchar(32) DEFAULT '' COMMENT '客户端',
  `device_type` varchar(32) DEFAULT '' COMMENT '设备类型',
  `ipaddr` varchar(128) DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) DEFAULT '' COMMENT '操作系统',
  `status` char(1) DEFAULT '0' COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) DEFAULT '' COMMENT '提示消息',
  `login_time` datetime DEFAULT NULL COMMENT '访问时间',
  PRIMARY KEY (`info_id`),
  KEY `idx_sys_logininfor_s` (`status`),
  KEY `idx_sys_logininfor_lt` (`login_time`)
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统访问记录';

-- ----------------------------
-- Records of sys_logininfor
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `menu_id` bigint NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `menu_name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `parent_id` bigint DEFAULT '0' COMMENT '父菜单ID',
  `level` int DEFAULT '1' COMMENT '关系树等级',
  `tree` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '关系树',
  `order_num` int DEFAULT '0' COMMENT '显示顺序',
  `path` varchar(200) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '路由地址',
  `component` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '组件路径',
  `query_param` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由参数',
  `is_frame` int DEFAULT '1' COMMENT '是否为外链（0是 1否）',
  `is_cache` int DEFAULT '0' COMMENT '是否缓存（0缓存 1不缓存）',
  `menu_type` char(1) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '菜单类型（M目录 C菜单 F按钮）',
  `visible` char(1) COLLATE utf8mb4_general_ci DEFAULT '0' COMMENT '显示状态（0显示 1隐藏）',
  `status` char(1) COLLATE utf8mb4_general_ci DEFAULT '0' COMMENT '菜单状态（0正常 1停用）',
  `perms` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '权限标识',
  `icon` varchar(100) COLLATE utf8mb4_general_ci DEFAULT '#' COMMENT '菜单图标',
  `created_dept` bigint DEFAULT NULL COMMENT '创建部门',
  `created_by` bigint DEFAULT NULL COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1647 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='菜单权限表';

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1, '系统管理', 0, 1, NULL, 2, 'system', NULL, '', 1, 0, 'M', '0', '0', '', 'ant-design:setting-filled', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '系统管理目录');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (2, '系统监控', 0, 1, NULL, 3, 'monitor', '', '', 1, 0, 'M', '0', '0', '', 'ant-design:bar-chart-outlined', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:58:26', '系统监控目录');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (3, '系统工具', 0, 1, NULL, 4, 'tool', '', '', 1, 0, 'M', '0', '0', '', 'ant-design:sliders-twotone', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 14:00:52', '系统工具目录');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (4, '首页', 0, 1, NULL, 1, 'analytics', '/dashboard/analytics/index', '', 1, 0, 'C', '0', '0', '', 'lucide:layout-dashboard', 103, 1, '2025-02-13 11:56:36', NULL, NULL, 'RuoYi-Vue-Plus官网地址');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (6, '租户管理', 0, 1, NULL, 2, 'tenant', '', '', 1, 0, 'M', '0', '0', '', 'ant-design:gold-twotone', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:56:09', '租户管理目录');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (100, '用户管理', 1, 1, NULL, 1, 'user', 'system/user/index', '', 1, 0, 'C', '0', '0', 'system:user:list', 'ant-design:user-outlined', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:31:59', '用户管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (101, '角色管理', 1, 1, NULL, 2, 'role', 'system/role/index', '', 1, 0, 'C', '0', '0', 'system:role:list', 'ant-design:team-outlined', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:37:39', '角色管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (102, '菜单管理', 1, 1, NULL, 3, 'menu', 'system/menu/index', '', 1, 0, 'C', '0', '0', 'system:menu:list', 'ant-design:menu-outlined', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:39:02', '菜单管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (103, '部门管理', 1, 1, NULL, 4, 'dept', 'system/dept/index', '', 1, 0, 'C', '0', '0', 'system:dept:list', 'ant-design:apartment-outlined', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:42:16', '部门管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (104, '岗位管理', 1, 1, NULL, 5, 'post', 'system/post/index', '', 1, 0, 'C', '0', '0', 'system:post:list', 'ant-design:solution-outlined', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:43:25', '岗位管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (105, '字典管理', 1, 1, NULL, 6, 'dict', 'system/dict/index', '', 1, 0, 'C', '0', '0', 'system:dict:list', 'ant-design:book-twotone', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:48:51', '字典管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (106, '参数设置', 1, 1, NULL, 7, 'config', 'system/config/index', '', 1, 0, 'C', '0', '0', 'system:config:list', 'ant-design:tool-outlined', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:50:51', '参数设置菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (107, '通知公告', 1, 1, NULL, 8, 'notice', 'system/notice/index', '', 1, 0, 'C', '0', '0', 'system:notice:list', 'ant-design:notification-outlined', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:51:37', '通知公告菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (108, '日志管理', 1, 1, NULL, 9, 'log', '', '', 1, 0, 'M', '0', '0', '', 'ant-design:file-text-outlined', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:53:07', '日志管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (109, '在线用户', 2, 1, NULL, 1, 'online', 'monitor/online/index', '', 1, 0, 'C', '0', '0', 'monitor:online:list', 'ant-design:api-filled', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:58:44', '在线用户菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (113, '缓存监控', 2, 1, NULL, 5, 'cache', 'monitor/cache/index', '', 1, 0, 'C', '0', '0', 'monitor:cache:list', 'ant-design:codepen-outlined', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:59:18', '缓存监控菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (115, '代码生成', 3, 1, NULL, 2, 'gen', 'tool/gen/index', '', 1, 0, 'C', '0', '0', 'tool:gen:list', 'ant-design:code-outlined', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 14:01:09', '代码生成菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (117, '服务监控', 2, 1, NULL, 5, 'Admin', 'monitor/server/index', '', 1, 0, 'C', '0', '0', 'monitor:server:list', 'ant-design:bar-chart-outlined', 103, 1, '2025-02-13 11:56:36', 1, '2025-03-17 10:47:57', 'Admin监控菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (118, '文件管理', 1, 1, NULL, 10, 'oss', 'system/oss/index', '', 1, 0, 'C', '0', '0', 'system:oss:list', 'ant-design:folder-outlined', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:55:15', '文件管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (120, '任务调度中心', 2, 1, NULL, 6, 'snailjob', 'system/job/index', '', 1, 0, 'C', '0', '0', 'system:job:list', 'ant-design:align-center-outlined', 103, 1, '2025-02-13 11:56:36', 11, '2025-03-19 17:08:27', 'SnailJob控制台菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (121, '租户管理', 6, 1, NULL, 1, 'tenant', 'system/tenant/index', '', 1, 0, 'C', '0', '0', 'system:tenant:list', 'ant-design:gold-twotone', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:56:31', '租户管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (122, '租户套餐管理', 6, 1, NULL, 2, 'tenantPackage', 'system/tenantPackage/index', '', 1, 0, 'C', '0', '0', 'system:tenantPackage:list', 'ant-design:unordered-list-outlined', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:58:05', '租户套餐管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (123, '客户端管理', 1, 1, NULL, 11, 'client', 'system/client/index', '', 1, 0, 'C', '0', '0', 'system:client:list', 'ant-design:group-outlined', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:54:54', '客户端管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (500, '操作日志', 108, 1, NULL, 1, 'operlog', 'system/operlog/index', '', 1, 0, 'C', '0', '0', 'monitor:operlog:list', 'ant-design:file-text-outlined', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:53:18', '操作日志菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (501, '登录日志', 108, 1, NULL, 2, 'logininfor', 'system/logininfor/index', '', 1, 0, 'C', '0', '0', 'monitor:logininfor:list', 'ant-design:file-text-outlined', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:53:29', '登录日志菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1001, '用户查询', 100, 1, NULL, 1, '', '', '', 1, 0, 'F', '0', '0', 'system:user:query', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1002, '用户新增', 100, 1, NULL, 2, '', '', '', 1, 0, 'F', '0', '0', 'system:user:add', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1003, '用户修改', 100, 1, NULL, 3, '', '', '', 1, 0, 'F', '0', '0', 'system:user:edit', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1004, '用户删除', 100, 1, NULL, 4, '', '', '', 1, 0, 'F', '0', '0', 'system:user:remove', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1005, '用户导出', 100, 1, NULL, 5, '', '', '', 1, 0, 'F', '0', '0', 'system:user:export', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1006, '用户导入', 100, 1, NULL, 6, '', '', '', 1, 0, 'F', '0', '0', 'system:user:import', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1007, '重置密码', 100, 1, NULL, 7, '', '', '', 1, 0, 'F', '0', '0', 'system:user:resetPwd', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1008, '角色查询', 101, 1, NULL, 1, '', '', '', 1, 0, 'F', '0', '0', 'system:role:query', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1009, '角色新增', 101, 1, NULL, 2, '', '', '', 1, 0, 'F', '0', '0', 'system:role:add', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1010, '角色修改', 101, 1, NULL, 3, '', '', '', 1, 0, 'F', '0', '0', 'system:role:edit', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1011, '角色删除', 101, 1, NULL, 4, '', '', '', 1, 0, 'F', '0', '0', 'system:role:remove', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1012, '角色导出', 101, 1, NULL, 5, '', '', '', 1, 0, 'F', '0', '0', 'system:role:export', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1013, '菜单查询', 102, 1, NULL, 1, '', '', '', 1, 0, 'F', '0', '0', 'system:menu:query', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1014, '菜单新增', 102, 1, NULL, 2, '', '', '', 1, 0, 'F', '0', '0', 'system:menu:add', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1015, '菜单修改', 102, 1, NULL, 3, '', '', '', 1, 0, 'F', '0', '0', 'system:menu:edit', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1016, '菜单删除', 102, 1, NULL, 4, '', '', '', 1, 0, 'F', '0', '0', 'system:menu:remove', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1017, '部门查询', 103, 1, NULL, 1, '', '', '', 1, 0, 'F', '0', '0', 'system:dept:query', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1018, '部门新增', 103, 1, NULL, 2, '', '', '', 1, 0, 'F', '0', '0', 'system:dept:add', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1019, '部门修改', 103, 1, NULL, 3, '', '', '', 1, 0, 'F', '0', '0', 'system:dept:edit', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1020, '部门删除', 103, 1, NULL, 4, '', '', '', 1, 0, 'F', '0', '0', 'system:dept:remove', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1021, '岗位查询', 104, 1, NULL, 1, '', '', '', 1, 0, 'F', '0', '0', 'system:post:query', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1022, '岗位新增', 104, 1, NULL, 2, '', '', '', 1, 0, 'F', '0', '0', 'system:post:add', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1023, '岗位修改', 104, 1, NULL, 3, '', '', '', 1, 0, 'F', '0', '0', 'system:post:edit', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1024, '岗位删除', 104, 1, NULL, 4, '', '', '', 1, 0, 'F', '0', '0', 'system:post:remove', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1025, '岗位导出', 104, 1, NULL, 5, '', '', '', 1, 0, 'F', '0', '0', 'system:post:export', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1026, '字典查询', 105, 1, NULL, 1, '#', '', '', 1, 0, 'F', '0', '0', 'system:dict:query', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1027, '字典新增', 105, 1, NULL, 2, '#', '', '', 1, 0, 'F', '0', '0', 'system:dict:add', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1028, '字典修改', 105, 1, NULL, 3, '#', '', '', 1, 0, 'F', '0', '0', 'system:dict:edit', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1029, '字典删除', 105, 1, NULL, 4, '#', '', '', 1, 0, 'F', '0', '0', 'system:dict:remove', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1030, '字典导出', 105, 1, NULL, 5, '#', '', '', 1, 0, 'F', '0', '0', 'system:dict:export', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1031, '参数查询', 106, 1, NULL, 1, '#', '', '', 1, 0, 'F', '0', '0', 'system:config:query', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1032, '参数新增', 106, 1, NULL, 2, '#', '', '', 1, 0, 'F', '0', '0', 'system:config:add', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1033, '参数修改', 106, 1, NULL, 3, '#', '', '', 1, 0, 'F', '0', '0', 'system:config:edit', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1034, '参数删除', 106, 1, NULL, 4, '#', '', '', 1, 0, 'F', '0', '0', 'system:config:remove', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1035, '参数导出', 106, 1, NULL, 5, '#', '', '', 1, 0, 'F', '0', '0', 'system:config:export', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1036, '公告查询', 107, 1, NULL, 1, '#', '', '', 1, 0, 'F', '0', '0', 'system:notice:query', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1037, '公告新增', 107, 1, NULL, 2, '#', '', '', 1, 0, 'F', '0', '0', 'system:notice:add', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1038, '公告修改', 107, 1, NULL, 3, '#', '', '', 1, 0, 'F', '0', '0', 'system:notice:edit', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1039, '公告删除', 107, 1, NULL, 4, '#', '', '', 1, 0, 'F', '0', '0', 'system:notice:remove', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1040, '操作查询', 500, 1, NULL, 1, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:operlog:query', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1041, '操作删除', 500, 1, NULL, 2, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:operlog:remove', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1042, '日志导出', 500, 1, NULL, 4, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:operlog:export', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1043, '登录查询', 501, 1, NULL, 1, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:logininfor:query', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1044, '登录删除', 501, 1, NULL, 2, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:logininfor:remove', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1045, '日志导出', 501, 1, NULL, 3, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:logininfor:export', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1046, '在线查询', 109, 1, NULL, 1, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:online:query', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1047, '批量强退', 109, 1, NULL, 2, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:online:batchLogout', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1048, '单条强退', 109, 1, NULL, 3, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:online:forceLogout', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1050, '账户解锁', 501, 1, NULL, 4, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:logininfor:unlock', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1055, '生成查询', 115, 1, NULL, 1, '#', '', '', 1, 0, 'F', '0', '0', 'tool:gen:query', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1056, '生成修改', 115, 1, NULL, 2, '#', '', '', 1, 0, 'F', '0', '0', 'tool:gen:edit', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1057, '生成删除', 115, 1, NULL, 3, '#', '', '', 1, 0, 'F', '0', '0', 'tool:gen:remove', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1058, '导入代码', 115, 1, NULL, 2, '#', '', '', 1, 0, 'F', '0', '0', 'tool:gen:import', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1059, '预览代码', 115, 1, NULL, 4, '#', '', '', 1, 0, 'F', '0', '0', 'tool:gen:preview', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1060, '生成代码', 115, 1, NULL, 5, '#', '', '', 1, 0, 'F', '0', '0', 'tool:gen:code', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1061, '客户端管理查询', 123, 1, NULL, 1, '#', '', '', 1, 0, 'F', '0', '0', 'system:client:query', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1062, '客户端管理新增', 123, 1, NULL, 2, '#', '', '', 1, 0, 'F', '0', '0', 'system:client:add', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1063, '客户端管理修改', 123, 1, NULL, 3, '#', '', '', 1, 0, 'F', '0', '0', 'system:client:edit', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1064, '客户端管理删除', 123, 1, NULL, 4, '#', '', '', 1, 0, 'F', '0', '0', 'system:client:remove', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1065, '客户端管理导出', 123, 1, NULL, 5, '#', '', '', 1, 0, 'F', '0', '0', 'system:client:export', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1201, '任务详情', 120, 1, NULL, 1, '#', '', '', 1, 0, 'F', '0', '0', 'system:job:view', '#', 103, 1, '2025-03-26 09:03:17', NULL, NULL, '查看');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1202, '任务新增', 120, 1, NULL, 2, '#', '', '', 1, 0, 'F', '0', '0', 'system:job:add', '#', 103, 1, '2025-03-26 09:09:51', NULL, NULL, '新增');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1203, '任务修改', 120, 1, NULL, 3, '#', '', '', 1, 0, 'F', '0', '0', 'system:job:update', '#', 103, 1, '2025-03-26 09:09:51', NULL, NULL, '修改');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1204, '任务状态修改', 120, 1, NULL, 4, '#', '', '', 1, 0, 'F', '0', '0', 'system:job:status', '#', 103, 1, '2025-03-26 09:09:51', NULL, NULL, '修改状态');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1205, '任务执行一次', 120, 1, NULL, 5, '#', '', '', 1, 0, 'F', '0', '0', 'system:job:exec', '#', 103, 1, '2025-03-26 09:09:51', NULL, NULL, '执行一次');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1206, '任务删除', 120, 1, NULL, 6, '#', '', '', 1, 0, 'F', '0', '0', 'system:job:delete', '#', 103, 1, '2025-03-26 09:09:51', NULL, NULL, '删除');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1600, '文件查询', 118, 1, NULL, 1, '#', '', '', 1, 0, 'F', '0', '0', 'system:oss:query', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1601, '文件上传', 118, 1, NULL, 2, '#', '', '', 1, 0, 'F', '0', '0', 'system:oss:upload', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1602, '文件下载', 118, 1, NULL, 3, '#', '', '', 1, 0, 'F', '0', '0', 'system:oss:download', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1603, '文件删除', 118, 1, NULL, 4, '#', '', '', 1, 0, 'F', '0', '0', 'system:oss:remove', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1606, '租户查询', 121, 1, NULL, 1, '#', '', '', 1, 0, 'F', '0', '0', 'system:tenant:query', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1607, '租户新增', 121, 1, NULL, 2, '#', '', '', 1, 0, 'F', '0', '0', 'system:tenant:add', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1608, '租户修改', 121, 1, NULL, 3, '#', '', '', 1, 0, 'F', '0', '0', 'system:tenant:edit', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1609, '租户删除', 121, 1, NULL, 4, '#', '', '', 1, 0, 'F', '0', '0', 'system:tenant:remove', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1610, '租户导出', 121, 1, NULL, 5, '#', '', '', 1, 0, 'F', '0', '0', 'system:tenant:export', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1611, '租户套餐查询', 122, 1, NULL, 1, '#', '', '', 1, 0, 'F', '0', '0', 'system:tenantPackage:query', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1612, '租户套餐新增', 122, 1, NULL, 2, '#', '', '', 1, 0, 'F', '0', '0', 'system:tenantPackage:add', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1613, '租户套餐修改', 122, 1, NULL, 3, '#', '', '', 1, 0, 'F', '0', '0', 'system:tenantPackage:edit', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1614, '租户套餐删除', 122, 1, NULL, 4, '#', '', '', 1, 0, 'F', '0', '0', 'system:tenantPackage:remove', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1615, '租户套餐导出', 122, 1, NULL, 5, '#', '', '', 1, 0, 'F', '0', '0', 'system:tenantPackage:export', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1620, '配置列表', 118, 1, NULL, 5, '#', '', '', 1, 0, 'F', '0', '0', 'system:ossConfig:list', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1621, '配置添加', 118, 1, NULL, 6, '#', '', '', 1, 0, 'F', '0', '0', 'system:ossConfig:add', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1622, '配置编辑', 118, 1, NULL, 6, '#', '', '', 1, 0, 'F', '0', '0', 'system:ossConfig:edit', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1623, '配置删除', 118, 1, NULL, 6, '#', '', '', 1, 0, 'F', '0', '0', 'system:ossConfig:remove', '#', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1624, '字典数据', 1, 1, NULL, 6, 'dict-data/:id', 'system/dict-data/index', '', 1, 0, 'C', '1', '0', 'system:dict:list', 'ant-design:book-outlined', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-25 13:48:35', '字典数据菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1626, 'OSS配置', 1, 1, NULL, 10, 'oss-config', 'system/oss-config/index', '', 1, 1, 'C', '1', '0', 'system:oss-config:list', 'ant-design:file-outlined', 103, 1, '2025-03-13 16:19:44', 1, '2025-03-13 16:19:44', '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1627, '生成配置', 3, 1, NULL, 20, 'gen-develop', 'tool/gen/develop', '', 1, 0, 'C', '1', '0', 'tool:gen:develop', 'ant-design:code-outlined', 103, 1, '2025-03-20 11:43:45', 1, '2025-03-26 15:35:17', '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1628, '代码生成的测试单表', 0, 1, '', 10, 'testDemo', 'LAYOUT', '', 1, 0, 'M', '0', '0', '', 'ant-design:home-twotone', 103, 1, '2025-03-23 14:43:49', 1, '2025-03-28 08:58:52', '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1629, '测试单表列表', 1628, 2, 'tr_1628 ', 10, 'index', 'gen/testDemo/index', '', 1, 0, 'C', '0', '0', 'gen:testDemo:list', 'ant-design:home-twotone', 103, 1, '2025-03-23 14:43:49', 1, '2025-03-25 17:29:41', '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1630, '测试单表详情', 1629, 3, 'tr_1628 tr_1629 ', 10, '', '', '', 1, 0, 'F', '0', '0', 'gen:testDemo:view', '', 103, 1, '2025-03-23 14:43:49', 1, '2025-03-25 17:29:57', '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1631, '编辑/新增测试单表', 1629, 3, 'tr_1628 tr_1629 ', 10, '', '', '', 1, 0, 'F', '0', '0', 'gen:testDemo:edit', '', 103, 1, '2025-03-23 14:43:49', 1, '2025-03-25 17:30:11', '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1632, '删除测试单表', 1629, 3, 'tr_1628 tr_1629 ', 10, '', '', '', 1, 0, 'F', '0', '0', 'gen:testDemo:delete', '', 103, 1, '2025-03-23 14:43:49', 1, '2025-03-25 17:30:24', '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1633, '导出测试单表', 1629, 3, 'tr_1628 tr_1629 ', 10, '', '', '', 1, 0, 'F', '0', '0', 'gen:testDemo:export', '', 103, 1, '2025-03-23 14:43:49', 1, '2025-03-25 17:30:38', '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `level`, `tree`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1646, '同步菜单', 121, 1, NULL, 1, '', '', '', 1, 0, 'F', '0', '0', 'system:tenant:syncmenu', '', 103, 1, '2026-03-08 11:06:09', 1, '2026-03-08 11:09:00', '');
COMMIT;

-- ----------------------------
-- Table structure for sys_notice
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice`;
CREATE TABLE `sys_notice` (
  `notice_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '公告ID',
  `tenant_id` varchar(20) DEFAULT '000000' COMMENT '租户编号',
  `notice_title` varchar(50) NOT NULL COMMENT '公告标题',
  `notice_type` char(1) NOT NULL COMMENT '公告类型（1通知 2公告）',
  `notice_content` longblob DEFAULT NULL COMMENT '公告内容',
  `status` char(1) DEFAULT '0' COMMENT '公告状态（0正常 1关闭）',
  `notice_range` tinyint(1) NOT NULL DEFAULT 1 COMMENT '通知范围（1全员 2指定机构 3指定用户）',
  `dept_ids` text   COMMENT '通知机构ID列表JSON',
  `user_ids` text   COMMENT '通知用户ID列表JSON',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`notice_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='通知公告表';

-- ----------------------------
-- Records of sys_notice
-- ----------------------------
BEGIN;
INSERT INTO `sys_notice` (`notice_id`, `tenant_id`, `notice_title`, `notice_type`, `notice_content`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1, '000000', '温馨提醒：2018-07-01 新版本发布啦', '2', 0xE696B0E78988E69CACE58685E5AEB9, '0', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '管理员');
INSERT INTO `sys_notice` (`notice_id`, `tenant_id`, `notice_title`, `notice_type`, `notice_content`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (2, '000000', '维护通知：2018-07-01 系统凌晨维护', '1', 0xE7BBB4E68AA4E58685E5AEB9, '0', 103, 1, '2025-02-13 11:56:36', NULL, NULL, '管理员');
COMMIT;



-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log` (
  `oper_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `tenant_id` varchar(20) DEFAULT '000000' COMMENT '租户编号',
  `title` varchar(50) DEFAULT '' COMMENT '模块标题',
  `business_type` int(2) DEFAULT 0 COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `method` varchar(100) DEFAULT '' COMMENT '方法名称',
  `request_method` varchar(10) DEFAULT '' COMMENT '请求方式',
  `operator_type` int(1) DEFAULT 0 COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
  `oper_name` varchar(50) DEFAULT '' COMMENT '操作人员',
  `dept_name` varchar(50) DEFAULT '' COMMENT '部门名称',
  `oper_url` varchar(255) DEFAULT '' COMMENT '请求URL',
  `oper_ip` varchar(128) DEFAULT '' COMMENT '主机地址',
  `oper_location` varchar(255) DEFAULT '' COMMENT '操作地点',
  `oper_param` varchar(4000) DEFAULT '' COMMENT '请求参数',
  `json_result` varchar(4000) DEFAULT '' COMMENT '返回参数',
  `status` int(1) DEFAULT 0 COMMENT '操作状态（0正常 1异常）',
  `error_msg` varchar(4000) DEFAULT '' COMMENT '错误消息',
  `oper_time` datetime DEFAULT NULL COMMENT '操作时间',
  `cost_time` bigint(20) DEFAULT 0 COMMENT '消耗时间',
  PRIMARY KEY (`oper_id`),
  KEY `idx_sys_oper_log_bt` (`business_type`),
  KEY `idx_sys_oper_log_s` (`status`),
  KEY `idx_sys_oper_log_ot` (`oper_time`)
) ENGINE=InnoDB AUTO_INCREMENT=10699 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='操作日志记录';

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_oss
-- ----------------------------
DROP TABLE IF EXISTS `sys_oss`;
CREATE TABLE `sys_oss` (
  `oss_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '对象存储主键',
  `tenant_id` varchar(20) DEFAULT '000000' COMMENT '租户编号',
  `file_name` varchar(255) NOT NULL DEFAULT '' COMMENT '文件名',
  `original_name` varchar(255) NOT NULL DEFAULT '' COMMENT '原名',
  `file_suffix` varchar(10) NOT NULL DEFAULT '' COMMENT '文件后缀名',
  `path` varchar(255) DEFAULT NULL COMMENT '存储路径',
  `url` varchar(500) NOT NULL COMMENT 'URL地址',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `service` varchar(20) NOT NULL DEFAULT 'minio' COMMENT '服务商',
  `md5` varchar(64)  NULL DEFAULT NULL COMMENT '文件MD5',
  `file_size` int NULL DEFAULT NULL COMMENT '文件大小',
  `file_crc16` int COMMENT '文件Crc16',
  `file_sum` int COMMENT '文件校验和',
  PRIMARY KEY (`oss_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='OSS对象存储表';

-- ----------------------------
-- Records of sys_oss
-- ----------------------------
BEGIN;
INSERT INTO `sys_oss` (`oss_id`, `tenant_id`, `file_name`, `original_name`, `file_suffix`, `path`, `url`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `service`) VALUES (6, '000000', 'd8fqgc4iypz4hjdjh4.png', 'image.png', '.png', 'resource/upload/2025-03-14/d8fqgc4iypz4hjdjh4.png', '/upload/2025-03-14/d8fqgc4iypz4hjdjh4.png', 103, '2025-03-14 13:17:25', 1, '2025-03-14 13:17:25', 1, 'local');
COMMIT;

-- ----------------------------
-- Table structure for sys_oss_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_oss_config`;
CREATE TABLE `sys_oss_config` (
  `oss_config_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` varchar(20) DEFAULT '000000' COMMENT '租户编号',
  `config_key` varchar(20) NOT NULL DEFAULT '' COMMENT '配置key',
  `access_key` varchar(255) DEFAULT '' COMMENT 'accessKey',
  `secret_key` varchar(255) DEFAULT '' COMMENT '秘钥',
  `bucket_name` varchar(255) DEFAULT '' COMMENT '桶名称',
  `prefix` varchar(255) DEFAULT '' COMMENT '前缀',
  `endpoint` varchar(255) DEFAULT '' COMMENT '访问站点',
  `domain` varchar(255) DEFAULT '' COMMENT '自定义域名',
  `is_https` char(1) DEFAULT 'N' COMMENT '是否https（Y=是,N=否）',
  `region` varchar(255) DEFAULT '' COMMENT '域',
  `access_policy` char(1) NOT NULL DEFAULT '1' COMMENT '桶权限类型(0=private 1=public 2=custom)',
  `status` char(1) DEFAULT '1' COMMENT '是否默认（0=是,1=否）',
  `ext1` varchar(255) DEFAULT '' COMMENT '扩展字段',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`oss_config_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='对象存储配置表';

-- ----------------------------
-- Records of sys_oss_config
-- ----------------------------
BEGIN;
INSERT INTO `sys_oss_config` (`oss_config_id`, `tenant_id`, `config_key`, `access_key`, `secret_key`, `bucket_name`, `prefix`, `endpoint`, `domain`, `is_https`, `region`, `access_policy`, `status`, `ext1`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (1, '000000', 'minio', 'ruoyi', 'ruoyi123', 'ruoyi', '', '127.0.0.1:9000', '', 'N', '', '1', '0', '', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-13 11:56:36', NULL);
INSERT INTO `sys_oss_config` (`oss_config_id`, `tenant_id`, `config_key`, `access_key`, `secret_key`, `bucket_name`, `prefix`, `endpoint`, `domain`, `is_https`, `region`, `access_policy`, `status`, `ext1`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (2, '000000', 'qiniu', 'XXXXXXXXXXXXXXX', 'XXXXXXXXXXXXXXX', 'ruoyi', '', 's3-cn-north-1.qiniucs.com', '', 'N', '', '1', '1', '', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-13 11:56:36', NULL);
INSERT INTO `sys_oss_config` (`oss_config_id`, `tenant_id`, `config_key`, `access_key`, `secret_key`, `bucket_name`, `prefix`, `endpoint`, `domain`, `is_https`, `region`, `access_policy`, `status`, `ext1`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (3, '000000', 'aliyun', 'XXXXXXXXXXXXXXX', 'XXXXXXXXXXXXXXX', 'ruoyi', '', 'oss-cn-beijing.aliyuncs.com', '', 'N', '', '1', '1', '', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-13 11:56:36', NULL);
INSERT INTO `sys_oss_config` (`oss_config_id`, `tenant_id`, `config_key`, `access_key`, `secret_key`, `bucket_name`, `prefix`, `endpoint`, `domain`, `is_https`, `region`, `access_policy`, `status`, `ext1`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (4, '000000', 'qcloud', 'XXXXXXXXXXXXXXX', 'XXXXXXXXXXXXXXX', 'ruoyi-1240000000', '', 'cos.ap-beijing.myqcloud.com', '', 'N', 'ap-beijing', '1', '1', '', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-13 11:56:36', NULL);
INSERT INTO `sys_oss_config` (`oss_config_id`, `tenant_id`, `config_key`, `access_key`, `secret_key`, `bucket_name`, `prefix`, `endpoint`, `domain`, `is_https`, `region`, `access_policy`, `status`, `ext1`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (5, '000000', 'image', 'ruoyi', 'ruoyi123', 'ruoyi', 'image', '127.0.0.1:9000', '', 'N', '', '1', '1', '', 103, 1, '2025-02-13 11:56:36', 1, '2025-02-13 11:56:36', NULL);
INSERT INTO `sys_oss_config` (`oss_config_id`, `tenant_id`, `config_key`, `access_key`, `secret_key`, `bucket_name`, `prefix`, `endpoint`, `domain`, `is_https`, `region`, `access_policy`, `status`, `ext1`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remark`) VALUES (7, '000000', 'test', 'test222', 'test222', 'test222', 'stet', 'test22', '222', 'Y', 'test', '2', '1', '', 103, 1, '2025-03-13 17:11:44', 1, '2025-03-13 17:17:53', 'sttt');
COMMIT;

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
  `post_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `tenant_id` varchar(20) DEFAULT '000000' COMMENT '租户编号',
  `dept_id` bigint(20) NOT NULL COMMENT '部门id',
  `post_code` varchar(64) NOT NULL COMMENT '岗位编码',
  `post_category` varchar(100) DEFAULT NULL COMMENT '岗位类别编码',
  `post_name` varchar(50) NOT NULL COMMENT '岗位名称',
  `post_sort` int(4) NOT NULL COMMENT '显示顺序',
  `status` char(1) NOT NULL COMMENT '状态（0正常 1停用）',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_by` bigint(20) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`post_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='岗位信息表';

-- ----------------------------
-- Records of sys_post
-- ----------------------------
BEGIN;
INSERT INTO `sys_post` (`post_id`, `tenant_id`, `dept_id`, `post_code`, `post_category`, `post_name`, `post_sort`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (1, '000000', 103, 'ceo', NULL, '董事长', 1, '0', 103, 1, '2025-02-13 11:56:36', NULL, NULL, NULL, NULL, '');
INSERT INTO `sys_post` (`post_id`, `tenant_id`, `dept_id`, `post_code`, `post_category`, `post_name`, `post_sort`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (2, '000000', 100, 'se', NULL, '项目经理', 2, '0', 103, 1, '2025-02-13 11:56:36', NULL, NULL, NULL, NULL, '');
INSERT INTO `sys_post` (`post_id`, `tenant_id`, `dept_id`, `post_code`, `post_category`, `post_name`, `post_sort`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (3, '000000', 100, 'hr', NULL, '人力资源', 3, '0', 103, 1, '2025-02-13 11:56:36', NULL, NULL, NULL, NULL, '');
INSERT INTO `sys_post` (`post_id`, `tenant_id`, `dept_id`, `post_code`, `post_category`, `post_name`, `post_sort`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (4, '000000', 100, 'user', NULL, '普通员工', 4, '0', 103, 1, '2025-02-13 11:56:36', NULL, NULL, NULL, NULL, '');
INSERT INTO `sys_post` (`post_id`, `tenant_id`, `dept_id`, `post_code`, `post_category`, `post_name`, `post_sort`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (5, '000000', 106, '333', '', '财务员', 10, '0', 103, 1, '2025-03-02 23:02:11', 1, '2025-03-02 23:07:04', 1, '2025-03-02 23:07:04', '333');
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `tenant_id` varchar(20) DEFAULT '000000' COMMENT '租户编号',
  `role_name` varchar(30) NOT NULL COMMENT '角色名称',
  `role_key` varchar(100) NOT NULL COMMENT '角色权限字符串',
  `role_sort` int(4) NOT NULL COMMENT '显示顺序',
  `data_scope` char(1) DEFAULT '1' COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `menu_check_strictly` tinyint(1) DEFAULT 1 COMMENT '菜单树选择项是否关联显示',
  `dept_check_strictly` tinyint(1) DEFAULT 1 COMMENT '部门树选择项是否关联显示',
  `status` char(1) NOT NULL COMMENT '角色状态（0正常 1停用）',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_by` bigint(20) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色信息表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` (`role_id`, `tenant_id`, `role_name`, `role_key`, `role_sort`, `data_scope`, `menu_check_strictly`, `dept_check_strictly`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (1, '000000', '超级管理员', 'superadmin', 1, '1', 1, 1, '0', 103, 1, '2025-02-13 11:56:36', NULL, NULL, NULL, NULL, '超级管理员');
INSERT INTO `sys_role` (`role_id`, `tenant_id`, `role_name`, `role_key`, `role_sort`, `data_scope`, `menu_check_strictly`, `dept_check_strictly`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (3, '000000', '本部门及以下', 'test1', 3, '4', 0, 0, '0', 0, 1, '2025-02-26 19:32:45', 1, '2025-03-26 15:20:20', NULL, NULL, '');
INSERT INTO `sys_role` (`role_id`, `tenant_id`, `role_name`, `role_key`, `role_sort`, `data_scope`, `menu_check_strictly`, `dept_check_strictly`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (4, '000000', '仅本人', 'test2', 4, '5', 0, 0, '0', 0, 1, '2025-02-26 19:32:48', 1, '2025-02-27 00:54:35', NULL, NULL, '4,100,1001,1002,1003,1004,1005,1006,1007');
INSERT INTO `sys_role` (`role_id`, `tenant_id`, `role_name`, `role_key`, `role_sort`, `data_scope`, `menu_check_strictly`, `dept_check_strictly`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (5, '000000', 'roletest2', 'roletest', 0, '', 0, 0, '0', 0, 1, '2025-02-26 19:32:52', 1, '2025-02-26 11:31:28', NULL, '2025-02-26 11:31:41', 'roletest22');
INSERT INTO `sys_role` (`role_id`, `tenant_id`, `role_name`, `role_key`, `role_sort`, `data_scope`, `menu_check_strictly`, `dept_check_strictly`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (6, '000000', 'test233', 'ddd', 3, '', 0, 0, '0', 103, 1, '2025-02-26 11:31:59', 1, '2025-02-26 11:31:59', NULL, '2025-02-26 11:32:09', 'test3');
INSERT INTO `sys_role` (`role_id`, `tenant_id`, `role_name`, `role_key`, `role_sort`, `data_scope`, `menu_check_strictly`, `dept_check_strictly`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (7, '000000', 'roletest2', 'roletest', 5, '1', 1, 0, '0', 103, 1, '2025-02-26 15:40:38', 1, '2025-02-27 00:54:40', NULL, NULL, '');
INSERT INTO `sys_role` (`role_id`, `tenant_id`, `role_name`, `role_key`, `role_sort`, `data_scope`, `menu_check_strictly`, `dept_check_strictly`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (8, '000000', 'eee', 'eee', 6, '2', 1, 1, '0', 0, 1, '2025-02-26 19:32:54', 1, '2025-02-27 03:22:02', NULL, NULL, '');
COMMIT;

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept` (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `dept_id` bigint(20) NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`role_id`,`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色和部门关联表';

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_dept` (`role_id`, `dept_id`) VALUES (8, 106);
INSERT INTO `sys_role_dept` (`role_id`, `dept_id`) VALUES (8, 107);
COMMIT;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色和菜单关联表';

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 5);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 100);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 101);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 102);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 103);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 104);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 105);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 106);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 107);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 108);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 118);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 123);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 500);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 501);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1001);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1002);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1003);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1004);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1005);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1006);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1007);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1008);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1009);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1010);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1011);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1012);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1013);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1014);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1015);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1016);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1017);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1018);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1019);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1020);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1021);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1022);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1023);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1024);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1025);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1026);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1027);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1028);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1029);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1030);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1031);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1032);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1033);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1034);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1035);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1036);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1037);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1038);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1039);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1040);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1041);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1042);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1043);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1045);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1050);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1061);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1062);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1064);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1065);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1500);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1501);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1502);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1503);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1504);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1505);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1506);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1507);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1508);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1509);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1510);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1511);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1600);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1601);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1602);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1603);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1620);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1621);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1622);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1623);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1624);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1626);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1628);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1629);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1630);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1631);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1632);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (3, 1633);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (4, 5);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (4, 1500);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (4, 1501);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (4, 1502);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (4, 1503);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (4, 1504);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (4, 1505);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (4, 1506);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (4, 1507);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (4, 1508);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (4, 1509);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (4, 1510);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (4, 1511);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (8, 1001);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (8, 1005);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (8, 1007);
COMMIT;

-- ----------------------------
-- Table structure for sys_social
-- ----------------------------
DROP TABLE IF EXISTS `sys_social`;
CREATE TABLE `sys_social` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `tenant_id` varchar(20) DEFAULT '000000' COMMENT '租户id',
  `auth_id` varchar(255) NOT NULL COMMENT '平台+平台唯一id',
  `source` varchar(255) NOT NULL COMMENT '用户来源',
  `open_id` varchar(255) DEFAULT NULL COMMENT '平台编号唯一id',
  `user_name` varchar(30) NOT NULL COMMENT '登录账号',
  `nick_name` varchar(30) DEFAULT '' COMMENT '用户昵称',
  `email` varchar(255) DEFAULT '' COMMENT '用户邮箱',
  `avatar` varchar(500) DEFAULT '' COMMENT '头像地址',
  `access_token` varchar(255) NOT NULL COMMENT '用户的授权令牌',
  `expire_in` int(11) DEFAULT NULL COMMENT '用户的授权令牌的有效期，部分平台可能没有',
  `refresh_token` varchar(255) DEFAULT NULL COMMENT '刷新令牌，部分平台可能没有',
  `access_code` varchar(255) DEFAULT NULL COMMENT '平台的授权信息，部分平台可能没有',
  `union_id` varchar(255) DEFAULT NULL COMMENT '用户的 unionid',
  `scope` varchar(255) DEFAULT NULL COMMENT '授予的权限，部分平台可能没有',
  `token_type` varchar(255) DEFAULT NULL COMMENT '个别平台的授权信息，部分平台可能没有',
  `id_token` varchar(2000) DEFAULT NULL COMMENT 'id token，部分平台可能没有',
  `mac_algorithm` varchar(255) DEFAULT NULL COMMENT '小米平台用户的附带属性，部分平台可能没有',
  `mac_key` varchar(255) DEFAULT NULL COMMENT '小米平台用户的附带属性，部分平台可能没有',
  `code` varchar(255) DEFAULT NULL COMMENT '用户的授权code，部分平台可能没有',
  `oauth_token` varchar(255) DEFAULT NULL COMMENT 'Twitter平台用户的附带属性，部分平台可能没有',
  `oauth_token_secret` varchar(255) DEFAULT NULL COMMENT 'Twitter平台用户的附带属性，部分平台可能没有',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_by` bigint(20) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='社会化关系表';

-- ----------------------------
-- Records of sys_social
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_tenant
-- ----------------------------
DROP TABLE IF EXISTS `sys_tenant`;
CREATE TABLE `sys_tenant` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` varchar(20) NOT NULL COMMENT '租户编号',
  `contact_user_name` varchar(20) DEFAULT NULL COMMENT '联系人',
  `contact_phone` varchar(20) DEFAULT NULL COMMENT '联系电话',
  `company_name` varchar(50) DEFAULT NULL COMMENT '企业名称',
  `license_number` varchar(30) DEFAULT NULL COMMENT '统一社会信用代码',
  `address` varchar(200) DEFAULT NULL COMMENT '地址',
  `intro` varchar(200) DEFAULT NULL COMMENT '企业简介',
  `domain` varchar(200) DEFAULT NULL COMMENT '域名',
  `remark` varchar(200) DEFAULT NULL COMMENT '备注',
  `package_id` bigint(20) DEFAULT NULL COMMENT '租户套餐编号',
  `admin_role_id` bigint(20) DEFAULT NULL COMMENT '管理员角色ID',
  `admin_dept_id` bigint(20) DEFAULT NULL COMMENT '管理员部门ID',
  `admin_user_id` bigint(20) DEFAULT NULL COMMENT '管理员用户ID',
  `expire_time` datetime DEFAULT NULL COMMENT '过期时间',
  `account_count` int(11) DEFAULT -1 COMMENT '用户数量（-1不限制）',
  `status` char(1) DEFAULT '0' COMMENT '租户状态（0正常 1停用）',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_by` bigint(20) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='租户表';

-- ----------------------------
-- Records of sys_tenant
-- ----------------------------
BEGIN;
INSERT INTO `sys_tenant` (
  `id`, `tenant_id`, `contact_user_name`, `contact_phone`, `company_name`, 
  `license_number`, `address`, `intro`, `domain`, `remark`, 
  `package_id`, `admin_role_id`, `admin_dept_id`, `admin_user_id`, `expire_time`, 
  `account_count`, `status`, `created_dept`, `created_by`, `created_at`, 
  `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (
    1, '000000', '管理组', '15888888888', 'XXX有限公司', 
    NULL, NULL, '多租户通用后台管理管理系统', NULL, '备注', 
    NULL, 1, 103, 1, NULL, 
    -1, '0', 103, 1, '2025-02-13 11:56:36', 
    NULL, NULL, NULL, NULL);

INSERT INTO `sys_tenant` (`id`, `tenant_id`, `contact_user_name`, `contact_phone`, `company_name`, `license_number`, `address`, `intro`, `domain`, `remark`, `package_id`, `admin_role_id`, `admin_dept_id`, `admin_user_id`, `expire_time`, `account_count`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (6, '100006', '李先生', '18600000000', '秀杰智联', '3332233', '地址', '企业介绍', 'xiujiezhilian.com', '备注', 9, 117, 10, 1, '2026-03-16 00:00:00', -1, '0', 103, 1, '2025-03-16 14:36:52', 1, '2025-03-16 14:37:11', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_tenant_package
-- ----------------------------
DROP TABLE IF EXISTS `sys_tenant_package`;
CREATE TABLE `sys_tenant_package` (
  `package_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '租户套餐id',
  `package_name` varchar(20) DEFAULT NULL COMMENT '套餐名称',
  `menu_ids` varchar(3000) DEFAULT NULL COMMENT '关联菜单id',
  `remark` varchar(200) DEFAULT NULL COMMENT '备注',
  `menu_check_strictly` tinyint(1) DEFAULT 1 COMMENT '菜单树选择项是否关联显示',
  `status` char(1) DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_by` bigint(20) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`package_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='租户套餐表';

-- ----------------------------
-- Records of sys_tenant_package
-- ----------------------------
BEGIN;
INSERT INTO `sys_tenant_package` (`package_id`, `package_name`, `menu_ids`, `remark`, `menu_check_strictly`, `status`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (1, 'test', '100,1001,1002,1003,1004,1005,1006,1007', 'test', 1, '0', 0, 0, '2025-03-16 11:56:06', 1, '2025-03-16 12:09:24', 0, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `tenant_id` varchar(20) DEFAULT '000000' COMMENT '租户编号',
  `dept_id` bigint(20) DEFAULT NULL COMMENT '部门ID',
  `user_name` varchar(30) NOT NULL COMMENT '用户账号',
  `nick_name` varchar(30) NOT NULL COMMENT '用户昵称',
  `user_type` varchar(10) DEFAULT 'sys_user' COMMENT '用户类型（sys_user系统用户）',
  `email` varchar(50) DEFAULT '' COMMENT '用户邮箱',
  `phonenumber` varchar(11) DEFAULT '' COMMENT '手机号码',
  `sex` char(1) DEFAULT '0' COMMENT '用户性别（0男 1女 2未知）',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像地址',
  `salt` varchar(100) DEFAULT NULL COMMENT '加密盐',
  `password` varchar(100) DEFAULT '' COMMENT '密码',
  `status` char(1) DEFAULT '0' COMMENT '帐号状态（0正常 1停用）',
  `login_ip` varchar(128) DEFAULT '' COMMENT '最后登录IP',
  `login_date` datetime DEFAULT NULL COMMENT '最后登录时间',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_by` bigint(20) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `uk_user_name` (`tenant_id`,`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息表';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` (`user_id`, `tenant_id`, `dept_id`, `user_name`, `nick_name`, `user_type`, `email`, `phonenumber`, `sex`, `avatar`, `salt`, `password`, `status`, `login_ip`, `login_date`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (1, '000000', 103, 'admin', '好心情', 'sys_user', 'lxj521w@163.com', '15888888887', '0', 'https://iot.xiujie.cn/wp-content/uploads/2024/09/96x96.png', '123456', '5f1d7a84db00d2fce00b31a7fc73224f', '0', '127.0.0.1', '2025-02-13 11:56:36', 103, 1, '2025-02-13 11:56:36', 1, '2025-03-26 13:44:18', NULL, NULL, '管理员');
INSERT INTO `sys_user` (`user_id`, `tenant_id`, `dept_id`, `user_name`, `nick_name`, `user_type`, `email`, `phonenumber`, `sex`, `avatar`, `salt`, `password`, `status`, `login_ip`, `login_date`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (3, '000000', 108, 'test', '本部门及以下 密码666666', 'sys_user', '', '', '0', NULL, 'ddd', '$2a$10$b8yUzN0C71sbz.PhNOCgJe.Tu1yWC3RNrTyjSQ8p1W0.aaUXUJ.Ne', '0', '127.0.0.1', '2025-02-13 11:56:36', 103, 1, '2025-02-13 11:56:36', 3, '2025-02-13 11:56:36', NULL, NULL, NULL);
INSERT INTO `sys_user` (`user_id`, `tenant_id`, `dept_id`, `user_name`, `nick_name`, `user_type`, `email`, `phonenumber`, `sex`, `avatar`, `salt`, `password`, `status`, `login_ip`, `login_date`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (4, '000000', 102, 'test1', '仅本人 密码666666', 'sys_user', '', '', '0', NULL, 'ddd', '$2a$10$b8yUzN0C71sbz.PhNOCgJe.Tu1yWC3RNrTyjSQ8p1W0.aaUXUJ.Ne', '0', '127.0.0.1', '2025-02-13 11:56:36', 103, 1, '2025-02-13 11:56:36', 4, '2025-02-13 11:56:36', NULL, NULL, NULL);
INSERT INTO `sys_user` (`user_id`, `tenant_id`, `dept_id`, `user_name`, `nick_name`, `user_type`, `email`, `phonenumber`, `sex`, `avatar`, `salt`, `password`, `status`, `login_ip`, `login_date`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (5, '000000', 103, 'test2', '好心情', '', 'li@xiujie.cn', '15888888887', '0', '', 'mjhhs', 'de396f3b10c2020d35161d8f45d0db52', '0', '', NULL, 103, 1, '2025-02-25 21:39:16', 1, '2025-02-25 22:09:49', 1, '2025-02-25 22:09:49', '');
INSERT INTO `sys_user` (`user_id`, `tenant_id`, `dept_id`, `user_name`, `nick_name`, `user_type`, `email`, `phonenumber`, `sex`, `avatar`, `salt`, `password`, `status`, `login_ip`, `login_date`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (6, '000000', 100, 'test3', '好心情', '', 'li@xiujie.cn', '15888888889', '0', '', '5nyxW', '8f2dd9b2f537dc0f3a4adf7238ab9939', '0', '', NULL, 103, 1, '2025-02-25 22:43:18', 1, '2025-02-25 23:27:08', 1, '2025-02-25 23:27:08', '');
INSERT INTO `sys_user` (`user_id`, `tenant_id`, `dept_id`, `user_name`, `nick_name`, `user_type`, `email`, `phonenumber`, `sex`, `avatar`, `salt`, `password`, `status`, `login_ip`, `login_date`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (10, '100006', 117, 'xiujiezhilian', '李先生', 'sys_user', '', '18600000000', '2', NULL, '5nyxW', '68c6e105523f29f1a97607f6739cd245', '0', '', NULL, 0, 0, '2025-03-16 14:36:52', 0, '2025-03-16 14:36:52', NULL, NULL, NULL);
INSERT INTO `sys_user` (`user_id`, `tenant_id`, `dept_id`, `user_name`, `nick_name`, `user_type`, `email`, `phonenumber`, `sex`, `avatar`, `salt`, `password`, `status`, `login_ip`, `login_date`, `created_dept`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`, `remark`) VALUES (11, '000000', 101, 'test001', '济南测试1', '', '563242931@qq.com', '', '0', '', 'mLwlN', '70824cefa9bc3da395ea16a48118e477', '0', '', NULL, 103, 1, '2025-03-17 13:52:46', 1, '2025-03-17 13:55:24', NULL, NULL, '');
COMMIT;

-- ----------------------------
-- Table structure for sys_user_online
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_online`;
CREATE TABLE `sys_user_online` (
  `online_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `tenant_id` varchar(20) DEFAULT '000000' COMMENT '租户编号',
  `uuid` varchar(50) DEFAULT '' COMMENT 'UUID',
  `user_name` varchar(50) DEFAULT '' COMMENT '用户账号',
  `client_key` varchar(32) DEFAULT '' COMMENT '客户端',
  `device_type` varchar(32) DEFAULT '' COMMENT '设备类型',
  `ipaddr` varchar(128) DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) DEFAULT '' COMMENT '操作系统',
  `token` text  COMMENT 'Token',
  `login_time` datetime DEFAULT NULL COMMENT '访问时间',
  `expire_time` datetime DEFAULT NULL COMMENT '过期时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`online_id`)
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户在线列表';

-- ----------------------------
-- Records of sys_user_online
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (1, '000000', '372da52666d74519949b93a9dae1b0c2', 'admin', 'web', 'web', '::1', '湖北省 黄石市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzcyZGE1MjY2NmQ3NDUxOTk0OWI5M2E5ZGFlMWIwYzIiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiQXV0aG9yaXR5SWRzIjpudWxsLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MDgzOTgyNywibmJmIjoxNzQwMjM1MDI3fQ.jdikYGZlmmeUzl6mEwL-tl1_4ELoqqHAfSTqKZKQPc0', '2025-02-22 22:37:08', '2025-03-01 22:37:07', '2025-03-26 19:56:24');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (2, '000000', 'fa31aef93ba4413e9f103c65768140f6', 'admin', 'web', 'web', '::1', '湖北省 黄石市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZmEzMWFlZjkzYmE0NDEzZTlmMTAzYzY1NzY4MTQwZjYiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiQXV0aG9yaXR5SWRzIjpudWxsLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MTA0ODMwMywibmJmIjoxNzQwNDQzNTAzfQ.TlyVDQW_vT09Ef_yzdCcZkJXzsnO1UfsHJDFW7gp0k8', '2025-02-25 08:31:44', '2025-03-04 08:31:43', '2025-03-26 19:56:24');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (3, '000000', 'f6f301146ca34c73bcfaa7f106e67638', 'admin', 'web', 'web', '::1', '湖北省 黄石市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZjZmMzAxMTQ2Y2EzNGM3M2JjZmFhN2YxMDZlNjc2MzgiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiQXV0aG9yaXR5SWRzIjpudWxsLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MTA0ODU4OSwibmJmIjoxNzQwNDQzNzg5fQ.wNcP6bncok4myQgvat2aCvCJaJ7LLMw7CmKeEg1WmE8', '2025-02-25 08:36:30', '2025-03-04 08:36:29', '2025-03-26 19:56:24');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (4, '000000', 'b7f64ee162734555a993f644d60eabd4', 'admin', 'web', 'web', '::1', '湖北省 黄石市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYjdmNjRlZTE2MjczNDU1NWE5OTNmNjQ0ZDYwZWFiZDQiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiQXV0aG9yaXR5SWRzIjpudWxsLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MTA0ODYyOCwibmJmIjoxNzQwNDQzODI4fQ.6u6gi9PJlxi1Xehwh1bFFiF21NNIZSbnR88XDyBR5U8', '2025-02-25 08:37:08', '2025-03-04 08:37:08', '2025-03-26 19:56:24');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (5, '000000', '1e8e02d0d9a04598a09653294cceba26', 'admin', 'web', 'web', '::1', '湖北省 黄石市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMWU4ZTAyZDBkOWEwNDU5OGEwOTY1MzI5NGNjZWJhMjYiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiQXV0aG9yaXR5SWRzIjpudWxsLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MTA0ODc0OSwibmJmIjoxNzQwNDQzOTQ5fQ.OZpIX4xqmuOQYmXWulMdnQEHPCUJ3BCFqkKdBnM2rLc', '2025-02-25 08:39:11', '2025-03-04 08:39:09', '2025-03-26 19:56:24');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (6, '000000', 'c6a613444af3465484b6efca728ae2d9', 'admin', 'web', 'web', '::1', '湖北省 黄石市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzZhNjEzNDQ0YWYzNDY1NDg0YjZlZmNhNzI4YWUyZDkiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiQXV0aG9yaXR5SWRzIjpudWxsLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MTA0ODc5OCwibmJmIjoxNzQwNDQzOTk4fQ.Ah036jhv63VKcNzkWG8nzqoWKdZOcg79AtKJE-HqgZo', '2025-02-25 08:39:59', '2025-03-04 08:39:58', '2025-03-26 19:56:24');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (7, '000000', 'f22257ed787c4c32b86d53c3a157d932', 'admin', 'web', 'web', '::1', '湖北省 黄石市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZjIyMjU3ZWQ3ODdjNGMzMmI4NmQ1M2MzYTE1N2Q5MzIiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiQXV0aG9yaXR5SWRzIjpudWxsLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MTA0ODg0MCwibmJmIjoxNzQwNDQ0MDQwfQ.iBQYlcjkd4vCkoBH2Hzh7KaAkTdo1bPKZFOVx1oqR6c', '2025-02-25 08:40:41', '2025-03-04 08:40:40', '2025-03-26 19:56:24');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (8, '000000', '7e96b52cde2d4cee809cb2dcff6a323e', 'admin', 'web', 'web', '::1', '湖北省 黄石市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiN2U5NmI1MmNkZTJkNGNlZTgwOWNiMmRjZmY2YTMyM2UiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiQXV0aG9yaXR5SWRzIjpudWxsLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MTA0ODk4NCwibmJmIjoxNzQwNDQ0MTg0fQ.DSQXfHqSBocR_Kh2TdSRSrYMb1DAZdv4nIncmk5-CEc', '2025-02-25 08:43:05', '2025-03-04 08:43:04', '2025-03-26 19:56:24');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (9, '000000', '124a82ef8ac34450bda6bbcc384a27a3', 'admin', 'web', 'web', '::1', '湖北省 黄石市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTI0YTgyZWY4YWMzNDQ1MGJkYTZiYmNjMzg0YTI3YTMiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiQXV0aG9yaXR5SWRzIjpudWxsLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MTA0OTA0OSwibmJmIjoxNzQwNDQ0MjQ5fQ.8UHL_ENPVPux9Ug7UnoMxkp33QMhNB58n4PBawIjcL0', '2025-02-25 08:44:10', '2025-03-04 08:44:09', '2025-03-26 19:56:24');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (10, '000000', '8c51495aefaa4d388b83da6fbe5caefc', 'admin', 'web', 'web', '::1', '湖北省 黄石市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiOGM1MTQ5NWFlZmFhNGQzODhiODNkYTZmYmU1Y2FlZmMiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiQXV0aG9yaXR5SWRzIjpudWxsLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MTA1MjQwMywibmJmIjoxNzQwNDQ3NjAzfQ.DNQpFd3ucwN_c4cLefdpG_waS-WWhg5QurJ7xwJckYs', '2025-02-25 09:40:05', '2025-03-04 09:40:03', '2025-03-26 19:56:24');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (11, '000000', 'dc1338d4f5c14e4b94ba373ddc8dbc06', 'admin', 'web', 'web', '::1', '湖北省 黄石市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZGMxMzM4ZDRmNWMxNGU0Yjk0YmEzNzNkZGM4ZGJjMDYiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiQXV0aG9yaXR5SWRzIjpudWxsLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MTA1MjQyNywibmJmIjoxNzQwNDQ3NjI3fQ.StsbjzjrKPq9pAXUm6UWx2qOwypgVdHzSu-Cs3yWieM', '2025-02-25 09:40:28', '2025-03-04 09:40:27', '2025-03-26 19:56:24');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (12, '000000', '3f66abc4f8884c0fba8072dcbb0bccbc', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiM2Y2NmFiYzRmODg4NGMwZmJhODA3MmRjYmIwYmNjYmMiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiQXV0aG9yaXR5SWRzIjpudWxsLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MTE0MTg1MCwibmJmIjoxNzQwNTM3MDUwfQ.Mj7KvtYesYq4paBj8JvWnYLyjyt0tRuZy6IEYwo1xOg', '2025-02-26 10:30:51', '2025-03-05 10:30:50', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (13, '000000', 'e541aac62e814b6a913017ba81408293', 'admin', 'web', 'web', '::1', '湖北省 黄石市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZTU0MWFhYzYyZTgxNGI2YTkxMzAxN2JhODE0MDgyOTMiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiQXV0aG9yaXR5SWRzIjpbMV0sIkxvZ2luQXQiOjE3NDA1ODgzNjEsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJ4aXVqaWUiLCJhdWQiOlsiWGl1amllQWRtaW4iXSwiZXhwIjoxNzQxMTkzMTYxLCJuYmYiOjE3NDA1ODgzNjF9.DbyqebD28u25EzmGF_e6ma470vcE8cpuwU_zmKU2z7s', '2025-02-27 00:46:02', '2025-03-06 00:46:01', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (14, '000000', '90ed1656170841ef929c67da646d55c7', 'admin', 'web', 'web', '::1', '湖北省 黄石市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiOTBlZDE2NTYxNzA4NDFlZjkyOWM2N2RhNjQ2ZDU1YzciLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiQXV0aG9yaXR5SWRzIjpbMV0sIkxvZ2luQXQiOjE3NDA1ODg1NzIsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJ4aXVqaWUiLCJhdWQiOlsiWGl1amllQWRtaW4iXSwiZXhwIjoxNzQxMTkzMzcyLCJuYmYiOjE3NDA1ODg1NzJ9.LEyrapVm5RQCtLEa1IZV1LX8DkoTQh325m-Jg6IeJms', '2025-02-27 00:49:33', '2025-03-06 00:49:32', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (15, '000000', '65dd5319a8b04c8e95ac6d55249a7042', 'admin', 'web', 'web', '::1', '湖北省 黄石市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNjVkZDUzMTlhOGIwNGM4ZTk1YWM2ZDU1MjQ5YTcwNDIiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOltdLCJMb2dpbkF0IjoxNzQwOTI2MDc5LCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MTUzMDg3OSwibmJmIjoxNzQwOTI2MDc5fQ.8UUN1pnB76FqrXjkQYg6OsbyWtySOBeAIYeGRsYK4iU', '2025-03-02 22:34:40', '2025-03-09 22:34:39', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (16, '000000', 'a850a398464e45b6924a2708f32b185d', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYTg1MGEzOTg0NjRlNDViNjkyNGEyNzA4ZjMyYjE4NWQiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOltdLCJMb2dpbkF0IjoxNzQxMzE2NzcxLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MTkyMTU3MSwibmJmIjoxNzQxMzE2NzcxfQ.5zacoD7wbEYBEwVxC5qhCQMnidhm_57OcjDqpEBtcfU', '2025-03-07 11:06:13', '2025-03-14 11:06:11', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (17, '000000', '438d65900342480e8ef11268559c4d90', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNDM4ZDY1OTAwMzQyNDgwZThlZjExMjY4NTU5YzRkOTAiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOltdLCJMb2dpbkF0IjoxNzQxOTIyMjkzLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MjUyNzA5MywibmJmIjoxNzQxOTIyMjkzfQ.KzleOM3NNz0pKx7RQHfPivx2ySg3uBY_BqZi10uChAw', '2025-03-14 11:18:14', '2025-03-21 11:18:13', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (18, '000000', 'af87c9692d2d456d9d8a54a560cfd48b', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYWY4N2M5NjkyZDJkNDU2ZDlkOGE1NGE1NjBjZmQ0OGIiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOltdLCJMb2dpbkF0IjoxNzQyMDk1NjU2LCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MjcwMDQ1NiwibmJmIjoxNzQyMDk1NjU2fQ.sRinD-FRx7TXzu8_G5cUrQkMNgJsyca6PZT3algJlnk', '2025-03-16 11:27:37', '2025-03-23 11:27:36', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (19, '100006', '08ca7a57548a4839ae0ef65584519b2d', 'xiujiezhilian', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMDhjYTdhNTc1NDhhNDgzOWFlMGVmNjU1ODQ1MTliMmQiLCJJRCI6MTAsIlVzZXJuYW1lIjoieGl1amllemhpbGlhbiIsIk5pY2tOYW1lIjoi5p2O5YWI55SfIiwiRGVwdElkIjoxMTcsIlRlbmFudElkIjoiMTAwMDA2IiwiUm9sZXMiOltdLCJMb2dpbkF0IjoxNzQyMTA3MTYzLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MjcxMTk2MywibmJmIjoxNzQyMTA3MTYzfQ.Y1yRakAG2CvyUbbvH-Z8aP5H4DVxlO4El694lXh6sOM', '2025-03-16 14:39:24', '2025-03-23 14:39:23', NULL);
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (20, '000000', 'ed9c97f07a2147bc9032346885c245b7', 'admin', 'web', 'web', '127.0.0.1', '内网IP', 'Chrome', 'Windows', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZWQ5Yzk3ZjA3YTIxNDdiYzkwMzIzNDY4ODVjMjQ1YjciLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOltdLCJMb2dpbkF0IjoxNzQyMTcyMDM2LCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0Mjc3NjgzNiwibmJmIjoxNzQyMTcyMDM2fQ.ig0moQb5pE4IrtHoEIsSWkV91pWM8sN0cAXupO9bA9Y', '2025-03-17 08:40:37', '2025-03-24 08:40:36', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (21, '000000', 'd9edaf372f424149ba1884a4edab0862', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZDllZGFmMzcyZjQyNDE0OWJhMTg4NGE0ZWRhYjA4NjIiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOltdLCJMb2dpbkF0IjoxNzQyMTc1NjgwLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0Mjc4MDQ4MCwibmJmIjoxNzQyMTc1NjgwfQ.Z_c5eh6dr0nZXbXVCZOUYQe0NGJkz9P7PnYwIWpG-lY', '2025-03-17 09:41:21', '2025-03-24 09:41:20', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (22, '000000', '3e98188952c245ad9f78b8764df6613a', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiM2U5ODE4ODk1MmMyNDVhZDlmNzhiODc2NGRmNjYxM2EiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOltdLCJMb2dpbkF0IjoxNzQyMTc1NzU5LCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0Mjc4MDU1OSwibmJmIjoxNzQyMTc1NzU5fQ.kpcuGYsgkVhmRkXtpk13L0aedZ6IT_0cDwxc_DrpIGw', '2025-03-17 09:42:40', '2025-03-24 09:42:39', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (23, '000000', 'f6e91e31e9ab4e58a1e803eac6505d5c', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZjZlOTFlMzFlOWFiNGU1OGExZTgwM2VhYzY1MDVkNWMiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOltdLCJMb2dpbkF0IjoxNzQyMTc5Njk5LCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0Mjc4NDQ5OSwibmJmIjoxNzQyMTc5Njk5fQ.g2pREc6J5n3RaYKWo2ns7YSeF7_Con-0-P6Bdbx1rL8', '2025-03-17 10:48:20', '2025-03-24 10:48:19', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (24, '000000', '8e5dc2536a774c9c9c13931317881983', 'test001', 'web', 'web', '127.0.0.1', '内网IP', 'Chrome', 'Windows', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiOGU1ZGMyNTM2YTc3NGM5YzljMTM5MzEzMTc4ODE5ODMiLCJJRCI6MTEsIlVzZXJuYW1lIjoidGVzdDAwMSIsIk5pY2tOYW1lIjoi5rWO5Y2X5rWL6K-VMSIsIkRlcHRJZCI6MTAxLCJUZW5hbnRJZCI6IjAwMDAwMCIsIlJvbGVzIjpbXSwiTG9naW5BdCI6MTc0MjE5MDk3MiwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDI3OTU3NzIsIm5iZiI6MTc0MjE5MDk3Mn0.7TLk6V_Ru11TBCWhH8NNQYTL5fOdUfODdpIUmF84Udo', '2025-03-17 13:56:12', '2025-03-24 13:56:12', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (25, '000000', '36bbaf667ee54f309af377d920ec26ee', 'admin', 'web', 'web', '::1', '湖北省 黄石市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzZiYmFmNjY3ZWU1NGYzMDlhZjM3N2Q5MjBlYzI2ZWUiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOltdLCJMb2dpbkF0IjoxNzQyNzM3OTk4LCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MzM0Mjc5OCwibmJmIjoxNzQyNzM3OTk4fQ.EA8hEolfJL9B4UlZVyErs4xK5acBO_Ya53gKIL1ZZ4U', '2025-03-23 21:53:19', '2025-03-30 21:53:18', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (26, '000000', '3781ffe36dfb4fbb841cf2e3b6b46b4c', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzc4MWZmZTM2ZGZiNGZiYjg0MWNmMmUzYjZiNDZiNGMiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOltdLCJMb2dpbkF0IjoxNzQyNzc1ODA2LCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MzM4MDYwNiwibmJmIjoxNzQyNzc1ODA2fQ.SdIfqIBupJ9G-QfgrtELSZm-ZUk9bWWjWxYSJN5UlUw', '2025-03-24 08:23:27', '2025-03-31 08:23:26', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (27, '000000', 'deaf2a12b4b84b72a566322e8461166d', 'admin', 'web', 'web', '127.0.0.1', '内网IP', 'Chrome', 'Windows', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZGVhZjJhMTJiNGI4NGI3MmE1NjYzMjJlODQ2MTE2NmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOltdLCJMb2dpbkF0IjoxNzQyNzk4MzU4LCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MzQwMzE1OCwibmJmIjoxNzQyNzk4MzU4fQ.JXSSZRsptw33WdmqGu75Ss-Cg-LBffotdkuVL4Z2Dw8', '2025-03-24 14:39:19', '2025-03-31 14:39:18', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (28, '000000', 'f42a707582f048c5881b0164ab42b865', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZjQyYTcwNzU4MmYwNDhjNTg4MWIwMTY0YWI0MmI4NjUiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOltdLCJMb2dpbkF0IjoxNzQyOTAwODM3LCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MzUwNTYzNywibmJmIjoxNzQyOTAwODM3fQ.CmdiB9RZ2E3FGJFjQBaX4Ob_VZ4jHZVVThRJZnZ2Qmo', '2025-03-25 19:07:17', '2025-04-01 19:07:17', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (29, '000000', '62bd513072bb48c3b2a806103088f2df', 'test001', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNjJiZDUxMzA3MmJiNDhjM2IyYTgwNjEwMzA4OGYyZGYiLCJJRCI6MTEsIlVzZXJuYW1lIjoidGVzdDAwMSIsIk5pY2tOYW1lIjoi5rWO5Y2X5rWL6K-VMSIsIkRlcHRJZCI6MTAxLCJUZW5hbnRJZCI6IjAwMDAwMCIsIlJvbGVzIjpbXSwiTG9naW5BdCI6MTc0Mjk1MDQyNCwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDM1NTUyMjQsIm5iZiI6MTc0Mjk1MDQyNH0.qFEnycWOrZmZvxu3R7GvYOxo78rWPEXl4l7BUkW4ibM', '2025-03-26 08:53:45', '2025-04-02 08:53:44', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (30, '000000', 'e4ad2972274a43b8ba46d0f51fb4e3a3', 'test001', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZTRhZDI5NzIyNzRhNDNiOGJhNDZkMGY1MWZiNGUzYTMiLCJJRCI6MTEsIlVzZXJuYW1lIjoidGVzdDAwMSIsIk5pY2tOYW1lIjoi5rWO5Y2X5rWL6K-VMSIsIkRlcHRJZCI6MTAxLCJUZW5hbnRJZCI6IjAwMDAwMCIsIlJvbGVzIjpbXSwiTG9naW5BdCI6MTc0Mjk1MjA5OSwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDM1NTY4OTksIm5iZiI6MTc0Mjk1MjA5OX0.ZrvFUxyi55-cDcXTn4yV-2y5bP0clMTrxs3YVFTh8t0', '2025-03-26 09:21:39', '2025-04-02 09:21:39', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (31, '000000', '2cac7f094bce4a70bb5c0306ac3e0a79', 'test001', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMmNhYzdmMDk0YmNlNGE3MGJiNWMwMzA2YWMzZTBhNzkiLCJJRCI6MTEsIlVzZXJuYW1lIjoidGVzdDAwMSIsIk5pY2tOYW1lIjoi5rWO5Y2X5rWL6K-VMSIsIkRlcHRJZCI6MTAxLCJUZW5hbnRJZCI6IjAwMDAwMCIsIlJvbGVzIjpbXSwiTG9naW5BdCI6MTc0Mjk1MjIzNCwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDM1NTcwMzQsIm5iZiI6MTc0Mjk1MjIzNH0.0uG7hw4CBE-lKHf59L5XqgCGNNAtl5vZWNXRue-sdeo', '2025-03-26 09:23:55', '2025-04-02 09:23:54', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (32, '000000', 'eb38600457aa4e13b6f3d876d10662b5', 'test001', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZWIzODYwMDQ1N2FhNGUxM2I2ZjNkODc2ZDEwNjYyYjUiLCJJRCI6MTEsIlVzZXJuYW1lIjoidGVzdDAwMSIsIk5pY2tOYW1lIjoi5rWO5Y2X5rWL6K-VMSIsIkRlcHRJZCI6MTAxLCJUZW5hbnRJZCI6IjAwMDAwMCIsIlJvbGVzIjpbXSwiTG9naW5BdCI6MTc0Mjk1MjUyMCwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDM1NTczMjAsIm5iZiI6MTc0Mjk1MjUyMH0.OC9yVtPmoTzqvEPKlBhau5OLdpXJFMmi61Z-Xc4w70g', '2025-03-26 09:28:41', '2025-04-02 09:28:40', '2025-03-26 19:56:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (33, '000000', '290c06bf45e84090887ebceca958ed59', 'test001', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMjkwYzA2YmY0NWU4NDA5MDg4N2ViY2VjYTk1OGVkNTkiLCJJRCI6MTEsIlVzZXJuYW1lIjoidGVzdDAwMSIsIk5pY2tOYW1lIjoi5rWO5Y2X5rWL6K-VMSIsIkRlcHRJZCI6MTAxLCJUZW5hbnRJZCI6IjAwMDAwMCIsIlJvbGVzIjpbXSwiTG9naW5BdCI6MTc0Mjk1Mjc1NCwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDM1NTc1NTQsIm5iZiI6MTc0Mjk1Mjc1NH0.24LAujGxi5KuPV6GkbhStAfqSdPwJaxq8W2GA2k2BJc', '2025-03-26 09:32:35', '2025-04-02 09:32:34', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (34, '000000', 'e6485c1a2ecd4a06877bd010a42cab38', 'test001', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZTY0ODVjMWEyZWNkNGEwNjg3N2JkMDEwYTQyY2FiMzgiLCJJRCI6MTEsIlVzZXJuYW1lIjoidGVzdDAwMSIsIk5pY2tOYW1lIjoi5rWO5Y2X5rWL6K-VMSIsIkRlcHRJZCI6MTAxLCJUZW5hbnRJZCI6IjAwMDAwMCIsIlJvbGVzIjpbXSwiTG9naW5BdCI6MTc0Mjk1MzAwMiwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDM1NTc4MDIsIm5iZiI6MTc0Mjk1MzAwMn0.slshwShs9hvaA5310ExfaepYE8DuWAvDsKLyu1yrD3Q', '2025-03-26 09:36:44', '2025-04-02 09:36:42', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (35, '000000', 'd1708df4f0654beaae6033a58097a147', 'test001', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZDE3MDhkZjRmMDY1NGJlYWFlNjAzM2E1ODA5N2ExNDciLCJJRCI6MTEsIlVzZXJuYW1lIjoidGVzdDAwMSIsIk5pY2tOYW1lIjoi5rWO5Y2X5rWL6K-VMSIsIkRlcHRJZCI6MTAxLCJUZW5hbnRJZCI6IjAwMDAwMCIsIlJvbGVzIjpbXSwiTG9naW5BdCI6MTc0Mjk1MzIxNywiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDM1NTgwMTcsIm5iZiI6MTc0Mjk1MzIxN30.HLT3nxxUdBxyf4EPiE7a1hNLuOzBEGLf7B5Ourq5Whg', '2025-03-26 09:40:18', '2025-04-02 09:40:17', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (36, '000000', '5149313a3cfb41068b6547366202f2b9', 'test001', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNTE0OTMxM2EzY2ZiNDEwNjhiNjU0NzM2NjIwMmYyYjkiLCJJRCI6MTEsIlVzZXJuYW1lIjoidGVzdDAwMSIsIk5pY2tOYW1lIjoi5rWO5Y2X5rWL6K-VMSIsIkRlcHRJZCI6MTAxLCJUZW5hbnRJZCI6IjAwMDAwMCIsIlJvbGVzIjpbXSwiTG9naW5BdCI6MTc0Mjk1NDQ4MiwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDM1NTkyODIsIm5iZiI6MTc0Mjk1NDQ4Mn0.ckbIzzlNwqLiSWpDAOBs1BFnyJAHhCopAjGzD60OeTs', '2025-03-26 10:01:22', '2025-04-02 10:01:22', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (37, '000000', '85591d3aebc04b349c5f724fb5c48d7e', 'test001', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiODU1OTFkM2FlYmMwNGIzNDljNWY3MjRmYjVjNDhkN2UiLCJJRCI6MTEsIlVzZXJuYW1lIjoidGVzdDAwMSIsIk5pY2tOYW1lIjoi5rWO5Y2X5rWL6K-VMSIsIkRlcHRJZCI6MTAxLCJUZW5hbnRJZCI6IjAwMDAwMCIsIlJvbGVzIjpbeyJSb2xlSWQiOjMsIkRhdGFTY29wZSI6IjQifV0sIkxvZ2luQXQiOjE3NDI5NTQ1NDUsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJ4aXVqaWUiLCJhdWQiOlsiWGl1amllQWRtaW4iXSwiZXhwIjoxNzQzNTU5MzQ1LCJuYmYiOjE3NDI5NTQ1NDV9.qzt3D6KbFRDx4ZcfNKH9FlN9ohdQMgwWmhT1j0ESZ40', '2025-03-26 10:02:26', '2025-04-02 10:02:25', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (38, '000000', '703df3504bd14feaa06702b0c91d648c', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Safari', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNzAzZGYzNTA0YmQxNGZlYWEwNjcwMmIwYzkxZDY0OGMiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOlt7IlJvbGVJZCI6MSwiRGF0YVNjb3BlIjoiMSJ9LHsiUm9sZUlkIjozLCJEYXRhU2NvcGUiOiI0In0seyJSb2xlSWQiOjQsIkRhdGFTY29wZSI6IjUifSx7IlJvbGVJZCI6NywiRGF0YVNjb3BlIjoiMSJ9LHsiUm9sZUlkIjo4LCJEYXRhU2NvcGUiOiIyIn1dLCJMb2dpbkF0IjoxNzQyOTY0MjQxLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MzU2OTA0MSwibmJmIjoxNzQyOTY0MjQxfQ.23oFo2Kq5La1c7HwPoZxMXEJoFDeHgAiTdJa0lKB-I8', '2025-03-26 12:44:02', '2025-04-02 12:44:01', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (39, '000000', '51e1bc7f26c54dc2b71787d446824cf8', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNTFlMWJjN2YyNmM1NGRjMmI3MTc4N2Q0NDY4MjRjZjgiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOlt7IlJvbGVJZCI6MSwiRGF0YVNjb3BlIjoiMSJ9LHsiUm9sZUlkIjozLCJEYXRhU2NvcGUiOiI0In0seyJSb2xlSWQiOjQsIkRhdGFTY29wZSI6IjUifSx7IlJvbGVJZCI6NywiRGF0YVNjb3BlIjoiMSJ9LHsiUm9sZUlkIjo4LCJEYXRhU2NvcGUiOiIyIn1dLCJMb2dpbkF0IjoxNzQyOTY0MzE5LCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MzU2OTExOSwibmJmIjoxNzQyOTY0MzE5fQ.eoKzh6LEc65ELPW0DQUC-2j77zMZhvToV7kHUA5m4gg', '2025-03-26 12:45:20', '2025-04-02 12:45:19', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (40, '000000', '9bf4ce18aeea4c3aac8b0322a1a4cae7', 'test001', 'web', 'web', '::1', '山东省 济南市', 'Safari', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiOWJmNGNlMThhZWVhNGMzYWFjOGIwMzIyYTFhNGNhZTciLCJJRCI6MTEsIlVzZXJuYW1lIjoidGVzdDAwMSIsIk5pY2tOYW1lIjoi5rWO5Y2X5rWL6K-VMSIsIkRlcHRJZCI6MTAxLCJUZW5hbnRJZCI6IjAwMDAwMCIsIlJvbGVzIjpbeyJSb2xlSWQiOjMsIkRhdGFTY29wZSI6IjQifV0sIkxvZ2luQXQiOjE3NDI5NjQ3MjEsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJ4aXVqaWUiLCJhdWQiOlsiWGl1amllQWRtaW4iXSwiZXhwIjoxNzQzNTY5NTIxLCJuYmYiOjE3NDI5NjQ3MjF9.HJr1AKntHwG1qGTDcxX7FcPanqMekU6rzfM1VErptYw', '2025-03-26 12:52:01', '2025-04-02 12:52:01', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (41, '000000', '0f3473c62de54fa6b9bfe2420b681361', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMGYzNDczYzYyZGU1NGZhNmI5YmZlMjQyMGI2ODEzNjEiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOlt7IlJvbGVJZCI6MSwiRGF0YVNjb3BlIjoiMSJ9LHsiUm9sZUlkIjozLCJEYXRhU2NvcGUiOiI0In0seyJSb2xlSWQiOjQsIkRhdGFTY29wZSI6IjUifSx7IlJvbGVJZCI6NywiRGF0YVNjb3BlIjoiMSJ9LHsiUm9sZUlkIjo4LCJEYXRhU2NvcGUiOiIyIn1dLCJMb2dpbkF0IjoxNzQyOTY2NDU4LCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MzU3MTI1OCwibmJmIjoxNzQyOTY2NDU4fQ.xHD8Wg1YwvSfoJNgmFoGjWTZ7tiq7lXxvIj1tY32r4w', '2025-03-26 13:20:59', '2025-04-02 13:20:58', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (42, '000000', '2265917556154efa90f885ce8cf5c642', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMjI2NTkxNzU1NjE1NGVmYTkwZjg4NWNlOGNmNWM2NDIiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOlt7IlJvbGVJZCI6MSwiRGF0YVNjb3BlIjoiMSJ9LHsiUm9sZUlkIjozLCJEYXRhU2NvcGUiOiI0In0seyJSb2xlSWQiOjQsIkRhdGFTY29wZSI6IjUifSx7IlJvbGVJZCI6NywiRGF0YVNjb3BlIjoiMSJ9LHsiUm9sZUlkIjo4LCJEYXRhU2NvcGUiOiIyIn1dLCJMb2dpbkF0IjoxNzQyOTY2NTE3LCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MzU3MTMxNywibmJmIjoxNzQyOTY2NTE3fQ.wpSDWaDH3VyyBi536HclxBO4gwKs-Jm_PniLkVattk0', '2025-03-26 13:21:58', '2025-04-02 13:21:57', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (43, '000000', 'c77057c3edd24801b076b417ba71fb0b', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzc3MDU3YzNlZGQyNDgwMWIwNzZiNDE3YmE3MWZiMGIiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOlt7IlJvbGVJZCI6MSwiRGF0YVNjb3BlIjoiMSJ9LHsiUm9sZUlkIjozLCJEYXRhU2NvcGUiOiI0In0seyJSb2xlSWQiOjQsIkRhdGFTY29wZSI6IjUifSx7IlJvbGVJZCI6NywiRGF0YVNjb3BlIjoiMSJ9LHsiUm9sZUlkIjo4LCJEYXRhU2NvcGUiOiIyIn1dLCJMb2dpbkF0IjoxNzQyOTY2Njk3LCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MzU3MTQ5NywibmJmIjoxNzQyOTY2Njk3fQ.krJAYw6JzQarakfmFK9sH9wHAC0MB2LzAJhF1Ao2s5I', '2025-03-26 13:24:58', '2025-04-02 13:24:57', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (44, '000000', 'f6180cce5b7d451ab8efab0bff6f3558', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZjYxODBjY2U1YjdkNDUxYWI4ZWZhYjBiZmY2ZjM1NTgiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOlt7IlJvbGVJZCI6MSwiRGF0YVNjb3BlIjoiMSJ9LHsiUm9sZUlkIjozLCJEYXRhU2NvcGUiOiI0In0seyJSb2xlSWQiOjQsIkRhdGFTY29wZSI6IjUifSx7IlJvbGVJZCI6NywiRGF0YVNjb3BlIjoiMSJ9LHsiUm9sZUlkIjo4LCJEYXRhU2NvcGUiOiIyIn1dLCJMb2dpbkF0IjoxNzQyOTY2NzkzLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MzU3MTU5MywibmJmIjoxNzQyOTY2NzkzfQ.AOI1o2mH2cdPeGcT0TliQYzECtQnGvOvqC9Xr45ue9c', '2025-03-26 13:26:34', '2025-04-02 13:26:33', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (45, '000000', '5940a304f97c418980f566cba2714f85', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNTk0MGEzMDRmOTdjNDE4OTgwZjU2NmNiYTI3MTRmODUiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOlt7IlJvbGVJZCI6MSwiRGF0YVNjb3BlIjoiMSJ9LHsiUm9sZUlkIjozLCJEYXRhU2NvcGUiOiI0In0seyJSb2xlSWQiOjQsIkRhdGFTY29wZSI6IjUifSx7IlJvbGVJZCI6NywiRGF0YVNjb3BlIjoiMSJ9LHsiUm9sZUlkIjo4LCJEYXRhU2NvcGUiOiIyIn1dLCJMb2dpbkF0IjoxNzQyOTY3MjI5LCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoieGl1amllIiwiYXVkIjpbIlhpdWppZUFkbWluIl0sImV4cCI6MTc0MzU3MjAyOSwibmJmIjoxNzQyOTY3MjI5fQ.r4i3zzxO1_THgKGEZE6W7GOCLcghDqqkHLwJOdzzM5A', '2025-03-26 13:33:49', '2025-04-02 13:33:49', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (46, '000000', '815b7fea7c6243b7a186c6eaf4c6be1f', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiODE1YjdmZWE3YzYyNDNiN2ExODZjNmVhZjRjNmJlMWYiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOlt7IlJvbGVJZCI6MSwiRGF0YVNjb3BlIjoiMSJ9XSwiTG9naW5BdCI6MTc0Mjk2OTE3MywiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDM1NzM5NzMsIm5iZiI6MTc0Mjk2OTE3M30.I8H1Pa9KNNFp-KKGLoKDLgETH6uFfjKjCuVKSjf53qw', '2025-03-26 14:06:14', '2025-04-02 14:06:13', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (47, '000000', 'bda62b527db24b778e9b51398138bd8f', 'test001', 'web', 'web', '::1', '山东省 济南市', 'Safari', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYmRhNjJiNTI3ZGIyNGI3NzhlOWI1MTM5ODEzOGJkOGYiLCJJRCI6MTEsIlVzZXJuYW1lIjoidGVzdDAwMSIsIk5pY2tOYW1lIjoi5rWO5Y2X5rWL6K-VMSIsIkRlcHRJZCI6MTAxLCJUZW5hbnRJZCI6IjAwMDAwMCIsIlJvbGVzIjpbeyJSb2xlSWQiOjMsIkRhdGFTY29wZSI6IjQifV0sIkxvZ2luQXQiOjE3NDI5NjkyMTIsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJ4aXVqaWUiLCJhdWQiOlsiWGl1amllQWRtaW4iXSwiZXhwIjoxNzQzNTc0MDEyLCJuYmYiOjE3NDI5NjkyMTJ9.RIMveaBNXMLS8nlgGlS2AT12sa6FX17rA6dTT3v84Xk', '2025-03-26 14:06:53', '2025-04-02 14:06:52', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (48, '000000', '37ab0b88742144bea4725a8ff20adbd6', 'test001', 'web', 'web', '::1', '山东省 济南市', 'Safari', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzdhYjBiODg3NDIxNDRiZWE0NzI1YThmZjIwYWRiZDYiLCJJRCI6MTEsIlVzZXJuYW1lIjoidGVzdDAwMSIsIk5pY2tOYW1lIjoi5rWO5Y2X5rWL6K-VMSIsIkRlcHRJZCI6MTAxLCJUZW5hbnRJZCI6IjAwMDAwMCIsIlJvbGVzIjpbeyJSb2xlSWQiOjMsIkRhdGFTY29wZSI6IjQifV0sIkxvZ2luQXQiOjE3NDI5NjkzMzEsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJ4aXVqaWUiLCJhdWQiOlsiWGl1amllQWRtaW4iXSwiZXhwIjoxNzQzNTc0MTMxLCJuYmYiOjE3NDI5NjkzMzF9.yurzDVWCuLu_JAlnqnPZylCt4mIJSY98-v5p01HYbgI', '2025-03-26 14:08:52', '2025-04-02 14:08:51', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (49, '000000', '29be7cc59af3419f86380a32322a8b05', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Safari', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMjliZTdjYzU5YWYzNDE5Zjg2MzgwYTMyMzIyYThiMDUiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOlt7IlJvbGVJZCI6MSwiRGF0YVNjb3BlIjoiMSJ9XSwiTG9naW5BdCI6MTc0Mjk3MDQwMSwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDM1NzUyMDEsIm5iZiI6MTc0Mjk3MDQwMX0.hhRyFEywZ2IiqhRB07I46Sg6lUQuEuGJXSjip2MfdwI', '2025-03-26 14:26:41', '2025-04-02 14:26:41', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (50, '000000', '486e585279b0475092744b835881ea96', 'test001', 'web', 'web', '::1', '山东省 济南市', 'Safari', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNDg2ZTU4NTI3OWIwNDc1MDkyNzQ0YjgzNTg4MWVhOTYiLCJJRCI6MTEsIlVzZXJuYW1lIjoidGVzdDAwMSIsIk5pY2tOYW1lIjoi5rWO5Y2X5rWL6K-VMSIsIkRlcHRJZCI6MTAxLCJUZW5hbnRJZCI6IjAwMDAwMCIsIlJvbGVzIjpbeyJSb2xlSWQiOjMsIkRhdGFTY29wZSI6IjQifV0sIkxvZ2luQXQiOjE3NDI5NzA0MjEsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJ4aXVqaWUiLCJhdWQiOlsiWGl1amllQWRtaW4iXSwiZXhwIjoxNzQzNTc1MjIxLCJuYmYiOjE3NDI5NzA0MjF9.66WPJN6OzSiPnfDlpl0aw9t0uWPQr0eNS9NeWNB-0EE', '2025-03-26 14:27:02', '2025-04-02 14:27:01', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (51, '000000', 'aa1b627b1f1d48dbb81927a5e8512d84', 'test001', 'web', 'web', '::1', '山东省 济南市', 'Safari', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYWExYjYyN2IxZjFkNDhkYmI4MTkyN2E1ZTg1MTJkODQiLCJJRCI6MTEsIlVzZXJuYW1lIjoidGVzdDAwMSIsIk5pY2tOYW1lIjoi5rWO5Y2X5rWL6K-VMSIsIkRlcHRJZCI6MTAxLCJUZW5hbnRJZCI6IjAwMDAwMCIsIlJvbGVzIjpbeyJSb2xlSWQiOjMsIkRhdGFTY29wZSI6IjQifV0sIkxvZ2luQXQiOjE3NDI5NzE0NDAsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJ4aXVqaWUiLCJhdWQiOlsiWGl1amllQWRtaW4iXSwiZXhwIjoxNzQzNTc2MjQwLCJuYmYiOjE3NDI5NzE0NDB9.GZcoyLgoQkHnROXFNR7rNeE0t-PpwFHNtVerxBU_xpk', '2025-03-26 14:44:01', '2025-04-02 14:44:00', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (52, '000000', '4e17facd20c84865801d7c1f02545b12', 'test001', 'web', 'web', '::1', '山东省 济南市', 'Safari', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNGUxN2ZhY2QyMGM4NDg2NTgwMWQ3YzFmMDI1NDViMTIiLCJJRCI6MTEsIlVzZXJuYW1lIjoidGVzdDAwMSIsIk5pY2tOYW1lIjoi5rWO5Y2X5rWL6K-VMSIsIkRlcHRJZCI6MTAxLCJUZW5hbnRJZCI6IjAwMDAwMCIsIlJvbGVzIjpbeyJSb2xlSWQiOjMsIkRhdGFTY29wZSI6IjQifV0sIkxvZ2luQXQiOjE3NDI5NzM2NDMsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJ4aXVqaWUiLCJhdWQiOlsiWGl1amllQWRtaW4iXSwiZXhwIjoxNzQzNTc4NDQzLCJuYmYiOjE3NDI5NzM2NDN9.BINdvTe-HBe-s1ehijXlYizhqJmKHnlMVw1_XCbdIlU', '2025-03-26 15:20:43', '2025-04-02 15:20:43', '2025-03-26 19:56:13');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (53, '000000', '53de63a69ad74fdaac6dbd1964483de7', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNTNkZTYzYTY5YWQ3NGZkYWFjNmRiZDE5NjQ0ODNkZTciLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOlt7IlJvbGVJZCI6MSwiRGF0YVNjb3BlIjoiMSJ9XSwiTG9naW5BdCI6MTc0Mjk3NDczMCwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDM1Nzk1MzAsIm5iZiI6MTc0Mjk3NDczMH0.4eZeWBPqUR6ZqARMHsx5NSn3oDDZ7ZAKUeUWGqYdnoQ', '2025-03-26 15:38:50', '2025-04-02 15:38:50', '2025-03-26 19:56:07');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (54, '000000', 'cb57f97515e84eb68b8b657ffe20eb24', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2I1N2Y5NzUxNWU4NGViNjhiOGI2NTdmZmUyMGViMjQiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOlt7IlJvbGVJZCI6MSwiRGF0YVNjb3BlIjoiMSJ9XSwiTG9naW5BdCI6MTc0Mjk4MjIxOCwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDM1ODcwMTgsIm5iZiI6MTc0Mjk4MjIxOH0.yv50Tkw-0KXP4df4nhESiRfvn5k0YbFAhuq1_ddXqH8', '2025-03-26 17:43:39', '2025-04-02 17:43:38', '2025-03-26 19:53:20');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (55, '000000', '93028e2a5a044226bc510c9f1d40364e', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiOTMwMjhlMmE1YTA0NDIyNmJjNTEwYzlmMWQ0MDM2NGUiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOlt7IlJvbGVJZCI6MSwiRGF0YVNjb3BlIjoiMSJ9XSwiTG9naW5BdCI6MTc0Mjk4NDEzNywiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDM1ODg5MzcsIm5iZiI6MTc0Mjk4NDEzN30.nOtUk456VTonSKFfdDCx34h5k96a5iAnlucrNKgzbkU', '2025-03-26 18:15:38', '2025-04-02 18:15:37', '2025-03-26 19:48:03');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (56, '000000', '052585d10dca406989295ac3a6e36b52', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMDUyNTg1ZDEwZGNhNDA2OTg5Mjk1YWMzYTZlMzZiNTIiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOlt7IlJvbGVJZCI6MSwiRGF0YVNjb3BlIjoiMSJ9XSwiTG9naW5BdCI6MTc0Mjk5MDU3MiwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDM1OTUzNzIsIm5iZiI6MTc0Mjk5MDU3Mn0.70WstfTTJ_Su7SOxhh9kyxqK77mLWesBdgcvJnNM_fA', '2025-03-26 20:02:53', '2025-04-02 20:02:52', '2025-03-26 20:04:00');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (57, '000000', '50f13441ae7a4f83b3d7d95999809e9a', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNTBmMTM0NDFhZTdhNGY4M2IzZDdkOTU5OTk4MDllOWEiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOlt7IlJvbGVJZCI6MSwiRGF0YVNjb3BlIjoiMSJ9XSwiTG9naW5BdCI6MTc0Mjk5MDY1NiwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDM1OTU0NTYsIm5iZiI6MTc0Mjk5MDY1Nn0.I43yKDR5FX4ZG2XVvQ-k-TBESumVEqx8IZO47CV0J6Y', '2025-03-26 20:04:16', '2025-04-02 20:04:16', NULL);
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (58, '000000', 'a7109d4a705044f2907f2e3b7c6c1adc', 'admin', 'web', 'web', '127.0.0.1', '内网IP', 'Chrome', 'Windows', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYTcxMDlkNGE3MDUwNDRmMjkwN2YyZTNiN2M2YzFhZGMiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOlt7IlJvbGVJZCI6MSwiRGF0YVNjb3BlIjoiMSJ9XSwiTG9naW5BdCI6MTc0MzA1ODE1MywiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDM2NjI5NTMsIm5iZiI6MTc0MzA1ODE1M30.KNDHIWezNSBFAKgsDDwodNx3L9LFwukuskRygdHAWx0', '2025-03-27 14:49:13', '2025-04-03 14:49:13', NULL);
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (59, '000000', 'a8ae3c2d058b4da5a8ef4cca0ab559c2', 'admin', 'web', 'web', '127.0.0.1', '内网IP', 'Chrome', 'Windows', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYThhZTNjMmQwNThiNGRhNWE4ZWY0Y2NhMGFiNTU5YzIiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOlt7IlJvbGVJZCI6MSwiRGF0YVNjb3BlIjoiMSJ9XSwiTG9naW5BdCI6MTc0MzA1ODc0NSwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDM2NjM1NDUsIm5iZiI6MTc0MzA1ODc0NX0.cNboyNDe_vcFU9KF1aeuucuwB8fFgTYVzvZi2FCQEx4', '2025-03-27 14:59:06', '2025-04-03 14:59:05', NULL);
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (60, '000000', '0e875fd0eaea45299e23b1ad95657cb1', 'admin', 'web', 'web', '::1', '山东省 济南市', 'Chrome', 'macOS', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMGU4NzVmZDBlYWVhNDUyOTllMjNiMWFkOTU2NTdjYjEiLCJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi5aW95b-D5oOFIiwiRGVwdElkIjoxMDMsIlRlbmFudElkIjoiMDAwMDAwIiwiUm9sZXMiOlt7IlJvbGVJZCI6MSwiRGF0YVNjb3BlIjoiMSJ9XSwiTG9naW5BdCI6MTc0MzA2ODM0MywiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InhpdWppZSIsImF1ZCI6WyJYaXVqaWVBZG1pbiJdLCJleHAiOjE3NDM2NzMxNDMsIm5iZiI6MTc0MzA2ODM0M30.PrTAkOp6YvWMhnXgKAo5hHFNB9U8SVU8q7Wck7q6QDM', '2025-03-27 17:39:04', '2025-04-03 17:39:03', NULL);
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (61, '000000', 'sss', 'test001', '1111', '', '', '', '', '', '', NULL, NULL, '2025-03-27 17:41:50');
INSERT INTO `sys_user_online` (`online_id`, `tenant_id`, `uuid`, `user_name`, `client_key`, `device_type`, `ipaddr`, `login_location`, `browser`, `os`, `token`, `login_time`, `expire_time`, `deleted_at`) VALUES (62, '000000', 'aa@123.com', '1111', '', '', '', '', '', '', '', NULL, NULL, '2025-03-27 17:41:52');
COMMIT;

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post` (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `post_id` bigint(20) NOT NULL COMMENT '岗位ID',
  PRIMARY KEY (`user_id`,`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户与岗位关联表';

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_post` (`user_id`, `post_id`) VALUES (6, 3);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role` (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`user_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户和角色关联表';

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_role` (`user_id`, `role_id`) VALUES (1, 1);
INSERT INTO `sys_user_role` (`user_id`, `role_id`) VALUES (3, 3);
INSERT INTO `sys_user_role` (`user_id`, `role_id`) VALUES (4, 4);
INSERT INTO `sys_user_role` (`user_id`, `role_id`) VALUES (6, 3);
INSERT INTO `sys_user_role` (`user_id`, `role_id`) VALUES (11, 3);
COMMIT;

-- ----------------------------
-- Table structure for test_demo
-- ----------------------------
DROP TABLE IF EXISTS `test_demo`;
CREATE TABLE `test_demo` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` varchar(20) DEFAULT '000000' COMMENT '租户编号',
  `dept_id` bigint(20) DEFAULT NULL COMMENT '部门id',
  `user_id` bigint(20) DEFAULT NULL COMMENT '用户id',
  `order_num` int(11) DEFAULT 0 COMMENT '排序号',
  `test_key` varchar(255) DEFAULT NULL COMMENT 'key键',
  `value` varchar(255) DEFAULT NULL COMMENT '值',
  `version` int(11) DEFAULT 0 COMMENT '版本',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint(20) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='测试单表';

-- ----------------------------
-- Records of test_demo
-- ----------------------------
BEGIN;
INSERT INTO `test_demo` (`id`, `tenant_id`, `dept_id`, `user_id`, `order_num`, `test_key`, `value`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (1, '000000', 102, 4, 1, '测试数据权限', '测试', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_demo` (`id`, `tenant_id`, `dept_id`, `user_id`, `order_num`, `test_key`, `value`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (2, '000000', 102, 3, 2, '子节点1', '111', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_demo` (`id`, `tenant_id`, `dept_id`, `user_id`, `order_num`, `test_key`, `value`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (3, '000000', 102, 3, 3, '子节点2', '222', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_demo` (`id`, `tenant_id`, `dept_id`, `user_id`, `order_num`, `test_key`, `value`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (4, '000000', 108, 4, 4, '测试数据', 'demo', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_demo` (`id`, `tenant_id`, `dept_id`, `user_id`, `order_num`, `test_key`, `value`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (5, '000000', 108, 3, 13, '子节点11', '1111', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_demo` (`id`, `tenant_id`, `dept_id`, `user_id`, `order_num`, `test_key`, `value`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (6, '000000', 108, 3, 12, '子节点22', '2222', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_demo` (`id`, `tenant_id`, `dept_id`, `user_id`, `order_num`, `test_key`, `value`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (7, '000000', 108, 3, 11, '子节点33', '3333', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_demo` (`id`, `tenant_id`, `dept_id`, `user_id`, `order_num`, `test_key`, `value`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (8, '000000', 108, 3, 10, '子节点44', '0', 0, 103, '2025-02-13 11:56:36', 1, '2025-03-26 14:15:36', 11, NULL, NULL);
INSERT INTO `test_demo` (`id`, `tenant_id`, `dept_id`, `user_id`, `order_num`, `test_key`, `value`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (9, '000000', 108, 3, 9, '子节点55', '5555', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, 1, '2025-03-24 11:35:10');
INSERT INTO `test_demo` (`id`, `tenant_id`, `dept_id`, `user_id`, `order_num`, `test_key`, `value`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (10, '000000', 108, 3, 8, '子节点66', '6666', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, 1, '2025-03-24 11:35:08');
INSERT INTO `test_demo` (`id`, `tenant_id`, `dept_id`, `user_id`, `order_num`, `test_key`, `value`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (11, '000000', 108, 3, 7, '子节点77', '7777', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, 1, '2025-03-24 11:34:19');
INSERT INTO `test_demo` (`id`, `tenant_id`, `dept_id`, `user_id`, `order_num`, `test_key`, `value`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (12, '000000', 108, 3, 6, '子节点88', '8888', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, 1, '2025-03-24 11:06:22');
INSERT INTO `test_demo` (`id`, `tenant_id`, `dept_id`, `user_id`, `order_num`, `test_key`, `value`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (13, '000000', 108, 3, 5, '子节点99', '9999', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, 1, '2025-03-24 11:15:37');
INSERT INTO `test_demo` (`id`, `tenant_id`, `dept_id`, `user_id`, `order_num`, `test_key`, `value`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (14, '000000', 11, 223, 33, '333', '1', 33, NULL, '2025-03-24 10:56:57', 1, '2025-03-24 10:58:37', 1, 1, '2025-03-24 11:06:19');
COMMIT;

-- ----------------------------
-- Table structure for test_tree
-- ----------------------------
DROP TABLE IF EXISTS `test_tree`;
CREATE TABLE `test_tree` (
  `id` bigint(20) NOT NULL COMMENT '主键',
  `tenant_id` varchar(20) DEFAULT '000000' COMMENT '租户编号',
  `parent_id` bigint(20) DEFAULT 0 COMMENT '父id',
  `dept_id` bigint(20) DEFAULT NULL COMMENT '部门id',
  `user_id` bigint(20) DEFAULT NULL COMMENT '用户id',
  `tree_name` varchar(255) DEFAULT NULL COMMENT '值',
  `version` int(11) DEFAULT 0 COMMENT '版本',
  `created_dept` bigint(20) DEFAULT NULL COMMENT '创建部门',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint(20) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='测试树表';

-- ----------------------------
-- Records of test_tree
-- ----------------------------
BEGIN;
INSERT INTO `test_tree` (`id`, `tenant_id`, `parent_id`, `dept_id`, `user_id`, `tree_name`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (1, '000000', 0, 102, 4, '测试数据权限', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_tree` (`id`, `tenant_id`, `parent_id`, `dept_id`, `user_id`, `tree_name`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (2, '000000', 1, 102, 3, '子节点1', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_tree` (`id`, `tenant_id`, `parent_id`, `dept_id`, `user_id`, `tree_name`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (3, '000000', 2, 102, 3, '子节点2', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_tree` (`id`, `tenant_id`, `parent_id`, `dept_id`, `user_id`, `tree_name`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (4, '000000', 0, 108, 4, '测试树1', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_tree` (`id`, `tenant_id`, `parent_id`, `dept_id`, `user_id`, `tree_name`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (5, '000000', 4, 108, 3, '子节点11', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_tree` (`id`, `tenant_id`, `parent_id`, `dept_id`, `user_id`, `tree_name`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (6, '000000', 4, 108, 3, '子节点22', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_tree` (`id`, `tenant_id`, `parent_id`, `dept_id`, `user_id`, `tree_name`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (7, '000000', 4, 108, 3, '子节点33', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_tree` (`id`, `tenant_id`, `parent_id`, `dept_id`, `user_id`, `tree_name`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (8, '000000', 5, 108, 3, '子节点44', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_tree` (`id`, `tenant_id`, `parent_id`, `dept_id`, `user_id`, `tree_name`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (9, '000000', 6, 108, 3, '子节点55', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_tree` (`id`, `tenant_id`, `parent_id`, `dept_id`, `user_id`, `tree_name`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (10, '000000', 7, 108, 3, '子节点66', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_tree` (`id`, `tenant_id`, `parent_id`, `dept_id`, `user_id`, `tree_name`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (11, '000000', 7, 108, 3, '子节点77', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_tree` (`id`, `tenant_id`, `parent_id`, `dept_id`, `user_id`, `tree_name`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (12, '000000', 10, 108, 3, '子节点88', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
INSERT INTO `test_tree` (`id`, `tenant_id`, `parent_id`, `dept_id`, `user_id`, `tree_name`, `version`, `created_dept`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_by`, `deleted_at`) VALUES (13, '000000', 10, 108, 3, '子节点99', 0, 103, '2025-02-13 11:56:36', 1, NULL, NULL, NULL, NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
