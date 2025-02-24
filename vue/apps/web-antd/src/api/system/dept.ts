import { requestClient } from '#/api/request';

export interface SysDeptMini {
  deptId: number;
  deptName: string;
}

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

export interface SysDeptTreeData {
  deptId: number;
  parentId: number; 
  key: string;
  deptName  : string;
  children?: SysDeptTreeData[];
}
export interface SysDeptTreeRes {
  items: SysDeptTreeData[];
} 

export async function getSysDeptTreeApi() {
  return requestClient.get<SysDeptTreeRes>('/system/dept/tree');
}