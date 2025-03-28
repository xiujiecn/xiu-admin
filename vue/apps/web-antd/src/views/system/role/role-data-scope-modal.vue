<!--
 * @description 角色数据权限设置弹窗组件
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script setup lang="ts">
import type { SysDeptListData } from '#/api/system/dept';

import { computed, nextTick, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { cloneDeep } from '@vben/utils';

import { useVbenForm } from '#/adapter/form';
import DeptSelect from '../dept/dept-select.vue';

import { getSysRoleViewApi, editSysRoleDataScopeApi } from '#/api/system/role';
import { dataScopeModalSchema } from './model';

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
    formItemClass: 'col-span-1',
  },
  layout: 'vertical',
  schema: dataScopeModalSchema,
  showDefaultActions: false,
  wrapperClass: 'grid-cols-1 gap-x-4',
});

const deptIds = ref<number[]>([]);
const checkStrictly = ref(false);

const [BasicModal, modalApi] = useVbenModal({
  fullscreenButton: false,
  onCancel: handleCancel,
  onConfirm: handleConfirm,
  async onOpenChange(isOpen) {
    if (!isOpen) {
      return null;
    }
    modalApi.setState({ confirmLoading: true, loading: true })
    const { id, update, view } = modalApi.getData() as ModalProps;
    isUpdate.value = update;
    isView.value = view;

    if (isUpdate.value || isView.value) {
      const record = await getSysRoleViewApi({ roleId: id });
      await nextTick();
      checkStrictly.value = record.deptCheckStrictly === 0;
      deptIds.value = record.deptIds;
      await formApi.setValues(record);
    }
    modalApi.setState({ confirmLoading: false, loading: false })
    if (view) {
      modalApi.setState({ showConfirmButton: false });
      formApi.setState({
        commonConfig: {
          componentProps: {
            readonly: true,
            "only-read": true,
          }
        }
      });
    } else {
      modalApi.setState({ showConfirmButton: true });
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
    modalApi.setState({ confirmLoading: true, loading: true })
    const { valid } = await formApi.validate();
    if (!valid) {
      return;
    }
    // formApi.getValues拿到的是一个readonly对象，不能直接修改，需要cloneDeep
    const data = cloneDeep(await formApi.getValues());
    await editSysRoleDataScopeApi(data);
    emit('reload');
    await handleCancel();
  } catch (error) {
    console.error(error);
  } finally {
    modalApi.setState({ confirmLoading: false, loading: false })
  }
}

async function handleCancel() {
  modalApi.close();
  await formApi.resetForm();
}
/**
 * 通过回调更新 无法通过v-model
 * @param deptIds 选中的部门ID数组
 * @param checkStrictly 部门选择是否严格模式 true 严格模式(解除父子联动) false 非严格模式(父子联动), 默认非严格模式, 后台 1：非严格模式(父子联动) 0：严格模式(解除父子联动)
 */
async function handleDeptChange(deptIds2: number[], checkStrictly2: boolean) {
  await nextTick();
  await formApi.setFieldValue('deptIds', deptIds2);
  await formApi.setFieldValue('deptCheckStrictly', checkStrictly2 ? 0 : 1);
  checkStrictly.value = checkStrictly2;
  deptIds.value = deptIds2;
}
</script>

<template>
  <BasicModal :close-on-click-modal="false" :title="title" class="min-h-[600px] w-[550px]">
    <BasicForm>
      <template #deptIds="slotProps">
        <div class="h-[600px] w-full">
          <DeptSelect :dept-ids="deptIds" :check-strictly="checkStrictly" @change="handleDeptChange" />
        </div>
      </template>
    </BasicForm>
  </BasicModal>
</template>
