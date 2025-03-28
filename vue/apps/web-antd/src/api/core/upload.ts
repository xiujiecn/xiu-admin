/**
 * @description 文件上传相关接口
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */

import type { AxiosRequestConfig } from '@vben/request';

import { requestClient } from '#/api/request';

/**
 * Axios上传进度事件
 */
export type AxiosProgressEvent = AxiosRequestConfig['onUploadProgress'];

/**
 * 通过单文件上传接口
 * @param file 上传的文件
 * @param onUploadProgress 上传进度事件 非必传
 * @returns 上传结果
 */
export function uploadApi(
  file: Blob | File,
  onUploadProgress?: AxiosProgressEvent,
) {
  return requestClient.upload(
    '/common/oss/upload',
    { file,fileType: 'file' },
    { onUploadProgress, timeout: 60_000 },
  );
}
/**
 * 默认上传结果
 */
export interface UploadResult {
  url: string;
  fileName: string;
  ossId: string;
}
