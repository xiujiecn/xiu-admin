/**
 * @description 系统二维码管理相关接口
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */

import { requestClient } from '#/api/request';

/**
 * @description 获取绑定二维码
 */
export interface QrBindRes {
  ticket: string;
  expireTime: string;
  qrUrl: string;
  scanned: boolean;
  openId?: string;
}
export async function getQrBindApi() {
  return requestClient.get<QrBindRes>('/system/qrcode/bind');
}

/**
 * @description 获取绑定二维码扫码状态
 */
export interface QrBindStatusRes extends QrBindRes {
  bound?: boolean;
  expired?: boolean;
}
export async function getQrBindStatusApi() {
  return requestClient.get<QrBindStatusRes>('/system/qrcode/bind/status');
}

/**
 * @description 获取登录二维码
 */
export interface QrLoginRes {
  tempUserId: string;
  ticket: string;
  expireTime: string;
  qrUrl: string;
  scanned: boolean;
  openId?: string;
  bound?: boolean;
  expired?: boolean;
  loginUserOut?: {
    id: number;
    username: string;
    nickName: string;
    avatar: string;
    tenantId: string;
    deptId: number;
  };
  token?: string;
}
export async function getQrLoginApi() {
  return requestClient.get<QrLoginRes>('/system/qrcode/login');
}

/**
 * @description 获取登录二维码扫码状态
 */
export async function getQrLoginStatusApi(tempUserId: string) {
  return requestClient.get<QrLoginRes>('/system/qrcode/login/status', { params: { tempUserId } });
}

/**
 * @description 登录并关联OpenId
 */
export interface QrLoginBindOpenIdParam {
  openId: string;
  tenantId: string;
  username: string;
  password: string;
  captchaID?: string;
  captchaValue: string;
}
export async function qrLoginBindOpenIdApi(params: QrLoginBindOpenIdParam) {
  return requestClient.post<QrLoginRes>('/system/qrcode/loginAndBindOpenId', params);
}

/**
 * @description 注册并关联OpenId
 */
export interface QrRegisterBindOpenIdParam {
  openId: string;
  userName: string;
  password: string;
  tenantId: string;
  captchaId?: string;
  captchaValue: string;
}
export async function qrRegisterBindOpenIdApi(params: QrRegisterBindOpenIdParam) {
  return requestClient.post<QrLoginRes>('/system/qrcode/registerAndBindOpenId', params);
}

/**
 * @description 扫码成功回调
 */
export interface QrScanCallbackParam {
  sceneId?: number;
  thirdId?: number;
  ticket?: string;
  openId?: string;
  event?: string;
}
export async function qrScanCallbackApi(params: QrScanCallbackParam) {
  return requestClient.post('/system/qrcode/scan/callback', params);
}

/**
 * @description 扫码后选择用户登录
 */
export interface QrLoginSelectUserParam {
  userName: string;
  tenantId: string;
  openId: string;
}
export async function qrLoginSelectUserApi(params: QrLoginSelectUserParam) {
  return requestClient.post<QrLoginRes>('/system/qrcode/loginSelectUser', params);
}

//删除绑定关系  
export async function deleteSocialBindingApi(id: string) {
  return requestClient.post(`/system/social/delete`, { id });
}
