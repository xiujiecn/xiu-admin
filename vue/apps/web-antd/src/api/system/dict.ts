import { requestClient } from '#/api/request';

export interface DictTypeListReq {
    page: number;
    pageSize: number;
    dictName: string;
    dictType: string;
}
export interface DictType {
    dictId: number; 
    tenantId: string;
    dictName: string;
    dictType: string;
    createdDept: number;
    createdBy: number;
    createdAt: string;
    remark: string;
}
export interface DictTypeListRes {
    items: DictType[];
    total: number;
}

export async function getDictTypeListApi(params: DictTypeListReq) {
    return requestClient.get<DictTypeListRes>('/dict/list', { params });
}

export interface DictDataListReq {
    page: number;
    pageSize: number;
    dictId: number;
}
export interface DictData {
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
export interface DictDataListRes {
    items: DictData[];
    type: DictType;
    total: number;
}

export async function getDictDataListApi(params: DictDataListReq) {
    return requestClient.get<DictDataListRes>('/dict-data/list/'+params.dictId, { params });
}


