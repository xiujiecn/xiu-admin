<script lang="ts" setup>
import { h } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SysMenuListData } from '#/api/system';

import { Page,useVbenDrawer } from '@vben/common-ui';

import { Button, message, Switch,Tag,Popconfirm  } from 'ant-design-vue';
import dayjs from 'dayjs';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysMenuListApi,deleteSysMenuApi } from '#/api/system'; 
import { Icon } from '@iconify/vue';
import {
  MdiPlus,
  MdiEdit,
  MdiDelete,
} from '@vben/icons';

import menuDrawer from './menu-drawer.vue';

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
        return await getSysMenuListApi({
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

const [MenuDrawer, drawerApi] = useVbenDrawer({
  connectedComponent: menuDrawer,
});

function handleAdd() {
  drawerApi.setData({update:false,view:false});
  drawerApi.open();
}

function handleView(row: SysMenuListData) {
  const { menuId } = row;
  drawerApi.setData({id: menuId, update:false,view:true});
  drawerApi.open();
}
function handleSubAdd(row: SysMenuListData) {
  const { menuId } = row;
  drawerApi.setData({ id: menuId, update: false,view:false });
  drawerApi.open();
}

async function handleEdit(record: SysMenuListData) {
  drawerApi.setData({ id: record.menuId, update: true,view:false });
  drawerApi.open();
}

async function handleDelete(row: SysMenuListData) {
  await deleteSysMenuApi({ menuId: row.menuId });
  await gridApi.query();
}

</script>

<template>
  <Page auto-content-height>
    <Grid table-title="菜单列表">
      <template #toolbar-tools>
        
        <Button class="mr-2 flex items-center " type="primary" :icon="h(MdiPlus)" @click="handleAdd">新增</Button>
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
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleView(row)">查看</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleEdit(row)">修改</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleSubAdd(row)">新增</Button>
          <Popconfirm placement="left" title="确定删除吗？" @confirm="handleDelete(row)">
            <Button class="mr-2 border-none p-0" :block="false" type="link"  danger>删除</Button>
          </Popconfirm>
        </div>
      </template>
    </Grid>
    <MenuDrawer @reload="gridApi.query()" />
  </Page>
</template>
