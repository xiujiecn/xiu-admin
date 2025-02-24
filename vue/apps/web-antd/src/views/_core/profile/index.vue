<script setup lang="ts">
import type { UserProfileRes } from '#/api/system/user';

import { onMounted, onUnmounted, ref } from 'vue';

import { Page } from '@vben/common-ui';
import { useUserStore } from '@vben/stores';

import { getUserProfileApi } from '#/api/system/user';
import { useAuthStore } from '#/store';

import { emitter } from './mitt';
import ProfilePanel from './profile-panel.vue';
import SettingPanel from './setting-panel.vue';

const profile = ref<UserProfileRes>();
async function loadProfile() {
  const resp = await getUserProfileApi();
  profile.value = resp;
}

onMounted(loadProfile);

const authStore = useAuthStore();
const userStore = useUserStore();
async function handleUploadFinish() {
  // 重新加载用户信息
  await loadProfile();
  // 更新store
  const userInfo = await authStore.fetchUserInfo();
  userStore.setUserInfo(userInfo);
}

onMounted(() => emitter.on('updateProfile', loadProfile));
onUnmounted(() => emitter.off('updateProfile'));
</script>

<template>
  <Page>
    <div class="flex flex-col gap-[16px] lg:flex-row">
      <!-- 左侧 -->
      <ProfilePanel :profile="profile" @upload-finish="handleUploadFinish" />
      <!-- 右侧 -->
      <SettingPanel
        v-if="profile"
        :profile="profile"
        class="flex-1 overflow-hidden"
      />
    </div>
  </Page>
</template>
