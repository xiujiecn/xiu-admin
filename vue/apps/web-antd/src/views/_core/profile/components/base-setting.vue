<!--
 * @description 个人中心-基本设置组件
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script setup lang="ts">
import type { Recordable } from '@vben/types';

import type { UserProfileRes } from '#/api/system/user';

import { onMounted } from 'vue';

import { DictEnum } from '@vben/constants';
import { useUserStore } from '@vben/stores';

import { pick } from 'lodash-es';

import { useVbenForm, z } from '#/adapter/form';
import { updateCurrentUserProfile } from '#/api/system/user';
import { useAuthStore } from '#/store';
import { getDictOptions } from '#/utils/dict';

import { emitter } from '../mitt';
import { message } from 'ant-design-vue';

const props = defineProps<{ profile: UserProfileRes }>();

const userStore = useUserStore();
const authStore = useAuthStore();

const [BasicForm, formApi] = useVbenForm({
  actionWrapperClass: 'text-left ml-[68px] mb-[16px]',
  commonConfig: {
    labelWidth: 60,
  },
  handleSubmit,
  resetButtonOptions: {
    show: false,
  },
  schema: [
    {
      component: 'Input',
      dependencies: {
        show: () => false,
        triggerFields: [''],
      },
      fieldName: 'userId',
      label: '用户ID',
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'nickName',
      label: '昵称',
      rules: 'required',
    },
    {
      component: 'RadioGroup',
      componentProps: {
        buttonStyle: 'solid',
        options: getDictOptions(DictEnum.SYS_USER_SEX),
        optionType: 'button',
      },
      defaultValue: '0',
      fieldName: 'sex',
      label: '性别',
      rules: 'required',
    },
  ],
  actionPosition:'left',
  submitButtonOptions: {
    // class: 'min-w-[88px] flex-none',
    content: '更新信息',
  },
});

function buttonLoading(loading: boolean) {
  formApi.setState((prev) => ({
    ...prev,
    submitButtonOptions: { ...prev.submitButtonOptions, loading },
  }));
}

async function handleSubmit(values: Recordable<any>) {
  try {
    buttonLoading(true);
    await updateCurrentUserProfile(values);
    message.success('更新成功');
    // 更新store
    const userInfo = await authStore.fetchUserInfo();
    userStore.setUserInfo(userInfo);
    // 左边reload
    emitter.emit('updateProfile');
  } catch (error) {
    console.error(error);
  } finally {
    buttonLoading(false);
  }
}

onMounted(() => {
  const data = pick(props.profile.user, [
    'userId',
    'nickName',
    'sex',
  ]);
  formApi.setValues(data);
});
</script>

<template>
  <div class="mt-[16px] md:w-full lg:w-1/2 2xl:w-2/5">
    <BasicForm />
  </div>
</template>
