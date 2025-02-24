import { requestClient } from '#/api/request';
import type { SysDeptMini } from './dept';
import type { SysRoleMini } from './role';
import type { SysPostMini } from './post';

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
    deptInfo: SysDeptMini;
  }
  export interface SysUserListRes {
    data: SysUserListData[];
    total: number;
  }
  
  export async function getSysUserListApi(params: SysUserListQuery) {
    return requestClient.get<SysUserListRes>('/system/user/list', { params });
  }
  

  export interface UserProfileRes {
    user: UserProfileData;
  }

  export interface UserProfileData {
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
    salt: string;
    password: string;
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
    deptInfo: SysDeptMini;
    roles: SysRoleMini[];
    posts: SysPostMini[];
  }
  export interface UpdateCurrentUserPasswordParam {
    oldPassword: string;
    newPassword: string;
  }

  export async function getUserProfileApi() {
    return requestClient.get<UserProfileRes>('/system/user/profile');
  }

  export async function updateCurrentUserAvatar(params: UserProfileData) {
    return requestClient.post<UserProfileRes>('/system/user/profile/avatar', { params });
  }

  export async function updateCurrentUserPassword(params: UpdateCurrentUserPasswordParam) {
    return requestClient.post<UserProfileRes>('/system/user/profile/password', { params });
  }
  export async function updateCurrentProfile(params: UserProfileData) {
    return requestClient.post<UserProfileRes>('/system/user/profile/password', { params });
  }