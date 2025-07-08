<script lang="ts" setup>
import {
  computed,
  ref,
  defineExpose,
  defineProps,
  defineEmits,
  withDefaults,
  watch,
} from 'vue';
import { cloneDeep } from 'lodash-es';
import { useVbenModal } from '@vben/common-ui';
import {
  Tabs,
  TabPane,
  Space,
  Tag,
  Button,
  message,
  Textarea,
} from 'ant-design-vue';
import { type GenCodesPreviewRes } from '#/api/gen_codes/gen_table';
import { AccessControl, useAccess } from '@vben/access';
const { hasAccessByCodes } = useAccess();
import hljs from 'highlight.js/lib/core';
import go from 'highlight.js/lib/languages/go';
import typescript from 'highlight.js/lib/languages/typescript';
import xml from 'highlight.js/lib/languages/xml';
import sql from 'highlight.js/lib/languages/sql';

import 'highlight.js/styles/atom-one-dark.css';
import 'highlight.js/lib/common';
import hljsVuePlugin from '@highlightjs/vue-plugin';

hljs.registerLanguage('go', go);
hljs.registerLanguage('ts', typescript);
hljs.registerLanguage('sql', sql);
hljs.registerLanguage('vue', xml);

const highlight = hljsVuePlugin.component;

const emit = defineEmits(['BuildPreview']);
interface Props {
  previewModel: GenCodesPreviewRes;
  showModal: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  previewModel: cloneDeep({ views: {} } as any),
  showModal: false,
});

watch(
  () => props.showModal,
  (showModal) => {
    if (showModal) {
      modalPerviewApi.open();
    } else {
      modalPerviewApi.close();
    }
  },
);

const content = ref('');
const views = computed(() => {
  let tmpViews: any = [];
  let i = 0;
  for (const [k, v] of Object.entries(props.previewModel.views)) {
    let item = v as any;
    item.name = k;
    switch (item.meth) {
      case 1:
        item.tag = { type: 'success', label: '创建文件' };
        break;
      case 2:
        item.tag = { type: 'warning', label: '覆盖文件' };
        break;
      case 3:
        item.tag = { type: 'info', label: '已存在跳过' };
        break;
      case 4:
        item.tag = { type: 'error', label: '不生成' };
        break;
      default:
        item.tag = { type: 'error', label: '未知状态' };
    }
    tmpViews[i] = item;
    i++;
  }
  return tmpViews;
});

function getFileExtension(path: string): string {
  const parts = path.split('.');
  if (parts.length > 1) {
    return parts[parts.length - 1] as string;
  }
  return '';
}

function handleCopy(code: string) {
  content.value = code;
  setTimeout(function () {
    const copyVal = document.getElementById('copy-code');
    if (copyVal) {
      copyVal.select();
      document.execCommand('copy');
      message.success('已复制');
    }
  }, 20);
}

const [ModalPerview, modalPerviewApi] = useVbenModal({
  cancelText: '关闭',
  confirmText: '生成代码',
  onConfirm: function () {
    if (hasAccessByCodes(['cpm:tool:gen:code'])) {
      emit('BuildPreview', '');
    } else {
      message.warning('您没有生成代码的权限');
    }
    modalPerviewApi.close();
  },
  onCancel: function () {
    modalPerviewApi.close();
  },
  onClosed: function () {
    // props.showModal.value = false;
  },
});
defineExpose({
  open: () => {
    modalPerviewApi.open();
  },
});
</script>

<template>
  <ModalPerview :width="800" :height="600" class="h-[600px] w-[800px]">
    <Tabs>
      <TabPane
        v-for="(view, index) in views"
        :key="index"
        :name="view.name"
        :tab="view.name"
        class="overflow-auto"
      >
        <Space class="flex">
          <Tag :color="view.tag.type" :icon="view.tag.icon"
            >{{ view.tag.label }} {{ view.path }}
          </Tag>
          <Button type="default" size="small" @click="handleCopy(view.code)"
            >复制</Button
          >
        </Space>
        <div>
          <highlight
            :class="'code-' + getFileExtension(view.path)"
            :code="view.content"
            :language="getFileExtension(view.path)"
            show-line-numbers
          />
        </div>
      </TabPane>
    </Tabs>
  </ModalPerview>
</template>
