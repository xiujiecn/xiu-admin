<script lang="ts" setup>
import { h } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions } from '#/adapter/vxe-table';

import { Page } from '@vben/common-ui';

import { Button, message, Switch,Tag  } from 'ant-design-vue';
import dayjs from 'dayjs';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysUserOnlineListApi, deleteSysUserOnlineApi } from '#/api/monitor/online';
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
    labelField: 'onlineId',
  },
  columns: [
    { align: 'left', title: 'ID', type: 'checkbox', width: 80 },
    { field: 'userName', title: '用户名称' },
    { field: 'uuid', title: 'UUID' },
    { field: 'clientKey', title: '客户端' },
    { field: 'deviceType', title: '设备类型' },
    { field: 'ipaddr', title: '地址' },
    { field: 'loginLocation', title: '登录地点' },
    { field: 'browser', title: '浏览器' },
    { field: 'os', title: '操作系统' },
    { field: 'loginTime', title: '登录时间' },
    { field: 'expireTime', title: '过期时间' },
    { title: '操作', width: 40, slots: { default: 'action' } }
  ],
  exportConfig: {},
  height: 'auto',
  keepSource: true,
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        message.success(`Query params: ${JSON.stringify(formValues)}`);
        return await getSysUserOnlineListApi({
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
        <Button class="mr-2 flex items-center" type="primary" disabled :icon="h(MdiDelete)">删除</Button>
      </template>
      <template #status="{ row }">
        <Tag :color="row.status == '0' ? 'green' : 'red'">{{ row.status == '0' ? '正常' : '关闭' }}</Tag>
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link"  danger>删除</Button>
        </div>
      </template>
    </Grid>
  </Page>
</template>
