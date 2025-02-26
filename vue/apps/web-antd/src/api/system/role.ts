
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

