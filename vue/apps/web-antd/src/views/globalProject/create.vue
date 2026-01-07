<script setup lang="ts" name="">
import { IconifyIcon } from '@vben/icons';

import { RadioButton, RadioGroup } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import RouterPath from '#/constants/routerPath';
import { router } from '#/router';

import { GPCreateForm } from './model';

const [Form] = useVbenForm({
  // 按钮配置
  actionPosition: 'left',
  submitButtonOptions: {
    content: '创建项目',
  },
  resetButtonOptions: {
    show: false,
  },
  // 表单配置
  schema: GPCreateForm,
  commonConfig: {
    hideRequiredMark: true,
    componentProps: {
      class: 'w-[450px]',
    },
  },
  layout: 'vertical',
  // 提交回调
  handleSubmit: async (values) => {
    console.log(values);
  },
});
</script>
<template>
  <div class="flex h-full flex-col gap-4 overflow-scroll bg-background p-8">
    <div class="flex items-center gap-4 pb-4">
      <div
        class="flex items-center text-xl"
        @click="router.push(RouterPath.SelectProject)"
      >
        <IconifyIcon mode="mask" icon="material-symbols:arrow-back-rounded" />
        <span class="ml-2 cursor-pointer">返回</span>
      </div>
      <span class="back"></span>
      <p class="text-2xl">创建项目</p>
    </div>
    <div>
      <Form>
        <template #status="props">
          <RadioGroup v-bind="props" button-style="solid">
            <RadioButton :value="1">1231231231231</RadioButton>
            <RadioButton :value="2">21231231231231</RadioButton>
            <RadioButton :value="3">12312312312313</RadioButton>
            <RadioButton :value="4">12312312312314</RadioButton>
          </RadioGroup>
        </template>
        <template #status1="props">
          <RadioGroup v-bind="props" button-style="solid">
            <RadioButton :value="1">1</RadioButton>
            <RadioButton :value="2">2</RadioButton>
            <RadioButton :value="3">3</RadioButton>
            <RadioButton :value="4">4</RadioButton>
          </RadioGroup>
        </template>
      </Form>
    </div>
  </div>
</template>
<style scoped>
.back {
  width: 1px;
  height: 80%;
  background: #595959;
}
</style>
