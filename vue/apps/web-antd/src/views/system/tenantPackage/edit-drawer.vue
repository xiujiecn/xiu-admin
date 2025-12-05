<script setup lang="ts">
import { computed, nextTick, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { addFullName, cloneDeep } from '@vben/utils';

import { useVbenForm } from '#/adapter/form';
import { Alert } from 'ant-design-vue';
import {
  addSysTenantPackageApi,
  editSysTenantPackageApi,
  getSysTenantPackageViewApi,
} from '#/api/system/tenant_package';
import { drawerSchema } from './model';
import MenuSelect from '../menu/menu-select.vue';

const emit = defineEmits<{ reload: [] }>();
interface ModalProps {
  id?: number | string;
  update: boolean;
  view: boolean;
}

const showMenuSelect = ref(false); // 菜单选择是否显示,控制加载顺序
const menuIds = ref<number[]>([]);

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
  schema: drawerSchema,
  showDefaultActions: false,
  wrapperClass: 'grid-cols-3 gap-x-4',
});

const [BasicDrawer, drawerApi] = useVbenDrawer({
  onCancel: handleCancel,
  onConfirm: handleConfirm,
  async onOpenChange(isOpen) {
    if (!isOpen) {
      menuIds.value = [];
      return null;
    }
    drawerApi.setState({ confirmLoading: true, loading: true });
    const { id, update, view } = drawerApi.getData() as ModalProps;
    isUpdate.value = update;
    isView.value = view;

    if (isUpdate.value || isView.value) {
      const record = await getSysTenantPackageViewApi({
        packageId: Number(id),
      });
      menuIds.value = record.menuIds.split(',').map(Number);
      await formApi.setValues(record);
    }
    showMenuSelect.value = true;
    drawerApi.setState({ confirmLoading: false, loading: false });

    if (view) {
      drawerApi.setState({ showConfirmButton: false });
      formApi.setState({
        commonConfig: {
          componentProps: {
            readonly: true,
            'only-read': true,
          },
        },
      });
    } else {
      drawerApi.setState({ showConfirmButton: true });
      formApi.setState({
        commonConfig: {
          componentProps: {
            readonly: false,
            'only-read': false,
          },
        },
      });
    }
  },
});
const refMenuSelect = ref<InstanceType<typeof MenuSelect> | null>(null);
async function handleConfirm() {
  try {
    const { valid } = await formApi.validate();
    if (!valid) {
      return;
    }
    drawerApi.setState({ confirmLoading: true, loading: true });
    // 这个用于提交
    // formApi.getValues拿到的是一个readonly对象，不能直接修改，需要cloneDeep
    const data = cloneDeep(await formApi.getValues());
    data.menuIds = refMenuSelect.value?.getData().join(',');
    await (isUpdate.value
      ? editSysTenantPackageApi(data)
      : addSysTenantPackageApi(data));
    emit('reload');
    await handleCancel();
  } catch (error) {
    console.error(error);
  } finally {
    drawerApi.setState({ confirmLoading: false, loading: false });
  }
}

async function handleCancel() {
  refMenuSelect.value?.clearData();
  drawerApi.close();
  await formApi.resetForm();
}

</script>

<template>
  <BasicDrawer :close-on-click-modal="false" :title="title" class="w-[600px]">
    <BasicForm>
      <template #menuIds="slotProps">
        <div class="h-[600px] w-full">
          <MenuSelect
            ref="refMenuSelect"
            :menu-ids="menuIds"
            v-if="showMenuSelect"
          />
        </div>
      </template>
      <template #tip>
        <div class="ml-7 w-full">
          <Alert
            message="私有桶使用自定义域名无法预览, 但可以正常上传/下载"
            show-icon
            type="warning"
          />
        </div>
      </template>
    </BasicForm>
  </BasicDrawer>
</template>
