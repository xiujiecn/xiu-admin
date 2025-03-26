<script lang="ts" setup>
import { h, ref, onMounted } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions, VxeGridListeners } from '#/adapter/vxe-table';
import type { DeepPartial } from '@vben/types';
import type { SysClient } from '#/api/system/client';

import { getVxePopupContainer } from '@vben/utils';
import { Page, useVbenDrawer } from '@vben/common-ui';
import { Button, message, Tag, Modal, Popconfirm, Switch } from 'ant-design-vue';
import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysGenTableListApi, deleteSysGenTableApi, getGenCodesSelectsApi } from '#/api/gen_codes/gen_table';
import { MdiPlus, MdiDelete, } from '@vben/icons';
import { useRouter } from 'vue-router';
import { querySchema, columns, setSelectListObj } from './model';
import viewDrawer from './view-drawer.vue';
import editDrawer from './edit-drawer.vue';
import { AccessControl, useAccess } from '@vben/access';
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
    labelField: 'tableId',
  },
  rowConfig: {
    keyField: 'tableId',
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
        return await getSysGenTableListApi({
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

async function handleDelete(row: any) {
  await deleteSysGenTableApi({ tableIds: [row.tableId] });
  message.success("删除成功");
  await handleRefresh();
}

async function handleRefresh() {
  await gridApi.query();
}

function handleMultiDelete() {
  const rows = gridApi.grid.getCheckboxRecords();
  const ids = rows.map((row) => row.tableId);

  Modal.confirm({
    title: '提示',
    okType: 'danger',
    content: `确认删除选中的${ids.length}条记录吗？`,
    onOk: async () => {
      await deleteSysGenTableApi({ tableIds: ids });
      message.success("删除成功");
      await handleRefresh();
    },
  });
}
onMounted(async () => {
  const res = await getGenCodesSelectsApi({});
  setSelectListObj(res);

});

const router = useRouter();
const handleClickDevelop = (tableId: number) => {
  router.push(`/tool/gen-develop?tableId=${tableId}`);
}
</script>

<template>
  <Page auto-content-height>
    <Grid table-title="代码生成">
      <template #toolbar-tools>
        <Button class="mr-2 flex items-center " type="primary" :icon="h(MdiPlus)" @click="handleAdd" v-access:code="'cpm:tool:gen:import'">导入生成</Button>
        <Button class="mr-2 flex items-center" type="primary" :disabled="!CheckboxChecked" :icon="h(MdiDelete)"
          @click="handleMultiDelete" v-access:code="'cpm:tool:gen:remove'">删除</Button>
      </template>

      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link"
            @click="handleClickDevelop(row.tableId)" v-access:code="'cpm:tool:gen:develop'">生成配置</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handlePreview(row)" v-access:code="'cpm:tool:gen:query'">查看</Button>
          <AccessControl :codes="['cpm:tool:gen:remove']" type="code">
            <Popconfirm title="确定删除吗？" :get-popup-container="getVxePopupContainer" placement="left"
              @confirm="handleDelete(row)" >
              <Button class="mr-2 border-none p-0" :block="false" type="link" danger
                v-access:code="'cpm:tool:gen:remove'">删除</Button>
            </Popconfirm>
          </AccessControl>
        </div>
      </template>
    </Grid>
    <ViewDrawer />
    <EditDrawer @reload="handleRefresh" />
  </Page>
</template>
