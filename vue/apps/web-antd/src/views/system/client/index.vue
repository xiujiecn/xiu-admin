<!--
 * @description 客户端管理页面
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 XiuAdmin CLI
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script lang="ts" setup>
import { h,ref } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions, VxeGridListeners } from '#/adapter/vxe-table';
import type { DeepPartial } from '@vben/types';
import type { SysClient } from '#/api/system/client';

import { getVxePopupContainer } from '@vben/utils';
import { Page, useVbenDrawer } from '@vben/common-ui';
import { Button, message,Tag, Modal, Popconfirm,Switch } from 'ant-design-vue';
import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysClientListApi, postSysClientDeleteApi, postSysClientStatusApi } from '#/api/system/client'; 
import { MdiPlus, MdiDelete,} from '@vben/icons';
import { AccessControl, useAccess } from '@vben/access';

import { querySchema, columns, } from './model';
import viewDrawer from './view-drawer.vue';
import editDrawer from './edit-drawer.vue';

const { hasAccessByCodes } = useAccess();
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
  schema: querySchema,
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
    labelField: 'id',
  },
  rowConfig: {
    keyField: 'id',
  },
  columns: columns,
  exportConfig: {},
  height: 'auto',
  keepSource: true,
  showOverflow: false,
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        // message.success(`Query params: ${JSON.stringify(formValues)}`);
        return await getSysClientListApi({
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

function handlePreview(record: SysClient) {
  drawerApi.setData({ record });
  drawerApi.open();
}



const [EditDrawer, editDrawerApi] = useVbenDrawer({
  connectedComponent: editDrawer,
});

function handleAdd() {
  editDrawerApi.setData({ update: false, view: false });
  editDrawerApi.open();
}

function handleEdit(row: SysClient) {
  editDrawerApi.setData({ id: row.id, update: true, view: false });
  editDrawerApi.open();
}

async function handleDelete(row: SysClient) {
  if(row.id === 1) {
    message.error("PC客户端不允许删除");
    return;
  }
  await postSysClientDeleteApi({ ids: [row.id] });
  message.success("删除成功");
  await handleRefresh();
}

async function handleRefresh() {
  await gridApi.query();
}


function handleMultiDelete() {
  const rows = gridApi.grid.getCheckboxRecords();
  const ids: string[] = [];
  for (const row of rows) {
    if (row.id != 1) {
      ids.push(row.id);
    }else {
      message.error("PC客户端不允许删除");
      return;
    }
  }

  Modal.confirm({
    title: '提示',
    okType: 'danger',
    content: `确认删除选中的${ids.length}条记录吗？`,
    onOk: async () => {
      await postSysClientDeleteApi({ ids: ids });
      message.success("删除成功");
      await handleRefresh();
    },
  });
}

async function handleStatusChange(row: SysClient) {
  // console.log('vue/apps/web-antd/src/views/system/client/index.vue handleStatusChange',row);
  if(row.id === 1) {
    message.error("PC客户端不允许禁用");
    return;
  }
  await postSysClientStatusApi({ id: row.id, status: row.status }); 
  await message.success("操作成功")
  await handleRefresh();
}

</script>

<template>
  <Page auto-content-height>
    <Grid table-title="客户端列表">
      <template #toolbar-tools>
        <Button class="mr-2 flex items-center " type="primary" :icon="h(MdiPlus)" @click="handleAdd" v-access:code="'cpm:system:client:add'">新增</Button>
        <Button class="mr-2 flex items-center" type="primary" :disabled="!CheckboxChecked" :icon="h(MdiDelete)" @click="handleMultiDelete" v-access:code="'cpm:system:client:remove'">删除</Button>
      </template>
      <template #status="{ row }">
        <!-- pc不允许禁用 禁用了直接登录不了 应该设置disabled -->
        <!-- 登录提示: 认证权限类型已禁用 -->
        <Switch
          v-model:checked="row.status" :checkedValue="'0'" :unCheckedValue="'1'" :disabled="row.id === 1 || !hasAccessByCodes(['cpm:system:client:edit'])"
          @change="handleStatusChange(row)"
        />
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handlePreview(row)" v-access:code="'cpm:system:client:query'">查看</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleEdit(row)" v-access:code="'cpm:system:client:edit'">修改</Button>
          <AccessControl :codes="['cpm:system:client:remove']" type="code">
            <Popconfirm title="确定删除吗？" v-if="row.id != 1" :get-popup-container="getVxePopupContainer" placement="left"  @confirm="handleDelete(row)" >  
              <Button class="mr-2 border-none p-0" :block="false" type="link"  danger  >删除</Button>
            </Popconfirm>
          </AccessControl>
        </div>
      </template>
    </Grid>
    <ViewDrawer />
    <EditDrawer @reload="handleRefresh" />
  </Page>
</template>
