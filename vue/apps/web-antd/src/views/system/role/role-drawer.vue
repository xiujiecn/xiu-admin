<script setup lang="ts">
import type { SysMenuListData } from '#/api/system/menu';

import { computed, nextTick, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { cloneDeep, eachTree } from '@vben/utils';

import { useVbenForm } from '#/adapter/form';
import { getSysMenuListApi } from '#/api/system/menu';

import { addSysRoleApi, editSysRoleApi, getSysRoleViewApi } from '#/api/system/role';
// import { MenuSelectTable } from '#/components/tree';

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

// const menuTree = ref<MenuOption[]>([]);
// async function setupMenuTree(id?: number | string) {
//   if (id) {
//     const resp = await getSysMenuListApi();
//     const menus = resp.items;
//     // i18n处理
//     eachTree(menus, (node) => {
//       node.label = $t(node.label);
//     });
//     // 设置菜单信息
//     menuTree.value =  menus;
//     // keys依赖于menu 需要先加载menu
//     await nextTick();
//     await formApi.setFieldValue('menuIds', resp.checkedKeys);
//   } else {
//     const resp = await menuTreeSelect();
//     // i18n处理
//     eachTree(resp, (node) => {
//       node.label = $t(node.label);
//     });
//     // 设置菜单信息
//     menuTree.value = resp;
//     // keys依赖于menu 需要先加载menu
//     await nextTick();
//     await formApi.setFieldValue('menuIds', []);
//   }
// }

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
    const menuIds: number[] = [];//menuSelectRef.value?.getCheckedKeys?.() ?? [];
    // formApi.getValues拿到的是一个readonly对象，不能直接修改，需要cloneDeep
    const data = cloneDeep(await formApi.getValues());
    data.menuIds = menuIds;
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
 * @param value 菜单选择是否严格模式
 */
function handleMenuCheckStrictlyChange(value: boolean) {
  formApi.setFieldValue('menuCheckStrictly', value);
}
</script>

<template>
  <BasicDrawer :close-on-click-modal="false" :title="title" class="w-[800px]">
    <BasicForm>
      <template #menuIds="slotProps">
        <div class="h-[600px] w-full">
          <!-- association为readonly 不能通过v-model绑定 -->
          <!-- <MenuSelectTable
            ref="menuSelectRef"
            :checked-keys="slotProps.value"
            :association="formApi.form.values.menuCheckStrictly"
            :menus="menuTree"
            @update:association="handleMenuCheckStrictlyChange"
          /> -->
        </div>
      </template>
    </BasicForm>
  </BasicDrawer>
</template>
