<!--
 * @description 500内部错误页面
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script lang="ts" setup>
import { Fallback } from '@vben/common-ui';

defineOptions({ name: 'Fallback500Demo' });
</script>

<template>
  <Fallback status="500" />
</template>
