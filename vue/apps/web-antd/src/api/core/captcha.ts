/**
 * @description 验证码相关接口
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */

import {  requestClient } from '#/api/request';

export namespace CaptchaApi {
  export interface GetCaptchaResult {
    captchaID: string;
    captchaImage: string;
  }
}

export async function getCaptchaApi() {
  return requestClient.get<CaptchaApi.GetCaptchaResult>('/system/captcha');
}
