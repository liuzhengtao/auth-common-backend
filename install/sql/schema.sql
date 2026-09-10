-- auth-common schema (MySQL)
-- 列名与 internal/dao/internal 保持一致

CREATE TABLE IF NOT EXISTS sys_dept (
  id bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  name varchar(64) NOT NULL COMMENT '部门名称',
  parent_id bigint NOT NULL DEFAULT 0 COMMENT '父节点id',
  tree_path varchar(255) DEFAULT NULL COMMENT '父节点id路径',
  sort int NOT NULL DEFAULT 0 COMMENT '显示顺序',
  status tinyint NOT NULL DEFAULT 1 COMMENT '状态(1:正常;0:禁用)',
  deleted tinyint NOT NULL DEFAULT 0 COMMENT '逻辑删除(1:已删除;0:未删除)',
  create_time datetime DEFAULT NULL,
  update_time datetime DEFAULT NULL,
  create_by bigint DEFAULT NULL,
  update_by bigint DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS sys_user (
  id bigint NOT NULL AUTO_INCREMENT,
  username varchar(64) NOT NULL COMMENT '用户名',
  nickname varchar(64) DEFAULT NULL COMMENT '昵称',
  gender tinyint DEFAULT NULL COMMENT '性别(1:男;2:女)',
  password varchar(100) NOT NULL COMMENT '密码(MD5(明文+盐))',
  dept_id bigint DEFAULT NULL COMMENT '部门ID',
  avatar varchar(255) DEFAULT NULL,
  mobile varchar(20) DEFAULT NULL,
  status tinyint NOT NULL DEFAULT 1 COMMENT '1:正常;0:禁用',
  email varchar(128) DEFAULT NULL,
  deleted tinyint NOT NULL DEFAULT 0 COMMENT '0:未删除;1:已删除',
  create_time datetime DEFAULT NULL,
  update_time datetime DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY uk_username (username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS sys_role (
  id bigint NOT NULL AUTO_INCREMENT,
  name varchar(64) NOT NULL COMMENT '角色名称',
  code varchar(64) NOT NULL COMMENT '角色编码',
  sort int NOT NULL DEFAULT 0,
  status tinyint NOT NULL DEFAULT 1 COMMENT '1-正常；0-停用',
  data_scope tinyint NOT NULL DEFAULT 0 COMMENT '0全部;1部门及子;2本部门;3本人',
  deleted tinyint NOT NULL DEFAULT 0,
  create_time datetime DEFAULT NULL,
  update_time datetime DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY uk_code (code)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS sys_menu (
  id bigint NOT NULL AUTO_INCREMENT,
  parent_id bigint NOT NULL DEFAULT 0 COMMENT '父菜单ID',
  tree_path varchar(255) DEFAULT NULL,
  name varchar(64) NOT NULL,
  type tinyint NOT NULL COMMENT '1菜单 2目录 3外链 4按钮',
  path varchar(255) DEFAULT NULL,
  component varchar(255) DEFAULT NULL,
  perm varchar(128) DEFAULT NULL COMMENT '权限标识',
  visible tinyint NOT NULL DEFAULT 1 COMMENT '1显示;0隐藏',
  sort int NOT NULL DEFAULT 0,
  icon varchar(64) DEFAULT NULL,
  redirect varchar(255) DEFAULT NULL,
  create_time datetime DEFAULT NULL,
  update_time datetime DEFAULT NULL,
  always_show tinyint DEFAULT 0,
  keep_alive tinyint DEFAULT 0,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS sys_user_role (
  user_id bigint NOT NULL,
  role_id bigint NOT NULL,
  PRIMARY KEY (user_id, role_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS sys_role_menu (
  role_id bigint NOT NULL,
  menu_id bigint NOT NULL,
  PRIMARY KEY (role_id, menu_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS sys_dict_type (
  id bigint NOT NULL AUTO_INCREMENT,
  name varchar(64) NOT NULL,
  code varchar(64) NOT NULL,
  status tinyint NOT NULL DEFAULT 0 COMMENT '0:正常;1:禁用',
  remark varchar(255) DEFAULT NULL,
  create_time datetime DEFAULT NULL,
  update_time datetime DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY uk_code (code)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS sys_dict (
  id bigint NOT NULL AUTO_INCREMENT,
  type_code varchar(64) NOT NULL,
  name varchar(64) NOT NULL,
  value varchar(64) NOT NULL,
  sort int NOT NULL DEFAULT 0,
  status tinyint NOT NULL DEFAULT 1 COMMENT '1:正常;0:禁用',
  defaulted tinyint NOT NULL DEFAULT 0,
  remark varchar(255) DEFAULT NULL,
  create_time datetime DEFAULT NULL,
  update_time datetime DEFAULT NULL,
  PRIMARY KEY (id),
  KEY idx_type_code (type_code)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
