/**
 * @description 请求客户端配置
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */

import type { RequestClientOptions, HttpResponse } from '@vben/request';

import { useAppConfig } from '@vben/hooks';
import { preferences } from '@vben/preferences';
import {
  authenticateResponseInterceptor,
  defaultResponseInterceptor,
  errorMessageResponseInterceptor,
  RequestClient,
} from '@vben/request';
import { useAccessStore } from '@vben/stores';

import { message, Modal } from 'ant-design-vue';

import { useAuthStore } from '#/store';

import { refreshTokenApi } from './core';

import { $t } from '@vben/locales';
import { isEmpty, isNull } from 'lodash-es';

const { apiURL } = useAppConfig(import.meta.env, import.meta.env.PROD);

/**
 * 是否已经处在登出过程中了 一个标志位
 * 主要是防止一个页面会请求多个api 都401 会导致登出执行多次
 */
let isLogoutProcessing = false;

function createRequestClient(baseURL: string, options?: RequestClientOptions) {
  const client = new RequestClient({
    ...options,
    baseURL,
    // 消息提示类型
    errorMessageMode: 'message',
    // 是否返回原生响应 比如：需要获取响应头时使用该属性
    isReturnNativeResponse: false,
    // 需要对返回数据进行处理
    isTransformResponse: true,
  });

  /**
   * 重新认证逻辑
   */
  async function doReAuthenticate() {
    console.warn('Access token or refresh token is invalid or expired. ');
    const accessStore = useAccessStore();
    const authStore = useAuthStore();
    accessStore.setAccessToken(null);
    if (
      preferences.app.loginExpiredMode === 'modal' &&
      accessStore.isAccessChecked
    ) {
      accessStore.setLoginExpired(true);
    } else {
      await authStore.logout();
    }
  }

  /**
   * 刷新token逻辑
   */
  async function doRefreshToken() {
    const accessStore = useAccessStore();
    const resp = await refreshTokenApi();
    const newToken = resp.data;
    accessStore.setAccessToken(newToken);
    return newToken;
  }

  function formatToken(token: null | string) {
    return token ? `Bearer ${token}` : null;
  }

  // 请求头处理
  client.addRequestInterceptor({
    fulfilled: async (config) => {
      const accessStore = useAccessStore();

      config.headers.Authorization = formatToken(accessStore.accessToken);
      config.headers['Accept-Language'] = preferences.app.locale;
      return config;
    },
  });

  // 处理返回的响应数据格式
  // client.addResponseInterceptor(
  //   defaultResponseInterceptor({
  //     codeField: 'code',
  //     dataField: 'data',
  //     successCode: 0,
  //   }),
  // );

  client.addResponseInterceptor<HttpResponse>({
    fulfilled: async (response) => {
      const { config, data: responseData, status, headers } = response;
      if (config.responseReturn === 'raw') {
        return response;
      }
      if (!responseData) {
        throw new Error($t('http.apiRequestFailed'));
      }

      // console.log('response', response);

      if (status >= 200 && status < 400) {
        if (config.responseReturn === 'body') {
          return responseData;
        } else if (headers['content-type']?.includes?.('application/json')) {
          if (config.responseType === 'blob' && Reflect.has(responseData, 'data')) 
          {
              // 这时候的data为blob类型
              const blob = responseData.data as unknown as Blob;
              // 拿到字符串转json对象
              responseData.data = JSON.parse(await blob.text());
          }
        } else if (headers['content-type']?.includes?.('application/octet-stream')) {
          // 下载文件
          return responseData;
        }
        const { code, data, message: msg, ...other } = responseData;
        // 业务状态码为200则请求成功
        const hasSuccess = Reflect.has(responseData, 'code') && code === 0;
        const hasLoginTimeout = Reflect.has(responseData, 'code') && code === 61;
        if (hasSuccess) {
          let successMsg = msg;

          if (isNull(successMsg) || isEmpty(successMsg)) {
            successMsg = $t(`http.operationSuccess`);
          }

          if (response.config.successMessageMode === 'modal') {
            Modal.success({
              content: successMsg,
              title: $t('http.successTip'),
            });
          } else if (response.config.successMessageMode === 'message') {
            message.success(successMsg);
          }
          // 分页情况下为code msg rows total 并没有data字段
          // 如果有data 直接返回data 没有data将剩余参数(...other)封装为data返回
          // 需要考虑data为null的情况(比如查询为空) 所以这里直接判断undefined
          if (data !== undefined) {
            return data;
          }
          // 没有data 将其他参数包装为data
          return other;
        }
        if (hasLoginTimeout) {
          // 已经在登出过程中 不再执行
          if (isLogoutProcessing) {
            return;
          }
          isLogoutProcessing = true;
          const _msg = $t('http.loginTimeout');
          const userStore = useAuthStore();
          userStore.logout().finally(() => {
            message.error(_msg);
            isLogoutProcessing = false;
          });
          // 不再执行下面逻辑
          return;
        }
      } else if (status == 401) {
        // 已经在登出过程中 不再执行
        if (isLogoutProcessing) {
          return;
        }
        isLogoutProcessing = true;
        const _msg = $t('http.loginTimeout');
        const userStore = useAuthStore();
        userStore.logout().finally(() => {
          message.error(_msg);
          isLogoutProcessing = false;
        });
        // 不再执行下面逻辑
        return;
      }
      console.log('addResponseInterceptor response', response);
      throw Object.assign({}, response, { response });
    },
  });
  // token过期的处理
  client.addResponseInterceptor(
    authenticateResponseInterceptor({
      client,
      doReAuthenticate,
      doRefreshToken,
      enableRefreshToken: preferences.app.enableRefreshToken,
      formatToken,
    }),
  );

  // 通用的错误处理,如果没有进入上面的错误处理逻辑，就会进入这里
  client.addResponseInterceptor(
    errorMessageResponseInterceptor((msg: string, error) => {
      console.log('errorMessageResponseInterceptor error', error);
      // 这里可以根据业务进行定制,你可以拿到 error 内的信息进行定制化处理，根据不同的 code 做不同的提示，而不是直接使用 message.error 提示 msg
      // 当前mock接口返回的错误字段是 error 或者 message
      const responseData = error?.response?.data ?? {};
      const errorMessage = responseData?.error ?? responseData?.message ?? '';
      // 如果没有错误信息，则会根据状态码进行提示
      message.error(errorMessage || msg);
    }),
  );
  return client;
}

export const requestClient = createRequestClient(apiURL, {
  responseReturn: 'data',
});

export const baseRequestClient = new RequestClient({ baseURL: apiURL });
