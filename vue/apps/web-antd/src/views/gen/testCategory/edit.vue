<template>
  <BasicDrawer :close-on-click-modal="false" :title="title" class="w-[600px]">
    <BasicForm></BasicForm>
  </BasicDrawer>
</template>
<script setup lang="ts">
import { computed, nextTick, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { addFullName, cloneDeep } from '@vben/utils';

import { useVbenForm } from '#/adapter/form';
import { Alert } from 'ant-design-vue';
import { Edit, View } from '#/api/gen/testCategory';
import { editSchema } from './model';

const emit = defineEmits<{ reload: [] }>();
interface ModalProps {
  id?: number | string;
  update: boolean;
  view: boolean;
}

const isUpdate = ref(false);
const isView = ref(false);

const title = computed(() => {
  if (isView.value) {
    return $t('pages.common.view');
  }
  return isUpdate.value ? $t('pages.common.edit') : $t('pages.common.add');
});

const [BasicForm, formApi] = useVbenForm({
  commonConfig: {
    componentProps: {
      class: 'w-full',
    },
    formItemClass: 'col-span-3',
  },
  layout: 'vertical',
  schema: editSchema,
  showDefaultActions: false,
  wrapperClass: 'grid-cols-3 gap-x-4',
});

const [BasicDrawer, drawerApi] = useVbenDrawer({
  onCancel: handleCancel,
  onConfirm: handleConfirm,
  async onOpenChange(isOpen) {
    if (!isOpen) {
      return null;
    }
    drawerApi.setState({ confirmLoading: true, loading: true })
    const { id, update, view, } = drawerApi.getData() as ModalProps;
    isUpdate.value = update;
    isView.value = view;
    if (isUpdate.value || isView.value) {
      const record = await View({ id: id });
      await formApi.setValues(record);
    }
    drawerApi.setState({ confirmLoading: false, loading: false })

    if (view) {
      drawerApi.setState({ showConfirmButton: false });
      formApi.setState({
        commonConfig: {
          componentProps: {
            readonly: true,
            "only-read": true,
          }
        }
      });
    } else {
      drawerApi.setState({ showConfirmButton: true });
      formApi.setState({
        commonConfig: {
          componentProps: {
            readonly: false,
            "only-read": false,
          }
        }
      });
    }
  },
});

async function handleConfirm() {
  try {
    drawerApi.setState({ confirmLoading: true, loading: true })
    const { valid } = await formApi.validate();
    if (!valid) {
      return;
    }
    const data = cloneDeep(await formApi.getValues());
    await (Edit(data));
    emit('reload');
    await handleCancel();
  } catch (error) {
    console.error(error);
  } finally {
    drawerApi.setState({ confirmLoading: false, loading: false })
  }
}

async function handleCancel() {
  drawerApi.close();
  await formApi.resetForm();
}

</script>