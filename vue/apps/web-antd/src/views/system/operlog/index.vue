
<script lang="ts" setup>
import { h } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SysOperLog } from '#/api/system/oper-log';
import { Page, useVbenDrawer } from '@vben/common-ui';
import { AccessControl, useAccess } from '@vben/access';
const { hasAccessByCodes } = useAccess();

import { Button, message,Tag  } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysOperLogList } from '#/api/system/oper-log';

import viewDrawer from './view-drawer.vue';

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
      fieldName: 'operIp',
      label: '操作地址',
    },
    {
      component: 'Input',
      fieldName: 'businessType',
      label: '系统模块',
    },
    {
      component: 'Input',
      fieldName: 'createdBy',
      label: '创建者',
    },


    {
      component: 'RangePicker',
      componentProps:{
        format:"YYYY-MM-DD",
        valueFormat:"YYYY-MM-DD",
      },
      // defaultValue: [dayjs().subtract(7, 'days'), dayjs()],
      fieldName: 'createdAt',
      label: '创建时间',
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
  },
  columns: [
    { type: 'checkbox', width: 40 },
    { field: 'operId', title: 'ID' },
    { field: 'businessType', title: '业务类型', slots: { default: 'businessType' } },
    { field: 'method', title: '操作' },
    { field: 'operName', title: '操作人员' },
    { field: 'operIp', title: 'IP地址' },
    { field: 'operLocation', title: 'IP信息' },
    { field: 'status', title: '操作状态', slots: { default: 'status' } },
    { field: 'operTime', title: '操作时间' },
    { field: 'costTime', title: '操作耗时', slots: { default: 'costTime' } },
    { title: '操作', width: 50, slots: { default: 'action' } }
  ],
  exportConfig: {},
  height: 'auto',
  keepSource: true,
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        // message.success(`Query params: ${JSON.stringify(formValues)}`);
        return await getSysOperLogList({
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

const [ViewDrawer, drawerApi] = useVbenDrawer({
  connectedComponent: viewDrawer,
});

function handlePreview(record: SysOperLog) {
  drawerApi.setData({ record });
  drawerApi.open();
}


</script>

<template>
  <Page auto-content-height>
    <Grid table-title="操作日志列表">
      <template #businessType="{ row }">
        {{  row.businessType == 1 ? '新增' : row.businessType == 2 ? '修改' : row.businessType == 3 ? '删除' : '其他' }}
      </template>
      <template #status="{ row }">
        <Tag :color="row.status == 0 ? 'green' : 'red'">{{ row.status == 0 ? '正常' : '关闭' }}</Tag>
      </template>
      <template #costTime="{ row }">
        {{ row.costTime }} ms
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handlePreview(row)" v-access:code="'cpm:monitor:operlog:query'">查看</Button>
        </div>
      </template>
    </Grid>
    <ViewDrawer />
  </Page>
</template>
