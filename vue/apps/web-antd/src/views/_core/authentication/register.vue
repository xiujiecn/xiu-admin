<!--
 * @description 注册页面
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script lang="ts" setup>
import type { VbenFormSchema } from '@vben/common-ui';
import type { Recordable } from '@vben/types';

import { computed, h, markRaw, onBeforeUnmount, ref } from 'vue';
import { useRouter } from 'vue-router';

import { AuthenticationRegister, z } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { message } from 'ant-design-vue';
import { md5 } from 'js-md5';

import { registerApi } from '#/api';

import RegisterCodeInput from './components/register-code-input.vue';

defineOptions({ name: 'Register' });

const loading = ref(false);
const router = useRouter();
let redirectTimer: ReturnType<typeof window.setTimeout> | undefined;

const contactSchema = z
  .string()
  .min(1, { message: '请输入手机号或邮箱' })
  .refine(
    (value) =>
      /^1[3-9]\d{9}$/.test(value) ||
      /^[^\s@]+@[^\s@][^\s.@]*\.[^\s@]+$/.test(value),
    {
      message: '请输入正确的手机号或邮箱',
    },
  );

const formSchema = computed((): VbenFormSchema[] => {
  return [
    {
      component: 'VbenInput',
      componentProps: {
        placeholder: '请输入手机号或邮箱',
      },
      fieldName: 'contact',
      label: '手机号或邮箱',
      rules: contactSchema,
    },
    {
      component: markRaw(RegisterCodeInput),
      dependencies: {
        componentProps(values) {
          return {
            contact: values.contact,
          };
        },
        triggerFields: ['contact'],
      },
      fieldName: 'code',
      label: '验证码',
      rules: z.string().min(1, { message: '请输入验证码' }),
    },
    {
      component: 'VbenInputPassword',
      componentProps: {
        passwordStrength: true,
        placeholder: $t('authentication.password'),
      },
      fieldName: 'password',
      label: $t('authentication.password'),
      rules: z.string().min(1, { message: $t('authentication.passwordTip') }),
    },
    {
      component: 'VbenInputPassword',
      componentProps: {
        placeholder: $t('authentication.confirmPassword'),
      },
      dependencies: {
        rules(values) {
          const { password } = values;
          return z
            .string({ required_error: $t('authentication.passwordTip') })
            .min(1, { message: $t('authentication.passwordTip') })
            .refine((value) => value === password, {
              message: $t('authentication.confirmPasswordTip'),
            });
        },
        triggerFields: ['password'],
      },
      fieldName: 'confirmPassword',
      label: $t('authentication.confirmPassword'),
    },
    {
      component: 'VbenInput',
      componentProps: {
        placeholder: '请输入公司名称',
      },
      fieldName: 'companyName',
      label: '公司名称',
      rules: z.string().min(1, { message: '请输入公司名称' }),
    },
    {
      component: 'VbenCheckbox',
      fieldName: 'agreePolicy',
      renderComponentContent: () => ({
        default: () =>
          h('span', [
            $t('authentication.agree'),
            h(
              'a',
              {
                class: 'vben-link ml-1 ',
                href: '',
              },
              `${$t('authentication.privacyPolicy')} & ${$t('authentication.terms')}`,
            ),
          ]),
      }),
      rules: z.boolean().refine((value) => !!value, {
        message: $t('authentication.agreeTip'),
      }),
    },
  ];
});

async function handleSubmit(value: Recordable<any>) {
  if (loading.value) {
    return;
  }
  loading.value = true;
  try {
    const password = md5(value.password);
    await registerApi({
      code: value.code,
      companyName: value.companyName,
      confirmPassword: password,
      contact: value.contact,
      password,
      username: value.contact,
    });
    message.success('注册成功，5秒后跳转到登录界面');
    redirectTimer = window.setTimeout(() => {
      router.push('/auth/login');
    }, 5000);
  } catch (error) {
    loading.value = false;
    throw error;
  }
}

onBeforeUnmount(() => {
  if (redirectTimer) {
    window.clearTimeout(redirectTimer);
  }
});
</script>

<template>
  <AuthenticationRegister
    :form-schema="formSchema"
    :loading="loading"
    @submit="handleSubmit"
  />
</template>
