<script setup lang="ts">
import type { UserProfileRes } from '#/api/system/user';

import { computed, onBeforeUnmount, ref } from 'vue';

import { Button, Image, Input, Modal, message } from 'ant-design-vue';

import { getCaptchaApi } from '#/api';
import {
  sendCurrentUserContactCode,
  updateCurrentUserContact,
} from '#/api/system/user';
import { useAuthStore } from '#/store';

import { emitter } from '../mitt';

const props = defineProps<{
  profile: UserProfileRes;
  type: 'email' | 'phone';
}>();

const authStore = useAuthStore();
const contact = ref('');
const code = ref('');
const captchaID = ref('');
const captchaImage = ref('');
const captchaValue = ref('');
const captchaOpen = ref(false);
const countdown = ref(0);
const sending = ref(false);
const saving = ref(false);
let countdownTimer: ReturnType<typeof window.setInterval> | undefined;

const title = computed(() => (props.type === 'email' ? '修改邮箱' : '修改电话'));
const currentValue = computed(() =>
  props.type === 'email'
    ? props.profile.user.email || '未绑定邮箱'
    : props.profile.user.phonenumber || '未绑定电话',
);
const placeholder = computed(() =>
  props.type === 'email' ? '请输入新邮箱' : '请输入新电话',
);
const canSend = computed(() => countdown.value <= 0);
const buttonText = computed(() =>
  canSend.value ? '获取验证码' : `${countdown.value}秒后重试`,
);

function isValidContact(value: string) {
  return props.type === 'email'
    ? /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value)
    : /^1[3-9]\d{9}$/.test(value);
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
}

async function openCaptcha() {
  const value = contact.value.trim();
  if (!isValidContact(value)) {
    message.warning(props.type === 'email' ? '请输入正确的邮箱' : '请输入正确的电话');
    return;
  }
  captchaOpen.value = true;
  await refreshCaptcha();
}

async function sendCode() {
  if (!captchaValue.value.trim()) {
    message.warning('请输入图片验证码');
    return;
  }
  sending.value = true;
  try {
    await sendCurrentUserContactCode({
      captchaID: captchaID.value,
      captchaValue: captchaValue.value,
      contact: contact.value.trim(),
    });
    captchaOpen.value = false;
    startCountdown();
    message.success('验证码已发送');
  } catch (error) {
    await refreshCaptcha();
    throw error;
  } finally {
    sending.value = false;
  }
}

async function saveContact() {
  const value = contact.value.trim();
  if (!isValidContact(value)) {
    message.warning(props.type === 'email' ? '请输入正确的邮箱' : '请输入正确的电话');
    return;
  }
  if (!code.value.trim()) {
    message.warning('请输入验证码');
    return;
  }
  saving.value = true;
  try {
    await updateCurrentUserContact({
      code: code.value.trim(),
      contact: value,
    });
    message.success('修改成功');
    contact.value = '';
    code.value = '';
    await authStore.fetchUserInfo();
    emitter.emit('updateProfile');
  } finally {
    saving.value = false;
  }
}

onBeforeUnmount(() => {
  if (countdownTimer) {
    window.clearInterval(countdownTimer);
  }
});
</script>

<template>
  <div class="contact-setting">
    <div class="contact-setting-title">{{ title }}</div>
    <div class="contact-form">
      <div class="contact-form-item">
        <div class="contact-form-label">当前{{ props.type === 'email' ? '邮箱' : '电话' }}</div>
        <div class="contact-setting-current">{{ currentValue }}</div>
      </div>
      <div class="contact-form-item">
        <div class="contact-form-label">新{{ props.type === 'email' ? '邮箱' : '电话' }}</div>
        <Input v-model:value="contact" :placeholder="placeholder" />
      </div>
      <div class="contact-form-item">
        <div class="contact-form-label">验证码</div>
        <div class="contact-code-row">
          <Input
            v-model:value="code"
            autocomplete="one-time-code"
            placeholder="请输入验证码"
          />
          <Button :disabled="!canSend" @click="openCaptcha">{{ buttonText }}</Button>
        </div>
      </div>
      <div class="contact-form-actions">
        <Button :loading="saving" type="primary" @click="saveContact">
          确认修改
        </Button>
      </div>
    </div>
    <Modal
      v-model:open="captchaOpen"
      :confirm-loading="sending"
      destroy-on-close
      title="图片验证码"
      @ok="sendCode"
    >
      <div class="contact-captcha">
        <Input
          v-model:value="captchaValue"
          allow-clear
          autocomplete="off"
          :maxlength="6"
          placeholder="请输入图片验证码"
          type="tel"
          @press-enter="sendCode"
        />
        <Image
          class="contact-captcha-image"
          :height="36"
          :preview="false"
          :src="captchaImage"
          :width="128"
          @click="refreshCaptcha"
        />
      </div>
    </Modal>
  </div>
</template>

<style scoped>
.contact-setting {
  margin-top: 20px;
  max-width: 420px;
}

.contact-setting-title {
  margin-bottom: 16px;
  font-weight: 600;
}

.contact-setting-current {
  color: hsl(var(--muted-foreground));
  line-height: 36px;
}

.contact-form {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.contact-form-item {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.contact-form-label {
  font-size: 14px;
  line-height: 22px;
}

.contact-code-row {
  display: grid;
  grid-template-columns: minmax(0, 1fr) 112px;
  gap: 8px;
  align-items: center;
}

.contact-form :deep(.ant-input),
.contact-form :deep(.ant-btn) {
  height: 36px;
}

.contact-form-actions {
  padding-top: 2px;
}

.contact-captcha {
  display: grid;
  grid-template-columns: minmax(0, 1fr) 128px;
  gap: 10px;
  width: 100%;
}

.contact-captcha-image {
  cursor: pointer;
}
</style>
