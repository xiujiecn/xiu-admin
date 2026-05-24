<!--
 * @description 用户编辑抽屉组件
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script setup lang="ts">
import type { SysRoleMini } from '#/api/system/role';
import type {  SysUserViewModel,  } from '#/api/system/user';
import type { SysPostMini } from '#/api/system/post';
import { computed, h, onMounted, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { addFullName, cloneDeep, getPopupContainer } from '@vben/utils';

import { message, Tag } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { getSysPostListApi } from '#/api/system/post';
import { getSysDeptTreeApi } from '#/api/system/dept';
import { getSysRoleListApi } from '#/api/system/role';
import {
  addSysUser,
  getSysUser,
  updateSysUser,
  EmptySysUserViewModel,
} from '#/api/system/user';
import { authScopeOptions } from '#/views/system/role/model';

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
    formItemClass: 'col-span-2',
    componentProps: {
      class: 'w-full',
    },
    labelWidth: 80,
  },
  schema: drawerSchema,
  showDefaultActions: false,
  wrapperClass: 'grid-cols-2',
});

/**
 * 生成角色的自定义label
 * 也可以用option插槽来做
 * renderComponentContent: () => ({
    option: ({value, label, [disabled, key, title]}) => '',
  }),
 */
function genRoleOptionlabel(role: SysRoleMini) {
  const found = authScopeOptions.find((item) => item.value === role.dataScope);
  if (!found) {
    return role.roleName;
  }
  return h('div', { class: 'flex items-center gap-[6px]' }, [
    h('span', null, role.roleName),
    h(Tag, { color: found.color }, () => found.label),
  ]);
}

/**
 * 岗位的加载
 */
async function setupPostOptions(deptId: number | string) {
  const postListResp = await getSysPostListApi({ deptId: Number(deptId) });
  let options: { label: string; value: number }[] = [];
  if (postListResp.items){
      options = postListResp.items?.map((item) => ({
      label: item.postName,
      value: item.postId,
    })) ;
  }
  const placeholder = options.length > 0 ? '请选择' : '该组织下暂无岗位';
  formApi.updateSchema([
    {
      componentProps: { options, placeholder },
      fieldName: 'postIds',
    },
  ]);
}

/**
 * 初始化组织选择
 */
async function setupDeptSelect() {
  // updateSchema
  const deptTree = await getSysDeptTreeApi({});
  // 选中后显示在输入框的值 即父节点 / 子节点
  addFullName(deptTree.items, 'deptName', ' / ');
  console.log('vue/apps/web-antd/src/views/system/user/user-drawer.vue setupDeptSelect',deptTree);
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
          /** 根据组织ID加载岗位 */
          await setupPostOptions(deptId);
          /** 变化后需要重新选择岗位 */
          formModel.postIds = [];
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

onMounted(async () => {

});

const [BasicDrawer, drawerApi] = useVbenDrawer({
  onCancel: handleCancel,
  onConfirm: handleConfirm,
  async onOpenChange(isOpen) {
    if (!isOpen) {
      // 需要重置岗位选择
      formApi.updateSchema([
        {
          componentProps: { options: [], placeholder: '请先选择岗位' },
          fieldName: 'postIds',
        },
      ]);
      return null;
    }
    drawerApi.setState({ confirmLoading: true, loading: true });

    const { id, update, view } = drawerApi.getData() as ModalProps;
    isUpdate.value = update;
    isView.value = view;
    /** update时 禁用用户名修改 不显示密码框 */
    formApi.updateSchema([
      { componentProps: { disabled: isUpdate.value }, fieldName: 'userName' },
      {
        dependencies: { show: () => !isUpdate.value, triggerFields: ['id'] },
        fieldName: 'password',
      },
    ]);
    // 更新 && 赋值
    let user: SysUserViewModel;
    let postIds: number[] = [];
    let roleIds: number[] = [];
    let posts: SysPostMini[] = [];
    let roles: SysRoleMini[] = [];
    const roleListResp = await getSysRoleListApi({});
      roles = roleListResp.items.map((item) => ({
        ...item,
        dataScope: item.dataScope.toString(),
      }));

    if(isUpdate.value || isView.value) {
      user = await getSysUser({ userId: Number(id) });
      postIds = user.posts?.map((item) => item.postId);
      roleIds = user.roles?.map((item) => item.roleId);
    } else {
      user = EmptySysUserViewModel;
    }
    
    

    const postOptions = (posts ?? []).map((item) => ({
      label: item.postName,
      value: item.postId,
    }));
    formApi.updateSchema([
      {
        componentProps: {
          // title用于选中后回填到输入框 默认为label
          optionLabelProp: 'title',
          options: roles.map((item) => ({
            label: genRoleOptionlabel(item),
            // title用于选中后回填到输入框 默认为label
            title: item.roleName,
            value: item.roleId,
          })),
        },
        fieldName: 'roleIds',
      },
      {
        componentProps: {
          options: postOptions,
        },
        fieldName: 'postIds',
      },
    ]);
    // 组织选择 && 初始密码
    await Promise.all([setupDeptSelect()]);//, loadDefaultPassword(isUpdate.value)]);
    if (user.userId > 0) {
      await Promise.all([
        // 添加基础信息
        formApi.setValues(user),
        // 添加角色和岗位
        formApi.setFieldValue('postIds', postIds),
        formApi.setFieldValue('roleIds', roleIds),
        // 更新时不会触发onSelect 需要手动调用
        setupPostOptions(Number(user.deptId)),
      ]);
    }
    drawerApi.setState({ confirmLoading: false, loading: false });

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
    drawerApi.setState({ confirmLoading: true, loading: true });
    const { valid } = await formApi.validate();
    if (!valid) {
      return;
    }
    const data = cloneDeep(await formApi.getValues());
    await (isUpdate.value ? updateSysUser(data) : addSysUser(data));
    isUpdate.value ? message.success('更新成功'): message.success('新增成功');
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
    <BasicForm />
  </BasicDrawer>
</template>
