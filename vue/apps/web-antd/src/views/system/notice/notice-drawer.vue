<!--
 * @description 通知公告编辑抽屉组件
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script setup lang="ts">
import type { SysConfig } from '#/api/system/config';
import { Select, Button } from 'ant-design-vue';
import { computed, nextTick, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { addFullName, cloneDeep } from '@vben/utils';

import { useVbenForm } from '#/adapter/form';

import { addSysNoticeApi, editSysNoticeApi, getSysNoticeApi } from '#/api/system/notice';
import { drawerSchema } from './model';
import { getSysDeptTreeApi } from '#/api/system/dept';
import { getSysUserListApi } from '#/api/system/user';
import SelectDevice from './selectDevice/index.vue';

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
    if (isUpdate.value || isView.value) {
      const record = await getSysNoticeApi({ noticeId: Number(id) });
      console.log('编辑模式 - 获取表单值:', record);
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
    // 这个用于提交
    // formApi.getValues拿到的是一个readonly对象，不能直接修改，需要cloneDeep
    const data = cloneDeep(await formApi.getValues());
    // console.log(data);
    await (isUpdate.value ? editSysNoticeApi(data) : addSysNoticeApi(data));
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





/** 获取组织树内容,并格式化 */
async function getDeptTree() {
  const treeRes = await getSysDeptTreeApi();
  const treeData = treeRes.items;
  addFullName(treeData, 'deptName', ' / ');
  return treeData;
}
//加载组织选项
async function loadDeptOptions() {
  const treeData = await getDeptTree();
  // console.log(treeData)
  formApi.updateSchema([
    {
      componentProps: {
        /** 显示内容,对应值 */
        fieldNames: { label: 'deptName', value: 'deptId' },
        showSearch: true,
        /** 结构树data */
        treeData: treeData,
        /** 默认展示全部 */
        treeDefaultExpandAll: true,
        treeLine: { showLeafIcon: false },
        treeNodeLabelProp: 'fullName',
        // 选中后显示在输入框的值
        displayRender: (label: any, selected: any, node: any) => {
          return node.props.dataRef.fullName || label;
        },
        // 模糊搜索
        filterTreeNode: (input: string, treeNode: any) => {
          const label = treeNode.deptName || treeNode.fullName || '';
          return label.toLowerCase().includes(input.toLowerCase());
        },
      },
      /** 要修改的字段 */
      fieldName: 'deptIdList',
    },
  ]);
}
loadDeptOptions()


//加载用户选项
async function loadUserOptions() {

  const userData = await getSysUserListApi({});
  formApi.updateSchema([
    {
      componentProps: {
        /** 显示内容,对应值 */
        fieldNames: { label: 'nickName', value: 'userId' },
        showSearch: true,
        /** 结构树data */
        options: userData.items,
        /** 默认展示全部 */
        treeDefaultExpandAll: true,
        // 选中后显示在输入框的值
        displayRender: "nickName"
      },
      /** 要修改的字段 */
      fieldName: 'userIdList',
    },
  ]);
}
loadUserOptions()


/** 选择用户 */
const refSelectDevice = ref();
const userList = ref([])
function openSelectUser() {
  refSelectDevice.value.openModal();
}
async function onUserSelected(userIdList: any) {
  userList.value = userIdList;
  const userIds = userIdList.map((item: any) => item.userId);
  // 设置用户相关字段
  await formApi.setFieldValue('userIdList', userIds);
}
</script>

<template>
  <BasicDrawer :close-on-click-modal="false" :title="title" class="w-[800px]">
    <BasicForm>
      <template #userIdList="slotProps">
        <Select :option="userList" mode="multiple" fieldNames="{ label: nickName, value: userId}" placeholder="请选择用户" v-bind="slotProps"
          readonly style="width: 100%;flex:1" />
        <Button style="margin-left: 8px;" v-if="!isUpdate && !isView" type="primary" @click="openSelectUser">选择用户</Button>
      </template>
    </BasicForm>
    <SelectDevice ref="refSelectDevice" @userSelected="onUserSelected" />

  </BasicDrawer>
</template>
