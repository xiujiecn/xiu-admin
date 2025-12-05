import { requestClient } from '#/api/request';

/** 创建文件夹 */
export function Create(params: { parentId: number; displayName: string }) {
  return requestClient.post<any>('gallery/gallery/createFolder', { ...params });
}

/** 获取文件夹列表 */
export function List(params: {
  parentId: number | string;
  type?: 'folder' | 'image';
  keyword?: string;
  page?: number;
  pageSize?: number;
}) {
  return requestClient.get<any>('gallery/gallery/list', { params });
}

/** 重命名 */
export function Rename(params: { galleryId: number; displayName: string }) {
  return requestClient.post<any>('gallery/gallery/rename', { ...params });
}

/** 批量删除 */
export function Delete(params: { galleryIds: number[] }) {
  return requestClient.post<any>('gallery/gallery/batchDelete', { ...params });
}

/** 上传图片 */
export function Upload(params: {
  parentId: number;
  file: File;
  displayName?: string;
  newFileType?: number;
  saveOriginalName?: string;
}) {
  return requestClient.upload<any>('gallery/gallery/upload', {
    ...params,
    file: params.file,
  });
}
