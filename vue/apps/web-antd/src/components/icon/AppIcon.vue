<template>
  <div
    class="flex items-center justify-center"
    :style="sizeStyle"
  >
    <!-- 如果找到图标，渲染其SVG -->
    <span
      v-if="iconSvg"
      class="app-icon icon"
      :style="{
        display: 'inline-block',
        width: '100%',
        aspectRatio: '1 / 1', // 保持1:1比例
        color: color,
        ...bodyStyle
      }"
      v-html="iconSvg"
    ></span>
    <!-- 如果未找到图标，显示一个占位符或空 -->
    <span
      v-else
      class="app-icon-placeholder icon"
      :style="{
        display: 'inline-block',
        width: '100%',
        aspectRatio: '1 / 1', // 保持1:1比例
        lineHeight: '1',
        textAlign: 'center',
        color: '#ccc',
        fontSize: '0.6em'
      }"
    >?</span>
  </div>
</template>

<script setup lang='ts' name=''>
import { iconManager } from '#/utils/IconManagers';
import { computed } from 'vue';

interface Props {
  /** 图标名称 (如: 'icon_power') */
  icon: string;
  /** 图标大小 (支持数字或带单位字符串) */
  size?: number | string;
  /** 图标颜色 */
  color?: string;
  /** 自定义内联样式 */
  bodyStyle?: Record<string, any>;
}

const props = withDefaults(defineProps<Props>(), {
  color: 'currentColor',
  bodyStyle: () => ({}),
});

const iconSvg = computed(() => {
  const icon = iconManager.getIconByKey(props.icon);
  return icon ? icon.svg : null;
});

// 计算 size 样式
const sizeStyle = computed(() => {
  if (!props.size) return {};

  // 如果是数字，添加 px 单位
  if (typeof props.size === 'number') {
    return {
      width: `${props.size}px`,
      height: `${props.size}px`
    };
  }

  // 如果是字符串，直接使用
  return {
    width: props.size,
    height: props.size
  };
});
</script>

<style scoped>
.app-icon {
  /* 确保SVG能继承父级颜色 */
  fill: currentColor;
  stroke: currentColor;
}
</style>
