<script lang="ts" setup>
import { h, ref } from 'vue';
import type { DeepPartial } from '@vben/types';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions, VxeGridListeners } from '#/adapter/vxe-table';
import type { SysNotice } from '#/api/system/notice';
import { getVxePopupContainer } from '@vben/utils';
import { Page, useVbenDrawer } from '@vben/common-ui';
import { DictEnum } from '@vben/constants';
import { getPopupContainer } from '@vben/utils';
import { getDictOptions } from '#/utils/dict';
import { Button, message, Tag, Popconfirm, Modal } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysNoticeListApi, deleteSysNoticeApi } from '#/api/system/notice';
import { renderDict } from '#/utils/render';

import {
  MdiPlus,
  MdiDelete,
} from '@vben/icons';
import noticeDrawer from './notice-drawer.vue';
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
      fieldName: 'noticeTitle',
      label: '标题',
    },
    {
    component: 'Select',
    componentProps: {
      getPopupContainer,
      options: getDictOptions(DictEnum.SYS_NOTICE_TYPE),
    },
    fieldName: 'noticeType',
    label: '公告类型',
  },
    {
      component: 'Input',
      fieldName: 'createdBy',
      label: '创建者',
    },


    {
      component: 'RangePicker',
      componentProps: {
        format: "YYYY-MM-DD",
        valueFormat: "YYYY-MM-DD",
      },
      // defaultValue: [dayjs().subtract(7, 'days'), dayjs()],
      fieldName: 'createdAt',
      label: '创建时间',
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
    labelField: 'noticeId',
  },
  columns: [
    { align: 'left', title: 'ID', type: 'checkbox', width: 80 },
    { field: 'noticeTitle', title: '标题' },
    {
    title: '公告类型',
    field: 'noticeType',
    width: 120,
    slots: {
      default: ({ row }) => {
        return renderDict(row.noticeType, DictEnum.SYS_NOTICE_TYPE);
      },
    },
  },
    { field: 'status', title: '状态', slots: { default: 'status' } },
    { field: 'createdByUser.userName', title: '创建者' },
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
        return await getSysNoticeListApi({
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


const [NoticeDrawer, noticeDrawerApi] = useVbenDrawer({
  connectedComponent: noticeDrawer,
});


function handleView(row: SysNotice) {
  const { noticeId } = row;
  noticeDrawerApi.setData({ id: noticeId, update: false, view: true });
  noticeDrawerApi.open();
}

function handleAdd() {
  noticeDrawerApi.setData({ update: false, view: false });
  noticeDrawerApi.open();
}

function handleEdit(row: SysNotice) {
  noticeDrawerApi.setData({ id: row.noticeId, update: true, view: false });
  noticeDrawerApi.open();
}

async function handleDelete(row: SysNotice) {
  await deleteSysNoticeApi({ noticeIds: [row.noticeId] });
  message.success("删除成功");
  await handleRefresh();
}
async function handleRefresh() {
  await gridApi.query();
}


function handleMultiDelete() {
  const rows = gridApi.grid.getCheckboxRecords();
  const ids = rows.map((row: SysNotice) => row.noticeId);
  Modal.confirm({
    title: '提示',
    okType: 'danger',
    content: `确认删除选中的${ids.length}条记录吗？`,
    onOk: async () => {
      await deleteSysNoticeApi({ noticeIds: ids });
      message.success("删除成功");
      await handleRefresh();
    },
  });
}
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="'公告列表'">
      <template #toolbar-tools>

        <Button class="mr-2 flex items-center " type="primary" :icon="h(MdiPlus)" @click="handleAdd" v-access:code="'cpm:system:notice:add'">新增</Button>
        <Button class="mr-2 flex items-center" type="primary" :disabled="!CheckboxChecked" :icon="h(MdiDelete)"
          @click="handleMultiDelete" v-access:code="'cpm:system:notice:remove'">删除</Button>
      </template>
      <template #status="{ row }">
        <Tag :color="row.status == '0' ? 'green' : 'red'">{{ row.status == '0' ? '正常' : '关闭' }}</Tag>
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleView(row)" v-access:code="'cpm:system:notice:query'">查看</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleEdit(row)" v-access:code="'cpm:system:notice:edit'">修改</Button>
          <Popconfirm :get-popup-container="getVxePopupContainer" placement="left" title="确定删除吗？"
            @confirm="handleDelete(row)" v-access:code="'cpm:system:notice:remove'"><Button class="mr-2 border-none p-0" :block="false" type="link"
              danger v-access:code="'cpm:system:notice:remove'">删除</Button></Popconfirm>
        </div>
      </template>

    </Grid>
    <NoticeDrawer @reload="handleRefresh" />
  </Page>
</template>
