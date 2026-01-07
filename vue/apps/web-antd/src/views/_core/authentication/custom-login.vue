<script setup lang="ts">
import type { AuthenticationProps, VbenFormSchema } from '@vben/common-ui';
import type { Recordable } from '@vben/types';

// import type { VbenFormSchema } from '@vben-core/form-ui';
// import type { AuthenticationProps } from './types';
import { computed, onMounted, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';

import { useVbenForm, useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

// import { useVbenForm } from '@vben-core/form-ui';
import { Button as VbenButton, Checkbox as VbenCheckbox } from 'ant-design-vue';

// import Title from './auth-title.vue';
// import ThirdPartyLogin from './third-party-login.vue';

interface Props extends AuthenticationProps {
  formSchema?: VbenFormSchema[];
}

defineOptions({
  name: 'AuthenticationLogin',
});

const props = withDefaults(defineProps<Props>(), {
  codeLoginPath: '/auth/code-login',
  forgetPasswordPath: '/auth/forget-password',
  formSchema: () => [],
  loading: false,
  qrCodeLoginPath: '/auth/qrcode-login',
  registerPath: '/auth/register',
  showCodeLogin: true,
  showForgetPassword: true,
  showQrcodeLogin: true,
  showRegister: true,
  showRememberMe: true,
  showThirdPartyLogin: true,
  submitButtonText: '',
  subTitle: '',
  title: '',
});

const emit = defineEmits<{
  submit: [Recordable<any>];
}>();

const [Form, formApi] = useVbenForm(
  reactive({
    commonConfig: {
      hideRequiredMark: true,
    },
    layout: 'vertical',
    schema: computed(() => props.formSchema),
    showDefaultActions: false,
  }),
);
const router = useRouter();

const REMEMBER_ME_KEY = `REMEMBER_ME_USERNAME_${location.hostname}`;

const localUsername = localStorage.getItem(REMEMBER_ME_KEY) || '';

const rememberMe = ref(!!localUsername);

async function handleSubmit() {
  const { valid } = await formApi.validate();
  const values = await formApi.getValues();
  if (valid) {
    localStorage.setItem(
      REMEMBER_ME_KEY,
      rememberMe.value ? values?.username : '',
    );
    emit('submit', values);
  }
}

function handleGo(path: string) {
  router.push(path);
}

onMounted(() => {
  if (localUsername) {
    formApi.setFieldValue('username', localUsername);
  }
});

defineExpose({
  getFormApi: () => formApi,
});

/** ====================  微信登录  ==================== */
const [Modal, modalApi] = useVbenModal({});
</script>

<template>
  <div @keydown.enter.prevent="handleSubmit">
    <div class="mb-7 sm:mx-auto sm:w-full sm:max-w-md">
      <h2
        class="mb-3 text-3xl font-bold leading-9 tracking-tight text-foreground lg:text-4xl"
      >
        欢迎回来
      </h2>
    </div>
    <!-- 微信快捷登录 -->
    <VbenButton
      size="large"
      type="primary"
      class="mb-6 w-full bg-green-600 hover:!bg-green-500 active:!bg-green-700"
      @click="modalApi.open()"
    >
      微信快捷登录
    </VbenButton>
    <!-- 表单登录 -->
    <Form />
    <!-- 提交按钮 -->
    <VbenButton
      :class="{
        'cursor-wait': loading,
      }"
      size="large"
      :loading="loading"
      aria-label="login"
      type="primary"
      class="mb-6 mt-6 w-full"
      @click="handleSubmit"
    >
      {{ submitButtonText || $t('common.login') }}
    </VbenButton>
    <!-- 记住账号/忘记密码 -->
    <div
      v-if="showRememberMe || showForgetPassword"
      class="mb-6 flex justify-between"
    >
      <div class="flex-center">
        <VbenCheckbox
          v-if="showRememberMe"
          v-model="rememberMe"
          name="rememberMe"
        >
          {{ $t('authentication.rememberMe') }}
        </VbenCheckbox>
      </div>

      <span
        v-if="showForgetPassword"
        class="vben-link text-sm font-normal"
        @click="handleGo(forgetPasswordPath)"
      >
        {{ $t('authentication.forgetPassword') }}
      </span>
    </div>

    <!-- 创建账号 -->
    <slot name="to-register">
      <div v-if="showRegister" class="mt-3 text-center text-sm">
        {{ $t('authentication.accountTip') }}
        <span
          class="vben-link text-sm font-normal"
          @click="handleGo(registerPath)"
        >
          {{ $t('authentication.createAccount') }}
        </span>
      </div>
    </slot>
  </div>
  <!-- 微信快捷登录 模态框 -->
  <Modal
    centered
    title="微信扫一扫"
    :footer="false"
    :fullscreen-button="false"
    class="h-[400px] w-[400px]"
  />
</template>
