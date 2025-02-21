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
    return requestClient.get<SysOperLogListRes>('/oper-log/list', { params });
}