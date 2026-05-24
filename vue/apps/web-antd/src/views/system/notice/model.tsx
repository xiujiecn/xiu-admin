/**
 * @description 通知公告模型定义
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 LiXiujie
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
 */
import type {
  VbenFormSchema,
} from '@vben/common-ui';
import { getDictOptions } from '#/utils/dict';
import { DictEnum } from '@vben/constants';
import { getPopupContainer } from '@vben/utils';
import { Tag } from 'ant-design-vue';
import type { DescItem } from '#/components/description';

const dictOptions = getDictOptions(DictEnum.NOTICE_RANGE);
/** 编辑表单 */
export const drawerSchema: VbenFormSchema[] = [
  {
    component: 'Input',
    dependencies: {
      show: () => false,
      triggerFields: [''],
    },
    fieldName: 'noticeId',
    label: '主键',
  },
  {
    component: 'Input',
    fieldName: 'noticeTitle',
    label: '公告标题',
    rules: 'required',
  },
  {
    component: 'RadioGroup',
    componentProps: {
      buttonStyle: 'solid',
      options: getDictOptions(DictEnum.SYS_NOTICE_STATUS),
      optionType: 'button',
    },
    defaultValue: '0',
    fieldName: 'status',
    label: '公告状态',
    rules: 'required',
    formItemClass: 'col-span-1',
  },
  {
    component: 'RadioGroup',
    componentProps: {
      buttonStyle: 'solid',
      options: getDictOptions(DictEnum.SYS_NOTICE_TYPE),
      optionType: 'button',
    },
    defaultValue: '1',
    fieldName: 'noticeType',
    label: '公告类型',
    rules: 'required',
    formItemClass: 'col-span-1',
  },
  {
    component: 'Textarea',
    componentProps: {
      width: '100%',
    },
    fieldName: 'noticeContent',
    label: '公告内容',
    rules: 'required',
  },
  /** 公告通知范围 */
  {
    fieldName: 'noticeRange',
    label: '公告通知范围',
    component: "Select",
    formItemClass: 'col-span-1',
    componentProps: {
      options: dictOptions
    },
    defaultValue: "1",
  },
  /** 组织树 */
  {
    fieldName: 'deptIdList',
    label: '通知组织',
    component: "TreeSelect",
    formItemClass: 'col-span-1',
    componentProps: {
      placeholder: '请选择组织',
      getPopupContainer,
      multiple: true, // 允许多选
      treeCheckable: true, // 显示复选框
    }, dependencies: {
      if(values) {
        return values.noticeRange == "2"
      },
      triggerFields: ['noticeRange'],
    },
    rules: "required"
  },
  /** 用户树 */
  {
    fieldName: 'userIdList',
    label: '通知用户',
    component: "Select",
    formItemClass: 'col-span-1',
    componentProps: {
      readonly: true, // 用户名字段始终只读，通过选择按钮填充
      placeholder: '请选择用户',
      mode: 'multiple', // 允许多选
    },
    dependencies: {
      if(values) {
        return values.noticeRange == "3"
      },
      triggerFields: ['noticeRange'],
    },
    rules: "required"
  },

  {
    component: 'Input',
    fieldName: 'remark',
    label: '备注',
  },
];
/** 查看表单 */
export const viewSchema: DescItem[] = [
  { field: 'noticeId', label: '公告ID' },
  { field: 'noticeTitle', label: '公告标题' },
  { field: 'noticeType', label: '公告类型', render: (val) => val == 1 ? '通知' : '公告' },
  { field: 'noticeContent', label: '公告内容' },
  { field: 'status', label: '公告状态', render: (val) => val == 0 ? '正常' : '停用' },
    {
    field: 'noticeRange',
    label: '公告通知范围',
    render: (value) => {
      const noticeRange=getDictOptions(DictEnum.NOTICE_RANGE)
      const item=noticeRange.filter(item=>item.dictValue==value)
      return item[0]["dictLabel"];
    }
  },
  {
    field: 'userIdList',
    label: '提及人员',
    render: (value) => {
      if(!value){return ""}
      return (<div>{value.map((v:any, idx:any) => <Tag key={idx}>{v}</Tag>)}</div>);
    }
  }, {
    field: 'deptIdList',
    label: '提及组织',
    render: (value) => {
      if(!value){return ""}
      return (<div>{value.map((v:any, idx:any) => <Tag key={idx}>{v}</Tag>)}</div>);
    }
  },
  { field: 'createdDept', label: '创建组织' },
  { field: 'createdBy', label: '创建者' },
  { field: 'createdAt', label: '创建时间' },
  { field: 'updatedBy', label: '更新者' },
  { field: 'updatedAt', label: '更新时间' },
  { field: 'remark', label: '备注' },
];

