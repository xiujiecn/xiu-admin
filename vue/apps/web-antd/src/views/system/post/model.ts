

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
      fieldName: 'postId',
      label: 'postId',
    },
    {
      component: 'TreeSelect',
      componentProps: {
        getPopupContainer,
      },
      fieldName: 'deptId',
      label: '所属部门',
      rules: 'selectRequired',
    },
    {
      component: 'Input',
      fieldName: 'postName',
      label: '岗位名称',
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'postCode',
      label: '岗位编码',
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'postCategory',
      label: '类别编码',
    },
    {
      component: 'InputNumber',
      fieldName: 'postSort',
      label: '岗位排序',
      rules: 'required',
    },
    {
      component: 'RadioGroup',
      componentProps: {
        buttonStyle: 'solid',
        options: getDictOptions(DictEnum.SYS_NORMAL_DISABLE),
        optionType: 'button',
      },
      defaultValue: '0',
      fieldName: 'status',
      label: '岗位状态',
      rules: 'required',
    },
    {
      component: 'Textarea',
      fieldName: 'remark',
      formItemClass: 'items-baseline',
      label: '备注',
    },
  ];
  