<script lang="ts" setup>
import { h,onMounted,ref, watch } from 'vue';
import type { DeepPartial } from '@vben/types';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions, VxeGridListeners } from '#/adapter/vxe-table';
import type { SysMenuListData } from '#/api/system';

import { Page } from '@vben/common-ui';

import {message,Tag, CheckboxGroup} from 'ant-design-vue';
import dayjs from 'dayjs';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysMenuListApi,deleteSysMenuApi } from '#/api/system'; 
import { Icon } from '@iconify/vue';
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
  },
  rowConfig: {
    keyField: 'menuId',
  },
  align: 'center',
  columns: [
    { align: 'left', title: '', type: 'checkbox', width: 40 },
    { field: 'menuName', title: '菜单名称' , treeNode: true, minWidth: 160,  align: 'left', },
    { field: 'menuType', title: '类型' ,width: 60, slots: { default: 'menuType' } },
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
        expandAll();
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

const [Grid,gridApi] = useVbenVxeGrid({
  gridOptions,
  gridEvents
});

watch(checkedMenuIds, (newVal) => {
  const rows = gridApi.grid.getCheckboxRecords();
  const ids = rows.map((row: SysMenuListData) => row.menuId);
  if (ids.toString() != newVal.toString()) {
    gridApi.grid.setCheckboxRowKey(newVal, true);
  }
});


function  handleCheckboxChange() {
  const rows = gridApi.grid.getCheckboxRecords();
  const ids = rows.map((row: SysMenuListData) => row.menuId);
  if (ids.toString() != checkedMenuIds.value.toString()) {
    checkedMenuIds.value = ids;
    emit('change',ids, checkStrictly.value);
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
  {label: '展开/折叠', value: 1},
  {label: '全选/全不选', value: 2},
  {label: '父子联动', value:3 }
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
  } else if (!newVal.includes(3) && oldVal.includes(3)) {
    gridApi.setGridOptions({
      checkboxConfig: {
        checkStrictly: true
      }
    });
    checkStrictly.value = true;
  }
});

</script>

<template>

    <CheckboxGroup class="w-full pl-3 plr-3 justify-between" :options="checkboxGroupOptions" v-model:value="checkboxGroupValue" />
    <Grid class="w-full m-0 p-0" >
      <template #menuType="{ row }"> 
        <Tag :color="getMenuTypeOptionsColor(row.menuType)">{{ getMenuTypeOptionsLabel(row.menuType) }}</Tag>
      </template>
    </Grid>
</template>
