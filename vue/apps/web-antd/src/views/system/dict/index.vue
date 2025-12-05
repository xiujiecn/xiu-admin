<!--
 * @description 字典类型管理页面
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script lang="ts" setup>
import { h, ref } from 'vue';
import type {  DeepPartial } from '@vben/types';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions,VxeGridListeners } from '#/adapter/vxe-table';
import type { SysDictTypeListModel } from '#/api/system/dict';
import { getVxePopupContainer } from '@vben/utils';
import { Page,useVbenDrawer } from '@vben/common-ui';
import { AccessControl, useAccess } from '@vben/access';
const { hasAccessByCodes } = useAccess();

import { Button, message, Popconfirm,Tag, Modal } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysDictTypeListApi, deleteSysDictTypeApi } from '#/api/system/dict';
import { useRoute, useRouter } from 'vue-router';
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
      component: 'Select',
      fieldName: 'isSys',
      label: '系统内置',
      componentProps: {
        options: [
          {
            label: '否',
            value: '1',
          },
          {
            label: '是',
            value: '0',
          },
        ],
      },
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
  },
  columns: [
    { type: 'checkbox', width: 40 },
    { field: 'dictId', title: 'ID' },
    { field: 'dictName', title: '字典名称' },
    { field: 'dictType', title: '字典类型' ,slots: { default: 'type' }, },
    { field: 'isSys', title: '系统内置', slots: { default: 'isSys' } },
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
       <template #isSys="{ row }">
        <Tag :color="row.isSys == '0' ? 'green' : 'red'">{{ row.isSys == '0' ? '是' : '否' }}</Tag>
      </template>
      <template #status="{ row }">
        <Tag :color="row.status == '0' ? 'green' : 'red'">{{ row.status == '0' ? '正常' : '停用' }}</Tag>
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleView(row)" v-access:code="'cpm:system:dict:query'">查看</Button>

          <!-- 修改按钮：系统内置时只有超级管理员可操作 -->
          <AccessControl
            :codes="row.isSys === '0' ? ['cpr:superadmin'] : ['cpm:system:dict:edit']"
            type="code"
          >
            <template #default="{ hasPermission }">
              <Button
                class="mr-2 border-none p-0"
                :block="false"
                type="link"
                :disabled="!hasPermission"
                @click="handleEdit(row)"
              >
                修改
              </Button>
            </template>
          </AccessControl>

          <!-- 删除按钮：系统内置时只有超级管理员可操作 -->
          <AccessControl
            :codes="row.isSys === '0' ? ['cpr:superadmin'] : ['cpm:system:dict:remove']"
            type="code"
          >
            <template #default="{ hasPermission }">
              <!-- 有权限时显示可点击的删除按钮 -->
              <Popconfirm
                v-if="hasPermission"
                :get-popup-container="getVxePopupContainer"
                placement="left"
                title="确定删除吗？"
                @confirm="handleDelete(row)"
              >
                <Button
                  class="mr-2 border-none p-0"
                  :block="false"
                  type="link"
                  danger
                >
                  删除
                </Button>
              </Popconfirm>
              <!-- 无权限时显示禁用的删除按钮，不包含点击事件 -->
              <Button
                v-else
                class="mr-2 border-none p-0"
                :block="false"
                type="link"
                danger
                disabled
              >
                删除
              </Button>
            </template>
          </AccessControl>
        </div>
      </template>
    </Grid>
    <DictDrawer @reload="handleRefresh"/>
  </Page>
</template>
