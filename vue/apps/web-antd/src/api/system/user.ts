import { requestClient } from '#/api/request';
export interface SysUserListQuery {
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
  export interface SysUserListData {
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
  export interface SysUserListRes {
    data: SysUserListData[];
    total: number;
  }
  
  export async function getSysUserListApi(params: SysUserListQuery) {
    return requestClient.get<SysUserListRes>('/user/list', { params });
  }
  