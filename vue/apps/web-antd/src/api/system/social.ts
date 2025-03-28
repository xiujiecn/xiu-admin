/**
 * @description 社交账号管理相关接口
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */

import { requestClient } from '#/api/request';

export interface SysSocialListReq {
  page?: number;
  pageSize?: number;
  userId?: number;
}

export interface SysSocialListRes {
  items: SysSocialListModel[];
  total: number;
}

export interface SysSocialListModel {
    id: number;
userId: number;
tenantId: string;
  authId: string;
  source: string;
  openId: string;
  userName: string;
  nickName: string;
  email: string;
  avatar: string;
  accessToken: string;
  expireIn: number;
  refreshToken: string;
  accessCode: string;
  unionId: string;
  scope: string;
  tokenType: string;
  idToken: string;
  macAlgorithm: string;
  macKey: string;
  code: string;
  oauthToken: string;
  oauthTokenSecret: string;
  createdDept: number;
  createdBy: number;
  createdAt: string;
}

export async function getSysSocialListApi(params: SysSocialListReq) {
    return requestClient.get<SysSocialListRes>('/system/social/list', { params });
}
