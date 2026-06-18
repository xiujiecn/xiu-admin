import { h, ref } from 'vue';
import { Tag } from 'ant-design-vue';
import type { VxeGridProps } from '#/adapter/vxe-table';
import { getPopupContainer } from '@vben/utils';
import type { DescItem } from '#/components/description';
import { cloneDeep } from 'lodash-es';
import type { VbenFormSchema } from '@vben/common-ui';
import dayjs from 'dayjs';
import { z } from '@vben/common-ui';
import { renderPopoverMemberSumma, type MemberSumma } from '#/utils';
import { TreeOption } from '#/api/gen/testTree';

export class State {
  public id = 0; // 主键
  public tenantId = '000000'; // 租户编号
  public parentId = 0; // 父id
  public deptId = 0; // 部门id
  public userId = 0; // 用户id
  public treeName = ''; // 值
  public level = 1; // 关系树等级
  public tree = ''; // 关系树
  public version = 0; // 版本
  public createdDept = 0; // 创建部门
  public createdAt = ''; // 创建时间
  public createdBy = 0; // 创建者
  public createdBySumma?: null | MemberSumma = null; // 创建者摘要信息
  public updatedAt = ''; // 更新时间
  public updatedBy = 0; // 更新者
  public updatedBySumma?: null | MemberSumma = null; // 更新者摘要信息
  public deletedBy = 0; // 删除人
  public deletedBySumma?: null | MemberSumma = null; // 删除人摘要信息
  public deletedAt = ''; // 删除时间

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) {
      return cloneDeep(state);
    }
    return new State(state);
  }
  return new State();
}

// 表单验证规则

// 表格搜索表单
export const querySchema: VbenFormSchema[] = [
  {
				fieldName: 'id',
				component: 'InputNumber',
				label: '主键',
				componentProps: {
					placeholder: '请输入主键',
					onUpdateValue: (e: any) => {
						console.log(e);
					},   
				},  rules:null,
				formItemClass: 'col-span-1',
			},  {
				fieldName: 'createdAt',
				component: 'RangePicker',
				label: '创建时间',
				componentProps: {
					type: 'datetimerange',
					clearable: true,
					valueFormat: 'FMDateRange',
					onUpdateValue: (e: any) => {
						console.log(e);
					},   
				},  rules:null,
				formItemClass: 'col-span-1',
			},];

// 表格列
export const columns: VxeGridProps['columns'] = [
  {
    title: '主键',
    field: 'id',
    align: 'left',
    width: -1,
    type: 'checkbox',
  },
  {
    title: '值',
    field: 'treeName',
    align: 'left',
    width: -1,
    treeNode: true,
 },
  {
    title: '关系树等级',
    field: 'level',
    align: 'left',
    width: -1,
 },
  {
    title: '创建部门',
    field: 'createdDept',
    align: 'left',
    width: -1,
 },
  {
    title: '创建时间',
    field: 'createdAt',
    align: 'left',
    width: -1,
 },
  {
    title: '创建者',
    field: 'createdBy',
    align: 'left',
    width: -1,
    slots: {
      default: ({ row }) =>  {
      return renderPopoverMemberSumma(row.createdBySumma);
    },
 },
 },
  {
    title: '更新时间',
    field: 'updatedAt',
    align: 'left',
    width: -1,
 },
  {
    title: '更新者',
    field: 'updatedBy',
    align: 'left',
    width: -1,
    slots: {
      default: ({ row }) =>  {
      return renderPopoverMemberSumma(row.updatedBySumma);
    },
 },
 },
  {
    title: '删除人',
    field: 'deletedBy',
    align: 'left',
    width: -1,
    slots: {
      default: ({ row }) =>  {
      return renderPopoverMemberSumma(row.deletedBySumma);
    },
 },
 },
  { title: '操作', width: 120, slots: { default: 'action' } },
];

// 表格列接口
export interface RowType {
  id: number;
  tenantId: string;
  parentId: number;
  deptId: number;
  userId: number;
  treeName: string;
  level: number;
  tree: string;
  version: number;
  createdDept: number;
  createdAt: string;
  createdBy: number;
  updatedAt: string;
  updatedBy: number;
  deletedBy: number;
  deletedAt: string;
};

// 查看字段列表
export const viewSchema: DescItem[] = [
  {  field: 'id',  label: '主键'},
  {  field: 'tenantId',  label: '租户编号'},
  {  field: 'parentId',  label: '父id'},
  {  field: 'deptId',  label: '部门id'},
  {  field: 'userId',  label: '用户id'},
  {  field: 'treeName',  label: '值'},
  {  field: 'level',  label: '关系树等级'},

  {  field: 'version',  label: '版本'},
  {  field: 'createdDept',  label: '创建部门'},
  {  field: 'createdAt',  label: '创建时间'},
  {  field: 'createdBy',  label: '创建者'},
  {  field: 'updatedAt',  label: '更新时间'},
  {  field: 'updatedBy',  label: '更新者'},
  {  field: 'deletedBy',  label: '删除人'},

];

// 编辑字段列表
export const editSchema: VbenFormSchema[] = [
 {
					fieldName: 'id',
					component: 'Input',
					label: '主键',
					dependencies: {   show: () => false,    triggerFields: [''],   },
					componentProps: {
						placeholder: '',
						onUpdateValue: (e: any) => {
							console.log(e);
						},   
					},
					formItemClass: 'col-span-1',
				},  {
				fieldName: 'tenantId',
				component: 'Input',
				label: '租户编号',
				componentProps: {
					placeholder: '请输入租户编号',
					onUpdateValue: (e: any) => {
						console.log(e);
					},   
				},  
				rules:null,
				formItemClass: 'col-span-1',
			},  {
			fieldName: 'parentId',
			component: 'TreeSelect',
			label: '父id',
			defaultValue: 0,
			componentProps: {
				allowClear: true,
				fieldNames: { label: 'treeName', value: 'id' },
				showSearch: true,
				treeData: [],
				treeDefaultExpandAll: true,
				treeLine: { showLeafIcon: false },
				treeNodeFilterProp: 'treeName',
			},
			rules:null,
			formItemClass: 'col-span-1',
		},
		  {
				fieldName: 'deptId',
				component: 'InputNumber',
				label: '部门id',
				componentProps: {
					placeholder: '请输入部门id',
					onUpdateValue: (e: any) => {
						console.log(e);
					},   
				},  rules:null,
				formItemClass: 'col-span-1',
			},  {
				fieldName: 'userId',
				component: 'InputNumber',
				label: '用户id',
				componentProps: {
					placeholder: '请输入用户id',
					onUpdateValue: (e: any) => {
						console.log(e);
					},   
				},  rules:null,
				formItemClass: 'col-span-1',
			},  {
				fieldName: 'treeName',
				component: 'Input',
				label: '值',
				componentProps: {
					placeholder: '请输入值',
					onUpdateValue: (e: any) => {
						console.log(e);
					},   
				},  
				rules:null,
				formItemClass: 'col-span-1',
			},  {
				fieldName: 'version',
				component: 'InputNumber',
				label: '版本',
				componentProps: {
					placeholder: '请输入版本',
					onUpdateValue: (e: any) => {
						console.log(e);
					},   
				},  rules:null,
				formItemClass: 'col-span-1',
			},];

// 关系树选项
export const treeOption = ref([]);

// 加载关系树选项
export async function loadTreeOption() {
  const res = await TreeOption();
  treeOption.value = res;
  return res;
}
