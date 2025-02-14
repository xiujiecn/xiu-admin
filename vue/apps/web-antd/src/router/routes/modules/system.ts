import type { RouteRecordRaw } from 'vue-router';

import { $t } from '#/locales';

const routes: RouteRecordRaw[] = [
  {
    meta: {
      icon: 'ic:baseline-view-in-ar',
      keepAlive: true,
      order: 1000,
      title: '系统管理',
    },
    name: 'System',
    path: '/system',
    children: [
      {
        meta: {
          title: '用户管理',
        },
        name: 'systemUser',
        path: '/system/system—user',
        component: () => import('#/views/demos/antd/index.vue'),
      },
      {
        meta: {
          title: '角色管理',
        },
        name: 'systemRole',
        path: '/system/system—role',
        component: () => import('#/views/demos/antd/index.vue'),
      },
    ],
  },
];

export default routes;
