
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

export const getSysClientListApi = (params: SysClientListReq) => {
  return requestClient.get<SysClientListRes>('/system/client/list', { params });
};
