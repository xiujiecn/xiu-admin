/**
 * @description 岗位管理相关接口
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */

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

export interface SysPostAddParam {
  deptId: number;
  postCode: string;
  postCategory: string;
  postName: string;
  postSort: number;
  status: string;
  remark: string;
}

export interface SysPostAddRes {
  postId: number;
}

export async function addSysPostApi(params: SysPostAddParam|{
  [x: string]: any;
}) {
  return requestClient.post<SysPostAddRes>('/system/post/add', { ...params });
}

export interface SysPostEditParam {
  postId: number;
  deptId?: number;
  postCode?: string;
  postCategory?: string;
  postName?: string;
  postSort?: number;
  status?: string;
  remark?: string;
}

export interface SysPostEditRes {
  postId: number;
}

export async function editSysPostApi(params: SysPostEditParam|{
  [x: string]: any;
}) {
  return requestClient.post<SysPostEditRes>('/system/post/edit', { ...params });
}

export interface SysPostDeleteParam {
  postId?: number;
  postIds?: number[];
}

export interface SysPostDeleteRes {
  postId: number;
} 

export async function deleteSysPostApi(params: SysPostDeleteParam) {
  return requestClient.post<SysPostDeleteRes>('/system/post/delete', { ...params });
}

export interface SysPostViewParam {
  postId: number;
}

export interface SysPostViewRes {
  postId: number;
  tenantId: string;
  deptId: number;
  postCode: string;
  postCategory: string;
  postName: string;
  postSort: number;
  status: string;
  remark: string;
  createdDept: number;
  createdBy: number;
  createdAt: string;
  deptInfo: SysDeptMini;
}

export async function getSysPostViewApi(params: SysPostViewParam) {
  return requestClient.get<SysPostViewRes>('/system/post/view', { params });
}

export interface SysPostExportParam extends SysPostListParam {
}

export interface SysPostExportRes {
}

export async function getSysPostExportApi(params: SysPostExportParam) {
  return requestClient.post<Blob>('/system/post/export',  params , {
    headers: { 'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8' },
    responseType: 'blob',
  });
}
