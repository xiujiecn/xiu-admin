<script lang="ts" setup>
import { ref, defineAsyncComponent, onMounted, shallowRef, markRaw,h } from 'vue';
import { message, Tabs, Card, Button, Space, Table, Modal } from 'ant-design-vue';
import { useVbenForm } from '#/adapter/form';
import { useRoute } from 'vue-router';
import { isJsonString } from '#/utils/is';
import EditMasterCell from './components/EditMasterCell.vue';
import EditSlaveCell from './components/EditSlaveCell.vue';
import { getSysGenTableViewApi, getGenCodesSelectsApi,type GenCodesPreviewRes, 
  type SysGenTableOptionsModel, type SysGenTableJoinModel, getGenCodesTableSelectApi, editSysGenTableApi, 
  postGenCodesPreviewApi, postGenCodesBuildApi, type SysGenTableViewRes } from '#/api/gen_codes/gen_table';
import { getSysMenuListApi } from '#/api/system/menu';
import { developBaseSchema, developJoinSchema, getSelectList, setSelectListObj, newState, genInfoObj } from './model';
import { $t } from '@vben/locales';
import { cloneDeep,merge } from 'lodash-es';
import {
  addFullName,
  getPopupContainer,
  listToTree,
} from '@vben/utils';
import PreviewTab from './components/PreviewTab.vue';

const loading = ref(true);
const route = useRoute();
const tableId = ref(route.query.tableId);
const genInfo = ref(newState(null));
const previewModel = ref<GenCodesPreviewRes>({ views: {} });
const showPreviewModal = ref(false);
async function setupMenuSelect() {
  const menuArray = await getSysMenuListApi();
  menuArray.items.forEach((item) => {
    item.menuName = $t(item.menuName);
  });
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

  baseFormApi.updateSchema([
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
      fieldName: 'options.menu.pid',
    },
  ]);
}
const loadSelectList = async () => {

  const genTypeOptions = getSelectList('genType');
  const genTypeValue = genTypeOptions?.[0]?.value;
  const dbOptions = getSelectList('db');

  baseFormApi.updateSchema([
    {
      componentProps: {
        options: genTypeOptions,
        defaultValue: genTypeValue,
      },
      fieldName: 'genType',
    },
    {
      componentProps: {
        options: dbOptions,
      },
      fieldName: 'dbName',
    },
    {
      componentProps: {
        options: [],
      },
      fieldName: 'tableName',
    },
  ]);
  baseFormApi.setFieldValue('genType', genTypeValue);

}


async function getTableInfo() {
  const res = await getSysGenTableViewApi({
    tableId: parseInt(tableId.value as string),
  });
  // 导入生成选项
  if (isJsonString(res.options)) {
    res.options = JSON.parse(res.options as string);
  }
  if (!(res.options as SysGenTableOptionsModel).headOps) {
    (res.options as SysGenTableOptionsModel).headOps = genInfoObj.options.headOps;
  }
  if (!(res.options as SysGenTableOptionsModel).columnOps) {
    (res.options as SysGenTableOptionsModel).columnOps = genInfoObj.options.columnOps;
  }
  if (!(res.options as SysGenTableOptionsModel).autoOps) {
    (res.options as SysGenTableOptionsModel).autoOps = genInfoObj.options.autoOps;
  }
  if (!(res.options as SysGenTableOptionsModel).join) {
    (res.options as SysGenTableOptionsModel).join = genInfoObj.options.join;
  }
  if (!(res.options as SysGenTableOptionsModel).menu) {
    (res.options as SysGenTableOptionsModel).menu = genInfoObj.options.menu;
  }
  if (!(res.options as SysGenTableOptionsModel).funcDict) {
    (res.options as SysGenTableOptionsModel).funcDict = genInfoObj.options.funcDict;
  }
  
  // 预设流程
  if (!(res.options as SysGenTableOptionsModel).presetStep) {
    (res.options as SysGenTableOptionsModel).presetStep = {
      formGridCols: 1,
    };
  }
  // 树表
  if (!(res.options as SysGenTableOptionsModel).tree) {
    (res.options as SysGenTableOptionsModel).tree = {
      titleColumn: "",
      styleType: 1,
    };
  }
  genInfo.value = res;
  baseFormApi.setValues(res);
  console.log('vue/apps/web-antd/src/views/tool/gen/develop.vue getTableInfo', res);
}

