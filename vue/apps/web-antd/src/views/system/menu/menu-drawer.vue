<!--
 * @description 菜单编辑抽屉组件
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script setup lang="ts">
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';
import {
  addFullName,
  cloneDeep,
  getPopupContainer,
  listToTree,
} from '@vben/utils';

import { useVbenForm } from '#/adapter/form';
import {
  addSysMenuApi,
  deleteSysMenuApi,
  getSysMenuViewApi,
  getSysMenuListApi,
  updateSysMenuApi,
} from '#/api/system/menu';

import { drawerSchema } from './model';

interface ModalProps {
  id?: number | string;
  update: boolean;
  view: boolean;
}

const emit = defineEmits<{ reload: [] }>();

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
    formItemClass: 'col-span-2',
    labelWidth: 90,
  },
  schema: drawerSchema,
  showDefaultActions: false,
  wrapperClass: 'grid-cols-2',
});

async function setupMenuSelect() {
  // menu
  const menuArray = await getSysMenuListApi();
  console.log('setupMenuSelect',menuArray);
  // support i18n
  menuArray.items.forEach((item) => {
    item.menuName = $t(item.menuName);
  });
  // const folderArray = menuArray.filter((item) => item.menuType === 'M');
  /**
   * 这里需要过滤掉按钮类型
   * 不允许在按钮下添加数据
   */
  const filteredList = menuArray.items.filter((item) => item.menuType !== 'F');
  const menuTree = listToTree(filteredList, { id: 'menuId', pid: 'parentId' });
  const fullMenuTree = [
    {
      menuId: 0,
      menuName: $t('menu.root'),
      children: menuTree,
    },
  ];
  addFullName(fullMenuTree, 'menuName', ' / ');

  formApi.updateSchema([
    {
      componentProps: {
        fieldNames: {
          label: 'menuName',
          value: 'menuId',
        },
        getPopupContainer,
        // 设置弹窗滚动高度 默认256
        listHeight: 300,
        showSearch: true,
        treeData: fullMenuTree,
        treeDefaultExpandAll: false,
        // 默认展开的树节点
        treeDefaultExpandedKeys: [0],
        treeLine: { showLeafIcon: false },
        // 筛选的字段
        treeNodeFilterProp: 'menuName',
        treeNodeLabelProp: 'fullName',
      },
      fieldName: 'parentId',
    },
  ]);
}

const [BasicDrawer, drawerApi] = useVbenDrawer({
  onCancel: handleCancel,
  onConfirm: handleConfirm,
  async onOpenChange(isOpen) {
    if (!isOpen) {
      return null;
    }
    drawerApi.setState({ confirmLoading: true, loading: true });
    const { id, update,view } = drawerApi.getData() as ModalProps;
    isUpdate.value = update;
    isView.value = view;
    // 加载菜单树选择
    await setupMenuSelect();
    if (id) {
      await formApi.setFieldValue('parentId', id);
      if (update || view) {
        const record = await getSysMenuViewApi({ menuId: Number(id) });
        await formApi.setValues(record);
      }
    }
    drawerApi.setState({ confirmLoading: false, loading: false });
    if (view) {
      drawerApi.setState({ showConfirmButton: false});
      formApi.setState({ commonConfig: { componentProps:{
        readonly:true,
        "only-read":true,
      } } });
    } else {
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
    drawerApi.setState({ confirmLoading: true, loading: true });
    const { valid } = await formApi.validate();
    if (!valid) {
      return;
    }
    const data = cloneDeep(await formApi.getValues());
    await (isUpdate.value ? updateSysMenuApi(data) : addSysMenuApi(data));
    emit('reload');
    await handleCancel();
  } catch (error) {
    console.error(error);
  } finally {
    drawerApi.setState({ confirmLoading: false, loading: false });
  }
}

async function handleCancel() {
  drawerApi.close();
  await formApi.resetForm();
}
</script>

<template>
  <BasicDrawer :close-on-click-modal="false" :title="title" class="w-[600px]">
    <BasicForm :disabled="isView" />
  </BasicDrawer>
</template>
<style scoped>
* {
  color: #000;
}


</style>
