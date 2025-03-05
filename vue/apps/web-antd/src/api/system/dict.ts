import { requestClient } from '#/api/request';

export interface SysDictTypeListReq {
    page: number;
    pageSize: number;
    dictName: string;
    dictType: string;
}
export interface SysDictTypeListModel {
    dictId: number;
    tenantId: string;
    dictName: string;
    dictType: string;
    createdDept: number;
    createdBy: number;
    createdAt: string;
    remark: string;
}
export interface SysDictTypeListRes {
    items: SysDictTypeListModel[];
    total: number;
}

export async function getSysDictTypeListApi(params: SysDictTypeListReq) {
    return requestClient.get<SysDictTypeListRes>('/system/dict/list', { params });
}

export interface SysDictDataListReq {
    page: number;
    pageSize: number;
    dictId: number;
    dictType: string;
}
export interface SysDictDataListModel    {
    dictCode: number;
    tenantId: string;
    dictSort: number;
    dictLabel: string;
    dictValue: string;
    dictType: string;
    cssClass: string;
    listClass: string;
    isDefault: string;
    createdDept: number;
    createdBy: number;
    createdAt: string;
    remark: string;
}
export interface SysDictDataListRes {
    items: SysDictDataListModel[];
    type: SysDictTypeListModel;  
    total: number;
}

export async function getSysDictDataListApi(params: SysDictDataListReq) {
    return requestClient.get<SysDictDataListRes>('/system/dict-data/list/'+params.dictId, { params });
}


export interface SysDictTypeAddReq {
    dictName: string;
    dictType: string;
    remark: string;
}
export interface SysDictTypeAddRes {
    dictId: number;
}

export async function addSysDictTypeApi(params: SysDictTypeAddReq | {
    [x: string]: any;
}) {
    return requestClient.post<SysDictTypeAddRes>('/system/dict-type/add', { ...params });
}

export interface SysDictTypeEditReq {
    dictId: number;
    dictName: string;
    dictType: string;
    remark: string;
}
export interface SysDictTypeEditRes {
    dictId: number;
}

export async function editSysDictTypeApi(params: SysDictTypeEditReq | {
    [x: string]: any;
}) {
    return requestClient.post<SysDictTypeEditRes>('/system/dict-type/edit', { ...params });
}

export interface SysDictTypeDeleteReq {
    dictIds: number[];
}
export interface SysDictTypeDeleteRes {
    dictIds: number[];
}   

export async function deleteSysDictTypeApi(params: SysDictTypeDeleteReq) {
    return requestClient.post<SysDictTypeDeleteRes>('/system/dict-type/delete', { ...params });
}

export interface SysDictTypeViewReq {
    dictId: number;
}
export interface SysDictTypeViewRes {
    dictId: number;
    tenantId: string;
    dictName: string;
    dictType: string;
    createdDept: number;
    createdBy: number;
    createdAt: string;
    updatedBy: number;
    updatedAt: string;
    remark: string; 
}

export async function getSysDictTypeViewApi(params: SysDictTypeViewReq) {
    return requestClient.get<SysDictTypeViewRes>('/system/dict-type/view', { params });
}

export interface SysDictDataAddReq {
    dictSort: number;
    dictLabel: string;
    dictValue: string;
    dictType: string;
    cssClass: string;
    listClass: string;
    isDefault: string;
    remark: string;
}
export interface SysDictDataAddRes {
    dictCode: number;
}   

export async function addSysDictDataApi(params: SysDictDataAddReq | {
    [x: string]: any;
}) {
    return requestClient.post<SysDictDataAddRes>('/system/dict-data/add', { ...params });
}

export interface SysDictDataEditReq {
    dictCode: number;
    dictSort: number;
    dictLabel: string;
    dictValue: string;
    dictType: string;
    cssClass: string;
    listClass: string;
    isDefault: string;
    remark: string;
}
export interface SysDictDataEditRes {
    dictCode: number;   
}

export async function editSysDictDataApi(params: SysDictDataEditReq | {
    [x: string]: any;
}) {
    return requestClient.post<SysDictDataEditRes>('/system/dict-data/edit', { ...params });
}

export interface SysDictDataDeleteReq {
    dictCodes: number[];
}
export interface SysDictDataDeleteRes {
    dictCodes: number[];
}   

export async function deleteSysDictDataApi(params: SysDictDataDeleteReq) {
    return requestClient.post<SysDictDataDeleteRes>('/system/dict-data/delete', { ...params });
}

export interface SysDictDataViewReq {
    dictCode: number;
}
export interface SysDictDataViewRes {
    dictCode: number;
    tenantId: string;
    dictSort: number;
    dictLabel: string;
    dictValue: string;
    dictType: string;
    cssClass: string;
    listClass: string;
    isDefault: string;
    createdDept: number;
    createdBy: number;
    createdAt: string;
    updatedBy: number;
    updatedAt: string;
    remark: string;
}

export async function getSysDictDataViewApi(params: SysDictDataViewReq) {
    return requestClient.get<SysDictDataViewRes>('/system/dict-data/view', { params });
}
