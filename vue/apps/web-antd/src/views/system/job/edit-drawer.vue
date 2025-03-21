<script setup lang="ts">
import { computed, nextTick, ref } from 'vue';

import { useVbenDrawer, useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { cloneDeep } from '@vben/utils';
import {Space, Button, Input } from 'ant-design-vue';
import { useVbenForm } from '#/adapter/form';
import { addSysJobApi, updateSysJobApi, viewSysJobApi } from '#/api/system/job';
import { drawerSchema } from './model';
import jobCronModel from './job-cron-model.vue';

const emit = defineEmits<{ reload: [] }>();
interface ModalProps {
  jobId: number | string;
  update: boolean;
  view: boolean;
}

const cronExpression = defineModel('cronExpression', {
  type: String,
  default: '',
});

const isUpdate = ref(false);
const isView = ref(false);

const title = computed(() => {
  if (isView.value) {
    return $t('pages.common.view');
  }
  return isUpdate.value ? $t('pages.common.edit') : $t('pages.common.add');
});

const [JobCronModel, jobCronModelApi] = useVbenModal({
  zIndex: 2001,
  connectedComponent: jobCronModel,
});

async function cronExpressionChanged(cron:any){
  cronExpression.value = cron;
  formApi.setValues({cronExpression: cron});
};

const [BasicForm, formApi] = useVbenForm({
  commonConfig: {
    componentProps: {
      class: 'w-full',
    },
    formItemClass: 'col-span-3',
  },
  layout: 'vertical',
  schema: drawerSchema,
  showDefaultActions: false,
  wrapperClass: 'grid-cols-3 gap-x-4',
});

const [BasicDrawer, drawerApi] = useVbenDrawer({
  onCancel: handleCancel,
  onConfirm: handleConfirm,
  async onOpenChange(isOpen) {
    if (!isOpen) {
      return null;
    }
    drawerApi.setState({confirmLoading:true,loading:true})
    const { jobId, update, view} = drawerApi.getData() as ModalProps;
    isUpdate.value = update;
    isView.value = view;
    cronExpression.value = "";
    if (isUpdate.value || isView.value) {
      const record = await viewSysJobApi({ jobId: Number(jobId) });
      cronExpression.value = record.cronExpression;
      record.misfirePolicy = record.misfirePolicy != null ? String(record.misfirePolicy): null;
      record.concurrent = record.concurrent != null ? String(record.concurrent): null;
      await formApi.setValues(record);
    }
    drawerApi.setState({confirmLoading:false,loading:false})

    if (view) {
      drawerApi.setState({ showConfirmButton: false});
      formApi.setState({ commonConfig: { componentProps:{
        readonly:true,
        "only-read":true,
      } } });
    }else{
      drawerApi.setState({ showConfirmButton: true});
      formApi.setState({ commonConfig: { componentProps:{
        readonly:false,
        "only-read":false,
      }} });
    }
  },
});

async function handleConfirm() {
  try {
    drawerApi.setState({confirmLoading:true,loading:true})
    const { valid } = await formApi.validate();
    if (!valid) {
      return;
    }
    // 这个用于提交
    // formApi.getValues拿到的是一个readonly对象，不能直接修改，需要cloneDeep
    const data = cloneDeep( await formApi.getValues());
    await (isUpdate.value ? updateSysJobApi (data) : addSysJobApi(data));
    emit('reload');
    await handleCancel();
  } catch (error) {
    console.error(error);
  } finally {
    drawerApi.setState({confirmLoading:false,loading:false})
  }
}

async function handleCancel() {
  drawerApi.close();
  await formApi.resetForm();
}

function cronExpressionChange(event: any) {
  console.log("cronExpressionChange", event.target.value);
  formApi.setValues({cronExpression: event.target.value});
}

</script>
<template>
  <BasicDrawer :close-on-click-modal="false" :title="title" class="w-[600px]">
    <BasicForm>
      <template #cronExpression>
        <Space>          
          <Input v-model:value="cronExpression" @change="cronExpressionChange" /> 
          <Button type="primary" @click="jobCronModelApi.open()">
            设置
          </Button>
        </Space>
      </template>
    </BasicForm>
  </BasicDrawer>
  <JobCronModel :close-on-click-modal="false" class="w-[800px]" @confirm="cronExpressionChanged($event)" />
</template>
