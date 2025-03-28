/**
 * @description 用户管理模型定义
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
    fieldName: 'userId',
  },
  {
    component: 'Input',
    fieldName: 'userName',
    label: '用户账号',
    rules: 'required',
  },
  // {
  //   component: 'InputPassword',
  //   fieldName: 'password',
  //   label: '用户密码',
  //   rules: 'required',
  // },
  {
    component: 'Input',
    fieldName: 'nickName',
    label: '用户昵称',
    rules: 'required',
  },
  {
    component: 'TreeSelect',
    // 在drawer里更新 这里不需要默认的componentProps
    defaultValue: undefined,
    fieldName: 'deptId',
    label: '所属部门',
    rules: 'selectRequired',
  },
  {
    component: 'Input',
    fieldName: 'phonenumber',
    label: '手机号码',
    defaultValue: undefined,
    rules: z
      .string()
      .regex(/^1[3-9]\d{9}$/, '请输入正确的手机号码')
      .optional()
      .or(z.literal('')),
  },
  {
    component: 'Input',
    fieldName: 'email',
    defaultValue: undefined,
    label: '邮箱',
    /**
     * z.literal 是 Zod 中的一种类型，用于定义一个特定的字面量值。
     * 它可以用于确保输入的值与指定的字面量完全匹配。
     * 例如，你可以使用 z.literal 来确保某个字段的值只能是特定的字符串、数字、布尔值等。
     * 即空字符串也可通过校验
     */
    rules: z.string().email('请输入正确的邮箱').optional().or(z.literal('')),
  },
  {
    component: 'RadioGroup',
    componentProps: {
      buttonStyle: 'solid',
      options: getDictOptions(DictEnum.SYS_USER_SEX),
      optionType: 'button',
    },
    defaultValue: '0',
    fieldName: 'sex',
    formItemClass: 'col-span-2 lg:col-span-1',
    label: '性别',
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
    formItemClass: 'col-span-2 lg:col-span-1',
    label: '状态',
  },
  {
    component: 'Select',
    componentProps: {
      getPopupContainer,
      mode: 'multiple',
      optionFilterProp: 'label',
      optionLabelProp: 'label',
      placeholder: '请先选择部门',
    },
    fieldName: 'postIds',
    help: '选择部门后, 将自动加载该部门下所有的岗位',
    label: '岗位',
  },
  {
    component: 'Select',
    componentProps: {
      getPopupContainer,
      mode: 'multiple',
      optionFilterProp: 'title',
      optionLabelProp: 'title',
    },
    fieldName: 'roleIds',
    label: '角色',
  },
  {
    component: 'Textarea',
    fieldName: 'remark',
    formItemClass: 'items-baseline',
    label: '备注',
  },
];