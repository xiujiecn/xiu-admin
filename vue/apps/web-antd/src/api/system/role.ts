
import { requestClient } from '#/api/request';

export interface RoleListQuery {
  page: number;
  pageSize: number;
  roleName: string;
  roleKey: string;
  status: string;
  createdAt: string;
}

export interface RoleListData {
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

export interface RoleListRes {
  data: RoleListData[];
  total: number;
}

export async function getRoleListApi(params: RoleListQuery) {
  return requestClient.get<RoleListRes>('/role/list', { params });
}
