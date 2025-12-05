<template>
  <Popover :autoAdjustOverflow="true" v-model:open="visible" trigger="click" :arrow="false">
    <template #content>
      <div class="flex flex-col gap-4 w-[310px]">
        <div>
          <Input v-model:value="filter" placeholder="请输入图标名称" @change="getIconList(pages.page)" />
        </div>
        <div class="grid grid-cols-6 gap-y-4 gap-x-8 justify-items-center">
          <template v-for="value in iconList" :key="value.key">
            <Tooltip :title="value.key">
              <AppIcon class="cursor-pointer size-4" :icon="value.key"
                :style="{ color: model === value.key ? '#409eff' : '' }" @click="clickIcon(value.key)" />
            </Tooltip>
          </template>
        </div>
        <div class="w-full flex items-center justify-center">
          <Pagination :current="pages.page" :defaultPageSize="pages.pageSize" :total="pages.total" size="small"
            :showSizeChanger="false" @change="getIconList" />
        </div>
      </div>
    </template>
    <!-- 触发 popover 的输入框 -->
    <Input :value="modelValue" @update:value="$emit('update:modelValue', $event)"> <template #suffix>
      <AppIcon :icon="modelValue || ''" class="size-4"></AppIcon>
    </template>
    </Input>
  </Popover>
</template>

<script setup lang='ts'>
import AppIcon from './AppIcon.vue';
import { iconManager, type Icon } from '#/utils/IconManagers';
import { Popover, Input, Pagination, Tooltip } from 'ant-design-vue';
import { ref, watch } from 'vue';
const visible = ref<boolean>(false);
const props = defineProps<{
  modelValue?: string;
}>();
const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void;
}>()
/** 内部输入框绑定值 */
const model = ref('');
// 同步外部 modelValue 到内部 model
watch(() => props.modelValue, newVal => {
  model.value = newVal || '';
});

/** 分页配置 */
const pages = ref({
  total: 0,
  page: 1,
  pageSize: 36
});
/** 筛选 */
const filter = ref('');

/** 图标列表 */
const iconList = ref<Icon[]>([]);
/** 获取图标列表 */
function getIconList(page = 1) {
  const res = iconManager.getIconsByPage(page, pages.value.pageSize, filter.value);
  iconList.value = res.icons;
  pages.value.total = res.total;
  pages.value.page = res.page;
}

getIconList();

/** 点击图标 */
function clickIcon(key: string) {
  visible.value = false;
  emit('update:modelValue', key); // 通知父组件更新 v-model 绑定的值
}
</script>

<style scoped></style>
