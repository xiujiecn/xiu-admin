<script lang="ts" setup>
import { h, ref } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions,VxeGridListeners  } from '#/adapter/vxe-table';
import type { DeepPartial } from '@vben/types';
import type { SysTenantPackageListData } from '#/api/system/tenant_package';
import { Page, useVbenDrawer } from '@vben/common-ui';
import { getVxePopupContainer } from '@vben/utils';
import { Button, message,Tag, Popconfirm,Switch,Modal   } from 'ant-design-vue';
import dayjs from 'dayjs';
import { AccessControl, useAccess } from '@vben/access';
import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysTenantPackageListApi,deleteSysTenantPackageApi,statusSysTenantPackageApi } from '#/api/system/tenant_package';
import {
  MdiPlus,
  MdiEdit,
  MdiDelete,
} from '@vben/icons';
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
    labelField: 'packageId',
  },
  columns: columns,
  exportConfig: {},
  height: 'auto',
  keepSource: true,
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        // message.success(`Query params: ${JSON.stringify(formValues)}`);
        return await getSysTenantPackageListApi({
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

function handlePreview(record: SysTenantPackageListData) {
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

function handleEdit(row: SysTenantPackageListData) {
  editDrawerApi.setData({ id: row.packageId, update: true, view: false });
  editDrawerApi.open();
}

async function handleDelete(row: SysTenantPackageListData) {
  await deleteSysTenantPackageApi({ packageIds: [row.packageId] });
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
      ids.push(row.packageId);
  }

  Modal.confirm({
    title: '提示',
    okType: 'danger',
    content: `确认删除选中的${ids.length}条记录吗？`,
    onOk: async () => {
      await deleteSysTenantPackageApi({ packageIds: ids.map(Number) });
      message.success("删除成功");
      await handleRefresh();
    },
  });
}

async function handleStatusChange(row: SysTenantPackageListData) {

  await statusSysTenantPackageApi({ packageId: row.packageId, status: row.status    }); 
  await message.success("操作成功")
  await handleRefresh();
}
</script>

<template>
  <Page auto-content-height>
    <Grid table-title="租户套餐列表">
      <template #toolbar-tools>
        
        <Button class="mr-2 flex items-center " type="primary" :icon="h(MdiPlus)" @click="handleAdd" v-access:code="'cpm:system:tenantPackage:add'">新增</Button>
        <Button class="mr-2 flex items-center" type="primary" :disabled="!CheckboxChecked" :icon="h(MdiDelete)" @click="handleMultiDelete" v-access:code="'cpm:system:tenantPackage:remove'">删除</Button>
      </template>
      <template #status="{ row }">
        <Switch
          v-model:checked="row.status" :checkedValue="'0'" :unCheckedValue="'1'" :disabled="!hasAccessByCodes(['cpm:system:tenantPackage:edit'])"
          @change="handleStatusChange(row)"
        />
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handlePreview(row)" v-access:code="'cpm:system:tenantPackage:query'">查看</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleEdit(row)" v-access:code="'cpm:system:tenantPackage:edit'"  >修改</Button>
          <AccessControl :codes="['cpm:system:tenantPackage:remove']" type="code">
            <Popconfirm title="确定删除吗？"  :get-popup-container="getVxePopupContainer" placement="left"  @confirm="handleDelete(row)" >  
              <Button class="mr-2 border-none p-0" :block="false" type="link"  danger >删除</Button>
            </Popconfirm>
          </AccessControl>
        </div>
      </template>
    </Grid>
    <ViewDrawer />
    <EditDrawer @reload="handleRefresh" />
  </Page>
</template>
