<script setup lang="ts">
import type { SysMenuListData } from '#/api/system/menu';

import { computed, nextTick, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { addFullName, cloneDeep } from '@vben/utils';

import { useVbenForm } from '#/adapter/form';
import MenuSelect from '../menu/menu-select.vue';

import { addSysPostApi, editSysPostApi, getSysPostViewApi } from '#/api/system/post';
import { getSysDeptTreeApi } from '#/api/system/dept';
import { drawerSchema } from './model';

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


async function getDeptTree() {
  const treeRes = await getSysDeptTreeApi();
  const treeData = treeRes.items;
  addFullName(treeData, 'deptName', ' / ');
  return treeData;
}

async function initDeptSelect() {
  // 需要动态更新TreeSelect组件 这里允许为空
  const treeData = await getDeptTree();
  formApi.updateSchema([
    {
      componentProps: {
        fieldNames: { label: 'deptName', value: 'deptId' },
        showSearch: true,
        treeData: treeData,
        treeDefaultExpandAll: true,
        treeLine: { showLeafIcon: false },
        // 选中后显示在输入框的值
        treeNodeLabelProp: 'fullName',
      },
      fieldName: 'postId',
    },
  ]);
}


const [BasicForm, formApi] = useVbenForm({
  commonConfig: {
    componentProps: {
      class: 'w-full',
    },
    formItemClass: 'col-span-1',
  },
  layout: 'vertical',
  schema: drawerSchema,
  showDefaultActions: false,
  wrapperClass: 'grid-cols-1 gap-x-4',
});

const [BasicDrawer, drawerApi] = useVbenDrawer({
  onCancel: handleCancel,
  onConfirm: handleConfirm,
  async onOpenChange(isOpen) {
    if (!isOpen) {
      return null;
    }
    drawerApi.setState({confirmLoading:true,loading:true})
    const { id, update, view } = drawerApi.getData() as ModalProps;
    isUpdate.value = update;
    isView.value = view;
    await initDeptSelect();
    if (isUpdate.value || isView.value) {
      const record = await getSysPostViewApi({ postId: Number(id) });
      await formApi.setValues(record);
    }
    drawerApi.setState({confirmLoading:false,loading:false})

    if (view) {
      drawerApi.setState({ showConfirmButton: false});
      formApi.setState({ commonConfig: { componentProps:{
        readonly:true,
        "only-read":true,
      } } });
    }else{
      drawerApi.setState({ showConfirmButton: true});
      formApi.setState({ commonConfig: { componentProps:{
        readonly:false,
        "only-read":false,
      }} });
    }

  },
});

async function handleConfirm() {
  try {
    drawerApi.setState({confirmLoading:true,loading:true})
    const { valid } = await formApi.validate();
    if (!valid) {
      return;
    }
    // 这个用于提交
    // formApi.getValues拿到的是一个readonly对象，不能直接修改，需要cloneDeep
    const data = cloneDeep(await formApi.getValues());
    await (isUpdate.value ? editSysPostApi (data) : addSysPostApi(data));
    emit('reload');
    await handleCancel();
  } catch (error) {
    console.error(error);
  } finally {
    drawerApi.setState({confirmLoading:false,loading:false})
  }
}

async function handleCancel() {
  drawerApi.close();
  await formApi.resetForm();
}

</script>

<template>
  <BasicDrawer :close-on-click-modal="false" :title="title" class="w-[600px]">
    <BasicForm>
    </BasicForm>
  </BasicDrawer>
</template>
