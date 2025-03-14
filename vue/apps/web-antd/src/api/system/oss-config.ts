import { requestClient } from '#/api/request';

export const sysOssConfigApi = {
  list: (params: any) => requestClient.get('/system/oss-config/list', { params }),
  view: (params: any) => requestClient.get('/system/oss-config/view', { params }),
  add: (params: any) => requestClient.post('/system/oss-config/add', { ...params }),
  edit: (params: any) => requestClient.post('/system/oss-config/edit', { ...params }),
  delete: (params: any) => requestClient.post('/system/oss-config/delete', { ...params }),
};

export interface SysOssConfigListParam {
  page: number;
  pageSize: number;
  configKey: string;
  bucketName: string;
  status: string;
}   
export interface SysOssConfigListRes {
  items: SysOssConfigListData[];
  total: number;
}
export interface SysOssConfigListData {
    ossConfigId: number;
    configKey: string;
    accessKey: string;
    secretKey: string;
    bucketName: string;
    prefix: string;
    endpoint: string;
    domain: string;
    isHttps: string;
    region: string;
    accessPolicy: string;
    status: string;
    ext1: string;
    createdDept: number;
    createdBy: number;
    createdAt: string;
    remark: string;
}
export interface SysOssConfigViewParam {
  ossConfigId: number;
}
export interface SysOssConfigViewRes {
  ossConfigId: number;
  configKey: string;
  accessKey: string;
  secretKey: string;
  bucketName: string;
  prefix: string;
  endpoint: string;
  domain: string;
  isHttps: string;
  region: string;
  accessPolicy: string;
  status: string;
  ext1: string;
  createdDept: number;
  createdBy: number;
  createdAt: string;
  remark: string;
}
export interface SysOssConfigAddParam {
    configKey: string;
    accessKey: string;
    secretKey: string;
    bucketName: string;
    prefix: string;
    endpoint: string;
    domain: string;
    isHttps: string;
    region: string;
    accessPolicy: string;
    status: string;
    ext1: string;
    remark: string;
}
export interface SysOssConfigAddRes {
  ossConfigId: number;
}
export interface SysOssConfigEditParam {
    ossConfigId: number;
    configKey?: string;
    accessKey?: string;
    secretKey?: string;
    bucketName?: string;
    prefix?: string;
    endpoint?: string;
    domain?: string;
    isHttps?: string;
    region?: string;
    accessPolicy?: string;
    status?: string;
    ext1?: string;
    remark?: string;
}
export interface SysOssConfigEditRes {
  ossConfigId: number;
}
export interface SysOssConfigDeleteParam {
  ossConfigIds: number[];
}
export interface SysOssConfigDeleteRes {
  ossConfigIds: number[];
}

export async function getSysOssConfigListApi(params: SysOssConfigListParam) {
  return requestClient.get<SysOssConfigListRes>('/system/oss-config/list', { params });
}
export async function getSysOssConfigViewApi(params: SysOssConfigViewParam) {
  return requestClient.get<SysOssConfigViewRes>('/system/oss-config/view', { params });
}
export async function addSysOssConfigApi(params: SysOssConfigAddParam |{
  [x: string]: any;
}) {
  return requestClient.post<SysOssConfigAddRes>('/system/oss-config/add', { ...params });
}
export async function editSysOssConfigApi(params: SysOssConfigEditParam|{
  [x: string]: any;
}) {
  return requestClient.post<SysOssConfigEditRes>('/system/oss-config/edit', { ...params });
}
export async function deleteSysOssConfigApi(params: SysOssConfigDeleteParam|{
  [x: string]: any;
}) {
  return requestClient.post<SysOssConfigDeleteRes>('/system/oss-config/delete', { ...params });
}
