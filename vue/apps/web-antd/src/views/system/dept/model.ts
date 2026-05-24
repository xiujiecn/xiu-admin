/**
 * @description 组织管理模型定义
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
      component: 'RadioGroup',
      componentProps: {
        buttonStyle: 'solid',
        options: getDictOptions(DictEnum.SYS_DEPT_TYPE),
        optionType: 'button',
      },
      defaultValue: '0',
      fieldName: 'deptType',
      label: '组织类型',
    },
    {
      component: 'Input',
      dependencies: {
        show: () => false,
        triggerFields: [''],
      },
      fieldName: 'deptId',
    },
    {
      component: 'TreeSelect',
      componentProps: {
        getPopupContainer,
      },
      dependencies: {
        show: (model) => model.parentId !== 0,
        triggerFields: ['parentId'],
      },
      fieldName: 'parentId',
      label: '上级组织',
      rules: 'selectRequired',
    },
    {
      component: 'Input',
      fieldName: 'deptName',
      label: '组织名称',rules: z.string()
      .min(1, '组织名称不能为空')
      .max(32, '组织名称最大长度32位'),
    },
    {
      component: 'InputNumber',
      fieldName: 'orderNum',
      label: '显示排序',
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'deptCategory',
      label: '组织编码',
    },
    {
      component: 'Select',
      componentProps: {
        // 选中了就只能修改 不能重置为无负责人
        allowClear: false,
        getPopupContainer,
      },
      fieldName: 'leader',
      label: '负责人',
    },
    {
      component: 'Input',
      fieldName: 'phone',
      label: '联系电话',
      rules: z
        .string()
        .regex(/^1[3,4578]\d{9}$/, { message: '请输入正确的手机号' })
        .optional()
        .or(z.literal('')),
    },
    {
      component: 'Input',
      fieldName: 'email',
      label: '邮箱',
      rules: z
        .string()
        .email({ message: '请输入正确的邮箱' })
        .optional()
        .or(z.literal('')),
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
      label: '状态',
    },

  ];
