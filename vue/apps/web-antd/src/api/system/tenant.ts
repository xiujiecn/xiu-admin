import { requestClient } from '#/api/request';

export interface SysTenantListParam {
  page: number;
  pageSize: number;
  tenantId: string;
  contactUserName: string;
  contactPhone: string;
}

export interface SysTenantListData {
    id: number;
    tenantId: string;
    contactUserName: string;
    contactPhone: string;
    companyName: string;
    licenseNumber: string;
    address: string;
    intro: string;
    domain: string;
    remark: string;
    packageId: number;
    expireTime: string;
    accountCount: number;
    status: string;
    createdDept: number;
    createdBy: number;
    createdAt: string;
}
export interface SysTenantListRes {
  items: SysTenantListData[];
  total: number;
}

export interface SysTenantPackageListParam {
  page: number;
  pageSize: number;
  tenantId: string;
  contactUserName: string;
  contactPhone: string;
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

export async function getSysTenantListApi(params: SysTenantListParam) {
  return requestClient.get<SysTenantListRes>('/system/tenant/list', { params });
}

export async function getSysTenantPackageListApi(params: SysTenantPackageListParam) {
  return requestClient.get<SysTenantPackageListRes>('/system/tenant/package/list', { params });
}
