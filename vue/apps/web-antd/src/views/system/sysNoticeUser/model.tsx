import { h, ref } from 'vue';
import { Tag } from 'ant-design-vue';
import type { VxeGridProps } from '#/adapter/vxe-table';
import { getPopupContainer } from '@vben/utils';
import type { DescItem } from '#/components/description';
import { cloneDeep } from 'lodash-es';
import type { VbenFormSchema } from '@vben/common-ui';
import { z } from '@vben/common-ui';
import { renderDict, renderPopoverMemberSumma, type MemberSumma } from '#/utils';
import { getDictOptions } from '#/utils/dict';
import { DictEnum } from '@vben/constants';

const dictOptions = getDictOptions(DictEnum.SYS_NOTICE_TYPE);
export class State {
  public id = 0; // ID
  public noticeId = 0; // 公告ID
  public tenantId = '000000'; // 租户编号
  public noticeTitle = ''; // 公告标题
  public noticeType = ''; // 公告类型
  public noticeContent = null; // 公告内容
  public status = 0; // 公告状态
  public isRead = 0; // 是否已读
  public createdDept = 0; // 创建机构
  public createdBy = 0; // 创建者
  public createdBySumma?: null | MemberSumma = null; // 创建者摘要信息
  public createdAt = ''; // 创建时间
  public updatedBy = 0; // 更新者
  public updatedBySumma?: null | MemberSumma = null; // 更新者摘要信息
  public updatedAt = ''; // 更新时间
  public remark = null; // 备注

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
    label: '公告ID',
    componentProps: {
      placeholder: '请输入ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    }, rules: null,
    formItemClass: 'col-span-1',
  },
  {
    fieldName: 'noticeType',
    component: 'Select',
    label: '公告类型',
    componentProps: {
      placeholder: '请选择公告类型',
      options: dictOptions,
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
    formItemClass: 'col-span-1',
  }, {
    fieldName: 'isRead',
    component: 'Select',
    label: '是否已读',
    componentProps: {
      placeholder: '请选择是否已读',
      options: [
        {
          label: '已读',
          value: 1,
        },
        {
          label: "未读",
          value: 0,
        },
      ],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
    formItemClass: 'col-span-1',
  },];

// 表格列
export const columns: VxeGridProps['columns'] = [
   {
    title: '',
    field: '',
    width: 40,
    type: 'checkbox',
  },  {
    title: 'ID',
    field: 'id',
    align: 'left',
    width: -1,
  },

  {
    title: '公告ID',
    field: 'noticeId',
    align: 'center',
    width: -1,
  }, {
    title: '公告类型',
    field: 'noticeType',
    align: 'center',
    slots: {
      default: (scope) => {
        return <Tag>{renderDict(scope.row.noticeType, DictEnum.NOTICE_RANGE)}</Tag>;
      },
    },
    width: -1,
  },
  {
    title: '公告标题',
    field: 'noticeTitle',
    align: 'center',
    width: -1,
  },

  {
    title: '公告内容',
    field: 'noticeContent',
    align: 'center',
    width: "50%",
  },
  {
    title: '创建者',
    field: 'createdBy',
    align: 'center',
    width: -1,
    slots: {
      default: ({ row }) => {
        return renderPopoverMemberSumma(row.createdBySumma);
      },
    },
  },
  { title: '操作', width: 120, slots: { default: 'action' } },
];

// 表格列接口
export interface RowType {
  id: number;
  noticeId: number;
  tenantId: string;
  noticeTitle: string;
  noticeType: string;
  noticeContent: string;
  status: string;
  isRead: number;
  createdDept: number;
  createdBy: number;
  createdAt: string;
  updatedBy: number;
  updatedAt: string;
  remark: string;
};

// 查看字段列表
export const viewSchema: DescItem[] = [
  { field: 'id', label: 'ID' },
  { field: 'noticeId', label: '公告ID' },
  { field: 'tenantId', label: '租户编号' },
  { field: 'noticeTitle', label: '公告标题' },
  { field: 'noticeType', label: '公告类型' },
  { field: 'noticeContent', label: '公告内容' },
  { field: 'status', label: '公告状态' },
  { field: 'isRead', label: '是否已读' },
  { field: 'createdDept', label: '创建机构' },
  { field: 'createdBy', label: '创建者' },
  { field: 'createdAt', label: '创建时间' },
  { field: 'updatedBy', label: '更新者' },
  { field: 'updatedAt', label: '更新时间' },
  { field: 'remark', label: '备注' },
];

// 编辑字段列表
export const editSchema: VbenFormSchema[] = [
  {
    fieldName: 'id',
    component: 'Input',
    label: 'ID',
    dependencies: { show: () => false, triggerFields: [''], },
    componentProps: {
      placeholder: '',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
    formItemClass: 'col-span-1',
  }, {
    fieldName: 'noticeId',
    component: 'InputNumber',
    label: '公告ID',
    componentProps: {
      placeholder: '请输入公告ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    }, rules: z.number({ required_error: '请输入公告ID', invalid_type_error: '无效数字' }),
    formItemClass: 'col-span-1',
  }, {
    fieldName: 'tenantId',
    component: 'Input',
    label: '租户编号',
    componentProps: {
      placeholder: '请输入租户编号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
    rules: 'required',
    formItemClass: 'col-span-1',
  }, {
    fieldName: 'noticeTitle',
    component: 'Input',
    label: '公告标题',
    componentProps: {
      placeholder: '请输入公告标题',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
    rules: 'required',
    formItemClass: 'col-span-1',
  }, {
    fieldName: 'noticeType',
    component: 'Input',
    label: '公告类型',
    componentProps: {
      placeholder: '请输入公告类型',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
    rules: 'required',
    formItemClass: 'col-span-1',
  }, {
    fieldName: 'noticeContent',
    component: 'Input',
    label: '公告内容',
    componentProps: {
      placeholder: '请输入公告内容',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
    rules: 'required',
    formItemClass: 'col-span-1',
  }, {
    fieldName: 'status',
    component: 'Input',
    label: '公告状态',
    componentProps: {
      placeholder: '请输入公告状态',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
    rules: 'required',
    formItemClass: 'col-span-1',
  }, {
    fieldName: 'isRead',
    component: 'InputNumber',
    label: '是否已读',
    componentProps: {
      placeholder: '请输入是否已读',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    }, rules: z.number({ required_error: '请输入是否已读', invalid_type_error: '无效数字' }),
    formItemClass: 'col-span-1',
  }, {
    fieldName: 'remark',
    component: 'Input',
    label: '备注',
    componentProps: {
      placeholder: '请输入备注',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
    rules: 'required',
    formItemClass: 'col-span-1',
  },];
