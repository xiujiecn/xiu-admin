<script lang="ts" setup>
  import { h, reactive, ref, computed, onMounted } from 'vue';
  import { Button, message,Tag, Modal, Popconfirm,Switch } from 'ant-design-vue';
  import type { VbenFormProps } from '#/adapter/form';
  import type { VxeTableGridOptions, VxeGridListeners } from '#/adapter/vxe-table';
  import type { DeepPartial } from '@vben/types';
  import { getVxePopupContainer } from '@vben/utils';
  import { Page, useVbenDrawer } from '@vben/common-ui';
  import { useVbenVxeGrid } from '#/adapter/vxe-table';
  import { AccessControl, useAccess } from '@vben/access';
  const { hasAccessByCodes } = useAccess();
  import { commonDownloadExcel } from '#/utils/file/download';
  import { List, Export, Delete, Status } from '#/api/gen/testCategory';
  import { MdiPlus, MdiExport, MdiDelete }  from '@vben/icons';
  import { columns, querySchema, type RowType } from './model';
  import editDrawer from './edit.vue';
  import viewDrawer from './view.vue';
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
  return await List({
  page: page.currentPage,
  pageSize: page.pageSize,
  ...formValues,
  });
  },
  },
  },
  toolbarConfig: {
  custom: true,
  export: false,
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
  function handlePreview(record: RowType) {
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
  function handleEdit(row: RowType) {
  editDrawerApi.setData({ id: row.id, update: true, view: false });
  editDrawerApi.open();
  }
  async function handleDelete(row: RowType) {
  await Delete({ id: [row.id] });
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
  ids.push(row.id);
  }
  if (ids.length === 0){
  message.error('请至少选择一项要删除的数据');
  return;
  }
  Modal.confirm({
  title: '提示',
  okType: 'danger',
  content: `确认删除选中的${ids.length}条记录吗？`,
  onOk: async () => {
  await Delete({ id: ids });
  message.success("删除成功");
  await handleRefresh();
  },
  });
  }
  async function handleExport() {
  const formValues = gridApi.formApi.form.values;
  await commonDownloadExcel(Export, '测试分类',{
  ...formValues,
  page: 1,
  pageSize: 2000,
  });
  message.success("导出成功");
  }
  async function handleStatusChange(row: RowType) {
  await Status({ id: row.id, status: row.status });
  await message.success("操作成功")
  await handleRefresh();
  }
</script>
<template>
  <Page auto-content-height>
    <Grid table-title="测试分类">
      <template #toolbar-tools>
        <Button class="mr-2 flex items-center " type="primary" :icon="h(MdiPlus)" @click="handleAdd" v-access:code="'cpm:gen:testCategory:edit'">
          新增
        </Button>
        <Button class="mr-2 flex items-center" type="primary" :disabled="!CheckboxChecked" :icon="h(MdiDelete)" @click="handleMultiDelete" v-access:code="'cpm:gen:testCategory:delete'">
          删除
        </Button>
        <Button class="mr-2 flex items-center" type="primary" :icon="h(MdiExport)" @click="handleExport" v-access:code="'cpm:gen:testCategory:export'">
          导出
        </Button>
      </template>
      <template #status="{ row }">
        <Switch
          v-model:checked="row.status" :checkedValue="'0'" :unCheckedValue="'1'"
          @change="handleStatusChange(row)"
           :disabled="!hasAccessByCodes(['cpm:gen:testCategory:status'])"
        />
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handlePreview(row)" v-access:code="'cpm:gen:testCategory:view'">
            查看
          </Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleEdit(row)" v-access:code="'cpm:gen:testCategory:edit'">
            修改
          </Button>
          <AccessControl :codes="['cpm:gen:testCategory:delete']" type="code">
            <Popconfirm title="确定删除吗？" :get-popup-container="getVxePopupContainer" placement="left"  @confirm="handleDelete(row)" >
              <Button class="mr-2 border-none p-0" :block="false" type="link"  danger >
                删除
              </Button>
            </Popconfirm>
          </AccessControl>
        </div>
      </template>
    </Grid>
    <EditDrawer @reload="handleRefresh" />
    <ViewDrawer />
  </Page>
</template>