<script lang="ts" setup>
import type { VbenFormProps } from '#/adapter/form';
import type { DeepPartial } from '@vben/types';
import type { VxeTableGridOptions, VxeGridListeners } from '#/adapter/vxe-table';
import type { SysJob } from '#/api/system/job'

import { h, ref } from 'vue';
import { getVxePopupContainer } from '@vben/utils';

import { Page, useVbenDrawer } from '@vben/common-ui';
import { getDictOptions } from '#/utils/dict';
import { DictEnum } from '@vben/constants';
import { Button, message, Tag, Modal, Popconfirm } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysJobListApi, deleteSysJobApi } from '#/api/system/job';
import viewDrawer from './view-drawer.vue';
import editDrawer from './edit-drawer.vue';
import {
  MdiDelete,
  MdiPlus,
} from '@vben/icons';

interface RowType {
  jobId: number;
  jobName: string;

  
}

const formOptions: VbenFormProps = {
  // 默认展开
  collapsed: false,
  fieldMappingTime: [['date', ['start', 'end']]],
  schema: [
    {
      component: 'Input',
      fieldName: 'jobName',
      label: '任务名称',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: getDictOptions(DictEnum.SYS_JOB_GROUP),
        placeholder: '请选择',
      },
      fieldName: 'jobGroup',
      label: '任务分组',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: getDictOptions(DictEnum.SYS_NORMAL_DISABLE),
        placeholder: '请选择',
      },
      fieldName: 'status',
      label: '任务状态',
    }
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
    labelField: 'jobId',
  },
  columns: [
    { align: 'left', title: 'ID', type: 'checkbox', width: 80 },
    { field: 'jobName', title: '任务名称' },
    { field: 'remark', title: '任务描述' },
    { field: 'jobGroup', title: '任务分组' },
    { field: 'invokeTarget', title: '任务方法名' },
    { field: 'cronExpression', title: 'corn执行表达式' },
    { field: 'status', title: '状态' },
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
        return await getSysJobListApi({
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

const [EditDrawer, drawerEditApi] = useVbenDrawer({
  connectedComponent: editDrawer,
});

function handlePreview(record: SysJob) {
  drawerApi.setData({ record });
  drawerApi.open();
}


async function handleDelete(row: SysJob) {
  await deleteSysJobApi({ jobIds: [row.jobId] });
  message.success("删除成功");
  await handleRefresh();
}

async function handleRefresh() {
  await gridApi.query();
}


function handleMultiDelete() {
  const rows = gridApi.grid.getCheckboxRecords();
  const ids = rows.map((row: SysJob) => row.jobId);
  Modal.confirm({
    title: '提示',
    okType: 'danger',
    content: `确认删除选中的${ids.length}条记录吗？`,
    onOk: async () => {
      await deleteSysJobApi({ jobIds: ids });
      message.success("删除成功");
      await handleRefresh();
    },
  });
}

function handleAdd() {
  drawerEditApi.setData({ record: {}, update: false, view: false });
  drawerEditApi.open();
}

function handleUpdate(row: SysJob) {
  drawerEditApi.setData({ jobId: row.jobId, update: true, view: false });
  drawerEditApi.open();
}

</script>

<template>
  <Page auto-content-height>
    <Grid table-title="定时任务列表">
      <template #toolbar-tools>
        <Button class="mr-2 flex items-center " type="primary" :icon="h(MdiPlus)" @click="handleAdd">新增</Button>
        <Button class="mr-2 flex items-center" type="primary" :disabled="!CheckboxChecked" :icon="h(MdiDelete)" @click="handleMultiDelete">删除</Button>
      </template>
      <template #status="{ row }">
        <Tag :color="row.status == '0' ? 'green' : 'red'">{{ row.status == '0' ? '正常' : '关闭' }}</Tag>
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handlePreview(row)">查看</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleUpdate(row)">编辑</Button>
          <Popconfirm :get-popup-container="getVxePopupContainer" placement="left" title="确定删除吗？"
            @confirm="handleDelete(row)">
            <Button class="mr-2 border-none p-0" :block="false" type="link" danger>删除</Button>
          </Popconfirm>
        </div>
      </template>
    </Grid>
    <ViewDrawer />
    <EditDrawer @reload="gridApi.query()" />
  </Page>
</template>
