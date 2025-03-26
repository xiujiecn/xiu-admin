<script lang="ts" setup>
import { h, ref } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions, VxeGridListeners } from '#/adapter/vxe-table';
import type { DeepPartial } from '@vben/types';
import { Page } from '@vben/common-ui';
import { getVxePopupContainer } from '@vben/utils';
import { Button, message, Switch,Tag, Modal, Popconfirm } from 'ant-design-vue';
import dayjs from 'dayjs';
import { AccessControl, useAccess } from '@vben/access';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysUserOnlineListApi, deleteSysUserOnlineApi, type SysUserOnline } from '#/api/monitor/online';
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
        // message.success(`Query params: ${JSON.stringify(formValues)}`);
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

async function handleDelete(row: SysUserOnline) {
  await deleteSysUserOnlineApi({ ids: [row.onlineId] });
  message.success("删除成功");
  await gridApi.query();
}


function handleMultiDelete() {
  const rows = gridApi.grid.getCheckboxRecords();
  const ids = rows.map((row) => row.onlineId);
  Modal.confirm({
    title: '提示',
    okType: 'danger',
    content: `确认删除选中的${ids.length}条记录吗？`,
    onOk: async () => {
      await deleteSysUserOnlineApi({ ids: ids.map(Number) });
      message.success("删除成功");
      await gridApi.query();
    },
  });
}

</script>

<template>
  <Page auto-content-height>
    <Grid table-title="在线用户">
      <template #toolbar-tools>
        <Button class="mr-2 flex items-center" type="primary"  :disabled="!CheckboxChecked" :icon="h(MdiDelete)"  v-access:code="'cpm:monitor:online:batchLogout'" @click="handleMultiDelete">批量删除</Button>
      </template>
      <template #status="{ row }">
        <Tag :color="row.status == '0' ? 'green' : 'red'">{{ row.status == '0' ? '正常' : '关闭' }}</Tag>
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <AccessControl :codes="['cpm:monitor:online:forceLogout']" type="code">
            <Popconfirm title="确定删除吗？" v-if="row.id != 1" :get-popup-container="getVxePopupContainer" placement="left"  @confirm="handleDelete(row)" >  
              <Button class="mr-2 border-none p-0" :block="false" type="link"  danger >删除</Button>
            </Popconfirm>
          </AccessControl>
        </div>
      </template>
    </Grid>
  </Page>
</template>
