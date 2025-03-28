/**
 * @description 角色管理相关接口
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */

import { requestClient } from '#/api/request';

export interface SysRoleListParam {
  page?: number;
  pageSize?: number;
  roleName?: string;
  roleKey?: string;
  status?: string;
  createdAt?: string[];
}

export interface SysRoleListData {
    roleId: number; 
    tenantId: string;
    roleName: string;
    roleKey: string;
    roleSort: number;
    dataScope: string;
    menuCheckStrictly: number;
    deptCheckStrictly: number;
    status: string;
    createdDept: number;
    createdAt: string;
    remark: string;
}

export interface SysRoleViewData {
  roleId: number; 
  tenantId: string;
  roleName: string;
  roleKey: string;
  roleSort: number;
  dataScope: string;
  menuCheckStrictly: number;
  deptCheckStrictly: number;
  status: string;
  createdDept: number;
  createdAt: string;
  remark: string;
  menuIds: number[];
  deptIds: number[];
}

export interface SysRoleListRes {
  items: SysRoleListData[];
  total: number;
}

export async function getSysRoleListApi(params: SysRoleListParam) {
  return requestClient.get<SysRoleListRes>('/system/role/list', { params });
}

export interface SysRoleMini {
  roleId: number;
  roleName: string;
  dataScope: string;
}
export interface SysRoleViewParam {
  roleId: number;
}

export interface SysRoleAddParam {
  roleName: string;
  roleKey: string;
  roleSort: number;
  menuCheckStrictly: number;
  deptCheckStrictly: number;
  status: string;
  remark: string;
  menuIds: number[];
}

export interface SysRoleEditParam {
  roleId: number;
  roleName: string;
  roleKey: string;
  roleSort: number;
  menuCheckStrictly: number;
  deptCheckStrictly: number;
  status: string;
  remark: string;
  menuIds: number[];
}

export interface SysRoleDeleteParam {
  roleId: number;
  roleIds: number[];
}

export interface SysRoleAddRes {
}

export interface SysRoleEditRes {
}

export interface SysRoleDeleteRes {
}

export interface SysRoleDataScopeEditParam {
  roleId: number;
  dataScope: string;
  deptCheckStrictly?: number;
  deptIds?: number[];
}
export interface SysRoleDataScopeEditRes {
}

export async function addSysRoleApi(params: SysRoleAddParam|any) {
  return requestClient.post<SysRoleAddRes>('/system/role/add', { ...params });
}

export async function editSysRoleApi(params: SysRoleEditParam|any) {
  return requestClient.post<SysRoleEditRes>('/system/role/edit', { ...params });
}

export async function deleteSysRoleApi(params: SysRoleDeleteParam|any) {
  return requestClient.post<SysRoleDeleteRes>('/system/role/delete', { ...params });
} 

export async function getSysRoleViewApi(params: SysRoleViewParam|any) {
  console.log('/Users/lixiujie/dev_test/dev_go/xiujie_iot/vue/apps/web-antd/src/api/system/role.ts', params);  
  return requestClient.get<SysRoleViewData>('/system/role/view', { params });
}

export async function editSysRoleDataScopeApi(params: SysRoleDataScopeEditParam|any) {
  return requestClient.post<SysRoleDataScopeEditRes>('/system/role/dataScope', { ...params });
}
