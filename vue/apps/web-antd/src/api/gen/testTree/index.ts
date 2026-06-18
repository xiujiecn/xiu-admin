import { requestClient } from '#/api/request';

// 获取测试树表列表
export function List(params:any) {
  return requestClient.get<any>('gen/testTree/list', { params });
}

// 删除/批量删除测试树表
export function Delete(params:any) {
  return requestClient.post<any>('gen/testTree/delete', { ...params });
}

// 添加/编辑测试树表
export function Edit(params:any) {
  return requestClient.post<any>('gen/testTree/edit', { ...params });
}

// 获取测试树表指定详情
export function View(params:any) {
  return requestClient.get<any>('gen/testTree/view', { params });
}

// 导出测试树表
export function Export(params:any) {
  return requestClient.post<Blob>('/gen/testTree/export',  { ...params } , {
    headers: { 'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8' },
    responseType: 'blob',
  });
}

// 获取测试树表关系树选项
export function TreeOption() {
  return requestClient.get<any>('gen/testTree/treeOption');
}