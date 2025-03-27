<!--
 * @description 登录页面
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 XiuAdmin CLI
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script lang="ts" setup>
import type { VbenFormSchema } from '@vben/common-ui';
import type { BasicOption, Recordable } from '@vben/types';
import type { TenantResp } from '#/api';

import { computed, markRaw, useTemplateRef, ref, onMounted } from 'vue';

import { AuthenticationLogin, SliderCaptcha, z } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { useAuthStore } from '#/store';

import ImgCaptcha from './components/img-captcha.vue';
import { tenantList } from '#/api/core/auth';


defineOptions({ name: 'Login' });

const authStore = useAuthStore();

const loginFormRef = useTemplateRef('loginFormRef');
const captchaCount = ref(0);

const MOCK_USER_OPTIONS: BasicOption[] = [
  {
    label: 'Super',
    value: 'super',
  },
  {
    label: 'Admin',
    value: 'admin',
  },
  {
    label: 'User',
    value: 'jack',
  },
];

const tenantInfo = ref<TenantResp>({
  tenantEnabled: true,
  voList: [],
});


async function getTenantList() {
  const resp = await tenantList();
  resp.tenantEnabled = true;
  tenantInfo.value = resp;

  // 选中第一个租户
  if (resp.tenantEnabled && resp.voList.length > 0) {
    const firstTenantId = resp.voList[0]!.tenantId;
    loginFormRef.value?.getFormApi().setFieldValue('tenantId', firstTenantId);
  }
}

onMounted(async () => {
  await getTenantList();
});

async function handleLogin(params: Recordable<any>,
  onSuccess?: () => Promise<void> | void,
  ) {
    try {
      await authStore.authLogin(params, onSuccess);
    } catch (error) {
      if (error instanceof Error) {
        // 刷新验证码
        // captchaCount.value = captchaCount.value+1;
      }
      loginFormRef.value?.getFormApi().setFieldValue('captcha', ['','']);
      captchaCount.value = captchaCount.value+1;
      console.error('登录失败', error);
    }
}
const formSchema = computed((): VbenFormSchema[] => {
  return [
    {
      component: 'VbenSelect',
      componentProps: {
        options: tenantInfo.value.voList?.map((item) => ({
          label: item.companyName,
          value: item.tenantId,
        })),
        placeholder: $t('authentication.selectAccount'),
      },
      fieldName: 'tenantId',
      defaultValue: '000000',
      dependencies: {
        if: () => tenantInfo.value.tenantEnabled,
        // 这里大致上是watch的一个效果
        componentProps: (model) => {
          localStorage.setItem(
            '__oauth_tenant_id',
            model?.tenantId ?? '000000',
          );
          return {};
        },
        triggerFields: ['', 'tenantId'],
      },
      label: $t('authentication.selectAccount'),
      rules: z
        .string()
        .min(1, { message: $t('authentication.selectAccount') })
        .optional()
        .default('admin'),
    },
    {
      component: 'VbenInput',
      componentProps: {
        placeholder: $t('authentication.usernameTip'),
      },
      dependencies: {
        trigger(values, form) {
          if (values.selectAccount) {
            const findUser = MOCK_USER_OPTIONS.find(
              (item) => item.value === values.selectAccount,
            );
            if (findUser) {
              form.setValues({
                password: '123456',
                username: findUser.value,
              });
            }
          }
        },
        triggerFields: ['selectAccount'],
      },
      fieldName: 'username',
      label: $t('authentication.username'),
      rules: z.string().min(1, { message: $t('authentication.usernameTip') }),
    },
    {
      component: 'VbenInputPassword',
      componentProps: {
        placeholder: $t('authentication.password'),
      },
      fieldName: 'password',
      label: $t('authentication.password'),
      rules: z.string().min(1, { message: $t('authentication.passwordTip') }),
    },
    {
      component: markRaw(ImgCaptcha),
      fieldName: 'captcha',
      defaultValue: [undefined, ''],
      componentProps: {
        captchaCount: captchaCount.value,
      },

      rules: z
      .array(z.string().optional()).refine((v) => !!v[0], {
        message: $t('authentication.verifyRequiredTip'),
      }),
    },
  ];
});
</script>

<template>
  <AuthenticationLogin
    ref="loginFormRef"
    :form-schema="formSchema"
    :loading="authStore.loginLoading"
    @submit="handleLogin"
  />
</template>
