import { requestClient } from '#/api/request';


export interface PostListQuery {
  page: number;
  pageSize: number;
  postCode: string;
  postName: string;
}

export interface PostListData {
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

export interface PostListRes {
  data: PostListData[];
  total: number;
}

export async function getPostListApi(params: PostListQuery) {
  return requestClient.get<PostListRes>('/post/list', { params });
}
