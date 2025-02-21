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

export async function getSysConfigListApi(params: GetSysConfigListReq) {
  return requestClient.get<GetSysConfigListRes>('/config/list', { params });
}
