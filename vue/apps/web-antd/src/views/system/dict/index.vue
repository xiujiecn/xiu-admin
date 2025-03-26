<script lang="ts" setup>
import { h, ref } from 'vue';
import type {  DeepPartial } from '@vben/types';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions,VxeGridListeners } from '#/adapter/vxe-table';
import type { SysDictTypeListModel } from '#/api/system/dict';
import { getVxePopupContainer } from '@vben/utils';
import { Page,useVbenDrawer } from '@vben/common-ui';

import { Button, message, Popconfirm,Tag, Modal } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysDictTypeListApi, deleteSysDictTypeApi } from '#/api/system/dict'; 
import { useRouter } from 'vue-router';
import {
  MdiPlus,
  MdiDelete,
} from '@vben/icons';

import dictDrawer from './dict-drawer.vue';

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
      fieldName: 'dictName',
      label: '字典名称',
    },
    {
      component: 'Input',
      fieldName: 'dictType',
      label: '字典类型',
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
};

const gridOptions: VxeTableGridOptions<RowType> = {
  checkboxConfig: {
    highlight: true,
    labelField: 'dictId',
  },
  columns: [
    { align: 'left', title: 'ID', type: 'checkbox', width: 80 },
    { field: 'dictName', title: '字典名称' },
    { field: 'dictType', title: '字典类型' ,slots: { default: 'type' }, },
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
        return await getSysDictTypeListApi({
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


const [Grid,gridApi] = useVbenVxeGrid({
  formOptions,
  gridOptions,
  gridEvents,
});
const router = useRouter();
const handleClickDictType = (dictId: number) => {
  router.push(`/system/dict-data/${dictId}`);
}

const [DictDrawer, dictDrawerApi] = useVbenDrawer({
  connectedComponent: dictDrawer,
});


function handleView(row: SysDictTypeListModel) {
  const { dictId } = row;
  dictDrawerApi.setData({id: dictId, update:false,view:true});
  dictDrawerApi.open();
}

function handleAdd() {
  dictDrawerApi.setData({update:false, view:false});
  dictDrawerApi.open();
}

function handleEdit(row: SysDictTypeListModel) {
  dictDrawerApi.setData({ id: row.dictId, update:true, view:false });
  dictDrawerApi.open();
}

async function handleDelete(row: SysDictTypeListModel) {
  await deleteSysDictTypeApi({ dictIds: [row.dictId] });
  message.success("删除成功");
  await handleRefresh();
}
async function handleRefresh() {
  await gridApi.query();
}


function handleMultiDelete() {
  const rows = gridApi.grid.getCheckboxRecords();
  const ids = rows.map((row: SysDictTypeListModel) => row.dictId);
  Modal.confirm({
    title: '提示',
    okType: 'danger',
    content: `确认删除选中的${ids.length}条记录吗？`,
    onOk: async () => {
      await deleteSysDictTypeApi({ dictIds: ids });
      message.success("删除成功");
      await handleRefresh();
    },
  });
}

</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="'字典类型列表'">
      <template #toolbar-tools>
        <Button class="mr-2 flex items-center " type="primary" :icon="h(MdiPlus)" @click="handleAdd" v-access:code="'cpm:system:dict:add'">新增</Button>
        <Button class="mr-2 flex items-center" type="primary" :disabled="!CheckboxChecked" :icon="h(MdiDelete)" @click="handleMultiDelete" v-access:code="'cpm:system:dict:remove'">删除</Button>
      </template>
      <template #type="{ row }">
        <Button type="link" :block="false" @click="handleClickDictType(row.dictId)" >{{ row.dictType }}</Button>
      </template>
      <template #status="{ row }">
        <Tag :color="row.status == '0' ? 'green' : 'red'">{{ row.status == '0' ? '正常' : '停用' }}</Tag>
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleView(row)" v-access:code="'cpm:system:dict:query'">查看</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleEdit(row)" v-access:code="'cpm:system:dict:edit'">修改</Button>
          <Popconfirm :get-popup-container="getVxePopupContainer" placement="left" title="确定删除吗？" @confirm="handleDelete(row)" v-access:code="'cpm:system:dict:remove'">
            <Button class="mr-2 border-none p-0" :block="false" type="link"  danger v-access:code="'cpm:system:dict:remove'">删除</Button>
          </Popconfirm>
        </div>
      </template>
    </Grid>
    <DictDrawer @reload="handleRefresh"/>
  </Page>
</template>
