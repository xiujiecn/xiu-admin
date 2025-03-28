import { requestClient } from '#/api/request';

// 获取测试分类列表
export function List(params:any) {
  return requestClient.get<any>('gen/testCategory/list', { params });
}

// 删除/批量删除测试分类
export function Delete(params:any) {
  return requestClient.post<any>('gen/testCategory/delete', { ...params });
}

// 添加/编辑测试分类
export function Edit(params:any) {
  return requestClient.post<any>('gen/testCategory/edit', { ...params });
}

// 修改测试分类状态
export function Status(params:any) {
  return requestClient.post<any>('gen/testCategory/status', { ...params });
}

// 获取测试分类指定详情
export function View(params:any) {
  return requestClient.get<any>('gen/testCategory/view', { params });
}

// 获取测试分类最大排序
export function MaxSort(params:any) {
  return requestClient.get<any>('gen/testCategory/maxSort', { params });
}

// 导出测试分类
export function Export(params:any) {
  return requestClient.post<Blob>('/gen/testCategory/export',  { ...params } , {
    headers: { 'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8' },
    responseType: 'blob',
  });
}