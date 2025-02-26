import { requestClient } from '#/api/request';

import type { SysDeptMini } from '#/api/system/dept';

export interface SysPostListParam {
  page?: number;
  pageSize?: number;
  postCode?: string;
  postName?: string;
  deptId?: number;
  belongDeptId?: number;
}

export interface SysPostListData {
    postId: number; 
    tenantId: string;
    deptId: number;
    postCode: string;
    postCategory: string;
    postName: string;
    postSort: number;
    status: string;
    remark: string;
    createdAt: string;
    deptInfo: SysDeptMini;
}

export interface SysPostListRes {
  items: SysPostListData[];
  total: number;
  page: number;
  pageSize: number;
}

export async function getSysPostListApi(params: SysPostListParam) {
  return requestClient.get<SysPostListRes>('/system/post/list', { params });
}

export interface SysPostMini {
  postId: number;
  deptId: number;
  postName: string;
}
