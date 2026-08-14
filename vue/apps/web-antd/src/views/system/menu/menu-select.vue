<!--
 * @description 菜单选择组件
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->

<script lang="ts" setup>
import { nextTick, ref, watch } from 'vue';
import type { Ref } from 'vue';
import type { DeepPartial } from '@vben/types';
import type { VxeGridListeners } from '#/adapter/vxe-table';
import { Tag, Checkbox, Select } from 'ant-design-vue';
import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysMenuListApi } from '#/api/system';
import type { SysMenuListData } from '#/api/system';
import { getMenuTypeOptionsLabel, getMenuTypeOptionsColor } from './model';

/** 右键菜单项 code */
const MenuCode = {
  SelectSelf: 'selectSelf',
  SelectSelfAndChildren: 'selectSelfAndChildren',
  UnselectSelf: 'unselectSelf',
  UnselectSelfAndChildren: 'unselectSelfAndChildren',
} as const;

/** 递归获取节点及其所有子孙节点 */
function getDescendants(tree: SysMenuListData[], parentId: number, field: keyof SysMenuListData, parentField: keyof SysMenuListData): SysMenuListData[] {
  const children = tree.filter((item) => item[parentField] === parentId);
  let result = [...children];
  children.forEach((child) => {
    result = result.concat(getDescendants(tree, child[field] as number, field, parentField));
  });
  return result;
}
/** ====================  tips:  ==================== */
/**
 * @description Tips
 * @author 幻天
 * 1.gridApi.setGridOptions({
    checkboxConfig: {
      checkStrictly: !params.target.checked,
    },
  });
  此方法:默认勾选指定行（只会在初始化时被触发一次，需要有 row-config.keyField）
 * 2.该组件在弹窗关闭时会被卸载,即触发onUnmounted
 * 3.菜单有缓存
 * 4.想要更新勾选项,必须通过重新挂载Grid来触发checkStrictly
 * 5.菜单关闭时,外部必须置空menuIds,否则第二次打开时会因为新menuIds正在请求,界面会显示旧的menuIds
  外部置空menuIds后,会重载一次无勾选的Grid,防止显示旧的menuIds
 */
/**
 * @description 现有流程
 * @author 幻天
 *
 * 1.第一次打开抽屉
 * 2.外部初始化,获取数据menuIds
 * 3.该组件开始挂载
 * 4.Grid开始挂载,请求菜单
 * 5.请求成功后,设置勾选项
 *
 * 1.多次打开抽屉
 * 2.外部已初始化,组件初始化
 * 3.外部获取menuIds
 * 4.内部监听menuIds变化,重新挂载Grid
 * 5.Grid开始挂载,设置勾选项
 *
 */

/**
 * @author 幻天
 * @todo flag默认为false,同时watch添加挂载监听一次,再改一些代码,是否可以避免外部置空menuIds导致的不必要挂载
 */

/** ====================    ==================== */
const flag = ref(true);
/** props */
const props = defineProps<{
  dataScopes?: Record<number, string>;
  menuIds?: number[];
}>();
const resList: Ref<SysMenuListData[]> = ref([]);
const dataScopeMap = ref<Record<number, string>>({});
const menuDataScopeOptions = [
  { label: '按角色数据权限', value: '0' },
  { label: '全部数据权限', value: '1' },
  { label: '本部门数据权限', value: '3' },
  { label: '本部门及以下数据权限', value: '4' },
  { label: '仅本人数据权限', value: '5' },
  { label: '本部门及以下或本人数据权限', value: '6' },
  { label: '本组织及本组织下一级数据权限', value: '7' },
  { label: '本组织下一级数据权限', value: '8' },
];

/** 监听 menuIds 变化并重新加载 Grid */
watch(
  () => props.menuIds,
  async () => {
    // 先销毁 Grid
    flag.value = false;
    // 等待 DOM 更新完成
    await nextTick();
    // 再次等待确保组件完全销毁
    await nextTick();
    // 重新创建 Grid
    flag.value = true;
  },
  { deep: true },
);
watch(
  () => props.dataScopes,
  (value) => {
    dataScopeMap.value = { ...(value || {}) };
  },
  { deep: true, immediate: true },
);

function getData() {
  const select = gridApi.grid.getCheckboxRecords().map((item) => item.menuId);
  const select2 = gridApi.grid
    .getCheckboxIndeterminateRecords()
    .map((item) => item.menuId);
  const select3 = select.concat(select2);
  return select3;
}
function getDataScopes() {
  const selectedMenuIds = getData();
  const result: Record<number, string> = {};
  selectedMenuIds.forEach((menuId) => {
    result[menuId] = dataScopeMap.value[menuId] || '0';
  });
  return result;
}
function handleDataScopeChange(menuId: number, value: string) {
  dataScopeMap.value[menuId] = value;
}
function clearData() {
  gridApi.grid.clearCheckboxRow();
  dataScopeMap.value = {};
}

