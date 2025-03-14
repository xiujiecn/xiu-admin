import { requestClient } from '#/api/request';
import type { AxiosRequestConfig } from '@vben/request';
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

export interface SysOssListModel {
  items: SysOssListData[];
  total: number;
}

export interface SysOssDownloadParam {
  ossId: number;
}

export interface SysOssDownloadModel {
  ossId: number;
}

export interface SysOssDeleteParam {
  ossIds: number[];
}

export interface SysOssDeleteModel {
  ossIds: number[];
}

export interface SysOssViewParam {
  ossId: number;
}

export interface SysOssViewModel {
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

export async function getSysOssListApi(params: SysOssListParam) {
  return requestClient.get<SysOssListModel>('/system/oss/list', { params });
}

export async function getSysOssDownloadApi(params: SysOssDownloadParam) {
  return requestClient.get<SysOssDownloadModel>('/system/oss/download', { params });
}

export async function getSysOssViewApi(params: SysOssViewParam) {
  return requestClient.get<SysOssViewModel>('/system/oss/view', { params });
}

export async function deleteSysOssApi(params: SysOssDeleteParam) {
  return requestClient.post<SysOssDeleteModel>('/system/oss/delete', { ...params });
}

export type ID = number | string;
export function ossDownload(
  ossId: ID,
  onDownloadProgress?: AxiosRequestConfig['onDownloadProgress'],
) {
  return requestClient.get<Blob>(`/common/oss/download?ossId=${ossId}`, {
    responseType: 'blob',
    timeout: 30 * 1000,
    isTransformResponse: false,
    onDownloadProgress,
  });
}

export function ossUpload(file: Blob | File) {
  const formData = new FormData();
  formData.append('file', file);
  return requestClient.post('/common/oss/upload', formData, {
    headers: { 'Content-Type': 'multipart/form-data;charset=UTF-8' },
    timeout: 30 * 1000,
  });
}
