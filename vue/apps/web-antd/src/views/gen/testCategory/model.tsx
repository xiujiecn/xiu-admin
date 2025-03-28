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
  public id = 0; // 分类ID
  public name = ''; // 分类名称
  public shortName = null; // 简称
  public description = null; // 描述
  public sort = 0; // 排序
  public remark = null; // 备注
  public status = 0; // 状态
  public createdDept = 0; // 创建部门
  public createdAt = ''; // 创建时间
  public createdBy = 0; // 创建者
  public createdBySumma?: null | MemberSumma = null; // 创建者摘要信息
  public updatedAt = ''; // 修改时间
  public updatedBy = 0; // 修改者
  public updatedBySumma?: null | MemberSumma = null; // 修改者摘要信息
  public deletedAt = ''; // 删除时间
  public deletedBy = 0; // 删除者
  public deletedBySumma?: null | MemberSumma = null; // 删除者摘要信息

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
    label: '分类ID',
    componentProps: {
      placeholder: '请输入分类ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  rules:null
},
  {    
			fieldName: 'status',    
			component: 'Select',    
			label: '状态',    
			defaultValue: null,    
			componentProps: {    
				placeholder: '请选择状态',    
				options: getDictOptions('sys_common_status'),    
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
];

// 表格列
export const columns: VxeGridProps['columns'] = [
  {
    title: '分类ID',
    field: 'id',
    align: 'left',
    width: -1,
    type: 'checkbox',
  },
  {
    title: '分类名称',
    field: 'name',
    align: 'left',
    width: -1,
 },
  {
    title: '简称',
    field: 'shortName',
    align: 'left',
    width: -1,
 },
  {
    title: '描述',
    field: 'description',
    align: 'left',
    width: -1,
 },
  {
    title: '排序',
    field: 'sort',
    align: 'left',
    width: -1,
 },
  {
    title: '备注',
    field: 'remark',
    align: 'left',
    width: -1,
 },
 {    
				title: '状态',    field: 'status',    align: 'left',    width: -1, 
				slots: {
      				default: ({ row }) => {
						return renderDict(row.status, 'sys_common_status');
					}
				},
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
    title: '修改时间',
    field: 'updatedAt',
    align: 'left',
    width: -1,
 },
  {
    title: '修改者',
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
    title: '删除者',
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
  name: string;
  shortName: string;
  description: string;
  sort: number;
  remark: string;
  status: number;
  createdDept: number;
  createdAt: string;
  createdBy: number;
  updatedAt: string;
  updatedBy: number;
  deletedAt: string;
  deletedBy: number;
};

// 查看字段列表
export const viewSchema: DescItem[] = [
  {  field: 'id',  label: '分类ID'},
  {  field: 'name',  label: '分类名称'},
  {  field: 'shortName',  label: '简称'},
  {  field: 'description',  label: '描述'},
  {  field: 'sort',  label: '排序'},
  {  field: 'remark',  label: '备注'},
  {
				field: 'status',
				label: '状态',
				render(row: any) {
					return renderDict(row.status, 'sys_common_status');
				},
			},
			  {  field: 'createdDept',  label: '创建部门'},
  {  field: 'createdAt',  label: '创建时间'},
  {  field: 'createdBy',  label: '创建者'},
  {  field: 'updatedAt',  label: '修改时间'},
  {  field: 'updatedBy',  label: '修改者'},
  {  field: 'deletedAt',  label: '删除时间'},
  {  field: 'deletedBy',  label: '删除者'},
];

// 编辑字段列表
export const editSchema: VbenFormSchema[] = [
  {
    fieldName: 'id',
    component: 'Input',
    label: '分类ID',
    dependencies: {   show: () => false,    triggerFields: [''],   },
    componentProps: {
      placeholder: '',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    fieldName: 'name',
    component: 'Input',
    label: '分类名称',
    componentProps: {
      placeholder: '请输入分类名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  rules:'required'
},
  {
    fieldName: 'shortName',
    component: 'Input',
    label: '简称',
    componentProps: {
      placeholder: '请输入简称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  rules:null
},
  {
    fieldName: 'description',
    component: 'Input',
    label: '描述',
    componentProps: {
      placeholder: '请输入描述',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  rules:null
},
  {
    fieldName: 'sort',
    component: 'InputNumber',
    label: '排序',
    componentProps: {
      placeholder: '请输入排序',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  rules:z.number({required_error: '请输入排序', invalid_type_error: '无效数字'})
},
  {
    fieldName: 'remark',
    component: 'Input',
    label: '备注',
    componentProps: {
      placeholder: '请输入备注',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  rules:null
},
  {    
			fieldName: 'status',    
			component: 'Select',    
			label: '状态',    
			defaultValue: null,    
			componentProps: {    
				placeholder: '请选择状态',    
				options: getDictOptions('sys_common_status'),    
				onUpdateValue: (e: any) => {    
					console.log(e);    
				},  
			},
			rules:null
		},
		];