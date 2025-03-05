<script lang="ts" setup>
import { h, ref } from 'vue';
import type {  DeepPartial } from '@vben/types';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions,VxeGridListeners } from '#/adapter/vxe-table';
import type { SysUserListData } from '#/api/system/user';
import { deleteSysUser } from '#/api/system/user';

import { Page, useVbenDrawer, useVbenModal } from '@vben/common-ui';

import { Button, message, Switch, Modal, Popconfirm } from 'ant-design-vue';
import dayjs from 'dayjs';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysUserListApi, updateSysUser } from '#/api/system'; 
import DeptTree from '#/components/dept/dept-tree.vue';

import {
  MdiPlus,
  MdiEdit,
  MdiDelete,
} from '@vben/icons';

import userDrawer from './user-drawer.vue';
import userResetPwdModal from './user-reset-pwd-modal.vue';

interface RowType {
  category: string;
  color: string;
  id: string;
  price: string;
  productName: string;
  releaseDate: string;
}

const selectDeptId = ref<string[]>([]);

const gridEvents: DeepPartial<VxeGridListeners> = {
  checkboxChange: handleCheckboxChange,
  checkboxAll: handleCheckboxChange,
}

const formOptions: VbenFormProps = {
  // 默认展开
  collapsed: false,
  fieldMappingTime: [['date', ['start', 'end']]],
  schema: [
    {
      component: 'Input',
      // defaultValue: '1',
      fieldName: 'userName',
      label: '用户名称',
    },
    {
      component: 'Input',
      fieldName: 'nickName',
      label: '用户昵称',
    },
    {
      component: 'Input',
      fieldName: 'email',
      label: '邮箱',
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
    labelField: 'userId',
  },
  columns: [
    { align: 'left', title: 'ID', type: 'checkbox', width: 80 },
    { field: 'userName', title: '用户名称' },
    { field: 'nickName', title: '用户昵称' },
    { field: 'deptInfo.deptName', title: '部门' },
    { field: 'email', title: '邮箱' },
    {
      field: 'status',
      slots: { default: 'open' },
      title: '状态',
      width: 100,
    },
    { field: 'createdAt', formatter: 'formatDateTime', title: '创建时间' },
    { title: '操作', width: 180, slots: { default: 'action' } }
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

        return await getSysUserListApi({
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
  gridEvents,
});
const [UserDrawer, userDrawerApi] = useVbenDrawer({
  connectedComponent: userDrawer,
});

function handleView(row: SysUserListData) {
  const { userId } = row;
  userDrawerApi.setData({id: userId, update:false,view:true});
  userDrawerApi.open();
}

function handleAdd() {
  userDrawerApi.setData({update:false, view:false});
  userDrawerApi.open();
}

function handleEdit(row: SysUserListData) {
  userDrawerApi.setData({ id: row.userId, update:true, view:false });
  userDrawerApi.open();
}

async function handleDelete(row: SysUserListData) {
  await deleteSysUser({ userId: row.userId });
  message.success("删除成功");
  await tableApi.query();
}

const CheckboxChecked = ref(false);

function handleCheckboxChange() {
  CheckboxChecked.value = tableApi.grid.getCheckboxRecords().length > 0;
}

function handleMultiDelete() {
  const rows = tableApi.grid.getCheckboxRecords();
  const ids = rows.map((row: SysUserListData) => row.userId);
  Modal.confirm({
    title: '提示',
    okType: 'danger',
    content: `确认删除选中的${ids.length}条记录吗？`,
    onOk: async () => {
      await deleteSysUser({ userIds: ids });
      message.success("删除成功");
      await tableApi.query();
    },
  });
}

const [UserResetPwdModal, userResetPwdModalApi] = useVbenModal({
  connectedComponent: userResetPwdModal,
});

function handleResetPassword(row: SysUserListData) {
  userResetPwdModalApi.setData({ record: row });
  userResetPwdModalApi.open();
}

async function handleStatusChange(row: SysUserListData) {
  await updateSysUser({ userId: row.userId, status: row.status });
  message.success('状态更新成功');
  await tableApi.query();
}

</script>

<template>
  <Page auto-content-height>
    <div class="flex h-full gap-[8px]">
    <DeptTree class="w-[240px]" 
      @select="()=> tableApi.reload()" 
      @reload="()=> tableApi.reload()"
      v-model:select-dept-id="selectDeptId" />

    <Grid class="flex-1" table-title="用户列表"    >
      <template #toolbar-tools>
        
        <Button class="mr-2 flex items-center " type="primary" :icon="h(MdiPlus)" @click="handleAdd">新增</Button>
        <Button class="mr-2 flex items-center" type="primary" :disabled="!CheckboxChecked" :icon="h(MdiDelete)" @click="handleMultiDelete">删除</Button>
      </template>
      <template #open="{ row }">
        <Switch v-model:checked="row.status" :checkedValue="'0'" :unCheckedValue="'1'" @change="handleStatusChange(row)" />
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleView(row)">查看</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleEdit(row)">修改</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleResetPassword(row)">重置密码</Button>
          <Popconfirm placement="left" title="确定删除吗？" @confirm="handleDelete(row)">
            <Button class="mr-2 border-none p-0" :block="false" type="link" v-if="row.userId != 1" danger >删除</Button>
          </Popconfirm>
        </div>
      </template>
    </Grid>
  </div>
  <UserDrawer @reload="tableApi.query()" />
  <UserResetPwdModal />
  </Page>
</template>
