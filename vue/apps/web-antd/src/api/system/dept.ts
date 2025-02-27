import { requestClient } from '#/api/request';

export interface SysDeptMini {
  deptId: number;
  deptName: string;
}

export interface SysDeptListParam {
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
  createdAt: string;
}

export interface SysDeptListRes {
  items: SysDeptListData[];
  total: number;
}


export async function getSysDeptListApi(params: SysDeptListParam) {
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

export interface SysDeptAddParam {
  parentId: number;
  ancestors: string;
  deptName: string;
  deptCategory: string;
  orderNum: number;
  leader: number;
  phone: string;
  email: string;
  status: string;
}

export interface SysDeptAddRes {
  deptId: number;
}

export async function addSysDeptApi(params: SysDeptAddParam | {
  [x: string]: any;
}) {
  return requestClient.post<SysDeptAddRes>('/system/dept/add', params);
}

export interface SysDeptEditParam {
  deptId: number;
  parentId?: number;
  ancestors?: string;
  deptName?: string;
  deptCategory?: string;
  orderNum?: number;
  leader?: number;
  phone?: string;
  email?: string;
  status?: string;
}

export interface SysDeptEditRes {
  deptId: number;
}

export async function editSysDeptApi(params: SysDeptEditParam | {
  [x: string]: any;
}) {
  return requestClient.post<SysDeptEditRes>('/system/dept/edit', params);
}

export interface SysDeptDeleteParam {
  deptId: number;
} 

export interface SysDeptDeleteRes {
  deptId: number;
}

export async function deleteSysDeptApi(params: SysDeptDeleteParam) {
  return requestClient.post<SysDeptDeleteRes>('/system/dept/delete', params);
}

export interface SysDeptViewParam {
  deptId: number;
} 

export interface SysDeptViewRes {
  deptId: number;
  tenantId: string;
  parentId: number;
  ancestors: string;
  deptName: string;
  deptCategory: string;
  orderNum: number;
  leader: number | null;
  phone: string;
  email: string;
  status: string;
  createdDept: number;
  createdBy: number;
  createdAt: string;
  updatedBy: number;
  updatedAt: string;
} 

export async function viewSysDeptApi(params: SysDeptViewParam) {
  return requestClient.get<SysDeptViewRes>('/system/dept/view', { params });
}
