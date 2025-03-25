import { h, ref } from 'vue';
import { Tag } from 'ant-design-vue';
import type { VxeGridProps } from '#/adapter/vxe-table';
import { getPopupContainer } from '@vben/utils';
import type { DescItem } from '#/components/description';
import { cloneDeep } from 'lodash-es';
import type { VbenFormSchema } from '@vben/common-ui';
import { z } from '@vben/common-ui';
import { renderDict, renderPopoverMemberSumma, type MemberSumma } from '#/utils';
import { DictEnum } from '@vben/constants';
import { getDictOptions } from '#/utils/dict';

export class State {
  public id = 0; // 主键
  public tenantId = '000000'; // 租户编号
  public deptId = 0; // 部门id
  public userId = 0; // 用户id
  public orderNum = 0; // 排序号
  public testKey = null; // key键
  public value = null; // 值
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
    },
  rules:null
},
  {
    fieldName: 'createdAt',
    component: 'RangePicker',
    label: '创建时间',
    componentProps: {
      type: 'daterange',
      clearable: true,
      valueFormat: 'YYYY-MM-DD HH:mm:ss',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  rules:null
},
  {
    fieldName: 'deptDeptName',
    component: 'Input',
    label: '部门名称',
    componentProps: {
      placeholder: '请输入部门名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  rules:null
},
];

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
    title: '租户编号',
    field: 'tenantId',
    align: 'left',
    width: -1,
 },
  {
    title: '部门id',
    field: 'deptId',
    align: 'left',
    width: -1,
 },
  {
    title: '用户id',
    field: 'userId',
    align: 'left',
    width: -1,
 },
  {
    title: '排序号',
    field: 'orderNum',
    align: 'left',
    width: -1,
 },
  {
    title: 'key键',
    field: 'testKey',
    align: 'left',
    width: -1,
 },
 {    
				title: '值',    field: 'value',    align: 'left',    width: -1, 
				slots: {
      				default: ({ row }) => {
						return renderDict(row.value, 'sys_yes_no_num');
					}
				},
			},
			  {
    title: '版本',
    field: 'version',
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
  {
    title: '部门名称',
    field: 'deptDeptName',
    align: 'left',
    width: -1,
 },
  { title: '操作', width: 120, slots: { default: 'action' } },
];

// 表格列接口
export interface RowType {
  id: number;
  tenantId: string;
  deptId: number;
  userId: number;
  orderNum: number;
  testKey: string;
  value: string;
  version: number;
  createdDept: number;
  createdAt: string;
  createdBy: number;
  updatedAt: string;
  updatedBy: number;
  deletedBy: number;
  deletedAt: string;
  deptDeptId: number;
  deptTenantId: string;
  deptParentId: number;
  deptAncestors: string;
  deptDeptName: string;
  deptDeptCategory: string;
  deptOrderNum: number;
  deptLeader: number;
  deptPhone: string;
  deptEmail: string;
  deptStatus: string;
  deptCreatedDept: number;
  deptCreatedBy: number;
  deptCreatedAt: string;
  deptUpdatedBy: number;
  deptUpdatedAt: string;
  deptDeletedBy: number;
  deptDeletedAt: string;
};

// 查看字段列表
export const viewSchema: DescItem[] = [
  {  field: 'id',  label: '主键'},
  {  field: 'tenantId',  label: '租户编号'},
  {  field: 'deptId',  label: '部门id'},
  {  field: 'userId',  label: '用户id'},
  {  field: 'orderNum',  label: '排序号'},
  {  field: 'testKey',  label: 'key键'},
  {
				field: 'value',
				label: '值',
				render(row: any) {
					return renderDict(row.value, 'sys_yes_no_num');
				},
			},
			  {  field: 'version',  label: '版本'},
  {  field: 'createdDept',  label: '创建部门'},
  {  field: 'createdAt',  label: '创建时间'},
  {  field: 'createdBy',  label: '创建者'},
  {  field: 'updatedAt',  label: '更新时间'},
  {  field: 'updatedBy',  label: '更新者'},
  {  field: 'deletedBy',  label: '删除人'},
  {  field: 'deletedAt',  label: '删除时间'},
  {  field: 'deptDeptId',  label: '部门id'},
  {  field: 'deptTenantId',  label: '租户编号'},
  {  field: 'deptParentId',  label: '父部门id'},
  {  field: 'deptAncestors',  label: '祖级列表'},
  {  field: 'deptDeptName',  label: '部门名称'},
  {  field: 'deptDeptCategory',  label: '部门类别编码'},
  {  field: 'deptOrderNum',  label: '显示顺序'},
  {  field: 'deptLeader',  label: '负责人'},
  {  field: 'deptPhone',  label: '联系电话'},
  {  field: 'deptEmail',  label: '邮箱'},
  {  field: 'deptStatus',  label: '部门状态（0正常 1停用）'},
  {  field: 'deptCreatedDept',  label: '创建部门'},
  {  field: 'deptCreatedBy',  label: '创建者'},
  {  field: 'deptCreatedAt',  label: '创建时间'},
  {  field: 'deptUpdatedBy',  label: '更新者'},
  {  field: 'deptUpdatedAt',  label: '更新时间'},
  {  field: 'deptDeletedBy',  label: '删除人'},
  {  field: 'deptDeletedAt',  label: '删除时间'},
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
    },
  rules:z.number({required_error: '请输入部门id', invalid_type_error: '无效数字'}	)
},
  {
    fieldName: 'userId',
    component: 'InputNumber',
    label: '用户id',
    componentProps: {
      placeholder: '请输入用户id',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  rules:null
},
  {
    fieldName: 'orderNum',
    component: 'InputNumber',
    label: '排序号',
    componentProps: {
      placeholder: '请输入排序号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  rules:null
},
  {
    fieldName: 'testKey',
    component: 'Input',
    label: 'key键',
    componentProps: {
      placeholder: '请输入key键',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  rules:null
},
  {    
			fieldName: 'value',    
			component: 'Select',    
			label: '值',    
			defaultValue: null,    
			componentProps: {    
				placeholder: '请选择值',    
				options: getDictOptions('sys_yes_no_num'),    
				onUpdateValue: (e: any) => {    
					console.log(e);    
				},  
			},
			rules:null
		},
		  {
    fieldName: 'version',
    component: 'InputNumber',
    label: '版本',
    componentProps: {
      placeholder: '请输入版本',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  rules:null
},
];