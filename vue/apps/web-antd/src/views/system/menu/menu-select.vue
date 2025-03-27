<!--
 * @description 菜单选择组件
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 XiuAdmin CLI
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script lang="ts" setup>
import { onMounted, ref, watch, nextTick } from 'vue';
import type { DeepPartial } from '@vben/types';
import type { VxeTableGridOptions, VxeGridListeners } from '#/adapter/vxe-table';
import type { SysMenuListData } from '#/api/system';

import { Tag, CheckboxGroup } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysMenuListApi } from '#/api/system';
import { getMenuTypeOptionsLabel, getMenuTypeOptionsColor } from './model';


interface RowType {
  category: string;
  color: string;
  id: string;
  price: string;
  productName: string;
  releaseDate: string;
}

const emit = defineEmits(['change']);
const props = defineProps<{
  menuIds?: number[];
  checkStrictly?: boolean;
}>();
const checkedMenuIds = ref<number[]>(props.menuIds || []);
const checkStrictly = ref(props.checkStrictly || false);

const gridEvents: DeepPartial<VxeGridListeners> = {
  checkboxChange: handleCheckboxChange,
  checkboxAll: handleCheckboxChange
}

onMounted(() => {
  expandAll();
});

const gridOptions: VxeTableGridOptions<RowType> = {
  checkboxConfig: {
    checkRowKeys: checkedMenuIds.value,
    checkStrictly: checkStrictly.value,
    reserve: true,
  },
  rowConfig: {
    keyField: 'menuId',
  },
  align: 'center',
  columns: [
    { align: 'left', title: '', type: 'checkbox', width: 40 },
    { field: 'menuName', title: '菜单名称', treeNode: true, minWidth: 160, align: 'left', },
    { field: 'menuType', title: '类型', width: 60, slots: { default: 'menuType' } },
    { field: 'perms', title: '权限标识' },
  ],
  exportConfig: {},
  height: 'auto',
  keepSource: true,
  pagerConfig: {
    enabled: false,
  },
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {

        return await getSysMenuListApi({
          page: page.currentPage,
          pageSize: page.pageSize,
          ...formValues,
        });
      },
      querySuccess: () => {
        nextTick();
        expandAll();

        gridApi.grid.setCheckboxRowKey(checkedMenuIds.value, true);
        // gridApi.grid.reloadData(gridApi.grid.getData());
        console.log('vue/apps/web-antd/src/views/system/menu/menu-select.vue querySuccess', checkedMenuIds.value);
      },
    },
  },
  toolbarConfig: {
  },
  treeConfig: {
    parentField: 'parentId',
    rowField: 'menuId',
    transform: true,
    expandAll: true,
  },
};

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions,
  gridEvents
});

watch(props, (newVal) => {
  nextTick();
  checkedMenuIds.value = newVal.menuIds || [];
  checkStrictly.value = newVal.checkStrictly || false;

  if (checkStrictly.value) {
    // checkboxGroupValue 中删除 3
    checkboxGroupValue.value.splice(checkboxGroupValue.value.indexOf(3), 1);
  }else {
    checkboxGroupValue.value.push(3);
  }
  gridApi.setGridOptions({
    checkboxConfig: {
      checkStrictly: checkStrictly.value,
      checkRowKeys: checkedMenuIds.value,
    }
  });
  nextTick();
  const rows = gridApi.grid.getCheckboxRecords();
  const ids = rows.map((row: SysMenuListData) => row.menuId);
  if (ids.toString() != checkedMenuIds.value.toString()) {
    gridApi.grid.setCheckboxRowKey(checkedMenuIds.value, true);
    console.log('vue/apps/web-antd/src/views/system/menu/menu-select.vue watch  props setCheckboxRowKey', checkedMenuIds.value, checkStrictly.value);
  }
  nextTick();
  // gridApi.reload();
  console.log('vue/apps/web-antd/src/views/system/menu/menu-select.vue watch props end', checkedMenuIds.value, checkStrictly.value);
});
// watch(checkedMenuIds, (newVal) => {
//   const rows = gridApi.grid.getCheckboxRecords();
//   const ids = rows.map((row: SysMenuListData) => row.menuId);
//   if (ids.toString() != newVal.toString()) {
//     gridApi.grid.setCheckboxRowKey(newVal, true);
//   }
// });


function handleCheckboxChange() {
  const rows = gridApi.grid.getCheckboxRecords();
  const ids = rows.map((row: SysMenuListData) => row.menuId);
  if (ids.toString() != props.menuIds?.toString() || checkStrictly.value != props?.checkStrictly) {
    checkedMenuIds.value = ids;
    emit('change', ids, checkStrictly.value);
  }
}

const expandAll = () => {
  gridApi.grid?.setAllTreeExpand(true);
};

const collapseAll = () => {
  gridApi.grid?.setAllTreeExpand(false);
};

const expandCollapseAll = (value: boolean) => {
  gridApi.grid?.setAllTreeExpand(value);
}


const checkboxGroupOptions = [
  { label: '展开/折叠', value: 1 },
  { label: '全选/全不选', value: 2 },
  { label: '父子联动', value: 3 }
]
const checkboxGroupValue = ref<number[]>([3]);
watch(checkboxGroupValue, (newVal, oldVal) => {
  if (newVal.includes(1) && !oldVal.includes(1)) {
    expandCollapseAll(true);
  } else if (!newVal.includes(1) && oldVal.includes(1)) {
    expandCollapseAll(false);
  }
  if (newVal.includes(2) && !oldVal.includes(2)) {
    gridApi.grid?.setAllCheckboxRow(true);
    handleCheckboxChange();
  } else if (!newVal.includes(2) && oldVal.includes(2)) {
    gridApi.grid?.setAllCheckboxRow(false);
    handleCheckboxChange();
  }
  if (newVal.includes(3) && !oldVal.includes(3)) {
    gridApi.setGridOptions({
      checkboxConfig: {
        checkStrictly: false
      }
    });
    checkStrictly.value = false;
    handleCheckboxChange();
  } else if (!newVal.includes(3) && oldVal.includes(3)) {
    gridApi.setGridOptions({
      checkboxConfig: {
        checkStrictly: true
      }
    });
    checkStrictly.value = true;
    handleCheckboxChange();
  }
});

</script>

<template>

  <CheckboxGroup class="w-full pl-3 plr-3 justify-between" :options="checkboxGroupOptions"
    v-model:value="checkboxGroupValue" />
  <Grid class="w-full m-0 p-0">
    <template #menuType="{ row }">
      <Tag :color="getMenuTypeOptionsColor(row.menuType)">{{ getMenuTypeOptionsLabel(row.menuType) }}</Tag>
    </template>
  </Grid>
</template>
