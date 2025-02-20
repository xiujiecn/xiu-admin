<script lang="ts" setup>
import { h,ref,onMounted } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions } from '#/adapter/vxe-table';

import { Page } from '@vben/common-ui';

import { Button, message, Switch,Tag  } from 'ant-design-vue';
import dayjs from 'dayjs';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getDictDataListApi } from '#/api'; 
import { useRouter } from 'vue-router';
// 获取url参数



const route = useRouter();
const dictId = parseInt(route.currentRoute.value.path.split('/').pop() || '0');
const dictName = ref('');
const dictType = ref('');

onMounted(() => {
  console.log("vue/apps/web-antd/src/views/system/dict-data/index.vue", dictId);
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
        message.success(`Query params: ${JSON.stringify(formValues)}`);
        return await getDictDataListApi({
          page: page.currentPage,
          pageSize: page.pageSize,
          dictId: dictId,
          ...formValues,
        });
      },
      querySuccess: ({ page, sort, sorts, filters, form, response }) => {
        dictName.value = response.data.type.dictName;
        dictType.value = response.data.type.dictType;
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

const [Grid] = useVbenVxeGrid({
  formOptions,
  gridOptions,
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
</script>

<template>
  <Page auto-content-height>
    <Grid>
      <template #toolbar-actions>
        
        <Button class="mr-2 flex items-center " type="primary" :icon="h(MdiPlus)">新增</Button>
        <Button class="mr-2 flex items-center bg-green-500"  disabled :icon="h(MdiEdit)">编辑</Button>
        <Button class="mr-2 flex items-center" type="primary" disabled :icon="h(MdiDelete)">删除</Button>
      </template>
      <template #label="{ row }">
        <Tag :color="labelColor(row)">{{ row.dictLabel }}</Tag>
      </template>
      <template #status="{ row }">
        <Tag :color="row.status == '0' ? 'green' : 'red'">{{ row.status == '0' ? '正常' : '停用' }}</Tag>
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link">查看</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link" >修改</Button>
          <Button class="mr-2 border-none p-0" :block="false" type="link"  danger>删除</Button>
        </div>
      </template>
    </Grid>
  </Page>
</template>
