import type { VxeTableGridOptions } from '@vben/plugins/vxe-table';

const BetterVben: {
  grid: {
    default: VxeTableGridOptions;
  };
  config: any;
} = {
  grid: {
    /**
     * 1.取消边框
     * 2.禁止调整列宽
     * 3.左对齐
     * 4.内容过长显示省略
     */
    default: {
      border: false,
      columnConfig: {
        resizable: false,
      },
      showOverflow: false,
      align: 'left',
    },
  },
  config: {
    cardGrid: [
      {
        field: 'list',
        slots: { default: 'list' },
      },
    ],
  },
};
export default BetterVben;
