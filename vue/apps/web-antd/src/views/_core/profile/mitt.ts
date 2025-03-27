import { mitt } from '@vben/utils';

type Events = {
  updateProfile: void;
};

export const emitter = mitt<Events>();

/**
 * @description 个人中心事件总线
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 XiuAdmin CLI
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */
