/**
 * @description 字典数据模型定义
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 XiuAdmin CLI
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
      fieldName: 'dictCode',
    },
    {
      component: 'Input',
      componentProps: {
        disabled: true,
      },
      fieldName: 'dictType',
      label: '字典类型',
    },
    {
      component: 'Input',
      fieldName: 'listClass',
      label: '标签样式',
    },
    {
      component: 'Input',
      fieldName: 'dictLabel',
      label: '数据标签',
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'dictValue',
      label: '数据键值',
      rules: 'required',
    },
    {
      component: 'Textarea',
      componentProps: {
        placeholder: '可使用tailwind类名 如bg-blue w-full h-full等',
      },
      fieldName: 'cssClass',
      formItemClass: 'items-baseline',
      help: '标签的css样式, 可添加已经编译的css类名',
      label: 'css类名',
    },
    {
      component: 'InputNumber',
      fieldName: 'dictSort',
      label: '显示排序',
      rules: 'required',
    },
    {
      component: 'Textarea',
      fieldName: 'remark',
      formItemClass: 'items-baseline',
      label: '备注',
    },
  ];
  