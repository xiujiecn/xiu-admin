import { requestClient } from '#/api/request';

// 获取测试单表列表
export function List(params:any) {
  return requestClient.get<any>('gen/testDemo/list', { params });
}

// 删除/批量删除测试单表
export function Delete(params:any) {
  return requestClient.post<any>('gen/testDemo/delete', { ...params });
}

// 添加/编辑测试单表
export function Edit(params:any) {
  return requestClient.post<any>('gen/testDemo/edit', { ...params });
}

// 获取测试单表指定详情
export function View(params:any) {
  return requestClient.get<any>('gen/testDemo/view', { params });
}

// 导出测试单表
export function Export(params:any) {
  return requestClient.post<Blob>('/gen/testDemo/export',  { ...params } , {
    headers: { 'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8' },
    responseType: 'blob',
  });
}