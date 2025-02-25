import { requestClient } from '#/api/request';
import type { SysDeptMini } from './dept';
import type { SysRoleMini } from './role';
import type { SysPostMini } from './post';
import type { UserInfo } from '@vben/types';
import type { Recordable } from '@vben/types';
import {md5} from 'js-md5';

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
    deptId: number | undefined;
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
    remark: string;
    deptInfo: SysDeptMini;
    roles: SysRoleMini[];
    posts: SysPostMini[];
  }



  export interface SysUserViewModel extends UserProfileData {

  }

  export const EmptySysUserViewModel: SysUserViewModel = {
    userId: 0,
    tenantId: '',
    deptId: undefined,
    userName: '',
    nickName: '',
    userType: '',
    email: '',
    phonenumber: '',
    sex: '0',
    avatar: '',
    status: '0',
    loginIp: '',
    loginDate: '',
    createdDept: 0,
    createdBy: 0,
    createdAt: '',
    remark: '',
    deptInfo: {
      deptId: 0,
      deptName: '',
    },
    roles: [],
    posts: [],
  }

  export interface SysUserViewParam {
    userId: number;
  }

  export interface SysUserAddParam {
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
    remark: string;
  }

  export interface SysUserUpdateParam {
    userId: number;
    deptId?: number;
    userName?: string;
    nickName?: string;
    userType?: string;
    email?: string;
    phonenumber?: string;
    sex?: string;
    avatar?: string;
    status?: string;
    updatedBy?: number;
    updatedAt?: string;
    remark?: string;
  }

  export interface SysUserDeleteParam {
    userId?: number;
    userIds?: number[];
  }

  export interface SysUserResetPasswordParam {
    userId: number;
    password: string;
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
    params.oldPassword = md5(params.oldPassword);
    params.newPassword = md5(params.newPassword);
    return requestClient.post('/system/user/profile/password', { ...params });
  }
  export async function updateCurrentUserProfile(params: UserProfileData|Recordable<any>) {
    return requestClient.post<UserInfo>('/system/user/profile/update', { ...params });
  }

  export async function getSysUser(params: SysUserViewParam) {
    return requestClient.get<SysUserViewModel>('/system/user/view', { params });
  }

  export async function addSysUser(params: SysUserAddParam|{
    [x: string]: any;
}) {
    return requestClient.post<SysUserViewModel>('/system/user/add', { ...params });
  }

  export async function updateSysUser(params: SysUserUpdateParam|{
    [x: string]: any;
}) {
    return requestClient.post<SysUserViewModel>('/system/user/update', { ...params });
  }

  export async function deleteSysUser(params: SysUserDeleteParam) {
    return requestClient.post('/system/user/delete', { ...params });
  } 

  export async function resetSysUserPassword(params: SysUserResetPasswordParam) {
    params.password = md5(params.password);
    return requestClient.post('/system/user/resetPassword', { ...params });
  }
  