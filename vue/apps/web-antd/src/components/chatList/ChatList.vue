<template>
  <div
    ref="chatContainerRef"
    class="flex h-full w-full flex-col gap-2 overflow-y-auto"
  >
    <Bubble
      class="bubble"
      v-for="(item, index) in chatHistory"
      :key="index"
      :placement="item.role === 'user' ? 'end' : 'start'"
      :content="item.content"
      :autoSize="{ minRows: 2, maxRows: 6 }"
      :messageRender="formart"
    >
      <template #footer>
        <slot v-if="item.role === 'user'" name="userFooter" :data="item">
          <p>{{ showTime ? item.createdAt : '' }}</p>
        </slot>
        <slot v-else name="cialloFooter" :data="item">
          <p>{{ showTime ? item.createdAt : '' }}</p>
        </slot>
      </template>
    </Bubble>
  </div>
</template>

<script setup lang="ts" name="">
import { Bubble } from 'ant-design-x-vue';
import { h, ref, watch } from 'vue';

/** ====================  setup  ==================== */
// 容器实例
const chatContainerRef = ref<any>(null);
// props
const {
  chatHistory,
  formart = (content: string) => {
    return h(
      'pre',
      {
        style: {
          margin: 0,
          fontFamily:
            'ui-monospace, SFMono-Regular, "SF Mono", Menlo, monospace',
          whiteSpace: 'pre-wrap',
          wordBreak: 'break-word',
        },
      },
      String(content),
    );
  },
  autoScroll = true,
  showTime = false,
} = defineProps<Props>();

/** ====================  暴露事件  ==================== */
// 添加对话记录
function addChatHistory(
  content: string,
  user: 'user' | 'ciallo',
  createdAt?: string,
) {
  chatHistory!.push({
    role: user,
    content,
    createdAt: createdAt == undefined ? new Date().toLocaleString() : createdAt,
  });

  autoScroll &&
    setTimeout(() => {
      scrollToBottom();
    });
}

// 滚动到底部
function scrollToBottom() {
  chatContainerRef.value.scrollTop = chatContainerRef.value.scrollHeight;
}

watch(
  () => chatHistory,
  () => {
    autoScroll &&
      setTimeout(() => {
        scrollToBottom();
      });
  },
);

defineExpose({
  addChatHistory,
  scrollToBottom,
});

/** ====================  types  ==================== */
interface Props {
  /** 对话记录 */
  chatHistory?: ChatHistory[];
  /** 格式化模板 */
  formart?: (content: string) => any;
  /** 自动滚动 */
  autoScroll?: boolean;
  /** 是否显示时间 */
  showTime?: boolean;
}
interface ChatHistory {
  /** 角色
   * @description user: 用户，ciallo: 模型
   */
  role: 'user' | 'ciallo';
  /** 内容 */
  content: string;
  /** 创建时间 */
  createdAt: string;
}
</script>

<style scoped></style>
