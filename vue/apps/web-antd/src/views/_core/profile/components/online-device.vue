<script setup lang="ts">
import type { Recordable } from '@vben/types';

import type { VxeGridProps } from '#/adapter/vxe-table';

import { Popconfirm } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysUserOnlineListApi, deleteSysUserOnlineApi } from '#/api/monitor/online';
import { useAccessStore, useUserStore } from '@vben/stores';
const userStore = useUserStore();
const gridOptions: VxeGridProps = {
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
  keepSource: true,
  pagerConfig: {
    enabled: false,
  },
  proxyConfig: {
    ajax: {
      query: async ({ page }) => {
        return await getSysUserOnlineListApi({
          page: page.currentPage,
          pageSize: page.pageSize,
          userName: userStore.userInfo?.userName,
          ipaddr: '',
          status: '',
        });
      },
    },
  },
  rowConfig: {
    keyField: 'tokenId',
  },
};

const [BasicTable, tableApi] = useVbenVxeGrid({ gridOptions });

async function handleForceOffline(row: Recordable<any>) {
  await deleteSysUserOnlineApi({
    id: row.onlineId,
  });
  await tableApi.query();
}
</script>

<template>
  <div>
    <BasicTable table-title="我的在线设备">
      <template #action="{ row }">
        <Popconfirm
          :title="`确认强制下线[${row.userName}]?`"
          placement="left"
          @confirm="handleForceOffline(row)"
        >
          <a-button danger size="small" type="link">强制下线</a-button>
        </Popconfirm>
      </template>
    </BasicTable>
  </div>
</template>
