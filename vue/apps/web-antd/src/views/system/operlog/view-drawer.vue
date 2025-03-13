<script setup lang="ts">
import type { SysOperLog } from '#/api/system/oper-log';

import { computed, nextTick, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { addFullName, cloneDeep } from '@vben/utils';

import { useVbenForm } from '#/adapter/form';
import { Description, useDescription } from '#/components/description';
import { viewSchema } from './model';


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
  const { record } = drawerApi.getData() as { record: SysOperLog };
  setDescProps({ data: record }, true);
}

</script>

<template>
  <BasicDrawer :footer="false" class="w-[600px]" title="查看日志">
    <Description @register="registerDescription">
    </Description>
  </BasicDrawer>
</template>
