<script setup lang="ts" name="">
import { Card } from 'ant-design-vue';

/** ====================  常量  ==================== */
defineProps<{ row: any }>();
</script>

<template>
  <div
    class="grid grid-cols-1 gap-4 lg:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4"
  >
    <Card
      v-for="item in row.list"
      :key="item.deviceName"
      class="w-full"
      :body-style="{ padding: '14px 14px 0 14px' }"
    >
      <div class="flex flex-col">
        <div class="mb-2 flex w-full items-center gap-2 text-left">
          <img
            src=""
            style="
              width: 62px;
              height: 62px;
              border-radius: 8px;
              object-fit: cover;
            "
          />
          <div
            class="flex flex-1 flex-col overflow-hidden text-nowrap text-left [&>p]:overflow-hidden [&>p]:text-ellipsis"
          >
            <p :title="item.deviceName" class="text-lg font-bold">
              {{ item.deviceName }}
            </p>
            <p>{{ item.onlineState }}</p>
          </div>
        </div>
        <div
          class="mb-2 grid grid-cols-2 gap-2 text-nowrap text-left [&>p]:overflow-hidden [&>p]:text-ellipsis"
        >
          <p :title="item.deviceType">设备类型:{{ item.deviceType }}</p>
          <p :title="item.alert">告警状态:{{ item.alert }}</p>
          <p :title="item.update" class="col-span-2">
            活跃时间:{{ item.update }}
          </p>
        </div>
      </div>
    </Card>
  </div>
</template>

<style scoped></style>
