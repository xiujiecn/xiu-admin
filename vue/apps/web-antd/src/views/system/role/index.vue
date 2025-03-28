<!--
 * @description 角色管理页面
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script lang="ts" setup>
import { h, ref } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SysRoleListData } from '#/api/system/role';
import { Page } from '@vben/common-ui';

import { Button, message, Switch,Tag, Modal, Popconfirm } from 'ant-design-vue';
import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysRoleListApi, deleteSysRoleApi } from '#/api/system/role'; 
import { getVxePopupContainer } from '@vben/utils';
import { useVbenDrawer,useVbenModal } from '@vben/common-ui';

import roleDrawer from './role-drawer.vue';
import roleDataScopeModal from './role-data-scope-modal.vue';
import { authScopeOptions } from './model';

import { AccessControl, useAccess } from '@vben/access';
const { hasAccessByCodes } = useAccess();

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

function getDataScopeOptionLabel(dataScope: string) {
  const found = authScopeOptions.find((item) => item.value === dataScope);
  if (!found) {
    return dataScope;
  }
  return found.label;
}

function getDataScopeOptionColor(dataScope: string) {
  const found = authScopeOptions.find((item) => item.value === dataScope);
  if (!found) {
    return 'default';
  }
  return found.color;
}


const formOptions: VbenFormProps = {
  // 默认展开
  collapsed: false,
  fieldMappingTime: [['date', ['start', 'end']]],
  schema: [
    {
      component: 'Input',
      fieldName: 'roleName',
      label: '角色名称',
    },
    {
      component: 'Input',
      fieldName: 'roleKey',
      label: '角色权限字符串',
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
};

const gridOptions: VxeTableGridOptions<RowType> = {
  checkboxConfig: {
    highlight: true,
    labelField: 'roleId',
  },
  columns: [
    { align: 'left', title: 'ID', type: 'checkbox', width: 80 },
    { field: 'roleName', title: '角色名称' },
    { field: 'roleKey', title: '角色权限字符串' },
    { field: 'dataScope', title: '数据范围', slots: { default: 'dataScope' }, minWidth: 90 },
    { field: 'roleSort', title: '排序',width: 60 },
    {
      field: 'status',
      slots: { default: 'status' },
      title: '状态',
      width: 100,
    },
    { field: 'createdAt', formatter: 'formatDateTime', title: '创建时间' },
    { title: '操作', width: 190, slots: { default: 'action' } }
  ],
  exportConfig: {},
  height: 'auto',
  keepSource: true,
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        // message.success(`Query params: ${JSON.stringify(formValues)}`);
        return await getSysRoleListApi({
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

const [Grid, tableApi] = useVbenVxeGrid({
  formOptions,
  gridOptions,
});

const [RoleDrawer, roleDrawerApi] = useVbenDrawer({
  connectedComponent: roleDrawer,
});

const [RoleDataScopeModal, roleDataScopeModalApi] = useVbenModal({
  connectedComponent: roleDataScopeModal,
});

function handleView(row: SysRoleListData) {
  const { roleId } = row;
  roleDrawerApi.setData({id: roleId, update:false,view:true});
  roleDrawerApi.open();
}

function handleAdd() {
  roleDrawerApi.setData({update:false, view:false});
  roleDrawerApi.open();
}

function handleEdit(row: SysRoleListData) {
  roleDrawerApi.setData({ id: row.roleId, update:true, view:false });
  roleDrawerApi.open();
}

async function handleDelete(row: SysRoleListData) {
  await deleteSysRoleApi({ roleId: row.roleId });
  message.success("删除成功");
  await tableApi.query();
}

const CheckboxChecked = ref(false);

function handleCheckboxChange() {
  CheckboxChecked.value = tableApi.grid.getCheckboxRecords().length > 0;
  console.log('vue/apps/web-antd/src/views/system/user/index.vue CheckboxChecked',CheckboxChecked.value);
}

function handleMultiDelete() {
  const rows = tableApi.grid.getCheckboxRecords();
  const ids = rows.map((row: SysRoleListData) => row.roleId);
  Modal.confirm({
    title: '提示',
    okType: 'danger',
    content: `确认删除选中的${ids.length}条记录吗？`,
    onOk: async () => {
      await deleteSysRoleApi({ roleIds: ids });
      message.success("删除成功");
      await tableApi.query();
    },
  });
}

function handleReload() {
  tableApi.query();
}

function handleDataScope(row: SysRoleListData) {
  roleDataScopeModalApi.setData({id: row.roleId, update:true, view:false});
  roleDataScopeModalApi.open();
}

</script>

<template>
  <Page auto-content-height>
    <Grid table-title="角色管理">
      <template #toolbar-tools>
        
          <Button class="mr-2 flex items-center " type="primary" :icon="h(MdiPlus)" @click="handleAdd" v-access:code="'cpm:system:role:add'">新增</Button>
        <Button class="mr-2 flex items-center" type="primary" :disabled="!CheckboxChecked" :icon="h(MdiDelete)" @click="handleMultiDelete" v-access:code="'cpm:system:role:remove'">删除</Button>
      </template>
      <template #open="{ row }">
        <Switch v-model:checked="row.status" :checkedValue="'0'" :unCheckedValue="'1'" />
      </template>
      <template #status="{ row }">
        <Tag :color="row.status == '0' ? 'green' : 'red'">{{ row.status == '0' ? '正常' : '停用' }}</Tag>
      </template>
      <template #dataScope="{ row }">
        <Tag :color="getDataScopeOptionColor(row.dataScope)">{{ getDataScopeOptionLabel(row.dataScope) }}</Tag>
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleView(row)" v-access:code="'cpm:system:role:query'">查看</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link" v-if="row.roleId != 1" @click="handleEdit(row)" v-access:code="'cpm:system:role:edit'" >修改</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link" v-if="row.roleId != 1" @click="handleDataScope(row)" v-access:code="'cpm:system:role:edit'">数据权限</Button>
          <AccessControl :codes="['cpm:system:role:remove']" type="code">
            <Popconfirm :get-popup-container="getVxePopupContainer" placement="left" title="确定删除吗？" @confirm="handleDelete(row)" v-if="row.roleId != 1" >
              <Button class="mr-2 border-none p-0" :block="false" type="link"  danger >删除</Button>
            </Popconfirm>
          </AccessControl>
        </div>
      </template>
    </Grid>
    <RoleDrawer @reload="handleReload" />
    <RoleDataScopeModal @reload="handleReload" />
  </Page>
</template>
