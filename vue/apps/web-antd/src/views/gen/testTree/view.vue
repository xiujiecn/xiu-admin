<script setup lang="ts">
  import type { SysClient } from '#/api/system/client';
  import { computed, nextTick, ref } from 'vue';
  import { useVbenDrawer } from '@vben/common-ui';
  import { $t } from '@vben/locales';
  import { addFullName, cloneDeep } from '@vben/utils';
  import { useVbenForm } from '#/adapter/form';
  import { Description, useDescription } from '#/components/description';
  import { viewSchema, type RowType } from './model';
  import { View } from '#/api/gen/testTree';
  const [BasicDrawer, drawerApi] = useVbenDrawer({
  onOpenChange: handleOpenChange,
  });
  const [registerDescription, { setDescProps }] = useDescription({
  column: 1,
  schema: viewSchema,
  });
  async function handleOpenChange(open: boolean) {
  if (!open) {
  return null;
  }
  const { record } = drawerApi.getData() as { record: RowType };
  const record2 = await View({ id: record.id });
  setDescProps({ data: record2 }, true);
  }
</script>
<template>
  <BasicDrawer :footer="false" class="w-[600px]" title="查看">
    <Description @register="registerDescription"></Description>
  </BasicDrawer>
</template>