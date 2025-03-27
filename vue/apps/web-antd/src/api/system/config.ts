/**
 * @description 系统配置相关接口
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 XiuAdmin CLI
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */

import { requestClient } from '#/api/request';

export interface GetSysConfigListReq {
  configName: string;
  configKey: string;
  configType: string;
  createdAt: string;
}

export interface SysConfig {
    configId: number;
    tenantId:string;
    configName:string;
    configKey:string;
    configValue:string;
    configType:string;
    createdDept:number;
    createdBy:number;
    createdAt:string;
    remark:string;
}

export interface GetSysConfigListRes {
  items: SysConfig[];
  total: number;
}

export interface AddSysConfigReq {
  configName: string;
  configKey: string;
  configValue: string;
  configType: string;
  remark: string;
} 

export interface AddSysConfigRes {
  configId: number;
}

export interface EditSysConfigReq {
  configId: number;
  configName: string;
  configKey: string;
  configValue: string;
  configType: string;
  remark: string;
}

export interface EditSysConfigRes {
  configId: number;
} 

export interface DeleteSysConfigReq {
  configIds: number[];
}

export interface DeleteSysConfigRes {
  configIds: number[];
} 

export interface GetSysConfigByIdReq {
  configId: number;
}

export interface GetSysConfigByIdRes {
  configId: number;
  tenantId: string;
  configName: string;
  configKey: string;
  configValue: string;
  configType: string;
  createdDept: number;
  createdBy: number;
  createdAt: string;
  remark: string;
}

export async function getSysConfigListApi(params: GetSysConfigListReq) {
  return requestClient.get<GetSysConfigListRes>('/system/config/list', { params });
}

export async function addSysConfigApi(params: AddSysConfigReq | {
  [x: string]: any;
}) {
  return requestClient.post<AddSysConfigRes>('/system/config/add', { ...params });
}

export async function editSysConfigApi(params: EditSysConfigReq | {
  [x: string]: any;
}) {
  return requestClient.post<EditSysConfigRes>('/system/config/edit', { ...params });
}

export async function deleteSysConfigApi(params: DeleteSysConfigReq) {
  return requestClient.post<DeleteSysConfigRes>('/system/config/delete', { ...params });
} 

export async function getSysConfigByIdApi(params: GetSysConfigByIdReq) {
  return requestClient.get<GetSysConfigByIdRes>('/system/config/view', { params });
}
