<script setup lang="ts">
import type { SysMenuListData } from '#/api/system/menu';

import { computed, nextTick, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { cloneDeep } from '@vben/utils';

import { useVbenForm } from '#/adapter/form';
import MenuSelect from '../menu/menu-select.vue';

import { addSysRoleApi, editSysRoleApi, getSysRoleViewApi } from '#/api/system/role';
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
  wrapperClass: 'grid-cols-2 gap-x-4',
});



const checkStrictly = ref(false);

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

    if (isUpdate.value || isView.value) {
      const record = await getSysRoleViewApi({roleId: id});
      await formApi.setValues(record);
      checkStrictly.value = record.menuCheckStrictly === 0;
    }
    // init菜单 注意顺序要放在赋值record之后 内部watch会依赖record
    // await setupMenuTree(id);

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

// const menuSelectRef = ref<InstanceType<typeof MenuSelectTable>>();
async function handleConfirm() {
  try {
    drawerApi.setState({confirmLoading:true,loading:true})
    const { valid } = await formApi.validate();
    if (!valid) {
      return;
    }
    // 这个用于提交
    // const menuIds: number[] = [];//menuSelectRef.value?.getCheckedKeys?.() ?? [];
    // formApi.getValues拿到的是一个readonly对象，不能直接修改，需要cloneDeep
    const data = cloneDeep(await formApi.getValues());
    // data.menuIds = menuIds;
    await (isUpdate.value ? editSysRoleApi (data) : addSysRoleApi(data));
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
/**
 * 通过回调更新 无法通过v-model
 * @param menuIds 选中的菜单ID数组
 * @param checkStrictly 菜单选择是否严格模式 true 严格模式(解除父子联动) false 非严格模式(父子联动), 默认非严格模式, 后台 1：非严格模式(父子联动) 0：严格模式(解除父子联动)
 */
async function handleMenuChange(menuIds: number[], checkStrictly: boolean) {
  await nextTick();
  await formApi.setFieldValue('menuIds', menuIds);
  await formApi.setFieldValue('menuCheckStrictly', checkStrictly ? 0 : 1);
}
</script>

<template>
  <BasicDrawer :close-on-click-modal="false" :title="title" class="w-[800px]">
    <BasicForm>
      <template #menuIds="slotProps">
        <div class="h-[600px] w-full">
          <MenuSelect ref="menuSelectRef" :menu-ids="slotProps.value" :check-strictly="checkStrictly" @change="handleMenuChange" />
        </div>
      </template>
    </BasicForm>
  </BasicDrawer>
</template>
