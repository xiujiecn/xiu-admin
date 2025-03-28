/**
 * @description 操作日志相关接口
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */

import { requestClient } from '#/api/request';

export interface SysOperLogListReq {
    operIp: string;
  businessType: string;
  method: string;
}

export interface SysOperLog {
    operId: number;
    tenantId: string;
    title: string;
    businessType: number;
    method: string;
    requestMethod: string;
    operatorType: number;
    operName: string;
    deptName: string;
    operUrl: string;
    operIp: string;
    operLocation: string;
    operParam: string;
    jsonResult: string;
    status: number;
    errorMsg: string;
    operTime: string;
    costTime: number;
}

export interface SysOperLogListRes {
  items: SysOperLog[];
  total: number;
}

export async function getSysOperLogList(params: SysOperLogListReq) {
    return requestClient.get<SysOperLogListRes>('/system/oper-log/list', { params });
}