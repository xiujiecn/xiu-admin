<!--
 * @description 部门管理编辑抽屉组件
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 XiuAdmin CLI
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script setup lang="ts">
import type { SysDeptTreeData } from '#/api/system/dept';
import type { SysUserMini } from '#/api/system/user';

import { computed, nextTick, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { addFullName , cloneDeep, listToTree} from '@vben/utils';

import { useVbenForm } from '#/adapter/form';

import { addSysDeptApi, editSysDeptApi, viewSysDeptApi, getSysDeptTreeApi } from '#/api/system/dept';
import { getSysUserListByDeptIdApi } from '#/api/system/user';
import { drawerSchema } from './model';

const emit = defineEmits<{ reload: [] }>();
interface ModalProps {
  id?: number | string;
  update: boolean;
  view: boolean;
}

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
      fieldName: 'parentId',
    },
  ]);
}

async function initDeptUsers(deptId: number | string) {
  const ret = await getSysUserListByDeptIdApi({deptId:Number(deptId)});
  const data = ret?.items || [];
  const options = data.map((user:SysUserMini) => ({
    label: `${user.userName} | ${user.nickName}`,
    value: user.userId,
  }));
  formApi.updateSchema([
    {
      componentProps: {
        disabled: data.length === 0,
        options,
        placeholder: data.length === 0 ? '该部门暂无用户' : '请选择部门负责人',
      },
      fieldName: 'leader',
    },
  ]);
}

async function setLeaderOptions() {
  formApi.updateSchema([
    {
      componentProps: {
        disabled: true,
        options: [],
        placeholder: '仅在更新时可选部门负责人',
      },
      fieldName: 'leader',
    },
  ]);
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
  wrapperClass: 'grid-cols-1 gap-x-4',
});

const [BasicDrawer, drawerApi] = useVbenDrawer({
  onCancel: handleCancel,
  onConfirm: handleConfirm,
  async onOpenChange(isOpen) {
    if (!isOpen) {
      return null;
    }
    drawerApi.setState({ confirmLoading: true, loading: true })
    const { id, update, view } = drawerApi.getData() as ModalProps;
    isUpdate.value = update;
    isView.value = view;
    await initDeptSelect();
    if (isUpdate.value || isView.value) {
      const record = await viewSysDeptApi({ deptId: Number(id) });
      record.leader = record.leader ? Number(record.leader) : null;
      await formApi.setValues(record);
      await initDeptUsers(record.deptId);
    }else {
      await setLeaderOptions();
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
    // 这个用于提交
    // formApi.getValues拿到的是一个readonly对象，不能直接修改，需要cloneDeep
    const data = cloneDeep(await formApi.getValues());
    await (isUpdate.value ? editSysDeptApi(data) : addSysDeptApi(data));
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

<template>
  <BasicDrawer :close-on-click-modal="false" :title="title" class="w-[600px]">
    <BasicForm>
    </BasicForm>
  </BasicDrawer>
</template>
