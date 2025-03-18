
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
    renderDictTags,
    renderHttpMethodTag,
    renderJsonPreview,
  } from '#/utils/render';

  export const querySchema : VbenFormSchema[] =  [
    {
      component: 'Input',
      fieldName: 'genType',
      label: '生成类型',
    },
    {
      component: 'Input',
      fieldName: 'varName',
      label: '实体命名',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: [
          {
            label: '生成成功',
            value: '0',
          },
          {
            label: '未开始',
            value: '1',
          },
          {
            label: '生成失败',
            value: '2',
          },
        ],
        placeholder: '请选择',
      },
      fieldName: 'status',
      label: '状态',
    },
  ];
  
  export const columns: VxeGridProps['columns'] = [
    { align: 'left', title: 'ID', type: 'checkbox', width: 80 },
    { field: 'genType', title: '生成类型' },
    { field: 'genTemplate', title: '生成模板' },
    { field: 'varName', title: '实体命名' },
    { field: 'tableComment', title: '生成名称' },
    { field: 'status', title: '状态', 
      slots: {
        default: ({ row }) => {
          return renderDict(row.status, DictEnum.SYS_GEN_STATUS);
        },
      }, },
    { field: 'createdAt', formatter: 'formatDateTime', title: '创建时间' },
    { title: '操作', width: 120, slots: { default: 'action' } }
  ]
  
  export const viewSchema: DescItem[] = [
    {
      field: 'tableId',
      label: 'ID',
    },
    {
      field: 'genType',
      label: '生成类型',
    },
    { 
      field: 'genTemplate',
      label: '生成模板',
    },
    {
      field: 'varName',
      label: '实体命名',
    },
    {
        field: 'tableComment',
      label: '生成名称',
    },
    {
      field: 'dbName',
      label: '数据库名称',
    },
    {
      field: 'tableName',
      label: '主表名称',
    },
    {
      field: 'daoName',
      label: '主表dao模型',
    },
    {
      field: 'masterColumns',
      label: '主表字段',
    },
    {
      field: 'addonName',
      label: '插件名称',
    },
    {
      field: 'status',
      label: '状态',
      render(value, { status }) {
        const operType = renderDict(status, DictEnum.SYS_GEN_STATUS);
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
      field: 'createdAt',
      label: '创建时间',
    },
    {
      field: 'createdBy',
      label: '创建人',
    },
    {
      field: 'updatedAt',
      label: '更新时间',
    },
    {
      field: 'updatedBy',
      label: '更新人',
    }
  ];

  export const drawerSchema: VbenFormSchema[] =  [
    {
      component: 'Input',
      dependencies: {
        show: () => false,
        triggerFields: [''],
      },
      fieldName: 'tableId',
      label: 'ID',
    },
    {
      component: 'Input',
      fieldName: 'genType',
      label: '生成类型',
    },
    {
      component: 'Input',
      fieldName: 'genTemplate',
      label: '生成模板',
    },
    
    
    
  ];
  