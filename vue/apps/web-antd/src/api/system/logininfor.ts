import { requestClient } from '#/api/request';

export interface SysLogininforListReq {
  ipaddr: string;
  userName: string;
  status: string;
  loginTime: string[];
}

export interface SysLogininfor {
    infoId: number;
    tenantId: string;
    userName: string;
    clientKey: string;
    deviceType: string;
    ipaddr: string;
    loginLocation: string;
    browser: string;
    os: string;
    status: string;
    msg: string;
    loginTime: string;
}

export interface SysLogininforListRes {
  items: SysLogininfor[];
  total: number;
}

export interface SysLogininforDeleteReq {
  infoIds: number[];
} 

export interface SysLogininforDeleteRes {
  infoIds: number[];
}

export async function getSysLogininforListApi(params: SysLogininforListReq) {
  return requestClient.get<SysLogininforListRes>('/system/logininfor/list', { params });
}

export async function deleteSysLogininforApi(params: SysLogininforDeleteReq) {
  return requestClient.post<SysLogininforDeleteRes>('/system/logininfor/delete', { ...params });
}
