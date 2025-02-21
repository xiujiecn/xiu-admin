import { requestClient } from '#/api/request';


export interface SysPostListQuery {
  page: number;
  pageSize: number;
  postCode: string;
  postName: string;
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
}

export interface SysPostListRes {
  data: SysPostListData[];
  total: number;
}

export async function getSysPostListApi(params: SysPostListQuery) {
  return requestClient.get<SysPostListRes>('/post/list', { params });
}
