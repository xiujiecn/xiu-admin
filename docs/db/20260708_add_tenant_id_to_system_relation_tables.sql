-- 变动日期时间：2026-07-08
-- 变动原因：补齐定时任务和系统关联表租户字段，避免多租户角色、菜单、岗位、任务数据串用。
-- 变动内容：sys_job、sys_role_dept、sys_role_menu、sys_user_role、sys_user_post 新增 tenant_id 字段；关联表主键纳入 tenant_id。

ALTER TABLE `sys_job`
  ADD COLUMN `tenant_id` varchar(20) DEFAULT '000000' COMMENT '租户编号' AFTER `job_id`;

UPDATE `sys_job`
SET `tenant_id` = '000000'
WHERE `tenant_id` IS NULL OR `tenant_id` = '';

UPDATE `sys_job` j
LEFT JOIN `sys_dept` d ON d.`dept_id` = j.`created_dept`
LEFT JOIN `sys_user` u ON u.`user_id` = j.`created_by`
SET j.`tenant_id` = COALESCE(NULLIF(d.`tenant_id`, ''), NULLIF(u.`tenant_id`, ''), '000000');

ALTER TABLE `sys_role_dept`
  ADD COLUMN `tenant_id` varchar(20) NOT NULL DEFAULT '000000' COMMENT '租户编号' FIRST;

UPDATE `sys_role_dept`
SET `tenant_id` = '000000'
WHERE `tenant_id` IS NULL OR `tenant_id` = '';

UPDATE `sys_role_dept` rd
LEFT JOIN `sys_role` r ON r.`role_id` = rd.`role_id`
LEFT JOIN `sys_dept` d ON d.`dept_id` = rd.`dept_id`
SET rd.`tenant_id` = COALESCE(NULLIF(r.`tenant_id`, ''), NULLIF(d.`tenant_id`, ''), '000000');

ALTER TABLE `sys_role_dept`
  DROP PRIMARY KEY,
  ADD PRIMARY KEY (`tenant_id`, `role_id`, `dept_id`);

ALTER TABLE `sys_role_menu`
  ADD COLUMN `tenant_id` varchar(20) NOT NULL DEFAULT '000000' COMMENT '租户编号' FIRST;

UPDATE `sys_role_menu`
SET `tenant_id` = '000000'
WHERE `tenant_id` IS NULL OR `tenant_id` = '';

UPDATE `sys_role_menu` rm
LEFT JOIN `sys_role` r ON r.`role_id` = rm.`role_id`
SET rm.`tenant_id` = COALESCE(NULLIF(r.`tenant_id`, ''), '000000');

ALTER TABLE `sys_role_menu`
  DROP PRIMARY KEY,
  ADD PRIMARY KEY (`tenant_id`, `role_id`, `menu_id`);

ALTER TABLE `sys_user_role`
  ADD COLUMN `tenant_id` varchar(20) NOT NULL DEFAULT '000000' COMMENT '租户编号' FIRST;

UPDATE `sys_user_role`
SET `tenant_id` = '000000'
WHERE `tenant_id` IS NULL OR `tenant_id` = '';

UPDATE `sys_user_role` ur
LEFT JOIN `sys_user` u ON u.`user_id` = ur.`user_id`
LEFT JOIN `sys_role` r ON r.`role_id` = ur.`role_id`
SET ur.`tenant_id` = COALESCE(NULLIF(u.`tenant_id`, ''), NULLIF(r.`tenant_id`, ''), '000000');

ALTER TABLE `sys_user_role`
  DROP PRIMARY KEY,
  ADD PRIMARY KEY (`tenant_id`, `user_id`, `role_id`);

ALTER TABLE `sys_user_post`
  ADD COLUMN `tenant_id` varchar(20) NOT NULL DEFAULT '000000' COMMENT '租户编号' FIRST;

UPDATE `sys_user_post`
SET `tenant_id` = '000000'
WHERE `tenant_id` IS NULL OR `tenant_id` = '';

UPDATE `sys_user_post` up
LEFT JOIN `sys_user` u ON u.`user_id` = up.`user_id`
LEFT JOIN `sys_post` p ON p.`post_id` = up.`post_id`
SET up.`tenant_id` = COALESCE(NULLIF(u.`tenant_id`, ''), NULLIF(p.`tenant_id`, ''), '000000');

ALTER TABLE `sys_user_post`
  DROP PRIMARY KEY,
  ADD PRIMARY KEY (`tenant_id`, `user_id`, `post_id`);
