<script setup lang="ts">
import { ref } from 'vue';

import { IconifyIcon } from '@vben/icons';

import { Button, Input, RadioButton, RadioGroup, Select } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { IndexGridColumns } from './model';

/** ====================  数据定义  ==================== */
/** 设备总数 */
const deviceSum = ref(0);
/** 查询表单 */
const searchForm = ref({
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
  mode: '',
});

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions: {
    border: false,
    columnConfig: {
      resizable: false,
    },
    headerCellStyle: {
      backgroundColor: '',
      display: 'none',
    },
    columns: IndexGridColumns,
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
          return res;
        },
      },
    },
  },
});
</script>
<template>
  <Grid class="p-4" grid-class="grid">
    <template #table-title>
      <div class="flex w-full flex-col gap-2">
        <p class="pb-2 text-2xl">设备({{ deviceSum }})</p>
        <div class="flex justify-between [&>*]:flex [&>*]:gap-2">
          <div>
            <Button v-model="searchForm.group" class="flex items-center">
              <IconifyIcon
                icon="material-symbols:devices-other-outline-rounded"
              />
              设备组
            </Button>
            <Button v-model="searchForm.query" class="flex items-center">
              <IconifyIcon
                icon="material-symbols:database-search-outline-rounded"
              />
              设备查询
            </Button>
            <Select v-model:value="searchForm.type" />
            <Select v-model:value="searchForm.online" placeholder="在线状态" />
            <Select v-model:value="searchForm.active" placeholder="活跃状态" />
            <Select v-model:value="searchForm.alert" placeholder="告警状态" />
            <Select v-model:value="searchForm.access" placeholder="接入类型" />
            <Input v-model:value="searchForm.search" placeholder="搜索设备" />
          </div>
          <div>
            <Select v-model:value="searchForm.sort" placeholder="排序方式" />
            <RadioGroup v-model:value="searchForm.mode">
              <RadioButton :value="1" label="1">
                <IconifyIcon mode="mask" icon="material-symbols:grid-view" />
              </RadioButton>
              <RadioButton :value="2" label="2">
                <IconifyIcon
                  mode="mask"
                  icon="material-symbols:format-list-bulleted"
                />
              </RadioButton>
              <RadioButton :value="3" label="3">
                <IconifyIcon
                  mode="mask"
                  icon="material-symbols:globe-location-pin"
                />
              </RadioButton>
            </RadioGroup>
          </div>
        </div>
      </div>
    </template>
  </Grid>
</template>

<style scoped lang="scss">
:deep(.grid) {
  .vxe-table--render-wrapper,
  .vxe-table--body-inner-wrapper {
    min-height: 50vh !important;
  }
}
</style>
