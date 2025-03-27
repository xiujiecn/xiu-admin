/**
 * @description 登录日志模型定义
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 XiuAdmin CLI
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */
import type {
    VbenFormSchema,
  } from '@vben/common-ui';
  import type { VxeGridProps } from '#/adapter/vxe-table';
  import type { DescItem } from '#/components/description';
  import { z } from '@vben/common-ui';
  import { getDictOptions } from '#/utils/dict';
  import { DictEnum } from '@vben/constants';
  import { getPopupContainer } from '@vben/utils';
  import { Tag } from 'ant-design-vue';
  import { h } from 'vue';

  import {
    renderBrowserIcon, renderOsIcon, 
    renderDict,
    renderHttpMethodTag,
    renderJsonPreview,
  } from '#/utils/render';

  
  export const viewSchema: DescItem[] = [
    {
      field: 'status',
      label: '登录状态',
      labelMinWidth: 80,
      render(value) {
        return renderDict(value, DictEnum.SYS_COMMON_STATUS);
      },
    },
    {
      field: 'clientKey',
      label: '登录平台',
      render(value) {
        if (value) {
          return value.toUpperCase();
        }
        return '';
      },
    },
    {
      field: 'ipaddr',
      label: '账号信息',
      render(_, data) {
        const { ipaddr, loginLocation, userName } = data;
        return `账号: ${userName} / ${ipaddr} / ${loginLocation}`;
      },
    },
    {
      field: 'loginTime',
      label: '登录时间',
    },
    {
      field: 'msg',
      label: '登录信息',
      render(_, data: any) {
        const { msg, status } = data;
        return (
          <span class={['font-bold', status === '0' ? '' : 'text-red-500']}>
            {msg}
          </span>
        );
      },
    },
    {
      field: 'os',
      label: '登录设备',
    },
    {
      field: 'browser',
      label: '浏览器',
    },
  ];

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
  