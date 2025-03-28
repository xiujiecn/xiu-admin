<!--
 * @description 图片验证码组件
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->

<script lang="ts" setup>
import {onMounted, ref, watch} from 'vue'
import { Input, Image } from 'ant-design-vue';
import { getCaptchaApi } from '#/api';
const emit = defineEmits(['blur', 'change']);
let imgCaptcha = ref<string>('');
// 加载时获取验证码
onMounted(()=>{
  getCaptcha();
});
const props = defineProps<{
  captchaCount?: number;
}>();

watch(
  () => props.captchaCount,
  () => {
    getCaptcha();
  },{});


const modelValue = defineModel<[string, string]>({
  default: () => [undefined, undefined],
});

function getCaptcha() {
  getCaptchaApi().then((res) => {
    imgCaptcha.value = res.captchaImage;
    modelValue.value[1] = res.captchaID;
  });
}
function onChange() {
  emit('change', modelValue.value);
}

</script>
<template>
  <div class="flex w-full gap-1">
    <Input
      placeholder="请输入验证码"
      class="w-[260px]"
      allow-clear
      :class="{ 'valid-success': !!modelValue[0] }"
      v-model:value="modelValue[0]"
      :maxlength="6"
      type="tel"
      @blur="emit('blur')"
      @change="onChange"
    />
    <Image class="flex-1" @click="getCaptcha" :preview="false" :src="imgCaptcha" :width="130" :height="38" style="cursor: pointer" />
  </div>
</template>
