import { baseRequestClient, requestClient } from '#/api/request';

export namespace CaptchaApi {
  export interface GetCaptchaResult {
    captchaID: string;
    captchaImage: string;
  }
}

export async function getCaptchaApi() {
  return requestClient.get<CaptchaApi.GetCaptchaResult>('/captcha');
}
