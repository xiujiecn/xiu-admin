<script lang="ts" setup>
import { h } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions } from '#/adapter/vxe-table';

import { Page } from '@vben/common-ui';

import { Button, message, Switch,Tag  } from 'ant-design-vue';
import dayjs from 'dayjs';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysLogininforListApi } from '#/api/system/logininfor'; 

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
      fieldName: 'ipaddr',
      label: '登录地址',
    },
    {
      component: 'Input',
      fieldName: 'userName',
      label: '用户名称',
    },
    // 登录状态
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: [
          {
            label: '成功',
            value: '0',
          },
          {
            label: '失败',
            value: '1',
          },
        ],
        placeholder: '请选择',
      },
      fieldName: 'status',
      label: '登录状态',
    },
    {
      component: 'RangePicker',
      componentProps:{
        format:"YYYY-MM-DD",
        valueFormat:"YYYY-MM-DD",
      },
      // defaultValue: [dayjs().subtract(7, 'days'), dayjs()],
      fieldName: 'createdAt',
      label: '登录时间',
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
    labelField: 'infoId',
  },
  columns: [
    { align: 'left', title: 'ID', type: 'checkbox', width: 80 },
    { field: 'userName', title: '用户名称' },
    { field: 'clientKey', title: '客户端' },
    { field: 'deviceType', title: '设备类型' },
    { field: 'ipaddr', title: '地址' },
    { field: 'loginLocation', title: '登录地点' },
    { field: 'browser', title: '浏览器' },
    { field: 'os', title: '操作系统' },
    { field: 'status', title: '登录状态' },
    { field: 'msg', title: '登录信息' },
    { field: 'loginTime', title: '登录时间' },
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
        return await getSysLogininforListApi({
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
      <template #status="{ row }">
        <Tag :color="row.status == '0' ? 'green' : 'red'">{{ row.status == '0' ? '正常' : '关闭' }}</Tag>
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
