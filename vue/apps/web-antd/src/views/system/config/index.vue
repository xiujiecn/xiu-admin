<script lang="ts" setup>
import { h } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions } from '#/adapter/vxe-table';

import { Page } from '@vben/common-ui';

import { Button, message, Switch,Tag  } from 'ant-design-vue';
import dayjs from 'dayjs';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysConfigListApi } from '#/api/system/config'; 

import {
  MdiPlus,
  MdiEdit,
  MdiDelete,
} from '@vben/icons';

interface RowType {
  category: string;
  color: string;
  id: string;
  price: string;
  productName: string;
  releaseDate: string;
}

const formOptions: VbenFormProps = {
  // 默认展开
  collapsed: false,
  fieldMappingTime: [['date', ['start', 'end']]],
  schema: [
    {
      component: 'Input',
      fieldName: 'configName',
      label: '参数名称',
    },
    {
      component: 'Input',
      fieldName: 'configKey',
      label: '参数键名',
    },
    {
      component: 'Input',
      fieldName: 'configValue',
      label: '参数键值',
    },
    

    {
      component: 'RangePicker',
      // defaultValue: [dayjs().subtract(7, 'days'), dayjs()],
      fieldName: 'createdAt',
      label: '创建时间',
      componentProps: {
        format: 'YYYY-MM-DD',
        valueFormat:"YYYY-MM-DD",
      },
    },
  ],
  // 控制表单是否显示折叠按钮
  showCollapseButton: true,
  // 是否在字段值改变时提交表单
  submitOnChange: true,
  // 按下回车时是否提交表单
  submitOnEnter: false,
};

const gridOptions: VxeTableGridOptions<RowType> = {
  checkboxConfig: {
    highlight: true,
    labelField: 'configId',
  },
  columns: [
    { align: 'left', title: 'ID', type: 'checkbox', width: 80 },
    { field: 'configName', title: '参数名称' },
    { field: 'configKey', title: '参数键名' },
    { field: 'configValue', title: '参数键值' },
    { field: 'configType', title: '系统内置', slots: { default: 'configType' } },
    { field: 'remark', title: '备注' },
    { field: 'createdAt', formatter: 'formatDateTime', title: '创建时间' },
    { title: '操作', width: 120, slots: { default: 'action' } }
  ],
  exportConfig: {},
  height: 'auto',
  keepSource: true,
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        message.success(`Query params: ${JSON.stringify(formValues)}`);
        return await getSysConfigListApi({
          page: page.currentPage,
          pageSize: page.pageSize,
          ...formValues,
        });
      },
    },
  },
  toolbarConfig: {
    custom: true,
    export: true,
    refresh: true,
    resizable: true,
    search: true,
    zoom: true,
  },
};

const [Grid] = useVbenVxeGrid({
  formOptions,
  gridOptions,
});
</script>

<template>
  <Page auto-content-height>
    <Grid>
      <template #toolbar-actions>
        
        <Button class="mr-2 flex items-center " type="primary" :icon="h(MdiPlus)">新增</Button>
        <Button class="mr-2 flex items-center bg-green-500"  disabled :icon="h(MdiEdit)">编辑</Button>
        <Button class="mr-2 flex items-center" type="primary" disabled :icon="h(MdiDelete)">删除</Button>
      </template>
      <template #configType="{ row }">
        <Tag :color="row.configType == 'Y' ? 'green' : 'red'">{{ row.configType == 'Y' ? '是' : '否' }}</Tag>
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link">查看</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link">修改</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link"  danger>删除</Button>
        </div>
      </template>
    </Grid>
  </Page>
</template>
