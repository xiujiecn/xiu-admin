<script lang="ts" setup>
import { useVbenModal} from '@vben/common-ui';

import type {
  VxeTableGridOptions,
  VxeGridListeners,
} from '#/adapter/vxe-table';
import type { VbenFormProps } from '#/adapter/form';

import { columns, type RowType } from './model';
import type { DeepPartial } from '@vben/types';
import { getSysUserListApi } from '#/api/system/user';
import { Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';

import { ref, defineProps } from 'vue';

const emit = defineEmits<{
  // 点击刷新按钮的事件
  userSelected:any;
}>();

const currentDeptId = ref(0);

// 组织选项数据
const deptOptions = ref<Array<{ label: string; value: number }>>([]);
// 全量组织选项数据（用于保存所有组织，不受搜索结果影响）
const allDeptOptions = ref<Array<{ label: string; value: number }>>([]);

// 表单配置 - 添加筛选条件
const formOptions: VbenFormProps = {
  // 默认展开
  collapsed: false,
  schema: [
    {
      component: 'Input',
      fieldName: 'userName',
      label: '用户名',
      componentProps: {
        placeholder: '请输入用户名',
      },
    },
    {
      component: 'Select',
      fieldName: 'deptId',
      label: '组织',
      componentProps: {
        placeholder: '请选择组织',
        allowClear: true,
        options: allDeptOptions,
      },
    },
  ],
  // 控制表单是否显示折叠按钮
  showCollapseButton: true,
  // 是否在字段值改变时提交表单
  submitOnChange: true,
  // 按下回车时是否提交表单
  submitOnEnter: true,
};
const gridOptions: VxeTableGridOptions<RowType> = {

  radioConfig: {
    highlight: true,
  },
  rowConfig: {
    keyField: 'userId',
  },
  columns: columns,
  exportConfig: {},
  height: 450, // 设置固定高度，适应弹窗大小
  keepSource: true,
  showOverflow: false,
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        console.log('🔍 查询参数:', {
          page: page.currentPage,
          pageSize: page.pageSize,
          deptId: currentDeptId.value,
          ...formValues,
        });

        try {
          // 处理组织ID参数
          let deptIdParam = currentDeptId.value;
          if (formValues.deptId !== undefined && formValues.deptId !== null) {
            deptIdParam = Number(formValues.deptId) || 0;
          }

          // 构建API请求参数
          const apiParams = {
            page: page.currentPage,
            pageSize: page.pageSize,
            deptId: deptIdParam,
            userName: formValues.userName || undefined,
          };

          console.log('🚀 发送给API的参数:', apiParams);
          console.log('📋 选择的组织ID:', deptIdParam);

          // 简化参数，只传必要的分页参数
          const result = await getSysUserListApi(apiParams as any);

          console.log('📡 API响应:', result);

          // 使用 any 类型来处理实际的API响应
          const apiResult = result as any;
          console.log('📊 数据结构:', {
            hasItems: !!apiResult.items,
            itemsLength: apiResult.items?.length || 0,
            total: apiResult.total,
            firstItem: apiResult.items?.[0]
          });

          // 提取组织信息并更新组织选项
          const items = apiResult.items || [];
          const deptMap = new Map<number, string>();

          console.log('👥 返回的用户数据:', items.map((user: any) => ({
            userName: user.userName,
            deptId: user.deptInfo?.deptId,
            deptName: user.deptInfo?.deptName
          })));

          items.forEach((user: any) => {
            if (user.deptInfo && user.deptInfo.deptId && user.deptInfo.deptName) {
              deptMap.set(user.deptInfo.deptId, user.deptInfo.deptName);
            }
          });

          // 更新组织选项（只在首次加载或选择"全部"时更新）
          const newDeptOptions = Array.from(deptMap.entries()).map(([deptId, deptName]) => ({
            label: deptName,
            value: deptId
          }));

          // 如果是首次加载（没有选择特定组织）或者当前组织选项为空，则更新全量组织选项
          if (deptIdParam === 0 || allDeptOptions.value.length === 0) {
            allDeptOptions.value = [
              { label: '全部', value: 0 },
              ...newDeptOptions.sort((a, b) => a.label.localeCompare(b.label))
            ];
            console.log('🏢 更新全量组织选项:', allDeptOptions.value);
          } else {
            // 如果选择了特定组织，合并新发现的组织到全量选项中
            const existingDeptIds = new Set(allDeptOptions.value.map(opt => opt.value));
            const newDepts = newDeptOptions.filter(opt => !existingDeptIds.has(opt.value));
            if (newDepts.length > 0) {
              const allOption = allDeptOptions.value.find(opt => opt.value === 0);
              const otherOptions = allDeptOptions.value.filter(opt => opt.value !== 0);
              allDeptOptions.value = [
                allOption!, // 保留"全部"选项
                ...otherOptions.concat(newDepts).sort((a, b) => a.label.localeCompare(b.label))
              ];
              console.log('🏢 合并新组织到全量选项:', allDeptOptions.value);
            }
          }

          // 如果选择了特定组织，进行前端精确过滤
          let filteredItems = items;
          if (deptIdParam > 0) {
            filteredItems = items.filter((user: any) => {
              return user.deptInfo && user.deptInfo.deptId === deptIdParam;
            });
            console.log('🔍 组织精确过滤后的用户:', filteredItems.length, '个用户');
          }

          // 转换数据格式以适配表格
          const transformedData = {
            items: filteredItems,
            total: filteredItems.length
          };

          console.log('🔄 转换后数据:', transformedData);
          console.log('🏢 组织选项:', deptOptions.value);
          return transformedData;
        } catch (error) {
          console.error('❌ API请求失败:', error);
          return {
            items: [],
            total: 0
          };
        }
      },
    },
  },
  toolbarConfig: {
    custom: false,
    export: false,
    refresh: true,
    resizable: true,
    search: true,
    zoom: false,
  },
};
const gridEvents: DeepPartial<VxeGridListeners> = {
  radioChange: handleRadioChange,
};
const RadioSelected = ref(false);
function handleRadioChange() {
  RadioSelected.value = !!gridApi.grid.getRadioRecord();
}
const [Grid, gridApi] = useVbenVxeGrid({
  formOptions,
  gridOptions,
  gridEvents,
});
//确认按钮回调函数
const onConfirm = () => {
  const selectedRecord = gridApi.grid.getCheckboxRecords();
  if (!selectedRecord) {
    console.warn('请选择一个用户');
    return;
  }

  console.log(selectedRecord)

  emit('userSelected',selectedRecord);
  modalApi.close();
};

