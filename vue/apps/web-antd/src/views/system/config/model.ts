/**
 * @description 系统配置模型定义
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
      fieldName: 'configId',
      label: '参数主键',
    },
    {
      component: 'Input',
      fieldName: 'configName',
      label: '参数名称',
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'configKey',
      label: '参数键名',
      rules: 'required',
    },
    {
      component: 'Textarea',
      formItemClass: 'items-baseline',
      fieldName: 'configValue',
      label: '参数键值',
      componentProps: {
        autoSize: true,
      },
      rules: 'required',
    },
    {
      component: 'RadioGroup',
      componentProps: {
        buttonStyle: 'solid',
        options: getDictOptions(DictEnum.SYS_YES_NO),
        optionType: 'button',
      },
      defaultValue: 'N',
      fieldName: 'configType',
      label: '是否内置',
      rules: 'required',
    },
    {
      component: 'Textarea',
      fieldName: 'remark',
      formItemClass: 'items-baseline',
      label: '备注',
    },
  ];
  