import { requestClient } from '#/api/request';

export interface SysGenTableListReq {
    page: number;
    pageSize: number;
    genType: number;
    varName: string;
    status: string;
}
export interface SysGenTableListModel {
    tableId: number;  // 生成ID
    genType: number;  // 生成类型
    genTemplate: number;  // 生成模板
    varName: string;  // 实体命名
    options: string;  // 配置选项
    dbName: string;  // 数据库名称
    tableName: string;  // 主表名称
    tableComment: string;  // 主表注释
    daoName: string;  // 主表dao模型
    masterColumns: string;  // 主表字段
    addonName: string;  // 插件名称
    status: string;  // 生成状态（0成功
    createdDept: number;  // 创建部门
    createdBy: number;  // 创建者
    createdAt: string;  // 创建时间
}
export interface SysGenTableListRes {
    items: SysGenTableListModel[];
    total: number;
}

export async function getSysGenTableListApi(params: SysGenTableListReq) {
    return requestClient.get<SysGenTableListRes>('/genCodes/genTable/list', { params });
}


export interface SysGenTableAddReq {
    genType: number;
    genTemplate: number;
    varName: string;
    options: string;
    dbName: string;
    tableName: string;
    tableComment: string;
    daoName: string;
    masterColumns: string;
    addonName: string;
}
export interface SysGenTableAddRes {
    tableId: number;
}

export interface SysGenTableDeleteReq {
    tableIds: number[];
}
export interface SysGenTableDeleteRes {
    tableIds: number[];
}

export interface SysGenTableViewReq {
    tableId: number;
}
export interface SysGenTableViewRes {
    tableId: number;
    genType: number;
    genTemplate: number;
    varName: string;
    options: string;
    dbName: string;
    tableName: string;
    tableComment: string;
    daoName: string;
    masterColumns: string;
    addonName: string;
    status: string;
    createdDept: number;
    createdBy: number;
    createdAt: string;
}

export async function addSysGenTableApi(params: SysGenTableAddReq | {
    [x: string]: any;
}) {
    return requestClient.post<SysGenTableAddRes>('/genCodes/genTable/add', { ...params });
}


export async function deleteSysGenTableApi(params: SysGenTableDeleteReq) {
    return requestClient.post<SysGenTableDeleteRes>('/genCodes/genTable/delete', { ...params });
}


export async function getSysGenTableViewApi(params: SysGenTableViewReq) {
    return requestClient.get<SysGenTableViewRes>('/genCodes/genTable/view', { params });
}

