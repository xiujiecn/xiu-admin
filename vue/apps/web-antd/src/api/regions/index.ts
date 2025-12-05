import { requestClient } from '#/api/request';

// 获取行政区划列表
export function List(params:any) {
  return requestClient.get<any>('regions/iotRegions/list', { params });
}

// 获取指定省份下行政区划树形结构
export function Tree(params:any) {
  return requestClient.get<any>('regions/iotRegions/tree', { params });
}
