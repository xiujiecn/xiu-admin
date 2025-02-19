import type { UserInfo } from '@vben/types';

import { requestClient } from '#/api/request';

/**
 * 获取用户信息
 */
export async function getUserInfoApi() {
  return requestClient.get<UserInfo>('/user/info');
}

export interface UserListQuery {
  tenantId: string;
  userId: number;
  deptId: number;
  userName: string;
  nickName: string;
  email: string;
  phonenumber: string;
  status: string;
  createdAt: string;
  page: number;
  pageSize: number;
}
export interface UserListData {
  userId: number; 
  tenantId: string;
  deptId: number;
  userName: string;
  nickName: string;
  userType: string;
  email: string;  
  phonenumber: string;
  sex: string;
  avatar: string;
  status: string;
  loginIp: string;
  loginDate: string;
  createdDept: number;
  createdBy: number;
  createdAt: string;
  updatedBy: number;
  updatedAt: string;
  deletedBy: number;
  deletedAt: string;
  remark: string;
}
export interface UserListRes {
  data: UserListData[];
  total: number;
}

export async function getUserListApi(params: UserListQuery) {
  return requestClient.get<UserListRes>('/user/list', { params });
}

