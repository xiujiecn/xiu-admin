import { requestClient } from '#/api/request';

export interface SysOssListParam {
  page: number;
  pageSize: number;
  fileName: string;
  originalName: string;
  fileSuffix: string;
}

export interface SysOssListData {
    ossId: number;
    tenantId: string;
    fileName: string;
    originalName: string;
    fileSuffix: string;
    url: string;
    createdDept: number;
    createdAt: string;
    createdBy: number;
    service: string;
}

export interface SysOssListRes {
  items: SysOssListData[];
  total: number;
}

export async function getSysOssListApi(params: SysOssListParam) {
  return requestClient.get<SysOssListRes>('/system/oss/list', { params });
}
