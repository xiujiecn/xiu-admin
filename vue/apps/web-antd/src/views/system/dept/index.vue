<!--
 * @description 部门管理页面
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script lang="ts" setup>
import { h } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SysDeptListData } from '#/api/system/dept';
import { AccessControl, useAccess } from '@vben/access';
const { hasAccessByCodes } = useAccess();

import { Page, useVbenDrawer } from '@vben/common-ui';
import { getVxePopupContainer } from '@vben/utils';
import { Button, message, Popconfirm,Tag  } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysDeptListApi, deleteSysDeptApi } from '#/api/system/dept'; 

import deptDrawer from './dept-drawer.vue';

import {
  MdiPlus,
} from '@vben/icons';

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
  collapsed: true,
  fieldMappingTime: [['date', ['start', 'end']]],
  schema: [
    {
      component: 'Input',
      fieldName: 'deptName',
      label: '部门名称',
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
  ],
  // 控制表单是否显示折叠按钮
  showCollapseButton: true,
  // 是否在字段值改变时提交表单
  submitOnChange: true,
  // 按下回车时是否提交表单
  submitOnEnter: false,
};

const gridOptions: VxeTableGridOptions<RowType> = {
  align: 'center',
  columns: [
    { field: 'deptName', title: '部门名称' , treeNode: true, minWidth: 240,  align: 'left', },
    { field: 'deptCategory', title: '部门编码', minWidth: 100 },
    {
      field: 'status',
      slots: { default: 'status' },
      title: '状态',
      width: 69,
      align: 'center',
    },
    { field: 'createdAt', formatter: 'formatDateTime', title: '创建时间' , width: 160},
    { title: '操作', width: 120, slots: { default: 'action' } }
  ],
  exportConfig: {},
  height: 'auto',
  keepSource: true,
  pagerConfig: {
    enabled: false,
  },
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        return await getSysDeptListApi({
          page: page.currentPage,
          pageSize: page.pageSize,
          ...formValues,
        });
      },
      querySuccess: () => {
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
  treeConfig: {
    parentField: 'parentId',
    rowField: 'deptId',
    transform: true,
    expandAll: true,
  },
};

const [Grid,gridApi] = useVbenVxeGrid({
  formOptions,
  gridOptions,
});
const expandAll = () => {
  gridApi.grid?.setAllTreeExpand(true);
};

const collapseAll = () => {
  gridApi.grid?.setAllTreeExpand(false);
};



const [DeptDrawer, deptDrawerApi] = useVbenDrawer({
  connectedComponent: deptDrawer,
});


function handleView(row: SysDeptListData) {
  const { deptId } = row;
  deptDrawerApi.setData({id: deptId, update:false,view:true});
  deptDrawerApi.open();
}

function handleAdd() {
  deptDrawerApi.setData({update:false, view:false});
  deptDrawerApi.open();
}

function handleEdit(row: SysDeptListData) {
  deptDrawerApi.setData({ id: row.deptId, update:true, view:false });
  deptDrawerApi.open();
}

async function handleDelete(row: SysDeptListData) {
  await deleteSysDeptApi({ deptId: row.deptId });
  message.success("删除成功");
  await handleRefresh();
}
async function handleRefresh() {
  await gridApi.query();
  expandAll();
}
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="'部门管理'">
      <template #toolbar-tools>
        
        <Button class="mr-2 flex items-center"  @click="expandAll">展开</Button>
        <Button class="mr-2 flex items-center"  @click="collapseAll">折叠</Button>
        <Button class="mr-2 flex items-center " type="primary" :icon="h(MdiPlus)" @click="handleAdd" v-access:code="'cpm:system:dept:add'">新增</Button>
      </template>
      <template #status="{ row }">
        <Tag :color="row.status == '0' ? 'green' : 'red'">{{ row.status == '0' ? '正常' : '停用' }}</Tag>
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleView(row)" v-access:code="'cpm:system:dept:query'">查看</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleEdit(row)" v-access:code="'cpm:system:dept:edit'">修改</Button>
          <AccessControl :codes="['cpm:system:dept:remove']" type="code">
            <Popconfirm :get-popup-container="getVxePopupContainer" placement="left" title="确定删除吗？" @confirm="handleDelete(row)" >
              <Button class="mr-2 border-none p-0" :block="false" type="link" danger >删除</Button>
            </Popconfirm>
          </AccessControl>
        </div>
      </template>
    </Grid>
    <DeptDrawer @reload="handleRefresh" />
  </Page>
</template>
