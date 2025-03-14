
import { requestClient } from '#/api/request';

export interface SysClientListReq {
  clientId: string;
  clientKey: string;
  clientSecret: string;
  status: string;
}

export interface SysClient {
    id: number;     
    clientId: string;
    clientKey: string;
    clientSecret: string;
    grantType: string;
    deviceType: string;
    activeTimeout: number;
    timeout: number;
    status: string;
    createdDept: number;
    createdBy: number;
    createdAt: string;
}

export interface SysClientListRes {
  items: SysClient[];
  total: number;
}

export interface SysClientViewReq {
  id: number;
}

export interface SysClientViewRes {
  id: number;     
  clientId: string;
  clientKey: string;
  clientSecret: string;
  grantType: string;
  grantTypeList?: string[];
  deviceType: string;
  activeTimeout: number;
  timeout: number;
  status: string;
  createdDept: number;
  createdBy: number;
  createdAt: string;
  updatedDept: number;
  updatedBy: number;
  updatedAt: string;
} 

export interface SysClientAddReq {
  clientId: string;
  clientKey: string;
  clientSecret: string;
  grantType: string;
  deviceType: string;
  activeTimeout: number;
  timeout: number;
  status: string;
} 

export interface SysClientAddRes {
  id: number;
} 

export interface SysClientEditReq {
  id: number;
  clientId?: string;
  clientKey?: string;
  clientSecret?: string; 
  grantType?: string;
  deviceType?: string;
  activeTimeout?: number;
  timeout?: number;
  status?: string;
} 

export interface SysClientEditRes {
  id: number;
}  

export interface SysClientDeleteReq {
  ids: number[];
}

export interface SysClientDeleteRes {
  ids: number[];
} 

export interface SysClientStatusReq {
  id: number;
  status: string;
}   

export interface SysClientStatusRes {
  id: number;
} 

export const getSysClientListApi = (params: SysClientListReq) => {
  return requestClient.get<SysClientListRes>('/system/client/list', { params });
};

export const getSysClientViewApi = (params: SysClientViewReq) => {
  return requestClient.get<SysClientViewRes>('/system/client/view', { params });
};

export const postSysClientAddApi = (params: SysClientAddReq|{
  [x: string]: any;
}) => {
  return requestClient.post<SysClientAddRes>('/system/client/add', { ...params });
};

export const postSysClientEditApi = (params: SysClientEditReq|{
  [x: string]: any;
}) => {
  return requestClient.post<SysClientEditRes>('/system/client/edit', { ...params });
};

export const postSysClientDeleteApi = (params: SysClientDeleteReq|{
  [x: string]: any;
}) => {
  return requestClient.post<SysClientDeleteRes>('/system/client/delete', { ...params });
};

export const postSysClientStatusApi = (params: SysClientStatusReq|{
  [x: string]: any;
}) => {
  return requestClient.post<SysClientStatusRes>('/system/client/status', { ...params });
};  