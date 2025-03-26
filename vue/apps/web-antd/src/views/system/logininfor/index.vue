<script lang="ts" setup>
import type { VbenFormProps } from '#/adapter/form';
import type { DeepPartial } from '@vben/types';
import type { VxeTableGridOptions, VxeGridListeners } from '#/adapter/vxe-table';
import type { SysLogininfor } from '#/api/system/logininfor'

import { h, ref } from 'vue';
import { getVxePopupContainer } from '@vben/utils';

import { Page, useVbenDrawer } from '@vben/common-ui';

import { Button, message, Tag, Modal, Popconfirm } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysLogininforListApi, deleteSysLogininforApi } from '#/api/system/logininfor';
import viewDrawer from './view-drawer.vue';
import {
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
      componentProps: {
        format: "YYYY-MM-DD",
        valueFormat: "YYYY-MM-DD",
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
        // message.success(`Query params: ${JSON.stringify(formValues)}`);
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

const gridEvents: DeepPartial<VxeGridListeners> = {
  checkboxChange: handleCheckboxChange,
  checkboxAll: handleCheckboxChange,
};

const CheckboxChecked = ref(false);
function handleCheckboxChange() {
  CheckboxChecked.value = gridApi.grid.getCheckboxRecords().length > 0;
}


const [Grid, gridApi] = useVbenVxeGrid({
  formOptions,
  gridOptions,
  gridEvents,
});


const [ViewDrawer, drawerApi] = useVbenDrawer({
  connectedComponent: viewDrawer,
});

function handlePreview(record: SysLogininfor) {
  drawerApi.setData({ record });
  drawerApi.open();
}


async function handleDelete(row: SysLogininfor) {
  await deleteSysLogininforApi({ infoIds: [row.infoId] });
  message.success("删除成功");
  await handleRefresh();
}

async function handleRefresh() {
  await gridApi.query();
}


function handleMultiDelete() {
  const rows = gridApi.grid.getCheckboxRecords();
  const ids = rows.map((row: SysLogininfor) => row.infoId);
  Modal.confirm({
    title: '提示',
    okType: 'danger',
    content: `确认删除选中的${ids.length}条记录吗？`,
    onOk: async () => {
      await deleteSysLogininforApi({ infoIds: ids });
      message.success("删除成功");
      await handleRefresh();
    },
  });
}
</script>

<template>
  <Page auto-content-height>
    <Grid table-title="登录日志列表">
      <template #toolbar-tools>
        <Button class="mr-2 flex items-center" type="primary" :disabled="!CheckboxChecked" :icon="h(MdiDelete)" @click="handleMultiDelete" v-access:code="'cpm:monitor:logininfor:remove'">删除</Button>
      </template>
      <template #status="{ row }">
        <Tag :color="row.status == '0' ? 'green' : 'red'">{{ row.status == '0' ? '正常' : '关闭' }}</Tag>
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handlePreview(row)" v-access:code="'cpm:monitor:logininfor:query'">查看</Button>
          <Popconfirm :get-popup-container="getVxePopupContainer" placement="left" title="确定删除吗？"
            @confirm="handleDelete(row)" v-access:code="'cpm:monitor:logininfor:remove'">
            <Button class="mr-2 border-none p-0" :block="false" type="link" danger v-access:code="'cpm:monitor:logininfor:remove'" >删除</Button>
          </Popconfirm>
        </div>
      </template>
    </Grid>
    <ViewDrawer />
  </Page>
</template>
