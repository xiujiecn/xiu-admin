<!--
 * @description 系统配置管理页面
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 XiuAdmin CLI
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script lang="ts" setup>
import { h, ref } from 'vue';
import type { DeepPartial } from '@vben/types';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions, VxeGridListeners } from '#/adapter/vxe-table';
import type { SysConfig } from '#/api/system/config';

import { getVxePopupContainer } from '@vben/utils';
import { Page, useVbenDrawer } from '@vben/common-ui';

import { Button, message, Tag, Popconfirm, Modal } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysConfigListApi, deleteSysConfigApi } from '#/api/system/config';

import {
  MdiPlus,
  MdiDelete,
} from '@vben/icons';
import configDrawer from './config-drawer.vue';

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
        valueFormat: "YYYY-MM-DD",
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
        // message.success(`Query params: ${JSON.stringify(formValues)}`);
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


const [ConfigDrawer, configDrawerApi] = useVbenDrawer({
  connectedComponent: configDrawer,
});


function handleView(row: SysConfig) {
  const { configId } = row;
  configDrawerApi.setData({ id: configId, update: false, view: true });
  configDrawerApi.open();
}

function handleAdd() {
  configDrawerApi.setData({ update: false, view: false });
  configDrawerApi.open();
}

function handleEdit(row: SysConfig) {
  configDrawerApi.setData({ id: row.configId, update: true, view: false });
  configDrawerApi.open();
}

async function handleDelete(row: SysConfig) {
  await deleteSysConfigApi({ configIds: [row.configId] });
  message.success("删除成功");
  await handleRefresh();
}
async function handleRefresh() {
  await gridApi.query();
}


function handleMultiDelete() {
  const rows = gridApi.grid.getCheckboxRecords();
  const ids = rows.map((row: SysConfig) => row.configId);
  Modal.confirm({
    title: '提示',
    okType: 'danger',
    content: `确认删除选中的${ids.length}条记录吗？`,
    onOk: async () => {
      await deleteSysConfigApi({ configIds: ids });
      message.success("删除成功");
      await handleRefresh();
    },
  });
}
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="'参数列表'">
      <template #toolbar-tools>

        <Button class="mr-2 flex items-center " type="primary" :icon="h(MdiPlus)" @click="handleAdd" v-access:code="'cpm:system:config:add'">新增</Button>
        <Button class="mr-2 flex items-center" type="primary" :disabled="!CheckboxChecked" :icon="h(MdiDelete)"
          @click="handleMultiDelete" v-access:code="'cpm:system:config:remove'">删除</Button>
      </template>
      <template #configType="{ row }">
        <Tag :color="row.configType == 'Y' ? 'green' : 'red'">{{ row.configType == 'Y' ? '是' : '否' }}</Tag>
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleView(row)" v-access:code="'cpm:system:config:query'">查看</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleEdit(row)" v-access:code="'cpm:system:config:edit'">修改</Button>
          <Popconfirm :get-popup-container="getVxePopupContainer" placement="left" title="确定删除吗？"
            @confirm="handleDelete(row)" v-access:code="'cpm:system:config:remove'">
            <Button class="mr-2 border-none p-0" :block="false" type="link"
              danger v-access:code="'cpm:system:config:remove'">删除</Button>
            </Popconfirm>
        </div>
      </template>
    </Grid>
    <ConfigDrawer @reload="handleRefresh" />
  </Page>
</template>
