/**
 * @description 代码生成器表格相关接口
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */

import { requestClient } from '#/api/request';
import type { DictOption } from '#/store/dict';

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
    createdDept: number;  // 创建机构
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
export interface SysGenTableEditReq {
    tableId: number;
    genType?: number;
    genTemplate?: number;
    varName?: string;
    options?: any;
    dbName?: string;
    tableName?: string;
    tableComment?: string;
    daoName?: string;
    masterColumns?: any;
    addonName?: string;
}

export interface SysGenTableEditRes {
    tableId: number;
}

export interface SysGenTableDeleteReq {
    tableIds: number[];
}
export interface SysGenTableDeleteRes {
    tableIds: number[];
}
export interface SysGenTableJoinModel {
    uuid: string;
    dbName: string;
    masterTableName: string;
    linkTable: string;
    alias: string;
    linkMode: number;
    field: string;
    masterField: string;
    daoName: string;
    columns: GenCodesColumn[];
}

export interface SysGenTableTreeModel {
    titleColumn: string;
    styleType: number;
}
export interface SysGenTableMenuModel {
    icon: string;
    pid: number;
    sort: number;
}
export interface SysGenTableFuncDictModel {
    valueColumn: string|null    ;
    labelColumn: string|null;
}
export interface SysGenTablePresetStepModel {
    formGridCols: number;
}
export interface SysGenTableOptionsModel {
    headOps?: string[];
    columnOps?: string[];
    autoOps?: string[];
    join?:any[];
    menu?:SysGenTableMenuModel;
    funcDict?: SysGenTableFuncDictModel;
    presetStep?:SysGenTablePresetStepModel;
    tree?:SysGenTableTreeModel;
}

export interface SysGenTableViewReq {
    tableId: number;
}


export interface SysGenTableViewRes {
    tableId: number;
    genType: number;
    genTemplate: number;
    varName: string;
    options: SysGenTableOptionsModel|string;
    dbName: string;
    tableName: string;
    tableComment: string;
    daoName: string;
    masterColumns: GenCodesColumn[]|string;
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

export async function editSysGenTableApi(params: SysGenTableEditReq) {
    return requestClient.post<SysGenTableEditRes>('/genCodes/genTable/edit', { ...params });
}

export async function getSysGenTableViewApi(params: SysGenTableViewReq) {
    return requestClient.get<SysGenTableViewRes>('/genCodes/genTable/view', { params });
}


export interface SelectItemModel {
    value: string|number;
    label: string;
    name: string;
    disabled?: boolean;
}

export interface SelectGenTypeModel extends SelectItemModel,DictOption {
    templates: SelectGenTemplateModel[];
}

export interface SelectGenTemplateModel extends SelectItemModel {
    isAddon: boolean;
}

export interface GenCodesSelect {
    genType: SelectGenTypeModel[];
    db: SelectItemModel[];
    status: SelectItemModel[];
    linkMode: SelectItemModel[];
    buildMethod: SelectItemModel[];
    formMode: SelectItemModel[];
    formRole: SelectItemModel[];
    whereMode: SelectItemModel[];
    addons: SelectItemModel[];
    tableAlign: SelectItemModel[];
    treeStyleType: SelectItemModel[];
    dictMode: SelectItemModel[];
}


export interface GenCodesSelectsReq {
}

export interface GenCodesSelectsRes extends GenCodesSelect {

}

export async function getGenCodesSelectsApi(params: GenCodesSelectsReq) {
    return requestClient.get<GenCodesSelectsRes>('/genCodes/genTable/selects', { params });
}

export interface DbTableSelectModel {
    value: string;
    label: string;
    name: string;
    daoName: string;
    defVarName: string;
    defAlias: string;
    defTableComment: string;
}
export interface GenCodesTableSelectReq {
    dbGroup: string;
}

export interface GenCodesTableSelectRes {
    items: DbTableSelectModel[];
}

export async function getGenCodesTableSelectApi(params: GenCodesTableSelectReq) {
    return requestClient.get<GenCodesTableSelectRes>('/genCodes/genTable/tableSelect', { params });
}

export interface DbColumnSelectModel {
    value: string;
    label: string;
    name: string;
}

export interface GenCodesColumnSelectReq {
    dbGroup: string;
    tableName: string;
}
export interface GenCodesColumnSelectRes {
    items: DbColumnSelectModel[];
}

export async function getGenCodesColumnSelectApi(params: GenCodesColumnSelectReq) {
    return requestClient.get<GenCodesColumnSelectRes>('/genCodes/genTable/columnSelect', { params });
}

export interface GenCodesColumn {
    id: number;
    name: string;
    dc: string;
    dataType: string;
    sqlType: string;
    length: number;
    isAllowNull: boolean;
    defaultValue: any;
    index: string;
    extra: string;
    goName: string;
    goType: string;
    tsName: string;
    tsType: string;
    isList: boolean;
    isExport: boolean;
    isSort: boolean;
    isQuery: boolean;
    queryWhere: string;
    isEdit: boolean;
    required: boolean;
    unique: boolean;
    formMode: string;
    formRole: string;
    formGridSpan: number;
    dictType: number;
    align: string;
    width: number;  
}
export interface GenCodesColumnListModel extends GenCodesColumn {

}

export interface GenCodesColumnListParam {
    dbGroup: string;
    tableName: string;
    isLink: number;
    alias: string;
}
export interface GenCodesColumnListRes {
    items: GenCodesColumnListModel[];
}

export async function getGenCodesColumnListApi(params: GenCodesColumnListParam) {
    return requestClient.get<GenCodesColumnListRes>('/genCodes/genTable/columnList', { params });
}

export interface GenCodesPreviewReq extends SysGenTableListModel {

}
export interface GenFile {
    content: string;
    path: string;
    meth: number;
    required: boolean;
}
export interface GenCodesPreviewRes {
    views: Record<string, GenFile>;
    config?: any;
}

export async function postGenCodesPreviewApi(params: GenCodesPreviewReq) {
    return requestClient.post<GenCodesPreviewRes>('/genCodes/genTable/preview', { ...params });
}

export interface GenCodesBuildReq extends GenCodesPreviewReq {

}
export interface GenCodesBuildRes {

}

export async function postGenCodesBuildApi(params: GenCodesBuildReq) {
    return requestClient.post<GenCodesBuildRes>('/genCodes/genTable/build', { ...params });
}