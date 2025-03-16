import { requestClient } from '#/api/request';
import {md5} from 'js-md5';
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
    updatedBy: number;
    updatedAt: string;  
}

export interface SysTenantListRes {
  items: SysTenantListData[];
  total: number;
}

export interface SysTenantAddParam {
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
  username: string;
  password: string;
}     

export interface SysTenantAddModel {
  id: number;
}

export interface SysTenantEditParam {
  id: number;
  tenantId?: string;
  contactUserName?: string;
  contactPhone?: string;
  companyName?: string;
  licenseNumber?: string;
  address?: string;
  intro?: string;
  domain?: string;
  remark?: string;
  packageId?: number;
  expireTime?: string;
  accountCount?: number;
  status?: string;
} 
export interface SysTenantEditModel {
  id: number;
}

export interface SysTenantDeleteParam {
  ids : number[];
}
export interface SysTenantDeleteModel {
  ids: number[];
}

export interface SysTenantStatusParam {
  id: number;
  status: string;
}
export interface SysTenantStatusModel {
  id: number;
} 

export interface SysTenantViewParam {
  id: number;
}
export interface SysTenantViewModel {
  id: number;
}

export async function getSysTenantListApi(params: SysTenantListParam) { 
  return requestClient.get<SysTenantListRes>('/system/tenant/list', { params });
}

export async function getSysTenantViewApi(params: SysTenantViewParam) {
  return requestClient.get<SysTenantViewModel>('/system/tenant/view', { params });
}

export async function addSysTenantApi(params: SysTenantAddParam|{
  [x: string]: any;
}) {
  params.password = md5(params.password);
  return requestClient.post<SysTenantAddModel>('/system/tenant/add', { ...params });
}

export async function editSysTenantApi(params: SysTenantEditParam|{
  [x: string]: any;
}) {
  return requestClient.post<SysTenantEditModel>('/system/tenant/edit', { ...params });
}

export async function deleteSysTenantApi(params: SysTenantDeleteParam|{
  [x: string]: any;
}) {
  return requestClient.post<SysTenantDeleteModel>('/system/tenant/delete', { ...params });
}

export async function statusSysTenantApi(params: SysTenantStatusParam) {
  return requestClient.post<SysTenantStatusModel>('/system/tenant/status', { ...params });
}



