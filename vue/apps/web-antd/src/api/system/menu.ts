import { requestClient } from '#/api/request';
export interface SysMenuListReq {
    menuName: string;
    status: string;
  }
  
export interface SysMenuListData {
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
  
export interface SysMenuListRes {
    data: SysMenuListData[];
    total: number;
  }
  
  export async function getSysMenuListApi(params: SysMenuListReq) {
    return requestClient.get<SysMenuListRes>('/menu/list', { params });
  }