const [BaseForm, baseFormApi] = useVbenForm({
  commonConfig: {
    componentProps: {
      class: 'w-full',
    },
  },
  layout: 'vertical',
  schema: developBaseSchema,
  wrapperClass: 'grid-cols-4 gap-x-4',
  showDefaultActions: false,
  handleValuesChange(values) {
    if (values.dbName) {
      joinFromList.value.forEach((item) => {
        item.Api.setFieldValue('dbName', values.dbName);
      });
    }
    if (values.tableName) {
      joinFromList.value.forEach((item) => {
        item.Api.setFieldValue('masterTableName', values.tableName);
      });
    }
  }
});

interface JoinFormItem {
  From: any;
  Api: any;
  uuid: string;
  linkTable: string;
}
const uuidIndex = ref(1);
const joinFromList = ref<JoinFormItem[]>([]);
async function AddJoinForm(values: SysGenTableJoinModel) {
  const [JoinForm, joinFormApi] = useVbenForm({
    // 所有表单项共用，可单独在表单内覆盖
    commonConfig: {
      // 所有表单项
      componentProps: {
        class: 'w-full',
      },
    },
    layout: 'vertical',
    schema: developJoinSchema,
    wrapperClass: 'grid-cols-10 gap-x-4',
    showDefaultActions: false,
    handleValuesChange: async (values) => {
      console.log('vue/apps/web-antd/src/views/tool/gen/develop.vue handleValuesChange', values);
      for (let j = 0; j < genInfo.value.options.join.length; j++) {
        if (genInfo.value.options.join[j].uuid === values.uuid) {
          if (genInfo.value.options.join[j].linkTable !== values.linkTable) {
            genInfo.value.options.join[j].columns = [];
          }
          genInfo.value.options.join[j].linkTable = values.linkTable;
          genInfo.value.options.join[j].alias = values.alias;
          genInfo.value.options.join[j].linkMode = values.linkMode;
          genInfo.value.options.join[j].field = values.field;
          genInfo.value.options.join[j].masterField = values.masterField;
          genInfo.value.options.join[j].daoName = values.daoName;
        }
      }

      for (let i = 0; i < joinFromList.value.length; i++) {
        if (joinFromList.value?.[i]?.uuid === values.uuid) {
          if (joinFromList.value?.[i]?.linkTable !== values.linkTable) {
            joinFromList.value[i].linkTable = values.linkTable;
          }
        }
      }
    }
  });
  let uuid = String(Date.now()*1000+uuidIndex.value);
  uuidIndex.value++;
  if (values.uuid) {
    uuid = values.uuid;
  }else {
    values.uuid = uuid;
    genInfo.value.options.join.push(values);
  }

  const baseValues = await baseFormApi.getValues();
  values.dbName = baseValues.dbName;
  values.masterTableName = baseValues.tableName;
  joinFormApi.setValues(values);
  joinFromList.value.push({
    From: markRaw(JoinForm),
    Api: joinFormApi,
    uuid: uuid,
    linkTable: values.linkTable,

  });
}
function InitJoinForm() {
  joinFromList.value = [];
  if ((genInfo.value.options as SysGenTableOptionsModel)) {
    const joinList = (genInfo.value.options as SysGenTableOptionsModel).join;
    for (let i = 0; i < (joinList as SysGenTableJoinModel[]).length; i++) {
      AddJoinForm(joinList?.[i] as SysGenTableJoinModel);
    }
  }
}

function handleAddJoin() {
  AddJoinForm({} as SysGenTableJoinModel);
}
function handleDeleteJoin(index: number) {
  joinFromList.value.splice(index, 1);
}
onMounted(async () => {
  const res = await getGenCodesSelectsApi({});
  await setSelectListObj(res);
  await setupMenuSelect();
  await loadSelectList();
  await getTableInfo();
  InitJoinForm();
  loading.value = false;
});

