import { requestClient } from '#/api/request';

export interface DeptListQuery {
  deptName: string;
  status: string;
}

export interface DeptListData {
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

export interface DeptListRes {
  data: DeptListData[];
  total: number;
}


export async function getDeptListApi(params: DeptListQuery) {
  return requestClient.get<DeptListRes>('/dept/list', { params });
}
