<script setup lang="ts">
import { ref } from 'vue';

import { IconifyIcon } from '@vben/icons';

import { Button, Input, RadioButton, RadioGroup, Select } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';

const searchForm = ref({});

const [Grid, gridApi] = useVbenVxeGrid({
  gridClass: 'p-8',
  gridOptions: {},
});
</script>
<template>
  <Grid>
    <template #table-title>
      <div class="flex w-full flex-col gap-2">
        <p class="text-2xl">设备(1)</p>
        <div class="flex justify-between [&>*]:flex [&>*]:gap-2">
          <div>
            <Button>设备组</Button>
            <Button>设备查询</Button>
            <Select></Select>
            <Select></Select>
            <Select></Select>
            <Select></Select>
            <Input></Input>
          </div>
          <div>
            <Button>123</Button>
            <RadioGroup>
              <RadioButton label="1">
                <IconifyIcon mode="mask" icon="material-symbols:grid-view" />
              </RadioButton>
              <RadioButton label="2">
                <IconifyIcon
                  mode="mask"
                  icon="material-symbols:format-list-bulleted"
                />
              </RadioButton>
              <RadioButton label="3">
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

<style scoped></style>
