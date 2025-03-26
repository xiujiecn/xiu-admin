import { requestClient } from '#/api/request';

export interface SysUserOnlineListReq { 
  page: number;
  pageSize: number;
  ipaddr: string;
  userName: string;
  status: string;
}

export interface SysUserOnline {
    onlineId: number;
    tenantId: string;
    uuid: string;
    userName: string;
    clientKey: string;
    deviceType: string;
    ipaddr: string;
    loginLocation: string;
    browser: string;
    os: string;
    token: string;
    loginTime: string;
    expireTime: string;
}

export interface SysUserOnlineListRes {
  items: SysUserOnline[];
  total: number;
}

export interface SysUserOnlineDeleteReq {
  ids: number[];
}

export interface SysUserOnlineDeleteRes {
}

export const getSysUserOnlineListApi = (params: SysUserOnlineListReq) => {
  return requestClient.get<SysUserOnlineListRes>('/system/user-online/list', { params });
};

export const deleteSysUserOnlineApi = (params: SysUserOnlineDeleteReq) => {
  return requestClient.post<SysUserOnlineDeleteRes>('/system/user-online/delete', { ...params });
};



