import type { RouteRecordRaw } from 'vue-router';

import { $t } from '#/locales';

const routes: RouteRecordRaw[] = [
  {
    meta: {
      icon: 'ic:baseline-view-in-ar',
      keepAlive: true,
      order: 1000,
      title: '租户管理',
    },
    name: 'Tenant',
    path: '/tenant',
    children: [
      {
        meta: {
          title: '租户管理',
        },
        name: 'tenant',
        path: '/tenant/tenant',
        component: () => import('#/views/demos/antd/index.vue'),
      },
      {
        meta: {
          title: '租户套餐管理',
        },
        name: 'tenantPackage`',
        path: '/tenant/tenant_package',
        component: () => import('#/views/demos/antd/index.vue'),
      },
    ],
  },
];

export default routes;
