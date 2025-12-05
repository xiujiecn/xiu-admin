/**
 * @description 认证相关接口
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */

import { baseRequestClient, requestClient } from '#/api/request';

/**
 * 登录类型
 * password 密码
 * sms 短信
 * social 第三方oauth
 * email 邮箱
 * xcx 小程序
 */
type GrantType = 'email' | 'password' | 'sms' | 'social' | 'xcx';

export namespace AuthApi {
  /** 登录接口参数 */
  export interface LoginParams {
    password?: string;
    username?: string;
  }

  /** 登录接口返回值 */
  export interface LoginResult {
    accessToken: string;
  }
  /** 注册接口参数 */
  export interface RegisterParams {
    tenantId?: string;
    username?: string;
    password?: string;
    captchaID?: string;
    captchaValue?: string;
  }
    export interface RegisterResult {
  }

  export interface RefreshTokenResult {
    data: string;
    status: number;
  }


  export interface BaseLoginParams {
    clientId?: string;
    grantType: GrantType;
    tenantId: string;
  }

  export interface OAuthLoginParams extends BaseLoginParams {
    socialCode: string;
    socialState: string;
    source: string;
  }

}


/**
 * @param companyName 租户/公司名称
 * @param domain 绑定域名(不带http(s)://) 可选
 * @param tenantId 租户id
 */
export interface TenantOption {
  companyName: string;
  domain?: string;
  tenantId: string;
}

/**
 * @param tenantEnabled 是否启用租户
 * @param voList 租户列表
 */
export interface TenantResp {
  tenantEnabled: boolean;
  voList: TenantOption[];
}

/**
 * 获取用户权限码
 */
export interface GetAccessCodesResult {
  data: string[];
}

/**
 * 登录
 */
export async function loginApi(data: AuthApi.LoginParams) {
  return requestClient.post<AuthApi.LoginResult>('/system/auth/login', data);
}
/**
 * 注册register
 */

export async function registerApi(data: AuthApi.RegisterParams) {
  return requestClient.post<AuthApi.RegisterResult>(
    '/system/user/register',
    data,
  );
}

/**
 * 刷新accessToken
 */
export async function refreshTokenApi() {
  return baseRequestClient.post<AuthApi.RefreshTokenResult>('/system/auth/refresh', {
    withCredentials: true,
  });
}

/**
 * 退出登录
 */
export async function logoutApi() {
  return requestClient.post('/system/auth/logout', {
    withCredentials: true,
  });
}

/**
 * 获取用户权限码
 */
export async function getAccessCodesApi() {
  return requestClient.get<GetAccessCodesResult>('/system/auth/codes');
}

/**
 * 绑定第三方账号
 * @param source 绑定的来源
 * @returns 跳转url
 */
export function authBinding(source: string, tenantId: string) {
  return requestClient.get<string>(`/auth/binding/${source}`, {
    params: {
      domain: window.location.host,
      tenantId,
    },
  });
}

/**
 * 取消绑定
 * @param id id
 */
export function authUnbinding(id: string) {
  return requestClient.delete<void>(`/auth/unlock/${id}`);
}

/**
 * oauth授权回调
 * @param data oauth授权
 * @returns void
 */
export function authCallback(data: AuthApi.OAuthLoginParams) {
  return requestClient.post<void>('/auth/social/callback', data);
}

// 获取租户列表
export function tenantList() {
  return requestClient.get<TenantResp>('/system/auth/getTenantList');
}

