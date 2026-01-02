import type { RouteMeta as IRouteMeta } from '@vben-core/typings';

import 'vue-router';

declare module 'vue-router' {
  // eslint-disable-next-line @typescript-eslint/no-empty-object-type
  interface RouteMeta extends IRouteMeta {}
}

export interface VbenAdminProAppConfigRaw {
  VITE_GLOB_API_URL: string;
  VITE_GLOB_CLIENT_ID: string; // 客户端ID
}

interface AuthConfig {
  dingding?: {
    clientId: string;
    corpId: string;
  };
}

export interface ApplicationConfig {
  apiURL: string;
  auth: AuthConfig;
  clientId: string; // 客户端ID
}

declare global {
  interface Window {
    _VBEN_ADMIN_PRO_APP_CONF_: VbenAdminProAppConfigRaw;
  }
}
