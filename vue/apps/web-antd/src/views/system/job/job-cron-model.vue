<!--
 * @description 定时任务Cron表达式设置弹窗组件
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script setup lang="ts">

import { useVbenModal } from '@vben/common-ui';

import JobCorn from '#/components/vue3cron/job-corn.vue';
import { ref } from 'vue';

const emit = defineEmits<{ 
  reload: [] ,
  confirm: [cron:string]
}>();

const cronExpression = ref<InstanceType<typeof JobCorn>>();
const [BasicModal, modalApi] = useVbenModal({
  
  onOpenChange: (isOpen) => {
    if (isOpen) {
      return null;
    }
    if (cronExpression.value) {
      emit('reload');
      modalApi.close();
      return null;
    }
  },
});
async function handleSubmit(apiData: any) {
    emit('confirm', apiData.cron);
    modalApi.close();
  }
async function handleCancel() {
    modalApi.close();
  }


</script>

<template>
    <BasicModal
      :close-on-click-modal="false"
      :footer="false"
      :fullscreen-button="false"
      title="选择Cron规则"
    >
      <div class="flex flex-col gap-4">
        <JobCorn
        @cancel="handleCancel"
        @confirm="handleSubmit"
        />
      </div>
    </BasicModal>
</template>
