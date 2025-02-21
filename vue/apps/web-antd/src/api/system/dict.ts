import { requestClient } from '#/api/request';

export interface SysDictTypeListReq {
    page: number;
    pageSize: number;
    dictName: string;
    dictType: string;
}
export interface SysDictType {
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
    items: SysDictType[];
    total: number;
}

export async function getSysDictTypeListApi(params: SysDictTypeListReq) {
    return requestClient.get<SysDictTypeListRes>('/dict/list', { params });
}

export interface SysDictDataListReq {
    page: number;
    pageSize: number;
    dictId: number;
}
export interface SysDictData {
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
    items: SysDictData[];
    type: SysDictType;
    total: number;
}

export async function getSysDictDataListApi(params: SysDictDataListReq) {
    return requestClient.get<SysDictDataListRes>('/dict-data/list/'+params.dictId, { params });
}


