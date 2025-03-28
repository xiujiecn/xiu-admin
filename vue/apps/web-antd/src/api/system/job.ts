/**
 * @description 定时任务相关接口
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */

import { requestClient } from '#/api/request';

export interface SysJobListReq {
  jobName: string;
  jobGroup: string;
  status: string;
  page: number;
  pageSize: number;
}

export interface SysJobSaveReq {
  jobId: number;
  jobName: string;
  jobGroup: string;
  jobParams: string;
  invokeTarget: string;
  cronExpression: string;
  misfirePolicy: string;
  concurrent: string;
  status: string;
  remark: string;
}


export interface SysJob {
  jobId: number;
  jobName: string;
  jobGroup: string;
  jobParams: string;
  invokeTarget: string;
  cronExpression: string;
  misfirePolicy: string | number | null;
  concurrent: string | number | null;
  status: string;
  remark: string;
}

export interface SysJobListRes {
  items: SysJob[];
  total: number;
}

export interface SysJobDeleteReq {
  jobIds: number[];
} 

export interface SysJobDeleteRes {
  jobIds: number[];
}

export async function getSysJobListApi(params: SysJobListReq) {
  return requestClient.get<SysJobListRes>('/system/job/list', { params });
}

export async function viewSysJobApi(params: {jobId: number}) {
  return requestClient.get<SysJob>('/system/job/view', { params });
}

export async function addSysJobApi(params: SysJobSaveReq|any) {
  return requestClient.post<SysJob>('/system/job/add', params);
}

export async function updateSysJobApi(params: SysJobSaveReq|any) {
  return requestClient.post<SysJob>('/system/job/update', params );
}
export async function updateStatusApi(params: SysJobSaveReq|any) {
  return requestClient.post<SysJob>('/system/job/status', params );
}

export async function deleteSysJobApi(params: SysJobDeleteReq) {
  return requestClient.post<SysJobDeleteRes>('/system/job/delete', { ...params });
}

export async function execOnceApi(params: {jobId: number}) {
  return requestClient.post<SysJob>('/system/job/exec', params);
}