async function packGenInfo() {
  let baseValue = cloneDeep(await baseFormApi.getValues());
  console.log('vue/apps/web-antd/src/views/tool/develop.vue packGenInfo', baseValue);
  let tmp = cloneDeep(genInfo.value) as any;
  Object.keys(baseValue).forEach((key) => {
    if (key in tmp && key !== 'options') {
      tmp[key] = baseValue[key];
    }
  });
  tmp.options.autoOps = baseValue.options.autoOps;
  tmp.options.columnOps = baseValue.options.columnOps;
  tmp.options.headOps = baseValue.options.headOps;
  tmp.options.menu = baseValue.options.menu;
  tmp.options.tree = baseValue.options.tree;
  return tmp;
}

async function handleSaveConfig() {
  // console.log('vue/apps/web-antd/src/views/tool/gen/develop.vue baseFormApi', await baseFormApi.getValues());
  // console.log('vue/apps/web-antd/src/views/tool/gen/develop.vue handleSaveConfig', await packGenInfo());
  Modal.confirm({
    content: `你确定要保存生成配置吗？`,
    async onOk() {
      const params = await packGenInfo();
      console.log('vue/apps/web-antd/src/views/tool/gen/develop.vue handleSaveConfig', params);
      await editSysGenTableApi(params).then(async (res) => {
        message.success('保存成功');
      });
    },
    title: '提示',
    type: 'warning',
  });
}

async function handlePreviewCode() {
  showPreviewModal.value = false;
  console.log('vue/apps/web-antd/src/views/tool/gen/develop.vue handlePreviewCode', genInfo.value);
  previewModel.value = await postGenCodesPreviewApi(await packGenInfo());
  showPreviewModal.value = true;
} 

async function handleBuildBtn() {
  console.log('vue/apps/web-antd/src/views/tool/gen/develop.vue handleSubmit', genInfo.value);
  const params = await packGenInfo();
  
  const autoOps = (params.options as SysGenTableOptionsModel).autoOps;
  const index = autoOps?.findIndex((item: any) => item === 'forcedCover');
  let content = '你确定要生成配置吗？';
  if (index !== undefined && index >= 0) {
    content = '你确定要强制覆盖生成配置吗？(强制覆盖后，将无法恢复原配置)';
  }
  Modal.confirm({
    content: content,
    async onOk() {
      loading.value = true;
      await postGenCodesBuildApi(params);
      message.success('生成成功');
      loading.value = false;
    },
    title: '提示',
    type: 'warning',
  });
}
async function handleBuildPreview(type: string) {
  console.log('vue/apps/web-antd/src/views/tool/gen/develop.vue handleBuildPreview', type);
  await handleBuildBtn();
}

</script>

<template>
  <Card class="m-4">
    <Tabs size="large">
      <template #rightExtra>
        <Space >
          <Button type="primary" @click="handlePreviewCode" :disabled="loading"   :loading="loading" v-access:code="'cpm:tool:gen:preview'">预览代码</Button>
          <Button class="bg-green-500" @click="handleBuildBtn" :disabled="loading" :loading="loading" v-access:code="'cpm:tool:gen:code'">提交生成</Button>
          <Button type="dashed" @click="handleSaveConfig" :disabled="loading" :loading="loading" v-access:code="'cpm:tool:gen:edit'">仅保存配置</Button>
        </Space>

      </template>
      <Tabs.TabPane tab="基本信息" key="1">
        <Card title="基本设置">
          <BaseForm />
        </Card>
        <Card title="关联表设置" class="mt-3">
          <template #extra><Button type="primary" @click="handleAddJoin">添加关联表</Button></template>
          <div v-for="(item, index) in joinFromList" :key="index" class="flex items-center gap-x-2">
            <component :is="item.From" :key="index" class="m-0" />
            <Button danger @click="handleDeleteJoin(index)">删除</Button>
          </div>
        </Card>
      </Tabs.TabPane>
      <Tabs.TabPane tab="主表字段" key="2" class="overflow-auto">
        <EditMasterCell  v-model:value="genInfo" :selectList="getSelectList" />
      </Tabs.TabPane>
      <Tabs.TabPane :tab="`关联字段-${item.linkTable}`"  class="overflow-auto" v-for="(item, _) in joinFromList" :key="item.uuid">
        <EditSlaveCell  v-model:value="genInfo" :selectList="getSelectList" :uuid="item.uuid" />
      </Tabs.TabPane>
    </Tabs>
  </Card>
  <PreviewTab  :previewModel="previewModel" :showModal="showPreviewModal" @BuildPreview="handleBuildPreview" />
</template>