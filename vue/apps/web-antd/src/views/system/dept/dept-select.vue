<script lang="ts" setup>
import { onMounted, ref, watch, nextTick } from 'vue';
import type { DeepPartial } from '@vben/types';
import type { VxeTableGridOptions, VxeGridListeners } from '#/adapter/vxe-table';
import type { SysDeptListData } from '#/api/system';
import { CheckboxGroup } from 'ant-design-vue';
import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysDeptListApi } from '#/api/system';


interface RowType {
  deptId: number;
  deptName: string;
  parentId: number;
  orderNum: number;
  status: string;
  children?: RowType[];
}

const emit = defineEmits(['change']);
const props = defineProps<{
  deptIds?: number[];
  checkStrictly?: boolean;
}>();
const checkedDeptIds = ref<number[]>(props.deptIds||[] );
const checkStrictly = ref(props.checkStrictly||false);

watch(props, (newVal) => {
  nextTick();
  checkedDeptIds.value = newVal.deptIds||[];
  checkStrictly.value = newVal.checkStrictly||false;
  gridApi.setGridOptions({
    checkboxConfig: {
      checkStrictly: checkStrictly.value
    }
  });
  nextTick();
  if (checkStrictly.value) {
    // checkboxGroupValue 中删除 3
    checkboxGroupValue.value.splice(checkboxGroupValue.value.indexOf(3), 1);
  }else {
    checkboxGroupValue.value.push(3);
  }
  const rows = gridApi.grid.getCheckboxRecords();
  const ids = rows.map((row: SysDeptListData) => row.deptId);
  if (ids.toString() != props.deptIds?.toString()) {
    gridApi.grid.setCheckboxRowKey(checkedDeptIds.value, true);
  }
});

const gridEvents: DeepPartial<VxeGridListeners> = {
  checkboxChange: handleCheckboxChange,
  checkboxAll: handleCheckboxChange
}

onMounted(() => {
  expandAll();
  console.log('vue/apps/web-antd/src/views/system/dept/dept-select.vue onMounted', checkedDeptIds.value, checkStrictly.value);
});

const gridOptions: VxeTableGridOptions<RowType> = {
  checkboxConfig: {
    checkRowKeys: checkedDeptIds.value,
    checkStrictly: checkStrictly.value,
  },
  rowConfig: {
    keyField: 'deptId',
  },
  align: 'center',
  columns: [
    { align: 'left', title: '', type: 'checkbox', width: 40 },
    { field: 'deptName', title: '部门名称', treeNode: true, minWidth: 160, align: 'left', },
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

        return await getSysDeptListApi({
          page: page.currentPage,
          pageSize: page.pageSize,
          ...formValues,
        });
      },
      querySuccess: () => {
        expandAll();
        nextTick();
        gridApi.grid.setCheckboxRowKey(checkedDeptIds.value, true);
        
      },
    },
  },
  toolbarConfig: {
  },
  treeConfig: {
    parentField: 'parentId',
    rowField: 'deptId',
    transform: true,
    expandAll: true,
  },
};

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions,
  gridEvents
});

watch(checkedDeptIds, (newVal) => {
  nextTick();
  const rows = gridApi.grid.getCheckboxRecords();
  const ids = rows.map((row: SysDeptListData) => row.deptId);
  if (ids.toString() != newVal.toString()) {
    gridApi.grid.setCheckboxRowKey(newVal, true);
  }
});


function handleCheckboxChange() {
  const rows = gridApi.grid.getCheckboxRecords();
  const ids = rows.map((row: SysDeptListData) => row.deptId);
  if (ids.toString() != checkedDeptIds.value.toString() || checkStrictly.value != props.checkStrictly) {
    checkedDeptIds.value = ids;
    emit('change', ids, checkStrictly.value);
  }
}

const expandAll = () => {
  gridApi.grid?.setAllTreeExpand(true);
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
  </Grid>
</template>
