<script setup lang="ts">
import type { SysTenantListData } from '#/api/system/tenant';

import { computed, nextTick, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { addFullName, cloneDeep } from '@vben/utils';

import { useVbenForm } from '#/adapter/form';
import { Description, useDescription } from '#/components/description';
import { getSysTenantPackageViewApi } from '#/api/system/tenant_package';
import { viewSchema } from './model';


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
  const { record } = drawerApi.getData() as { record: SysTenantListData };
  const record2 = await getSysTenantPackageViewApi ({ packageId: Number(record.id) });
  setDescProps({ data: record2 }, true);
}

</script>

<template>
  <BasicDrawer :footer="false" class="w-[600px]" title="查看">
    <Description @register="registerDescription">
    </Description>
  </BasicDrawer>
</template>
