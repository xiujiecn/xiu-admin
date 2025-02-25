import { requestClient } from '#/api/request';
export interface SysMenuListReq {
  menuName: string;
  status: string;
}

export interface SysMenuListData {
  menuId: number;
  menuName: string;
  parentId: number;
  orderNum: number;
  path: string;
  component: string;
  queryParam: string;
  isFrame: number;
  isCache: number;
  menuType: string;
  visible: string;
  status: string;
  perms: string;
  icon: string;
  createdDept: number;
  createdBy: number;
  createdAt: string;
  remark: string;
}

export interface SysMenuListRes {
  items: SysMenuListData[];
  total: number;
}

export interface SysMenuViewReq {
  menuId: number;
}

export interface SysMenuViewRes extends SysMenuListData {}

export interface SysMenuAddReq {
  menuName: string;
  parentId: number;
  orderNum: number;
  path: string;
  component: string;
  queryParam: string;
  isFrame: number;
  isCache: number;
  menuType: string;
  visible: string;
  status: string;
  perms: string;
  icon: string;
  remark: string;
}

export interface SysMenuUpdateReq {
  menuId: number;
  menuName?: string;
  parentId?: number;
  orderNum?: number;
  path?: string;
  component?: string;
  queryParam?: string;
  isFrame?: number;
  isCache?: number;
  menuType?: string;
  visible?: string;
  status?: string;
  perms?: string;
  icon?: string;
  remark?: string;
}

export interface SysMenuDeleteReq {
  menuId: number;
}

export async function getSysMenuListApi(params?: SysMenuListReq) {
  return requestClient.get<SysMenuListRes>('/system/menu/list', { params });
}

export async function getSysMenuViewApi(params: SysMenuViewReq) {
  return requestClient.get<SysMenuViewRes>('/system/menu/view', { params });
}

export async function updateSysMenuApi(params: SysMenuUpdateReq|any) {
  return requestClient.post('/system/menu/update', { ...params });
}

export async function deleteSysMenuApi(params: SysMenuDeleteReq|any) {
  return requestClient.post('/system/menu/delete', { ...params });
}

export async function addSysMenuApi(params: SysMenuAddReq|any) {
  return requestClient.post('/system/menu/add', { ...params });
}