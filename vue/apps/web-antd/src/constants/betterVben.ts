const BetterVben = {
  modal: {
    width: {
      /**
       * @description 弹窗默认宽度
       * @description 600px
       */
      default: 'w-[600px] ',
      /**
       * @description 表格弹窗宽度
       * @description 1200px
       */
      gird: 'w-[1200px]',
    },
  },
  form: {
    /**
     * @description 表单默认配置
     * @description 1.组件默认宽度占满
     * @description 2.栅格布局一列
     * @description 3.label与表单在一行
     * @description 4.隐藏默认操作按钮
     * @description 5.label宽度80px
     * @description 6.label左对齐
     */
    default: {
      commonConfig: {
        componentProps: {
          class: 'w-full',
        },
        formItemClass: 'col-span-1',
        labelWidth: 80,
        labelClass: 'justify-start',
      },
      showDefaultActions: false,
      wrapperClass: 'grid-cols-1 gap-x-4',
    },
    '3row': {
      wrapperClass: 'grid-cols-3 gap-x-8',
      commonConfig: {
        componentProps: {
          class: 'w-full',
        },
        formItemClass: 'col-span-1',
        labelWidth: 100,
        labelClass: 'justify-start ml-2',
      },
    },
  },
  grid: {
    formOptions: {
      /**
       * @description 表单默认配置
       * @description 1.关闭展开/折叠按钮
       * @description 2.字段值更改时提交表单
       * @description 3.回车不提交表单
       * @description 4.grid布局,4列
       * @description 5.label宽度80px
       * @description 6.label左对齐
       * @description 7.紧凑模式,移除预留的校验信息的空间
       */
      default: {
        showCollapseButton: false,
        submitOnChange: true,
        submitOnEnter: false,
        compact: true,
        commonConfig: {
          formItemClass: 'col-span-1',
          labelWidth: 80,
          labelClass: 'justify-start ml-4',
        },
        wrapperClass: 'grid-cols-4 gap-x-4',
      },
    },
    gridOptions: {
      /**
       * @description 表格默认配置
       * @description 1.自适应高度
       * @description 2.保持原始值的状态
       * @description 3.过长显示省略号
       */
      default: {
        height: 'auto',
        keepSource: true,
        showOverflow: true,
      },
      /**
       * @description 模态框表格默认配置
       * @description 1.600px高度
       * @description 2.保持原始值的状态
       * @description 3.过长显示省略号
       */
      modalDefault: {
        height: '600px',
        keepSource: true,
        showOverflow: true,
      },
      /**
       * @description 模态框表格默认配置
       * @description 1.300px高度
       * @description 2.保持原始值的状态
       * @description 3.过长显示省略号
       */
      modalShort: {
        height: '300px',
        keepSource: true,
        showOverflow: true,
      },
    },
  },
};
export default BetterVben;
