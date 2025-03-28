<!--
 * @description 定时任务查看抽屉组件
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script setup lang="ts">
import { useVbenDrawer } from '@vben/common-ui';
import { Description, useDescription } from '#/components/description';
import { viewSchema } from './model';
import type { SysJob } from '#/api/system/job';

const [BasicDrawer, drawerApi] = useVbenDrawer({
  onOpenChange: handleOpenChange,
});

const [registerDescription, { setDescProps }] = useDescription({
  column: 1,
  schema: viewSchema,
});

function handleOpenChange(open: boolean) {
  if (!open) {
    return null;
  }
  const { record } = drawerApi.getData() as { record: SysJob };
  console.log('record', record);
  setDescProps({ data: record }, true);
}

</script>

<template>
  <BasicDrawer :footer="false" class="w-[600px]" title="查看任务">
    <Description @register="registerDescription">
    </Description>
  </BasicDrawer>
</template>
