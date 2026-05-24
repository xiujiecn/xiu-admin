<!--
 * @description 角色编辑抽屉组件
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
import { useVbenForm } from '#/adapter/form';
import MenuSelect from '../menu/menu-select.vue';
import { Spin } from 'ant-design-vue';
import { getSysDeptTreeApi } from '#/api/system/dept';
import { addFullName, cloneDeep, getPopupContainer } from '@vben/utils';

import {
  addSysRoleApi,
  editSysRoleApi,
  getSysRoleViewApi,
} from '#/api/system/role';
import { drawerSchema } from './model';
/** emit */
const emit = defineEmits<{ reload: [] }>();
/** 是否更新 */
const isUpdate = ref(false);
/** 是否预览 */
const isView = ref(false);
/** 标题 */
const title = computed(() => {
  if (isView.value) {
    return $t('pages.common.view');
  }
  return isUpdate.value ? $t('pages.common.edit') : $t('pages.common.add');
});
/** ====================  变量  ==================== */
/** 是否加载菜单组件 */
const showMenuSelect = ref(false);

/** 选中的菜单id */
const menuIds = ref<number[]>([]);

/** ====================  组件实例  ==================== */
/** 菜单选择组件实例 */
const refMenuSelect = ref<InstanceType<typeof MenuSelect> | null>(null);

/** 表单实例 */
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

/** 抽屉组件实例 */
const [BasicDrawer, drawerApi] = useVbenDrawer({
  onCancel: handleCancel,
  onConfirm: handleConfirm,
  async onOpenChange(isOpen) {
    if (!isOpen) {
      menuIds.value = [];
      return null;
    }
    setupDeptSelect();

    drawerApi.setState({ confirmLoading: true, loading: true });
    const { id, update, view } = drawerApi.getData();
    isUpdate.value = update;
    isView.value = view;
    // 获取数据
    if (isUpdate.value || isView.value) {
      const record = await getSysRoleViewApi({ roleId: id });
      await formApi.setValues(record);

      // 获取菜单id列表
      menuIds.value = record.menuIds;
    }
    showMenuSelect.value = true;
    drawerApi.setState({ confirmLoading: false, loading: false });
    // 根据是否更新调整表达状态
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

async function handleConfirm() {
  try {
    const { valid } = await formApi.validate();
    if (!valid) {
      return;
    }
    // 设置加载状态
    drawerApi.setState({ confirmLoading: true, loading: true });

    // 获取表单数据
    const data = cloneDeep(await formApi.getValues());
    // 获取菜单id
    data.menuIds = refMenuSelect.value?.getData();
    // 编辑数据
    await (isUpdate.value ? editSysRoleApi(data) : addSysRoleApi(data));
    emit('reload');
    await handleCancel();
  } catch (error) {
    console.error(error);
  } finally {
    // 关闭加载状态
    drawerApi.setState({ confirmLoading: false, loading: false });
  }
}

async function handleCancel() {
  refMenuSelect.value?.clearData();
  drawerApi.close();
  await formApi.resetForm();
}


/**
 * 初始化组织选择
 */
 async function setupDeptSelect() {
  // updateSchema
  const deptTree = await getSysDeptTreeApi({ deptType: 1 });
  // 选中后显示在输入框的值 即父节点 / 子节点
  addFullName(deptTree.items, 'deptName', ' / ');
  console.log('role-drawer.vue setupDeptSelect',deptTree);
  formApi.updateSchema([
    {
      componentProps: (formModel) => ({
        class: 'w-full',
        fieldNames: {
          key: 'deptId',
          value: 'deptId',
          children: 'children',
          label: 'deptName',
        },
        getPopupContainer,
        async onSelect(deptId: number | string) {

        },
        placeholder: '请选择',
        showSearch: true,
        treeData: deptTree.items,
        treeDefaultExpandAll: true,
        treeLine: { showLeafIcon: false },
        // 筛选的字段
        treeNodeFilterProp: 'label',
        // 选中后显示在输入框的值
        treeNodeLabelProp: 'fullName',
      }),
      fieldName: 'deptId',
    },
  ]);
}

</script>

<template>
  <BasicDrawer :close-on-click-modal="false" :title="title" class="w-[800px]">
    <BasicForm>
      <template #menuIds>
        <div class="h-[600px] w-full">
          <MenuSelect
            ref="refMenuSelect"
            :menuIds="menuIds"
            v-if="showMenuSelect"
          />
        </div>
      </template>
    </BasicForm>
  </BasicDrawer>
</template>
