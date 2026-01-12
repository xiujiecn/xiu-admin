<script setup lang="ts" name="">
import type { FormType } from '../types';

import { IconifyIcon } from '@vben/icons';

import { Button, Input, RadioButton, RadioGroup, Select } from 'ant-design-vue';
/** ====================  常量  ==================== */
const emit = defineEmits(['changeMode']);
/** ====================  数据  ==================== */
const searchForm = defineModel<FormType>('form', { default: {} });
</script>

<template>
  <div class="flex justify-between [&>*]:flex [&>*]:gap-2">
    <div>
      <Button v-model="searchForm.group" class="flex items-center">
        <IconifyIcon icon="material-symbols:devices-other-outline-rounded" />
        设备组
      </Button>
      <Button v-model="searchForm.query" class="flex items-center">
        <IconifyIcon icon="material-symbols:database-search-outline-rounded" />
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
      <RadioGroup
        v-model:value="searchForm.mode"
        @change="
          (e) => {
            emit('changeMode', e.target.value);
          }
        "
      >
        <RadioButton :value="1">
          <IconifyIcon mode="mask" icon="material-symbols:grid-view" />
        </RadioButton>
        <RadioButton :value="2">
          <IconifyIcon
            mode="mask"
            icon="material-symbols:format-list-bulleted"
          />
        </RadioButton>
        <RadioButton :value="3">
          <IconifyIcon mode="mask" icon="material-symbols:globe-location-pin" />
        </RadioButton>
      </RadioGroup>
    </div>
  </div>
</template>

<style scoped></style>
