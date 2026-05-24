
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
    renderDict,
    renderHttpMethodTag,
    renderJsonPreview,
  } from '#/utils/render';

  const accessPolicyOptions = [
    { color: 'orange', label: '私有', value: '0' },
    { color: 'green', label: '公开', value: '1' },
    { color: 'blue', label: '自定义', value: '2' },
  ];

  export const querySchema : VbenFormSchema[] =  [
    {
      component: 'Input',
      fieldName: 'configKey',
      label: '配置名称',
    },
    {
      component: 'Input',
      fieldName: 'bucketName',
      label: '桶名称',
    },
    {
      component: 'Select',
      componentProps: {
        options: getDictOptions(DictEnum.SYS_YES_NO),
      },
      fieldName: 'status',
      label: '是否默认',
    },
  ];
  export const columns: VxeGridProps['columns'] = [
    { type: 'checkbox', width: 60 },
  {
    title: '配置名称',
    field: 'configKey',
  },
  {
    title: '访问站点',
    field: 'endpoint',
    showOverflow: true,
  },
  {
    title: '桶名称',
    field: 'bucketName',
  },
  {
    title: '域',
    field: 'region',
  },
  {
    title: '权限桶类型',
    field: 'accessPolicy',
    slots: {
      default: ({ row }) => {
        const current = accessPolicyOptions.find(
          (item) => item.value === row.accessPolicy,
        );
        if (current) {
          return <Tag color={current.color}>{current.label}</Tag>;
        }
        return '未知类型';
      },
    },
  },
  {
    title: '是否默认',
    field: 'status',
    slots: {
      default: ({ row }) => {
        return renderDict(row.status, DictEnum.SYS_YES_NO_NUM);
      },
    },
  },
  {
    field: 'action',
    fixed: 'right',
    slots: { default: 'action' },
    title: '操作',
    width: 120,
  },
  ]
  
  export const viewSchema: DescItem[] = [
    {
      field: 'ossConfigId',
      label: '配置ID',
    },
    {
      field: 'configKey',
      label: '配置名称',
    },
    {
      field: 'accessKey',
      label: 'accessKey',
    },
    {
      field: 'secretKey',
      label: 'secretKey',
    },
    {
      field: 'domain',
      label: '域',
    },
    {
      field: 'isHttps',
      label: '是否https',
      render(value, { isHttps }) {
        const operType = renderDict(isHttps, DictEnum.SYS_YES_NO);
        return (
          <div class="flex items-center">
            <Tag>{value}</Tag>
            {operType}
          </div>
        )
        ;
      },
    },
    {
      field: 'region',
      label: '区域',
    },
    {
      field: 'accessPolicy',
      label: '权限桶类型',
      render(value, { accessPolicy }) {
        const operType = renderDict(accessPolicy, DictEnum.SYS_OSS_ACCESS_POLICY);
        return (
          <div class="flex items-center">
            <Tag>{value}</Tag>
            {operType}
          </div>
        )
        ;
      },
    },
    {
      field: 'status',
      label: '是否默认',
      render(value, { status }) {
        const operType = renderDict(status, DictEnum.SYS_YES_NO_NUM);
        return (
          <div class="flex items-center">
            <Tag>{value}</Tag>
            {operType}
          </div>
        )
        ;
      },
    },
    {
      field: 'ext1',
      label: '扩展字段',
    },
    {
      field: 'remark',
      label: '备注',
    },
    {
      field: 'createdDept',
      label: '创建组织',
    },
    {
      field: 'createdBy',
      label: '创建者',
    },
    {
      field: 'createdAt',
      label: '创建时间',
    },
    {
      field: 'updatedBy',
      label: '更新者',
    },
    {
      field: 'updatedAt',
      label: '更新时间',
    },
  ];

  export const drawerSchema: VbenFormSchema[] =  [
    {
      component: 'Input',
      dependencies: {
        show: () => false,
        triggerFields: [''],
      },
      fieldName: 'ossConfigId',
    },
    {
      component: 'Divider',
      componentProps: {
        orientation: 'center',
      },
      fieldName: 'divider1',
      hideLabel: true,
      renderComponentContent: () => ({
        default: () => '基本信息',
      }),
    },
    {
      component: 'Input',
      fieldName: 'configKey',
      label: '配置名称',
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'endpoint',
      label: '服务地址',
      renderComponentContent: (formModel) => ({
        addonBefore: () => (formModel.isHttps === 'Y' ? 'https://' : 'http://'),
      }),
      rules: z
        .string()
        .refine((domain) => domain && !/^https?:\/\/.*/.test(domain), {
          message: '请输入正确的域名, 不需要http(s)',
        }),
    },
    {
      component: 'Input',
      fieldName: 'domain',
      label: '自定义域名',
    },
    {
      component: 'Input',
      fieldName: 'tip',
      label: '占位作为提示使用',
      hideLabel: true,

    },
    {
      component: 'Divider',
      componentProps: {
        orientation: 'center',
      },
      fieldName: 'divider2',
      hideLabel: true,
      renderComponentContent: () => ({
        default: () => '认证信息',
      }),
    },
    {
      component: 'Input',
      fieldName: 'accessKey',
      label: 'accessKey',
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'secretKey',
      label: 'secretKey',
      rules: 'required',
    },
    {
      component: 'Divider',
      componentProps: {
        orientation: 'center',
      },
      fieldName: 'divider3',
      hideLabel: true,
      renderComponentContent: () => ({
        default: () => '其他信息',
      }),
    },
    {
      component: 'Input',
      fieldName: 'bucketName',
      label: '桶名称',
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'prefix',
      label: '前缀',
    },
    {
      component: 'RadioGroup',
      componentProps: {
        buttonStyle: 'solid',
        options: accessPolicyOptions,
        optionType: 'button',
      },
      defaultValue: '0',
      fieldName: 'accessPolicy',
      formItemClass: 'col-span-3 lg:col-span-2',
      label: '权限桶类型',
    },
    {
      component: 'RadioGroup',
      componentProps: {
        buttonStyle: 'solid',
        options: getDictOptions(DictEnum.SYS_YES_NO),
        optionType: 'button',
      },
      defaultValue: 'N',
      fieldName: 'isHttps',
      formItemClass: 'col-span-3 lg:col-span-1',
      label: '是否https',
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'region',
      label: '区域',
    },
    {
      component: 'Textarea',
      fieldName: 'remark',
      formItemClass: 'items-baseline',
      label: '备注',
    },
  ];
  