<!--
 * @description 个人中心-个人信息面板
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script setup lang="ts">
import type { UserProfileRes } from '#/api/system/user';

import { computed } from 'vue';

import { preferences, usePreferences } from '@vben/preferences';

import {
  Card,
  Descriptions,
  DescriptionsItem,
  Tag,
  Tooltip,
} from 'ant-design-vue';

import { updateCurrentUserAvatar } from '#/api/system/user';
import { CropperAvatar } from '#/components/cropper';

const props = defineProps<{ profile?: UserProfileRes }>();

defineEmits<{
  // 头像上传完毕
  uploadFinish: [];
}>();

const avatar = computed(
  () => props.profile?.user.avatar ?? preferences.app.defaultAvatar,
);

const { isDark } = usePreferences();
const poetrySrc = computed(() => {
  const color = isDark.value ? 'white' : 'gray';
  return `https://v2.jinrishici.com/one.svg?font-size=12&color=${color}`;
});
const postName = computed(() => {
  return props.profile?.user.posts.find((item) => item.deptId === props.profile?.user.deptId)?.postName;
});
</script>

<template>
  <Card :loading="!profile" class="h-full lg:w-1/3">
    <div v-if="profile" class="flex flex-col items-center gap-[24px]">
      <div class="flex flex-col items-center gap-[20px]">
        <Tooltip title="点击上传头像">
          <CropperAvatar
            :show-btn="false"
            :upload-api="updateCurrentUserAvatar"
            :value="avatar"
            width="120"
            @change="$emit('uploadFinish')"
          />
        </Tooltip>
        <div class="flex flex-col items-center gap-[8px]">
          <span class="text-foreground text-xl font-bold">
            {{ profile.user.nickName ?? '未知' }}
          </span>
          <!-- https://www.jinrishici.com/doc/#image -->
          <img :src="poetrySrc" />
        </div>
      </div>
      <div class="px-[24px]">
        <Descriptions :column="1">
          <DescriptionsItem label="用户账号">
            {{ profile.user.userName }}
          </DescriptionsItem>
          <DescriptionsItem label="手机号码">
            {{ profile.user.phonenumber || '未绑定手机号' }}
          </DescriptionsItem>
          <DescriptionsItem label="用户邮箱">
            {{ profile.user.email || '未绑定邮箱' }}
          </DescriptionsItem>
          <DescriptionsItem label="所属公司">
            <Tag color="processing">
              {{ profile.user.companyInfo?.deptName ?? '未分配公司' }}
            </Tag>
          </DescriptionsItem>
          <DescriptionsItem label="所属组织">
            <Tag color="processing">
              {{ profile.user.deptInfo?.deptName ?? '未分配组织' }}
            </Tag>
            <Tag v-if="postName" color="processing">
              {{ postName }}
            </Tag>
          </DescriptionsItem>
          <DescriptionsItem label="上次登录">
            {{ profile.user.loginDate }}
          </DescriptionsItem>
        </Descriptions>
      </div>
    </div>
  </Card>
</template>
