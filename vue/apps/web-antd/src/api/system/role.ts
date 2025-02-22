
import { requestClient } from '#/api/request';

export interface SysRoleListQuery {
  page: number;
  pageSize: number;
  roleName: string;
  roleKey: string;
  status: string;
  createdAt: string;
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
  data: SysRoleListData[];
  total: number;
}

export async function getSysRoleListApi(params: SysRoleListQuery) {
  return requestClient.get<SysRoleListRes>('/system/role/list', { params });
}
