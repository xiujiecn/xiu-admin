/**
 * @description 定时任务模型定义
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */
import type {
    VbenFormSchema,
  } from '@vben/common-ui';
import type { DescItem } from '#/components/description';
import { getDictOptions } from '#/utils/dict';
import { DictEnum } from '@vben/constants';


import {
  renderDict,
} from '#/utils/render';

export const viewSchema: DescItem[] = [
  {
    field: 'jobName',
    label: '任务名称',
  },
  {
    field: 'remark',
    label: '任务描述',
  },
  {
    field: 'jobGroup',
    label: '任务分组',
    render(_, data) {
      return renderDict(data.jobGroup, DictEnum.SYS_JOB_GROUP);
    },
  },
  {
    field: 'invokeTarget',
    label: '调用方法',
  },
  {
    field: 'jobParams',
    label: '执行参数',
  },
  {
    field: 'cronExpression',
    label: 'cron表达式'
  },
  {
    field: 'misfirePolicy',
    label: '计划执行策略',
    render(_, data) {
      return renderDict(data.misfirePolicy, DictEnum.SYS_MISSFIRE_POLICY);
    },
  },
  {
    field: 'concurrent',
    label: '是否并发执行',
    render(_, data) {
      return renderDict(data.concurrent, DictEnum.SYS_JOB_CONCURRENT);
    },
  },
  {
    field: 'status',
    label: '状态',
    render(_, data) {
      return renderDict(data.status, DictEnum.SYS_NORMAL_DISABLE);
    },
  }
];

export const drawerSchema: VbenFormSchema[] =  [
  {
    component: 'Input',
    dependencies: {
      show: () => false,
      triggerFields: [''],
    },
    fieldName: 'jobId',
    label: '参数主键',
  },
  {
    component: 'Input',
    fieldName: 'jobName',
    label: '任务名称',
    rules: 'required'
  },
  {
    component: 'Input',
    fieldName: 'remark',
    label: '任务描述',
  },
  {
    component: 'Select',
    fieldName: 'jobGroup',
    label: '任务分组',
    componentProps: {
      buttonStyle: 'solid',
      options: getDictOptions(DictEnum.SYS_JOB_GROUP),
      optionType: 'button',
    },
    rules: 'required'
  },
  {
    component: 'Input',
    fieldName: 'invokeTarget',
    label: '调用方法',
    rules: 'required'
  },
  {
    component: 'Input',
    fieldName: 'jobParams',
    label: '执行参数',
  },
  {
    component: '#cronExpression',
    fieldName: 'cronExpression',
    label: 'cron表达式',
    rules: 'required'
  },
  {
    component: 'RadioGroup',
    componentProps: {
      buttonStyle: 'solid',
      options: getDictOptions(DictEnum.SYS_MISSFIRE_POLICY),
      optionType: 'button',
    },
    defaultValue: '1',
    fieldName: 'misfirePolicy',
    label: '计划执行策略'
  },
  {
    component: 'RadioGroup',
    componentProps: {
      buttonStyle: 'solid',
      options: getDictOptions(DictEnum.SYS_JOB_CONCURRENT),
      optionType: 'button',
    },
    defaultValue: '0',
    fieldName: 'concurrent',
    label: '是否并发执行'
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
    label: '状态'
  }
];
  