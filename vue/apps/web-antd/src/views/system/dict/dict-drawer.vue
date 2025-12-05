<!--
 * @description 字典类型编辑抽屉组件
 * @Link  https://github.com/xiujie
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script setup lang="ts">
import type { SysMenuListData } from '#/api/system/menu';

import { computed, nextTick, ref } from 'vue';
import { AccessControl } from '@vben/access';
import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { addFullName, cloneDeep } from '@vben/utils';
import { Radio, RadioGroup } from  'ant-design-vue';
import { useVbenForm } from '#/adapter/form';
  import { getDictOptions } from '#/utils/dict';
  import { DictEnum } from '@vben/constants';
import { addSysDictTypeApi, editSysDictTypeApi, getSysDictTypeViewApi } from '#/api/system/dict';
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
    if (isUpdate.value || isView.value) {
      const record = await getSysDictTypeViewApi({ dictId: Number(id) });
      console.log('record', record);
      // 回显 isSys 字段，接口返回 '0' 显示“是”，'1' 显示“否”
      if (record && typeof record.isSys !== 'undefined') {
        record.isSys = record.isSys === '0' ? 'Y' : 'N';
      }
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

    // 转换 isSys 字段：Y -> 0（是），N -> 1（否）
    if (data.isSys) {
      data.isSys = data.isSys === 'Y' ? '0' : '1';
    }

    await (isUpdate.value ? editSysDictTypeApi (data) : addSysDictTypeApi(data));

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
      <!-- 这里判断用户身份，并设置字段的可编辑性  isSys    ['cpr:superadmin']    -->
        <template #isSys="slotProps">
          <AccessControl :codes="['cpr:superadmin']" type="code" v-slot="{ hasPermission }">
            <div>
              <RadioGroup
                v-bind="slotProps"
                :options="getDictOptions(DictEnum.SYS_YES_NO)"
                buttonStyle="solid"
                optionType="button"
                :disabled="!hasPermission || isView"
                @change="(value) => {
                  if (!isView && slotProps.onChange) {
                    slotProps.onChange(value);
                  }
                }"
              />
            </div>
          </AccessControl>
        </template>
    </BasicForm>
  </BasicDrawer>
</template>
