import type { VxeGridProps } from '@vben/plugins/vxe-table';

export const IndexGridColumns: VxeGridProps['columns'] = [
  {
    type: 'checkbox',
    width: 40,
  },
  {
    field: 'deviceName',
    title: '设备名称',
  },
  {
    field: 'deviceType',
    title: '设备类型',
  },
  {
    field: 'onlineState',
    title: '在线状态',
  },
  {
    field: 'update',
    title: '活跃时间',
  },
  {
    field: 'alert',
    title: '告警状态',
  },
];
