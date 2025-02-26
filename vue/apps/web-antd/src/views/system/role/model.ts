

import type {
    VbenFormSchema,
  } from '@vben/common-ui';
  import { z } from '@vben/common-ui';
  import { getDictOptions } from '#/utils/dict';
  import { DictEnum } from '@vben/constants';
  import { getPopupContainer } from '@vben/utils';
  
  
export const authScopeOptions = [
    { color: 'green', label: '全部数据权限', value: '1' },
    { color: 'default', label: '自定数据权限', value: '2' },
    { color: 'orange', label: '本部门数据权限', value: '3' },
    { color: 'cyan', label: '本部门及以下数据权限', value: '4' },
    { color: 'error', label: '仅本人数据权限', value: '5' },
    { color: 'default', label: '部门及以下或本人数据权限', value: '6' },
];

export const drawerSchema: VbenFormSchema[] =  [
    {
      component: 'Input',
      dependencies: {
        show: () => false,
        triggerFields: [''],
      },
      fieldName: 'roleId',
      label: '角色ID',
    },
    {
      component: 'Input',
      fieldName: 'roleName',
      label: '角色名称',
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'roleKey',
      help: '如: test simpleUser等',
      label: '权限标识',
      rules: 'required',
    },
    {
      component: 'InputNumber',
      fieldName: 'roleSort',
      label: '角色排序',
      rules: 'required',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: false,
        options: getDictOptions(DictEnum.SYS_NORMAL_DISABLE),
        getPopupContainer,
      },
      defaultValue: '0',
      fieldName: 'status',
      help: '修改后, 拥有该角色的用户将自动下线.',
      label: '角色状态',
      rules: 'required',
    },
    {
      component: 'InputNumber',
      dependencies: {
        show: () => false,
        triggerFields: [''],
      },
      fieldName: 'menuCheckStrictly',
      label: '菜单树父子关联',
    },
    {
      component: 'CheckboxGroup',
      defaultValue: [],
      fieldName: 'menuIds',
      label: '菜单权限',
      formItemClass: 'col-span-2',
    },
    {
      component: 'Textarea',
      defaultValue: '',
      fieldName: 'remark',
      formItemClass: 'items-baseline col-span-2',
      label: '备注',
    },
  ];

  export const dataScopeModalSchema: VbenFormSchema[] = [
    {
      component: 'Input',
      fieldName: 'roleId',
      label: '角色ID',
      dependencies: {
        show: () => false,
        triggerFields: [''],
      },
    },
    {
      component: 'Input',
      fieldName: 'roleName',
      label: '角色名称',
      componentProps: {
        readonly: true,
      },
    },
    {
      component: 'Select',
      fieldName: 'dataScope',
      label: '数据权限',
      componentProps: {
        options: authScopeOptions,
      },
    },
    {
      component: 'CheckboxGroup',
      defaultValue: [],
      fieldName: 'deptIds',
      label: '部门权限',
      // formItemClass: 'col-span-2',
      dependencies: {
        show: (values) => values?.dataScope === '2',
        triggerFields: ['dataScope'],
      },
    },
    {
      component: 'InputNumber',
      dependencies: {
        show: () => false,
        triggerFields: [''],
      },
      
      fieldName: 'deptCheckStrictly',
      label: '部门树父子关联',
    },
  ];