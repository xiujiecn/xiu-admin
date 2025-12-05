/**
 * @description 字典类型模型定义
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */
import type {
    VbenFormSchema,
  } from '@vben/common-ui';
  import { z } from '@vben/common-ui';
  import { getDictOptions } from '#/utils/dict';
  import { DictEnum } from '@vben/constants';
  import { getPopupContainer } from '@vben/utils';

  export const drawerSchema: VbenFormSchema[] =  [
    {
      component: 'Input',
      dependencies: {
        show: () => false,
        triggerFields: [''],
      },
      fieldName: 'dictId',
      label: 'dictId',
    },
    {
      component: 'Input',
      fieldName: 'dictName',
      label: '字典名称',rules: z.string()
      .min(1, '字典名称不能为空')
      .max(32, '字典名称最大长度32位'),
    },
    {
      component: 'Input',
      fieldName: 'dictType',
      help: '使用英文/下划线命名, 如:sys_normal_disable',
      label: '字典类型',
      rules: z
        .string()
        .regex(/^[a-z_]+$/i, { message: '字典类型只能使用英文/下划线命名' }),
    },
    //添加是否系统内置 字段
     {
      component: 'RadioGroup',
      componentProps: {
        buttonStyle: 'solid',
        options: getDictOptions(DictEnum.SYS_YES_NO),
        optionType: 'button',
      },
      defaultValue: 'N', // 默认为"否"，提交时会转换为 1
      fieldName: 'isSys',
      label: '是否系统内置',
      rules: 'required',
      //根据用户身份决定这个字段的值是否可以修改； cpr:superadmin
    },
    {
      component: 'Textarea',
      fieldName: 'remark',
      formItemClass: 'items-baseline',
      label: '备注',
    },
  ];
