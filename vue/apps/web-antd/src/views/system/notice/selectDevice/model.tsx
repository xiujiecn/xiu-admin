import { h, ref } from 'vue';
import { Tag } from 'ant-design-vue';
import type { VxeGridProps } from '#/adapter/vxe-table';
import { getPopupContainer } from '@vben/utils';
import type { DescItem } from '#/components/description';
import { cloneDeep } from 'lodash-es';
import type { VbenFormSchema } from '@vben/common-ui';
import { z } from '@vben/common-ui';
import { renderDict } from '#/utils';
import { DictEnum } from '@vben/constants';
import { getDictOptions } from '#/utils/dict';

// 表单验证规则

// 表格列
export const columns: VxeGridProps['columns'] = [
  {
    type: 'checkbox', // 多选按钮
    width: 60,
    align: 'center',
  },
  {
    title: '用户ID',
    field: 'userId',
    align: 'center',
    width: 80,
  },
  {
    title: '用户名',
    field: 'userName',
    align: 'center',
    width: 120,
  },
  {
    title: '用户昵称',
    field: 'nickName',
    align: 'center',
    width: -1,
  },
  {
    title: '手机号',
    field: 'phonenumber',
    align: 'center',
    width: 120,
  },
  {
    title: '邮箱',
    field: 'email',
    align: 'center',
    width: 150,
  },
  {
    title: '机构',
    field: 'deptInfo.deptName',
    align: 'center',
    width: 120,
  },
  {
    title: '状态',
    field: 'status',
    align: 'center',
    width: -1,
    slots: { default: 'status' },
  },
];

// 表格列接口
export interface RowType {
  userId: number;
  tenantId: string;
  deptId: number;
  userName: string;
  nickName: string;
  userType: string;
  email: string;
  phonenumber: string;
  sex: string;
  avatar: string;
  status: string;
  loginIp: string;
  loginDate: string;
  createdDept: number;
  createdBy: number;
  createdAt: string;
  updatedBy: number;
  updatedAt: string;
  deletedBy: number;
  deletedAt: string;
  remark: string;
  deptInfo: {
    deptId: number;
    deptName: string;
  };
}
