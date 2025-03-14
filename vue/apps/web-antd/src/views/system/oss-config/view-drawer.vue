<script setup lang="ts">
import type { SysOssConfigListData } from '#/api/system/oss-config';

import { computed, nextTick, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { addFullName, cloneDeep } from '@vben/utils';

import { useVbenForm } from '#/adapter/form';
import { Description, useDescription } from '#/components/description';
import { viewSchema } from './model';
import { getSysOssConfigViewApi } from '#/api/system/oss-config';


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
  const { record } = drawerApi.getData() as { record: SysOssConfigListData };
  console.log("vue/apps/web-antd/src/views/system/oss-config/view-drawer.vue record", record);
  const record2 = await getSysOssConfigViewApi({ ossConfigId: Number(record.ossConfigId) });
  console.log("vue/apps/web-antd/src/views/system/oss-config/view-drawer.vue record2", record2);
  setDescProps({ data: record2 }, true);
}

</script>

<template>
  <BasicDrawer :footer="false" class="w-[600px]" title="查看">
    <Description @register="registerDescription">
    </Description>
  </BasicDrawer>
</template>
