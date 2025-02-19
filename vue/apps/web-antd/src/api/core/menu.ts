import type { RouteRecordStringComponent } from '@vben/types';

import { requestClient } from '#/api/request';

/**
 * 获取用户所有菜单
 */
export async function getAllMenusApi() {
  return requestClient.get<RouteRecordStringComponent[]>('/menu/all');
}

export interface MenuListReq {
  menuName: string;
  status: string;
}

export interface MenuListData {
  menuId:number;
  menuName: string;
  parentId:number;
  orderNum:number;
  path:string;
  component:string;
  queryParam:string;
  isFrame:number;
  isCache:number;
  menuType:string;
  visible:string;
  status:string;
  perms:string;
  icon:string;
  createdDept:number;
  createdBy:number;
  createdAt:string;
  remark:string;
}

export interface MenuListRes {
  data: MenuListData[];
  total: number;
}

export async function getMenuListApi(params: MenuListReq) {
  return requestClient.get<MenuListRes>('/menu/list', { params });
}
