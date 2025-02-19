<script lang="ts" setup>
import { h } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions } from '#/adapter/vxe-table';

import { Page } from '@vben/common-ui';

import { Button, message, Switch,Tag  } from 'ant-design-vue';
import dayjs from 'dayjs';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getMenuListApi } from '#/api'; 
import { Icon } from '@iconify/vue';
import {
  MdiPlus,
  MdiEdit,
  MdiDelete,
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
  collapsed: false,
  fieldMappingTime: [['date', ['start', 'end']]],
  schema: [
    {
      component: 'Input',
      fieldName: 'menuName',
      label: '菜单名称',
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
    { field: 'menuName', title: '菜单名称' , treeNode: true, minWidth: 160,  align: 'left', },
    { field: 'icon', title: '图标',width: 60,slots: { default: 'icon' }, },
    { field: 'orderNum', title: '排序' ,width: 60,},
    { field: 'perms', title: '权限标识' },
    { field: 'path', title: '路由地址' },
    { field: 'component', title: '组件路径' },
    {
      field: 'status',
      slots: { default: 'status' },
      title: '状态',
      width: 69,
      align: 'center',
    },
    { field: 'createdAt', formatter: 'formatDateTime', title: '创建时间' , width: 160},
    { title: '操作', width: 160, slots: { default: 'action' } }
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
        message.success(`Query params: ${JSON.stringify(formValues)}`);
        return await getMenuListApi({
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
  treeConfig: {
    parentField: 'parentId',
    rowField: 'menuId',
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
</script>

<template>
  <Page auto-content-height>
    <Grid>
      <template #toolbar-actions>
        
        <Button class="mr-2 flex items-center " type="primary" :icon="h(MdiPlus)">新增</Button>
        <Button class="mr-2 flex items-center"  @click="expandAll">展开</Button>
        <Button class="mr-2 flex items-center"  @click="collapseAll">折叠</Button>
      </template>
      <template #icon="{ row }">
        <Icon :icon="row.icon" />
      </template>
      <template #status="{ row }">
        <Tag :color="row.status == '0' ? 'green' : 'red'">{{ row.status == '0' ? '正常' : '停用' }}</Tag>
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link">查看</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link">修改</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link">新增</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link"  danger>删除</Button>
        </div>
      </template>
    </Grid>
  </Page>
</template>
