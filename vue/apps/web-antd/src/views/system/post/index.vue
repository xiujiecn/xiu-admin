<script lang="ts" setup>
import { h, ref } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SysPostListData } from '#/api';

import { Page, useVbenDrawer } from '@vben/common-ui';
import { getVxePopupContainer } from '@vben/utils';

import { Button, message, Switch,Tag, Modal, Popconfirm } from 'ant-design-vue';
import dayjs from 'dayjs';
import { commonDownloadExcel } from '#/utils/file/download';
import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysPostListApi, deleteSysPostApi, getSysPostExportApi } from '#/api'; 
import DeptTree from '#/components/dept/dept-tree.vue';
import postDrawer from './post-drawer.vue';

import {
  MdiPlus,
  MdiEdit,
  MdiDelete,
  MdiExport,
} from '@vben/icons';

interface RowType {
  category: string;
  color: string;
  id: string;
  price: string;
  productName: string;
  releaseDate: string;
}

const selectDeptId = ref<string[]>([]);

const formOptions: VbenFormProps = {
  // 默认展开
  collapsed: false,
  fieldMappingTime: [['date', ['start', 'end']]],
  schema: [
    {
      component: 'Input',
      fieldName: 'postCode',
      label: '岗位编码',
    },
    {
      component: 'Input',
      fieldName: 'postName',
      label: '岗位名称',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: [
          {
            label: '正常',
            value: '0',
          },
          {
            label: '停用',
            value: '1',
          },
        ],
        placeholder: '请选择',
      },
      fieldName: 'status',
      label: '状态',
    },
    {
      component: 'RangePicker',
      // defaultValue: [dayjs().subtract(7, 'days'), dayjs()],
      fieldName: 'createdAt',
      label: '创建时间',
      componentProps: {
        format: 'YYYY-MM-DD',
        valueFormat:"YYYY-MM-DD",
      },
    },
  ],
  // 控制表单是否显示折叠按钮
  showCollapseButton: true,
  // 是否在字段值改变时提交表单
  submitOnChange: true,
  // 按下回车时是否提交表单
  submitOnEnter: false,
  handleReset: async () => {
    selectDeptId.value = [];
    await tableApi.formApi.resetForm();
    await tableApi.reload();
  },
};

const gridOptions: VxeTableGridOptions<RowType> = {
  checkboxConfig: {
    highlight: true,
    labelField: 'postId',
  },
  columns: [
    { align: 'left', title: 'ID', type: 'checkbox', width: 80 },
    { field: 'postCode', title: '岗位编码' },
    { field: 'postName', title: '岗位名称' },
    { field: 'deptInfo.deptName', title: '部门' },
    { field: 'postCategory', title: '岗位类别' },
    {
      field: 'status',
      slots: { default: 'status' },
      title: '状态',
      width: 100,
    },
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
        let deptId:number = 0;
        if(selectDeptId.value.length > 0) {
          deptId = Number(selectDeptId.value[0]);
        }
        return await getSysPostListApi({
          page: page.currentPage,
          pageSize: page.pageSize,
          deptId: deptId,
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

const [Grid, tableApi ] = useVbenVxeGrid({
  formOptions,
  gridOptions,
});


const [PostDrawer, postDrawerApi] = useVbenDrawer({
  connectedComponent: postDrawer,
});

function handleView(row: SysPostListData) {
  const { postId } = row;
  postDrawerApi.setData({id: postId,  update:false,view:true});
  postDrawerApi.open();
}

function handleAdd() {
  const deptId = selectDeptId.value?.[0] || '';
  postDrawerApi.setData({deptId: deptId, update:false, view:false});
  postDrawerApi.open();
}

function handleEdit(row: SysPostListData) {
  postDrawerApi.setData({ id: row.postId, update:true, view:false });
  postDrawerApi.open();
}

async function handleDelete(row: SysPostListData) {
  await deleteSysPostApi({ postId: row.postId });
  message.success("删除成功");
  await tableApi.query();
}

async function handleExport() {
  // 导出
  const formValues = tableApi.formApi.form.values;
  // const res = await getSysPostExportApi({
  //   ...formValues,
  //   page: 1,
  //   pageSize: 2000,
  // });
  commonDownloadExcel(getSysPostExportApi, '岗位列表',{
    ...formValues,
    page: 1,
    pageSize: 2000,
  });
  message.success("导出成功");
}


function handleMultiDelete() {
  const rows = tableApi.grid.getCheckboxRecords();
  const ids = rows.map((row: SysPostListData) => row.postId);
  Modal.confirm({
    title: '提示',
    okType: 'danger',
    content: `确认删除选中的${ids.length}条记录吗？`,
    onOk: async () => {
      await deleteSysPostApi({ postIds: ids });
      message.success("删除成功");
      await tableApi.query();
    },
  });
}

function handleReload() {
  tableApi.query();
}

function handleStatusChange(row: SysPostListData) {
  console.log(row);
}

</script>

<template>
  <Page auto-content-height>
    <div class="flex h-full gap-[8px]">
    <DeptTree class="w-[240px]" 
      @select="()=> tableApi.reload()" 
      @reload="()=> tableApi.reload()"
      v-model:select-dept-id="selectDeptId" />
    <Grid table-title="岗位列表">
      <template #toolbar-tools>
        
        <Button class="mr-2 flex items-center " type="primary" :icon="h(MdiPlus)" @click="handleAdd">新增</Button>
        <Button class="mr-2 flex items-center" type="primary" disabled :icon="h(MdiDelete)" @click="handleMultiDelete">删除</Button>
        <Button class="mr-2 flex items-center "  :icon="h(MdiExport)" @click="handleExport">导出</Button>
      </template>
      <template #open="{ row }">
        <Switch v-model:checked="row.status" :checkedValue="'0'" :unCheckedValue="'1'"  @change="handleStatusChange(row)" />
      </template>
      <template #status="{ row }">
        <Tag :color="row.status == '0' ? 'green' : 'red'">{{ row.status == '0' ? '正常' : '停用' }}</Tag>
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleView(row)">查看</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleEdit(row)">修改</Button>
          <Popconfirm
            :get-popup-container="getVxePopupContainer"
            placement="left"
            title="确认删除？"
            @confirm="handleDelete(row)"
          >
            <Button class="mr-2 border-none p-0" :block="false" type="link" v-if="row.userId != 1" danger>删除</Button>
          </Popconfirm>
        </div>
      </template>
    </Grid>
    </div>
    <PostDrawer @reload="handleReload" />
  </Page>
</template>
