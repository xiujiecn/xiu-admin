/**
 * @description 通知公告相关接口
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */

import { requestClient } from '#/api/request';
import type { SysUserMini } from '#/api/system/user';
export interface SysNoticeListReq {
  noticeTitle: string;
  noticeType: string;
  createdBy: string;
}

export interface SysNotice {
    noticeId: number;
    tenantId:string;
    noticeTitle:string;
    noticeType:string;
    noticeContent:string;
    status:string;
    createdDept:number;
    createdBy:number;
    createdAt:string;
    remark:string;
    createdByUser:SysUserMini;
}

export interface SysNoticeListRes {
  items: SysNotice[];
  total: number;
}

export interface SysNoticeAddReq {
  noticeTitle: string;
  noticeType: string;
  noticeContent: string;
  status: string;
  remark: string;
}

export interface SysNoticeEditReq {
  noticeId: number;
  noticeTitle?: string;
  noticeType?: string;
  noticeContent?: string;
  status?: string;
  remark?: string;
}

export interface SysNoticeDeleteReq {
  noticeIds: number[];
}

export interface SysNoticeViewReq {
  noticeId : number;
}

export interface SysNoticeViewRes extends SysNotice {
  userIdList: any;
  deptIdList: any;
}

export interface SysNoticeAddRes {
  noticeId: number;
}

export interface SysNoticeEditRes  {
  noticeId: number;
}

export interface SysNoticeDeleteRes  {
  noticeIds: number[];
}



export async function getSysNoticeListApi(params: SysNoticeListReq) {
  return requestClient.get<SysNoticeListRes>('/system/notice/list', { params });
}

export async function addSysNoticeApi(params: SysNoticeAddReq | {
  [x: string]: any;
}) {
  return requestClient.post<SysNoticeAddRes>('/system/notice/add', { ...params });
}


export async function editSysNoticeApi(params: SysNoticeEditReq| {
  [x: string]: any;
}) {
  return requestClient.post<SysNoticeEditRes>('/system/notice/edit', { ...params });
}

export async function deleteSysNoticeApi(params: SysNoticeDeleteReq) {
  return requestClient.post<SysNoticeDeleteRes>('/system/notice/delete', { ...params });
}

export async function getSysNoticeApi(params: SysNoticeViewReq) {
  return requestClient.get<SysNoticeViewRes>('/system/notice/view', { params });
}
