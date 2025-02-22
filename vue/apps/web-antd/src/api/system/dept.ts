import { requestClient } from '#/api/request';

export interface SysDeptListQuery {
  deptName: string;
  status: string;
}

export interface SysDeptListData {
  deptId: number;
  tenantId: string; 
  parentId: number; 
  ancestors : string;
  deptName  : string;
  deptCategory  : string;
  orderNum: number; 
  leader: number; 
  phone: string; 
  email: string; 
  status: string;
}

export interface SysDeptListRes {
  data: SysDeptListData[];
  total: number;
}


export async function getSysDeptListApi(params: SysDeptListQuery) {
  return requestClient.get<SysDeptListRes>('/system/dept/list', { params });
}
