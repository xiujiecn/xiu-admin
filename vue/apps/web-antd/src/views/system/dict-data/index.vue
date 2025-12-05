<!--
 * @description 字典数据管理页面
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script lang="ts" setup>
import { h,ref,onMounted } from 'vue';
import type {  DeepPartial } from '@vben/types';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions,VxeGridListeners } from '#/adapter/vxe-table';
import type { SysDictDataListModel } from '#/api';
import { getVxePopupContainer } from '@vben/utils';
import { Page,useVbenDrawer } from '@vben/common-ui';
import { Button, message,Tag, Modal,Popconfirm } from 'ant-design-vue';
import { AccessControl, useAccess } from '@vben/access';
const { hasAccessByCodes } = useAccess();

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysDictDataListApi,deleteSysDictDataApi } from '#/api';
import { useRouter } from 'vue-router';
import dictDataDrawer from './dict-data-drawer.vue';



const route = useRouter();
const dictId = parseInt(route.currentRoute.value.path.split('/').pop() || '0');
const dictName = ref('');
const dictType = ref('');
const dictIsSys = ref('1'); // 默认为非系统内置，避免在加载完成前误判权限

onMounted(() => {
})
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
      fieldName: 'dictName',
      defaultValue: dictName,
      label: '字典名称',
    },
    // 字典类型
    {
      component: 'Input',
      fieldName: 'dictType',
      defaultValue: dictType,
      label: '字典类型',
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
    labelField: 'dictCode',
  },
  columns: [
    { align: 'left', title: 'ID', type: 'checkbox', width: 80 },
    { field: 'dictLabel', title: '字典标签', slots: { default: 'label' } },
    { field: 'dictValue', title: '字典值' },
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
         const res = await getSysDictDataListApi({
          page: page.currentPage,
          pageSize: page.pageSize,
          dictId: dictId,
          ...formValues,
        })
        // console.log('res',res);
        return res;
      },
      querySuccess: ({ page, sort, sorts, filters, form, response }) => {
        dictName.value = response.type.dictName;
        dictType.value = response.type.dictType;
        dictIsSys.value = response.type.isSys || '1'; // 获取字典类型的系统内置状态
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

const CheckboxChecked = ref(false);
const gridEvents: DeepPartial<VxeGridListeners> = {
  checkboxChange: handleCheckboxChange,
  checkboxAll: handleCheckboxChange,
};

function handleCheckboxChange() {
  CheckboxChecked.value = gridApi.grid.getCheckboxRecords().length > 0;
}
const [Grid,gridApi] = useVbenVxeGrid({
  // formOptions,
  gridOptions,
  gridEvents,
});
const labelColor = (row: any) => {
  if (row.listClass == 'primary') {
    return 'blue';
  }
  if (row.listClass == 'success') {
    return 'green';
  }
  if (row.listClass == 'warning') {
    return 'yellow';
  }
  if (row.listClass == 'danger') {
    return 'red';
  }
  if (row.listClass == 'secondary') {
    return 'purple';
  }
  if (row.listClass == 'info') {
    return 'cyan';
  }
  return 'default';
}

const [DictDrawer, dictDrawerApi] = useVbenDrawer({
  connectedComponent: dictDataDrawer,
});


function handleView(row: SysDictDataListModel) {
  const { dictCode } = row;
  dictDrawerApi.setData({id: dictCode, update:false,view:true,dictType:dictType.value});
  dictDrawerApi.open();
}

function handleAdd() {
  dictDrawerApi.setData({update:false, view:false,dictType:dictType.value});
  dictDrawerApi.open();
}

function handleEdit(row: SysDictDataListModel) {
  dictDrawerApi.setData({ id: row.dictCode, update:true, view:false,dictType:dictType.value });
  dictDrawerApi.open();
}

async function handleDelete(row: SysDictDataListModel) {
  await deleteSysDictDataApi({ dictCodes: [row.dictCode] });
  message.success("删除成功");
  await handleRefresh();
}
async function handleRefresh() {
  await gridApi.query();
}


function handleMultiDelete() {
  const rows = gridApi.grid.getCheckboxRecords();
  const ids = rows.map((row: SysDictDataListModel) => row.dictCode);
  Modal.confirm({
    title: '提示',
    okType: 'danger',
    content: `确认删除选中的${ids.length}条记录吗？`,
    onOk: async () => {
      await deleteSysDictDataApi({ dictCodes: ids });
      message.success("删除成功");
      await handleRefresh();
    },
  });
}

</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="'['+dictName+']' +'['+dictType+']' + '字典数据'">
      <template #toolbar-tools>
        <!-- 新增按钮：系统内置时只有超级管理员可操作 -->
        <AccessControl
          :codes="dictIsSys === '0' ? ['cpr:superadmin'] : ['cpm:system:dict:add']"
          type="code"
        >
          <template #default="{ hasPermission }">
            <Button
              class="mr-2 flex items-center"
              type="primary"
              :icon="h(MdiPlus)"
              :disabled="!hasPermission"
              @click="handleAdd"
            >
              新增
            </Button>
          </template>
        </AccessControl>

        <!-- 批量删除按钮：系统内置时只有超级管理员可操作 -->
        <AccessControl
          :codes="dictIsSys === '0' ? ['cpr:superadmin'] : ['cpm:system:dict:remove']"
          type="code"
        >
          <template #default="{ hasPermission }">
            <Button
              class="mr-2 flex items-center"
              type="primary"
              :disabled="!CheckboxChecked || !hasPermission"
              :icon="h(MdiDelete)"
              @click="handleMultiDelete"
            >
              删除
            </Button>
          </template>
        </AccessControl>
      </template>
      <template #label="{ row }">
        <Tag :color="labelColor(row)">{{ row.dictLabel }}</Tag>
      </template>
      <template #status="{ row }">
        <Tag :color="row.status == '0' ? 'green' : 'red'">{{ row.status == '0' ? '正常' : '停用' }}</Tag>
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleView(row)" v-access:code="'cpm:system:dict:query'">查看</Button>

          <!-- 修改按钮：系统内置时只有超级管理员可操作 -->
          <AccessControl
            :codes="dictIsSys === '0' ? ['cpr:superadmin'] : ['cpm:system:dict:edit']"
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
            :codes="dictIsSys === '0' ? ['cpr:superadmin'] : ['cpm:system:dict:remove']"
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
