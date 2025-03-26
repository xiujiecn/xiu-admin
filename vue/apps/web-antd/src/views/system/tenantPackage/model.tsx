import type { VbenFormSchema } from '@vben/common-ui';
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { DescItem } from '#/components/description';
import { z } from '@vben/common-ui';
import { getDictOptions } from '#/utils/dict';
import { DictEnum } from '@vben/constants';
import { getPopupContainer } from '@vben/utils';
import { Tag } from 'ant-design-vue';
import dayjs from 'dayjs';
import { h } from 'vue';
import {
  renderDict,
  renderDictTags,
  renderHttpMethodTag,
  renderJsonPreview,
} from '#/utils/render';
const defaultExpireTime = dayjs()
  .add(365, 'days')
  .startOf('day')
  .format('YYYY-MM-DD HH:mm:ss');

export const querySchema: VbenFormSchema[] = [
  {
    component: 'Input',
    fieldName: 'packageName',
    label: '套餐名称',
  },
];
export const columns: VxeGridProps['columns'] = [
  { align: 'left', title: 'ID', type: 'checkbox', width: 80 },
  {
    title: '套餐名称',
    field: 'packageName',
  },
  {
    title: '备注',
    field: 'remark',
  },
  {
    title: '状态',
    field: 'status',
    slots: { default: 'status' },
    width: 65,
  },
  { title: '操作', width: 120, slots: { default: 'action' } },
];

export const viewSchema: DescItem[] = [
  { field: 'packageId', label: '套餐ID' },
  { field: 'packageName', label: '套餐名称' },
  { field: 'menuIds', label: '关联菜单' },
  { field: 'remark', label: '备注' },
  { field: 'menuCheckStrictly', label: '菜单树选择项是否关联显示' },
  { field: 'status', label: '状态' ,
    render(value, { status }) {
      const operType = renderDict(status, DictEnum.SYS_NORMAL_DISABLE);
      return (
        <div class="flex items-center">
          <Tag>{value}</Tag>
          {operType}
        </div>
      );
    },
  },
  { field: 'createdDept', label: '创建部门' },
  { field: 'createdBy', label: '创建者' },
  { field: 'createdAt', label: '创建时间' },
  { field: 'updatedBy', label: '更新者' },
  { field: 'updatedAt', label: '更新时间' },
];

export const drawerSchema: VbenFormSchema[] = [
  {
    component: 'Input',
    dependencies: {
      show: () => false,
      triggerFields: [''],
    },
    fieldName: 'packageId',
  },
  {
    component: 'Radio',
    dependencies: {
      show: () => false,
      triggerFields: [''],
    },
    fieldName: 'menuCheckStrictly',
  },
  {
    component: 'Input',
    fieldName: 'packageName',
    label: '套餐名称',
    rules: 'required',
  },
  {
    component: 'menuIds',
    defaultValue: [],
    fieldName: 'menuIds',
    label: '关联菜单',
  },
  {
    component: 'Textarea',
    fieldName: 'remark',
    formItemClass: 'items-baseline',
    label: '备注',
  },
];
