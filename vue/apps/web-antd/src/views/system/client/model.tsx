import type { VbenFormSchema } from '@vben/common-ui';
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { DescItem } from '#/components/description';
import { z } from '@vben/common-ui';
import { getDictOptions } from '#/utils/dict';
import { DictEnum } from '@vben/constants';
import { getPopupContainer } from '@vben/utils';
import { Tag } from 'ant-design-vue';
import { h } from 'vue';
import {
  renderDict,
  renderDictTags,
  renderHttpMethodTag,
  renderJsonPreview,
} from '#/utils/render';

export const querySchema: VbenFormSchema[] = [
  {
    component: 'Input',
    fieldName: 'clientId',
    label: '客户端ID',
  },
  {
    component: 'Input',
    fieldName: 'clientKey',
    label: '客户端Key',
  },
  {
    component: 'Input',
    fieldName: 'clientSecret',
    label: '客户端Secret',
  },
];
export const columns: VxeGridProps['columns'] = [
  { align: 'left', title: 'ID', type: 'checkbox', width: 80 },
  { field: 'clientId', title: '客户端ID' },
  { field: 'clientKey', title: '客户端Key' },
  { field: 'clientSecret', title: '客户端Secret' },
  {
    title: '授权类型',
    field: 'grantTypeList',
    slots: {
      default: ({ row }) => {
        row.grantTypeList = row.grantType.split(',');
        if (!row.grantTypeList) {
          return '无';
        }
        return renderDictTags(
          row.grantTypeList,
          getDictOptions(DictEnum.SYS_GRANT_TYPE),
          true,
          4,
        );
      },
    },
  },
  { field: 'deviceType', title: '设备类型' },
  { field: 'activeTimeout', title: '活跃超时时间' },
  { field: 'timeout', title: '固定超时' },
  { field: 'status', title: '状态', slots: { default: 'status' } },
  { field: 'createdAt', formatter: 'formatDateTime', title: '创建时间' },
  { title: '操作', width: 120, slots: { default: 'action' } },
];

export const viewSchema: DescItem[] = [
  {
    field: 'id',
    label: 'ID',
  },
  {
    field: 'clientId',
    label: '客户端ID',
  },
  {
    field: 'clientKey',
    label: '客户端Key',
  },
  {
    field: 'clientSecret',
    label: '客户端Secret',
  },
  {
    field: 'grantType',
    label: '授权类型',
    render(value, { grantType }) {
      // 将grantType转换为数组
      const grantTypeList = grantType.split(',');
      const operType = grantTypeList.map((item: string) =>
        renderDict(item, DictEnum.SYS_GRANT_TYPE),
      );
      return <div class="flex items-center">{operType}</div>;
    },
  },
  {
    field: 'deviceType',
    label: '设备类型',
    render(value, { deviceType }) {
      const operType = renderDict(deviceType, DictEnum.SYS_DEVICE_TYPE);
      return (
        <div class="flex items-center">
          <Tag>{value}</Tag>
          {operType}
        </div>
      );
    },
  },
  {
    field: 'activeTimeout',
    label: 'Token活跃超时时间',
  },
  {
    field: 'timeout',
    label: 'Token固定超时时间',
  },
  {
    field: 'status',
    label: '状态',
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
  {
    field: 'createdAt',
    label: '创建时间',
  },
  {
    field: 'createdBy',
    label: '创建人',
  },
  {
    field: 'updatedAt',
    label: '更新时间',
  },
  {
    field: 'updatedBy',
    label: '更新人',
  },
];

export const drawerSchema: VbenFormSchema[] = [
  {
    component: 'Input',
    dependencies: {
      show: () => false,
      triggerFields: [''],
    },
    fieldName: 'id',
    label: 'id',
  },
  {
    component: 'Input',
    componentProps: {
      disabled: true,
    },
    dependencies: {
      show: () => false,
      triggerFields: [''],
    },
    fieldName: 'clientId',
    label: '客户端ID',
  },
  {
    component: 'Input',
    fieldName: 'clientKey',
    label: '客户端key',
    rules: 'required',
  },
  {
    component: 'Input',
    fieldName: 'clientSecret',
    label: '客户端密钥',
    rules: 'required',
  },
  {
    component: 'Select',
    componentProps: {
      getPopupContainer,
      mode: 'multiple',
      optionFilterProp: 'label',
      options: getDictOptions(DictEnum.SYS_GRANT_TYPE),
    },
    fieldName: 'grantTypeList',
    label: '授权类型',
    rules: 'selectRequired',
  },
  {
    component: 'Select',
    componentProps: {
      allowClear: false,
      getPopupContainer,
      options: getDictOptions(DictEnum.SYS_DEVICE_TYPE),
    },
    fieldName: 'deviceType',
    label: '设备类型',
    rules: 'selectRequired',
  },
  {
    component: 'InputNumber',
    componentProps: {
      addonAfter: '秒',
      placeholder: '请输入',
    },
    defaultValue: 1800,
    fieldName: 'activeTimeout',
    formItemClass: 'col-span-2 lg:col-span-1',
    help: '指定时间无操作则过期(单位：秒), 默认30分钟(1800秒)',
    label: 'Token活跃超时时间',
    rules: 'required',
  },
  {
    component: 'InputNumber',
    componentProps: {
      addonAfter: '秒',
    },
    defaultValue: 604_800,
    fieldName: 'timeout',
    formItemClass: 'col-span-2 lg:col-span-1 ',
    help: '指定时间必定过期(单位：秒)，默认七天(604800秒)',
    label: 'Token固定超时时间',
    rules: 'required',
  },
  {
    component: 'RadioGroup',
    componentProps: {
      buttonStyle: 'solid',
      options: getDictOptions(DictEnum.SYS_NORMAL_DISABLE),
      optionType: 'button',
    },
    defaultValue: '0',
    fieldName: 'status',
    label: '状态',
  },
];
