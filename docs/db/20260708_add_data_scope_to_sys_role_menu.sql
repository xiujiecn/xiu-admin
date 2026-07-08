-- 变动日期时间：2026-07-08
-- 变动原因：支持角色菜单维度覆盖数据权限范围。
-- 变动内容：sys_role_menu 新增 data_scope 字段，默认 0 表示按角色数据权限。

ALTER TABLE `sys_role_menu`
  ADD COLUMN `data_scope` char(1) NOT NULL DEFAULT '0' COMMENT '数据范围（0按角色数据权限 1全部数据权限 3本部门数据权限 4本部门及以下数据权限 5仅本人数据权限 6本部门及以下或本人数据权限 7本组织及本组织下一级数据权限 8本组织下一级数据权限）' AFTER `menu_id`;

UPDATE `sys_role_menu`
SET `data_scope` = '0'
WHERE `data_scope` IS NULL OR `data_scope` = '';