defineExpose({
  getData,
  getDataScopes,
  clearData,
});
/** ====================  列表实例  ==================== */
const gridOptions = {
  checkboxConfig: {
    // 关闭父子联动
    checkStrictly: true,
  },
  rowConfig: {
    keyField: 'menuId',
  },
  align: 'center' as const,
  columns: [
    { align: 'left' as const, title: '', type: 'checkbox', width: 40 },
    {
      field: 'menuName',
      title: '菜单名称',
      treeNode: true,
      minWidth: 160,
      align: 'left' as const,
    },
    {
      field: 'menuType',
      title: '类型',
      width: 60,
      slots: { default: 'menuType' },
    },
    { field: 'perms', title: '权限标识' },
    {
      field: 'dataScope',
      title: '数据范围',
      width: 210,
      slots: { default: 'dataScope' },
    },
  ],
  height: 'auto' as const,
  keepSource: true,
  pagerConfig: {
    enabled: false,
  },
  proxyConfig: {
    ajax: {
      query: async ({ page }: any, formValues: any) => {
        const res = await getSysMenuListApi({
          page: page.currentPage,
          pageSize: page.pageSize,
          ...formValues,
        });
        resList.value = res.items;
        return res;
      },
      querySuccess: () => {
        gridApi.setGridOptions({
          checkboxConfig: {
            checkRowKeys: props.menuIds,
          },
        });
      },
    },
  },
  treeConfig: {
    parentField: 'parentId',
    rowField: 'menuId',
    transform: true,
    expandAll: true,
  },
  menuConfig: {
    body: {
      options: [
        [
          { code: MenuCode.SelectSelf, name: '选择本级' },
          { code: MenuCode.SelectSelfAndChildren, name: '选择本级及下级' },
        ],
        [
          { code: MenuCode.UnselectSelf, name: '不选择本级' },
          { code: MenuCode.UnselectSelfAndChildren, name: '不选择本级及下级' },
        ],
      ],
    },
  },
};

/** 设置行及其子孙节点的勾选状态 */
function setRowAndDescendantsChecked(row: SysMenuListData, checked: boolean) {
  const descendants = getDescendants(resList.value, row.menuId, 'menuId', 'parentId');
  [row, ...descendants].forEach((r) => {
    gridApi.grid.setCheckboxRow(r, checked);
  });
}

/** 右键菜单点击事件 */
const gridEvents: DeepPartial<VxeGridListeners> = {
  menuClick: ({ menu, row }: any) => {
    if (!row) return;
    const { code } = menu;
    switch (code) {
      case MenuCode.SelectSelf:
        gridApi.grid.clearCheckboxRow();
        gridApi.grid.setCheckboxRow(row, true);
        break;
      case MenuCode.SelectSelfAndChildren:
        gridApi.grid.clearCheckboxRow();
        setRowAndDescendantsChecked(row, true);
        break;
      case MenuCode.UnselectSelf:
        gridApi.grid.setCheckboxRow(row, false);
        break;
      case MenuCode.UnselectSelfAndChildren:
        setRowAndDescendantsChecked(row, false);
        break;
    }
  },
};

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions,
  gridEvents,
});

/** ====================  多选组配置  ==================== */
const checkboxStatus = ref({
  all: false,
  expand: true,
  linkage: false,
});
/** 全选事件 */
function checkboxAllChange(params: any) {
  gridApi.grid.setAllCheckboxRow(params.target.checked);
}
/** 展开折叠事件 */
function checkboxExpandChange(params: any) {
  gridApi.grid?.setAllTreeExpand(params.target.checked);
}
/** 父子联动事件 */
function checkboxLinkageChange(params: any) {
  gridApi.setGridOptions({
    checkboxConfig: {
      checkStrictly: !params.target.checked,
    },
  });
}
</script>

<template>
  <div class="flex h-full flex-col">
    <div class="flex justify-between px-2">
      <Checkbox
        :disabled="!flag"
        @change="checkboxAllChange"
        v-model:checked="checkboxStatus.all"
        >全选/全不选</Checkbox
      >
      <Checkbox
        :disabled="!flag"
        @change="checkboxExpandChange"
        v-model:checked="checkboxStatus.expand"
        >展开/折叠</Checkbox
      >
      <Checkbox
        :disabled="!flag"
        @change="checkboxLinkageChange"
        v-model:checked="checkboxStatus.linkage"
        >父子联动</Checkbox
      >
    </div>
    <Grid v-if="flag" class="m-0 w-full flex-1 p-0">
      <template #menuType="{ row }">
        <Tag :color="getMenuTypeOptionsColor(row.menuType)">{{
          getMenuTypeOptionsLabel(row.menuType)
        }}</Tag>
      </template>
      <template #dataScope="{ row }">
        <Select
          :value="dataScopeMap[row.menuId] || '0'"
          @change="(value) => handleDataScopeChange(row.menuId, value as string)"
          :options="menuDataScopeOptions"
          size="small"
          class="w-full text-left"
        />
      </template>
    </Grid>
  </div>
</template>
