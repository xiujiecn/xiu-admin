import { requestClient } from '#/api/request';

export interface SysTenantPackageListParam {
  page: number;
  pageSize: number;
  tenantId?: string;
  contactUserName?: string;
  contactPhone?: string;
}

export interface SysTenantPackageListData {
    packageId: number;
    packageName: string;
    menuIds: string;
    remark: string;
    menuCheckStrictly: number;
    status: string;
    createdDept: number;
    createdBy: number;
    createdAt: string;
    updatedBy: number;
    updatedAt: string;
}
export interface SysTenantPackageListRes {
  items: SysTenantPackageListData[];
  total: number;
}

export interface SysTenantPackageViewParam {
  packageId: number;
}

export interface SysTenantPackageViewRes   {
  packageId: number;
  packageName: string;
  menuIds: string;
  remark: string;
  menuCheckStrictly: number;
  status: string;
  createdDept: number;
  createdBy: number;
  createdAt: string;
  updatedBy: number;
  updatedAt: string;
}

export interface SysTenantPackageAddParam {
  packageName: string;
  menuIds: string;
  remark: string;
  menuCheckStrictly: number;
  status: string;
}

  export interface SysTenantPackageAddRes {
  packageId: number;
}

export interface SysTenantPackageEditParam {
  packageId: number;
  packageName: string;
  menuIds: string;
  remark: string;
  menuCheckStrictly: number;
  status: string;
}

export interface SysTenantPackageEditRes {
  packageId: number;
}

export interface SysTenantPackageDeleteParam {  
  packageIds: number[];
}

export interface SysTenantPackageDeleteRes {
  packageIds: number[];
}

export interface SysTenantPackageStatusParam {
  packageId: number;
  status: string;
}

export interface SysTenantPackageStatusRes {
  packageId: number;
}

export async function getSysTenantPackageListApi(params: SysTenantPackageListParam) {
  return requestClient.get<SysTenantPackageListRes>('/system/tenant/package/list', { params });
}

export async function getSysTenantPackageViewApi(params: SysTenantPackageViewParam) {
  return requestClient.get<SysTenantPackageViewRes>('/system/tenant/package/view', { params });
}


export async function addSysTenantPackageApi(params: SysTenantPackageAddParam|{
  [x: string]: any;
}) {
  return requestClient.post<SysTenantPackageAddRes>('/system/tenant/package/add', { ...params });
}

export async function editSysTenantPackageApi(params: SysTenantPackageEditParam|{
  [x: string]: any;
}) {
  return requestClient.post<SysTenantPackageEditRes>('/system/tenant/package/edit', { ...params });
} 

export async function deleteSysTenantPackageApi(params: SysTenantPackageDeleteParam) {
  return requestClient.post<SysTenantPackageDeleteRes>('/system/tenant/package/delete', { ...params });
}   

export async function statusSysTenantPackageApi(params: SysTenantPackageStatusParam) {
  return requestClient.post<SysTenantPackageStatusRes>('/system/tenant/package/status', { ...params });
}   