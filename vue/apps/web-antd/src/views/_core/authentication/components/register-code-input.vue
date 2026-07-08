<script lang="ts" setup>
import { computed, nextTick, onBeforeUnmount, ref, watch } from 'vue';

import { Button, Image, Input, Modal, message } from 'ant-design-vue';

import { getCaptchaApi, sendRegisterCodeApi } from '#/api';

const modelValue = defineModel<string>({ default: '' });

const props = withDefaults(
  defineProps<{
    contact?: string;
  }>(),
  {
    contact: '',
  },
);

const countdown = ref(0);
const captchaValue = ref('');
const captchaID = ref('');
const captchaImage = ref('');
const captchaOpen = ref(false);
const captchaInputRef = ref<InstanceType<typeof Input>>();
const sending = ref(false);
let countdownTimer: ReturnType<typeof window.setInterval> | undefined;

const canSend = computed(() => countdown.value <= 0);
const buttonText = computed(() =>
  canSend.value ? '获取验证码' : `${countdown.value}秒后重试`,
);

function isValidContact(value: string) {
  return /^1[3-9]\d{9}$/.test(value) || /^[^\s@]+@[^\s@][^\s.@]*\.[^\s@]+$/.test(value);
}

function startCountdown() {
  countdown.value = 60;
  if (countdownTimer) {
    window.clearInterval(countdownTimer);
  }
  countdownTimer = window.setInterval(() => {
    countdown.value -= 1;
    if (countdown.value <= 0 && countdownTimer) {
      window.clearInterval(countdownTimer);
      countdownTimer = undefined;
    }
  }, 1000);
}

async function refreshCaptcha() {
  const res = await getCaptchaApi();
  captchaImage.value = res.captchaImage;
  captchaID.value = res.captchaID;
  captchaValue.value = '';
  await focusCaptchaInput();
}

async function focusCaptchaInput() {
  await nextTick();
  captchaInputRef.value?.focus?.();
}

async function handleOpenCaptcha() {
  if (!canSend.value) {
    return;
  }
  if (!isValidContact(props.contact.trim())) {
    message.warning('请先输入正确的手机号或邮箱');
    return;
  }
  captchaOpen.value = true;
  await refreshCaptcha();
  await focusCaptchaInput();
}

async function handleSendCode() {
  if (!canSend.value) {
    return;
  }
  if (!captchaValue.value.trim()) {
    message.warning('请输入图片验证码');
    return;
  }
  sending.value = true;
  try {
    await sendRegisterCodeApi({
      captchaID: captchaID.value,
      captchaValue: captchaValue.value,
      contact: props.contact.trim(),
      tenantId: '000000',
    });
    startCountdown();
    captchaOpen.value = false;
    message.success('验证码已发送');
  } catch (error) {
    await refreshCaptcha();
    throw error;
  } finally {
    sending.value = false;
  }
}

watch(
  () => props.contact,
  () => {
    captchaValue.value = '';
  },
);

refreshCaptcha();

onBeforeUnmount(() => {
  if (countdownTimer) {
    window.clearInterval(countdownTimer);
  }
});
</script>

<template>
  <div class="register-code-input register-code-input-main">
    <Input
      v-model:value="modelValue"
      autocomplete="one-time-code"
      placeholder="请输入验证码"
    />
    <Button :disabled="!canSend" html-type="button" @click="handleOpenCaptcha">
      {{ buttonText }}
    </Button>
  </div>
  <Modal
    v-model:open="captchaOpen"
    :confirm-loading="sending"
    destroy-on-close
    title="图片验证码"
    @ok="handleSendCode"
  >
    <div class="register-captcha-modal">
      <Input
        ref="captchaInputRef"
        v-model:value="captchaValue"
        allow-clear
        autocomplete="off"
        :maxlength="6"
        placeholder="请输入图片验证码"
        type="tel"
        @press-enter="handleSendCode"
      />
      <Image
        class="register-captcha-image"
        :height="36"
        :preview="false"
        :src="captchaImage"
        :width="128"
        @click="refreshCaptcha"
      />
    </div>
  </Modal>
</template>

<style scoped>
.register-code-input {
  display: grid;
  grid-template-columns: minmax(0, 1fr) 112px;
  gap: 8px;
  width: 100%;
}

.register-code-input :deep(.ant-input),
.register-code-input :deep(.ant-btn) {
  height: 36px;
}

.register-code-input :deep(.ant-btn) {
  width: 112px;
  padding-inline: 8px;
}

.register-code-input :deep(.ant-btn > span) {
  line-height: 1;
}

.register-captcha-modal {
  display: grid;
  grid-template-columns: minmax(0, 1fr) 128px;
  gap: 10px;
  width: 100%;
}

.register-captcha-image {
  cursor: pointer;
}
</style>
