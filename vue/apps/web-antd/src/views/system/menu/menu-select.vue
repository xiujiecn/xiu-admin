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
import { Tag, Checkbox } from 'ant-design-vue';
import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysMenuListApi } from '#/api/system';
import { getMenuTypeOptionsLabel, getMenuTypeOptionsColor } from './model';
import { cloneDeep } from 'lodash-es';
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
  menuIds?: number[];
}>();
const resList = ref([]);

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

function getData() {
  const select = gridApi.grid.getCheckboxRecords().map((item) => item.menuId);
  const select2 = gridApi.grid
    .getCheckboxIndeterminateRecords()
    .map((item) => item.menuId);
  const select3 = select.concat(select2);
  return select3;
}
function clearData() {
  gridApi.grid.clearCheckboxRow();
}

// onUnmounted(() => {
//   console.log('unmounted');
// });
// onMounted(() => {
//   console.log('mounted');
// });

defineExpose({
  getData,
  clearData,
});
/** ====================  列表实例  ==================== */
const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions: {
    checkboxConfig: {
      // 关闭父子联动
      checkStrictly: true,
    },
    rowConfig: {
      keyField: 'menuId',
    },
    align: 'center',
    columns: [
      { align: 'left', title: '', type: 'checkbox', width: 40 },
      {
        field: 'menuName',
        title: '菜单名称',
        treeNode: true,
        minWidth: 160,
        align: 'left',
      },
      {
        field: 'menuType',
        title: '类型',
        width: 60,
        slots: { default: 'menuType' },
      },
      { field: 'perms', title: '权限标识' },
    ],
    height: 'auto',
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
  },
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
  <Grid v-if="flag" class="m-0 w-full p-0">
    <template #menuType="{ row }">
      <Tag :color="getMenuTypeOptionsColor(row.menuType)">{{
        getMenuTypeOptionsLabel(row.menuType)
      }}</Tag>
    </template>
  </Grid>
</template>
