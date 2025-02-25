import type { RouteRecordStringComponent } from '@vben/types';
import { $t } from '@vben/locales';

// 非后台返回路由
export const localRoutes : RouteRecordStringComponent[] = [
    {
        path: '/profile',
        name: 'Profile',
        component: '/_core/profile/index',
        meta: {
            title: $t('ui.widgets.profile'),
            icon: 'mingcute:profile-line',
            hideInMenu: true,
        }
    },
]