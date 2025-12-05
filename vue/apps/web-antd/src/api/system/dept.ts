/**
 * @description 机构管理相关接口
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */

import { requestClient } from '#/api/request';

export interface SysDeptMini {
  deptId: number;
  deptName: string;
}

export interface SysDeptListParam {
  deptName: string;
  status: string;
  page?: number;
  pageSize?: number;
}

export interface SysDeptListData {
  deptId: number;
  tenantId: string;
  parentId: number;
  ancestors: string;
  deptName: string;
  deptCategory: string;
  orderNum: number;
  leader: number;
  phone: string;
  email: string;
  status: string;
  createdAt: string;
}

export interface SysDeptListRes {
  items: SysDeptListData[];
  total: number;
}


export async function getSysDeptListApi(params: SysDeptListParam) {
  return requestClient.get<SysDeptListRes>('/system/dept/list', { params });
}

export interface SysDeptTreeData {
  deptId: number;
  parentId: number;
  key: string;
  deptName: string;
  children?: SysDeptTreeData[];
}
export interface SysDeptTreeRes {
  items: SysDeptTreeData[];
}

export async function getSysDeptTreeApi() {
  return requestClient.get<SysDeptTreeRes>('/system/dept/tree');
}

export interface SysDeptAddParam {
  parentId: number;
  ancestors: string;
  deptName: string;
  deptCategory: string;
  orderNum: number;
  leader: number;
  phone: string;
  email: string;
  status: string;
}

export interface SysDeptAddRes {
  deptId: number;
}

export async function addSysDeptApi(params: SysDeptAddParam | {
  [x: string]: any;
}) {
  return requestClient.post<SysDeptAddRes>('/system/dept/add', params);
}

export interface SysDeptEditParam {
  deptId: number;
  parentId?: number;
  ancestors?: string;
  deptName?: string;
  deptCategory?: string;
  orderNum?: number;
  leader?: number;
  phone?: string;
  email?: string;
  status?: string;
}

export interface SysDeptEditRes {
  deptId: number;
}

export async function editSysDeptApi(params: SysDeptEditParam | {
  [x: string]: any;
}) {
  return requestClient.post<SysDeptEditRes>('/system/dept/edit', params);
}

export interface SysDeptDeleteParam {
  deptId: number;
}

export interface SysDeptDeleteRes {
  deptId: number;
}

export async function deleteSysDeptApi(params: SysDeptDeleteParam) {
  return requestClient.post<SysDeptDeleteRes>('/system/dept/delete', params);
}

export interface SysDeptViewParam {
  deptId: number;
}

export interface SysDeptViewRes {
  deptId: number;
  tenantId: string;
  parentId: number;
  ancestors: string;
  deptName: string;
  deptCategory: string;
  orderNum: number;
  leader: number | null;
  phone: string;
  email: string;
  status: string;
  createdDept: number;
  createdBy: number;
  createdAt: string;
  updatedBy: number;
  updatedAt: string;
}

export async function viewSysDeptApi(params: SysDeptViewParam) {
  return requestClient.get<SysDeptViewRes>('/system/dept/view', { params });
}

/**
 * 批量获取机构名称
 * @param deptIds 机构ID数组
 * @returns 机构ID到名称的映射对象
 */
export async function batchGetDeptNames(deptIds: number[]): Promise<Record<number, string>> {
  console.log('📡 调用批量机构名称API:', deptIds);

  if (!deptIds || deptIds.length === 0) {
    return {};
  }

  try {
    // 获取所有机构列表，使用合理的页面大小（不超过2000）
    const response = await getSysDeptListApi({
      deptName: '',
      status: '',
      page: 1,
      pageSize: 2000
    });

    if (!response || !response.items) {
      console.error('机构列表API返回格式错误:', response);
      return {};
    }

    const deptList = response.items;

    // 创建ID到名称的映射，只包含请求的ID
    const nameMap: Record<number, string> = {};
    deptIds.forEach(id => {
      const dept = deptList.find(d => d.deptId === id);
      nameMap[id] = dept ? dept.deptName : `机构${id}`;
    });

    console.log('🎯 批量机构名称查询结果:', nameMap);
    return nameMap;
  } catch (error) {
    console.error('批量获取机构名称失败:', error);
    // 失败时返回ID作为名称的映射
    const fallbackMap: Record<number, string> = {};
    deptIds.forEach(id => {
      fallbackMap[id] = String(id);
    });
    return fallbackMap;
  }
}
