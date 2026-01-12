<script setup lang="ts">
import type { FormType } from './types';

import type { VxeTableGridOptions } from '#/adapter/vxe-table';

import { computed, ref } from 'vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import BetterVben from '#/constants/betterVben';

import Card from './components/Card.vue';
import Form from './components/Form.vue';
import { IndexGridColumns } from './model';

/** ====================  常量  ==================== */
const GridMode = computed<VxeTableGridOptions>(() => {
  return {
    columns:
      form.value.mode === 1 ? BetterVben.config.cardGrid : IndexGridColumns,
    showHeader: form.value.mode !== 1,
    border: form.value.mode === 1 ? 'none' : 'default',
  };
});

/** ====================  数据  ==================== */
/** 查询表单 */
const form = ref<FormType>({
  /** 设备组 */
  group: false,
  /** 设备查询 */
  query: false,
  /** 设备类型 */
  type: '',
  /** 在线状态 */
  online: '',
  /** 活跃状态 */
  active: '',
  /** 告警状态 */
  alert: '',
  /** 接入类型 */
  access: '',
  /** 搜索设备 */
  search: '',
  /** 排序方式 */
  sort: '',
  /** 展示模式 */
  mode: 1,
});

/** 设备总数 */
const deviceSum = ref(0);

/** 选中的设备 */
const selectRow = ref<any>([]);

/** ====================  方法  ==================== */
/** 切换视图 */
function changeGridMode(value: 1 | 2 | 3) {
  form.value.mode = value;
  gridApi.setGridOptions({
    ...GridMode.value,
  });
  gridApi.reload();
}

/** ====================  配置  ==================== */
/** 表格配置 */
const gridOptions: VxeTableGridOptions = {
  ...BetterVben.grid.default,
  ...GridMode.value,
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = {
          total: 100,
          page: 1,
          pageSize: 20,
          items: [
            {
              deviceName: '设备名称',
              deviceType: '设备类型',
              onlineState: '在线状态',
              update: '活跃时间',
              alert: '告警状态',
            },
          ],
        };
        deviceSum.value = res.total;
        // 卡片模式需要包装数据
        if (form.value.mode === 1) {
          return {
            ...res,
            items: [{ list: res.items }],
          };
        }
        return res;
      },
    },
  },
};

/** 表格事件 */
const gridEvents: VxeGridListeners = {
  checkboxAll: () => {},
  checkboxChange: () => {},
};

/** ====================  组件  ==================== */
const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions,
  gridEvents,
});
</script>

<template>
  <Grid class="p-4" grid-class="grid">
    <template #table-title>
      <div class="flex w-full flex-col gap-2">
        <p class="pb-2 text-2xl">设备({{ deviceSum }})</p>
        <Form v-model:form="form" @change-mode="changeGridMode" />
      </div>
    </template>

    <!-- 卡片视图 -->
    <template #list="{ row }">
      <Card :row="row" />
    </template>
  </Grid>
</template>