function openModal() {
  modalApi.open();
}

defineExpose({
  openModal,
});

const [Modal, modalApi] = useVbenModal({ onConfirm: onConfirm ,
  onCancel: () => {
    modalApi.close();
  },
  onOpenChange: (isOpen) => {
    const {deptId} = modalApi.getData() || {};
    if (deptId) {
      currentDeptId.value = deptId;
    }
    if (isOpen) {
      return null;
    }
    modalApi.close();
  },
});
</script>
<template>
  <div>
    <Modal class="w-[1000px] h-[600px]" title="选择用户">
      <div class="table-wrapper">
        <Grid table-title="用户列表">
          <template #status="{ row }">
            <Tag :color="row.status === '0' ? 'green' : 'red'">
              {{ row.status === '0' ? '正常' : '停用' }}
            </Tag>
          </template>
          <template #action="{ row }">
            <div class="flex items-center justify-center">
              <!-- 用户选择弹窗中的操作列，可以留空或添加查看按钮 -->
              <span class="text-gray-400">-</span>
            </div>
          </template>
        </Grid>
      </div>
    </Modal>
  </div>
</template>
<style scoped>
/* 表格包装器 */
.table-wrapper {
  width: 100%;
  height: 100%;
  padding: 0;
  margin: 0;
}

/* 表格容器样式 */
.table-wrapper :deep(.vxe-grid) {
  border: none;
  height: 100%;
}

/* 表格主体样式 */
.table-wrapper :deep(.vxe-table) {
  height: 100%;
}

/* 表格内容区域 */
.table-wrapper :deep(.vxe-table--body-wrapper) {
  max-height: 400px;
  overflow-y: auto;
}

/* 表格工具栏 */
.table-wrapper :deep(.vxe-toolbar) {
  padding: 8px 16px;
  border-bottom: 1px solid #f0f0f0;
  background: #fafafa;
}

/* 表格分页器 */
.table-wrapper :deep(.vxe-pager) {
  padding: 8px 16px;
  border-top: 1px solid #f0f0f0;
  background: #fafafa;
}

/* 确保表格字段充满宽度 */
.table-wrapper :deep(.vxe-table--header),
.table-wrapper :deep(.vxe-table--body) {
  width: 100% !important;
}

/* 美化滚动条 */
.table-wrapper :deep(.vxe-table--body-wrapper)::-webkit-scrollbar {
  width: 6px;
}

.table-wrapper :deep(.vxe-table--body-wrapper)::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.table-wrapper :deep(.vxe-table--body-wrapper)::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

.table-wrapper :deep(.vxe-table--body-wrapper)::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}
</style>
