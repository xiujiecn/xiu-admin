import type { VbenFormSchema } from '@vben/common-ui';
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { DescItem } from '#/components/description';
import type {
  GenCodesSelect,
  DbTableSelectModel,
} from '#/api/gen_codes/gen_table';
import {
  Button,
  Input,
  type TableColumnsType,
  InputNumber,
  Space,
  Checkbox,
  Select,
  Cascader,
} from 'ant-design-vue';
import { cloneDeep } from 'lodash-es';
import { isJsonString } from '#/utils/is';
import { z } from '@vben/common-ui';
import { getDictOptions } from '#/utils/dict';
import { DictEnum } from '@vben/constants';
import { getPopupContainer } from '@vben/utils';
import { Tag } from 'ant-design-vue';
import { h, defineEmits, ref, computed } from 'vue';

import {
  renderDict,
  renderIcon,
  renderTooltip,
  renderDictTag,
  renderDictTags,
  renderHttpMethodTag,
  renderJsonPreview,
} from '#/utils/render';
import { getGenCodesTableSelectApi, getGenCodesColumnSelectApi } from '#/api/gen_codes/gen_table';
export let selectListObj: GenCodesSelect = {
  genType: [],
  db: [],
  status: [],
  linkMode: [],
  buildMethod: [],
  formMode: [],
  formRole: [],
  whereMode: [],
  addons: [],
  tableAlign: [],
  treeStyleType: [],
  dictMode: [],
};
export let tableListObj: DbTableSelectModel[] = [];

export function getSelectList<K extends keyof GenCodesSelect>(key: K): GenCodesSelect[K] {
  return selectListObj[key];
}
export function setSelectList(key: keyof GenCodesSelect, value: any) {
  selectListObj[key] = value;
}

export function setSelectListObj(value: GenCodesSelect) {
  selectListObj = value;
}
// 格式化列字段
export function formatColumns(columns: any) {
  if (columns === undefined || columns.length === 0) {
    columns = [];
  }

  if (isJsonString(columns)) {
    columns = JSON.parse(columns);
  }

  if (columns.length > 0) {
    for (let i = 0; i < columns.length; i++) {
      if (!columns[i].formGridSpan) {
        columns[i].formGridSpan = 1;
      }
      if (!columns[i].align) {
        columns[i].align = 'left';
      }
      if (!columns[i].width || columns[i].width < 1) {
        columns[i].width = null;
      }
    }
  }
  return columns;
}


export const genInfoObj = {
  id: 0,
  genType: 0,
  genTemplate: null,
  varName: '',
  options: {
    headOps: ['add', 'batchDel', 'export'],
    columnOps: ['edit', 'del', 'view', 'status', 'check'],
    autoOps: ['genMenuPermissions', 'runDao', 'runService'],
    join: [],
    menu: {
      pid: 0,
      icon: 'ant-design:home-twotone',
      sort: 1,
    },
    tree: {
      titleColumn: null,
      pidColumn: 'parent_id',
      levelColumn: null,
      treeColumn: null,
      styleType: 1,
    },
    funcDict: {
      valueColumn: null,
      labelColumn: null,
    },
    presetStep: {
      formGridCols: 1,
    },
  },
  dbName: '',
  tableName: '',
  tableComment: '',
  daoName: '',
  masterColumns: [],
  addonName: null,
  status: 2,
  createdAt: '',
  updatedAt: '',
};  
export function newState(state:any) {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(genInfoObj);
}

export interface GenTypeOption {
  color: string;
  label: string;
  value: string;
}

export const formGridColsOptions = [
  {
    value: 1,
    label: '一行一列',
  },
  {
    value: 2,
    label: '一行两列',
  },
  {
    value: 3,
    label: '一行三列',
  },
  {
    value: 4,
    label: '一行四列',
  },
];

export const formGridSpanOptions = [
  {
    value: 1,
    label: '占一列位置',
  },
  {
    value: 2,
    label: '占两列位置',
  },
  {
    value: 3,
    label: '占三列位置',
  },
  {
    value: 4,
    label: '占四列位置',
  },
];

export const querySchema: VbenFormSchema[] = [
  {
    component: 'Input',
    fieldName: 'genType',
    label: '生成类型',
  },
  {
    component: 'Input',
    fieldName: 'varName',
    label: '实体命名',
  },
  {
    component: 'Select',
    componentProps: {
      allowClear: true,
      options: [
        {
          label: '生成成功',
          value: '0',
        },
        {
          label: '未开始',
          value: '1',
        },
        {
          label: '生成失败',
          value: '2',
        },
      ],
      placeholder: '请选择',
    },
    fieldName: 'status',
    label: '状态',
  },
];

export const columns: VxeGridProps['columns'] = [
  { align: 'left', title: 'ID', type: 'checkbox', width: 80 },
  { field: 'genType', title: '生成类型' },
  { field: 'genTemplate', title: '生成模板' },
  { field: 'varName', title: '实体命名' },
  { field: 'tableComment', title: '生成名称' },
  {
    field: 'status',
    title: '状态',
    slots: {
      default: ({ row }) => {
        return renderDict(row.status, DictEnum.SYS_GEN_STATUS);
      },
    },
  },
  { field: 'createdAt', formatter: 'formatDateTime', title: '创建时间' },
  { title: '操作', width: 160, slots: { default: 'action' } },
];

export const viewSchema: DescItem[] = [
  {
    field: 'tableId',
    label: 'ID',
  },
  {
    field: 'genType',
    label: '生成类型',
    render(value) {
      const operType = selectListObj.genType.find(
        (item) => item.value === value,
      )?.label;
      return (
        <div class="flex items-center">
          <Tag>{operType}</Tag>
        </div>
      );
    },
  },
  {
    field: 'genTemplate',
    label: '生成模板',
    render(value, { genType }) {
      const operType = selectListObj.genType
        .find((item) => item.value === genType)
        ?.templates.find((item) => item.value === value)?.label;
      return (
        <div class="flex items-center">
          <Tag>{operType}</Tag>
        </div>
      );
    },
  },
  {
    field: 'varName',
    label: '实体命名',
  },
  {
    field: 'tableComment',
    label: '生成名称',
  },
  {
    field: 'dbName',
    label: '数据库名称',
  },
  {
    field: 'tableName',
    label: '主表名称',
  },
  {
    field: 'daoName',
    label: '主表dao模型',
  },
  {
    field: 'masterColumns',
    label: '主表字段',
  },
  {
    field: 'addonName',
    label: '插件名称',
  },
  {
    field: 'status',
    label: '状态',
    render(value, { status }) {
      const operType = renderDict(status, DictEnum.SYS_GEN_STATUS);
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
    fieldName: 'tableId',
    label: 'ID',
  },
  {
    component: 'Select',
    componentProps: {
      options: selectListObj.genType,
      defaultValue: 0,
    },
    fieldName: 'genType',
    label: '生成类型',
  },
  {
    component: 'Select',
    fieldName: 'dbName',
    label: '数据库名称',
    componentProps: {
      options: selectListObj.db,
      defaultValue: '',
    },
  },
  {
    component: 'Select',
    fieldName: 'tableName',
    label: '数据库表',
    componentProps: {
      options: tableListObj,
      showSearch: true,
    },
    dependencies: {
      async componentProps(values) {
        if (values.dbName) {
          const res = await getGenCodesTableSelectApi({
            dbGroup: values.dbName,
          });
          tableListObj = res.items;
        }
        return {
          options: tableListObj,
        };
      },
      triggerFields: ['dbName'],
    },
  },
  {
    component: 'Select',
    fieldName: 'genTemplate',
    label: '生成模板',
    componentProps: {
      options: [],
    },
    dependencies: {
      async componentProps(values) {
        if (values.genType !== undefined) {
          const templates = selectListObj.genType.find(
            (item) => item.value === values.genType,
          )?.templates;
          return {
            options: templates,
          };
        }
        return {
          options: selectListObj.genType[0]?.templates || [],
        };
      },
      triggerFields: ['genType', 'tableName', 'dbName'],
    },
  },
  {
    component: 'Select',
    fieldName: 'addons',
    label: '选择插件',
    componentProps: {
      options: selectListObj.addons,
    },
    dependencies: {
      show: (values) => {
        return (
          selectListObj.genType
            .find((item) => item.value === values.genType)
            ?.templates.find((item) => item.value === values.genTemplate)
            ?.isAddon || false
        );
      },
      triggerFields: ['genTemplate'],
    },
  },
  {
    component: 'Input',
    fieldName: 'tableComment',
    label: '菜单名称',
    dependencies: {
      trigger: (values, formApi) => {
        const tableComment = tableListObj.find(
          (item) => item.value === values.tableName,
        )?.defTableComment;
        formApi.setFieldValue('tableComment', tableComment);
      },
      triggerFields: ['tableName'],
    },
  },
  {
    component: 'Input',
    fieldName: 'varName',
    label: '实体命名',
    dependencies: {
      trigger: (values, formApi) => {
        console.log(
          'vue/apps/web-antd/src/views/tool/gen/model.tsx tableComment trigger ',
          values,
          formApi.values,
        );
        const varName = tableListObj.find(
          (item) => item.value === values.tableName,
        )?.defVarName;
        formApi.setFieldValue('varName', varName);
      },
      triggerFields: ['tableName'],
    },
  },
  {
    component: 'Input',
    componentProps: {
      class: 'hidden',
      hideLabel: true,
    },
    dependencies: {
      // show: () => false,
      trigger: (values, formApi) => {
        const daoName = tableListObj.find(
          (item) => item.value === values.tableName,
        )?.daoName;
        formApi.setFieldValue('daoName', daoName);
        console.log('vue/apps/web-antd/src/views/tool/gen/model.tsx daoName trigger ', daoName);
      },
      triggerFields: ['tableName'],
    },
    fieldName: 'daoName',
    label: '',
  },
];

export const developBaseSchema: VbenFormSchema[] = [
  {
    component: 'Input',
    dependencies: {
      show: () => false,
      triggerFields: [''],
    },
    fieldName: 'tableId',
    label: 'ID',

  },

  {
    component: 'Select',
    componentProps: {
      options: selectListObj.genType,
      defaultValue: 0,
    },
    fieldName: 'genType',
    label: '生成类型',
  },
  {
    component: 'Input',
    fieldName: 'varName',
    label: '实体命名',
  },
  {
    component: 'Select',
    fieldName: 'dbName',
    label: '数据库名称',
    componentProps: {
      options: selectListObj.db,
      defaultValue: '',
    },
  },
  {
    component: 'Select',
    fieldName: 'tableName',
    label: '数据库表',
    componentProps: {
      options: tableListObj,
      showSearch: true,
    },
    dependencies: {
      async componentProps(values) {
        if (values.dbName) {
          const res = await getGenCodesTableSelectApi({
            dbGroup: values.dbName,
          });
          tableListObj = res.items;
        }
        return {
          options: tableListObj,
        };
      },
      triggerFields: ['dbName'],
    },
  },
  {
    component: 'CheckboxGroup',
    componentProps: {
      name: 'cname',
      options: [
        {
          label: '新增表单按钮',
          value: 'add',
        },
        {
          label: '批量删除按钮',
          value: 'batchDel',
        },
        {
          label: '导出按钮',
          value: 'export',
        },
      ],
    },
    fieldName: 'options.headOps',
    label: '表格头部按钮组',
    formItemClass: 'col-span-4',
  },
  {
    component: 'CheckboxGroup',
    componentProps: {
      name: 'cname',
      options: [
        {
          label: '编辑',
          value: 'edit',
        },
        {
          label: '状态修改',
          value: 'status',
        },
        {
          label: '删除',
          value: 'del',
        },
        {
          label: '详情页',
          value: 'view',
        },
        {
          label: '开启勾选列',
          value: 'check',
        },
        {
          label: '不过滤权限',
          value: 'notFilterAuth',
        },
      ],
    },
    fieldName: 'options.columnOps',
    label: '表格列操作',
    formItemClass: 'col-span-4',
  },
  {
    component: 'CheckboxGroup',
    componentProps: {
      name: 'cname',
      options: [
        {
          label: '生成菜单权限',
          value: 'genMenuPermissions',
        },
        {
          label: '生成前运行 [gf gen dao]',
          value: 'runDao',
        },
        {
          label: '生成后运行 [gf gen service]',
          value: 'runService',
        },
        {
          label: '生成字典选项',
          value: 'genFuncDict',
        },
        {
          label: '新增修改模态框',
          value: 'genEditModal',
        },
        {
          label: '强制覆盖',
          value: 'forcedCover',
        },
      ],
    },
    fieldName: 'options.autoOps',
    label: '高级设置',
    formItemClass: 'col-span-4',
  },
  {
    component: 'Select',
    fieldName: 'options.tree.titleColumn',
    label: '树名称字段',
    componentProps: {
      allowClear: true,
      options: [],
      placeholder: '请选择树节点显示字段',
    },
    dependencies: {
      show: (values) => values.genType === 1,
      triggerFields: ['genType'],
    },
    formItemClass: 'col-span-1',
  },
  {
    component: 'Select',
    fieldName: 'options.tree.pidColumn',
    label: '树父级字段',
    componentProps: {
      allowClear: true,
      options: [],
      placeholder: '请选择父级ID字段',
    },
    dependencies: {
      show: (values) => values.genType === 1,
      triggerFields: ['genType'],
    },
    formItemClass: 'col-span-1',
  },
  {
    component: 'Select',
    fieldName: 'options.tree.levelColumn',
    label: '树层级字段',
    componentProps: {
      allowClear: true,
      options: [],
      placeholder: '请选择层级字段',
    },
    dependencies: {
      show: (values) => values.genType === 1,
      triggerFields: ['genType'],
    },
    formItemClass: 'col-span-1',
  },
  {
    component: 'Select',
    fieldName: 'options.tree.treeColumn',
    label: '树路径字段',
    componentProps: {
      allowClear: true,
      options: [],
      placeholder: '请选择关系路径字段',
    },
    dependencies: {
      show: (values) => values.genType === 1,
      triggerFields: ['genType'],
    },
    formItemClass: 'col-span-1',
  },
  {
    component: 'TreeSelect',
    defaultValue: 0,
    fieldName: 'options.menu.pid',
    label: '上级菜单',
    rules: 'selectRequired',
    formItemClass: 'col-span-2',
  },

  {
    component: 'IconPicker',
    dependencies: {
      // 类型不为按钮时显示
      show: (values) => values.menuType !== 'F',
      triggerFields: ['menuType'],
    },
    fieldName: 'options.menu.icon',
    help: '点击搜索图标跳转到iconify & 粘贴',
    label: '菜单图标',
    componentProps: {
      placeholder: '请选择菜单图标',
      class: 'w-[300px]',
    },
    formItemClass: 'col-span-2',
  },
  {
    component: 'Input',
    fieldName: 'tableComment',
    label: '菜单名称',
    rules: 'required',
  },
  {
    component: 'InputNumber',
    fieldName: 'options.menu.sort',
    help: '排序, 数字越小越靠前',
    label: '菜单排序',
    rules: 'required',
  },
];
export const developJoinSchema: VbenFormSchema[] = [
  {
    component: 'Input',
    fieldName: 'uuid',
    label: 'uuid',
    dependencies: {
      show: () => false,
      triggerFields: [''],
    },
  },
  {
    component: 'Input',
    fieldName: 'dbName',
    label: '数据库名称',
    dependencies: {
      show: () => false,
      triggerFields: [''],
    },
  },
  {
    component: 'Input',
    fieldName: 'masterTableName',
    label: '主表名称',
    dependencies: {
      show: () => false,
      triggerFields: [''],
    },
  },
  {
    component: 'Select',
    componentProps: {
      options: tableListObj,
      defaultValue: 0,
    },
    fieldName: 'linkTable',
    label: '关联表',
    formItemClass: 'col-span-2',
    dependencies: {
      async componentProps(values) {
        if (values.dbName) {
          const res = await getGenCodesTableSelectApi({
            dbGroup: values.dbName,
          });
          tableListObj = res.items;
        }
        return {
          options: tableListObj,
        };
      },
      triggerFields: ['dbName'],
    },
  },
  {
    component: 'Input',
    fieldName: 'alias',
    label: '别名',
    formItemClass: 'col-span-2',
  },
  {
    component: 'Select',
    fieldName: 'linkMode',
    label: '关联方式',
    formItemClass: 'col-span-2',
    componentProps: {
      options: selectListObj.linkMode, 
    },
    dependencies: {
      async componentProps(values) {
        return {
          options: selectListObj.linkMode,
        };
      },  
      triggerFields: ['dbName', 'linkTable'],
    },
  },
  {
    component: 'Select',
    fieldName: 'field',
    label: '关联字段',
    formItemClass: 'col-span-2',
    componentProps: {
      options:[]
    },
    dependencies: {
      async componentProps(values) {
        if (values.dbName && values.linkTable) {
          const res = await getGenCodesColumnSelectApi({
            dbGroup: values.dbName,
            tableName: values.linkTable,
          });
          return {
            options: res.items,
          };
        }
        return {
          options: [],
        };
      },
      triggerFields: ['dbName', 'linkTable'],
    },
  },
  {
    component: 'Select',
    fieldName: 'masterField',
    label: '主表关联字段',
    formItemClass: 'col-span-2',
    componentProps: {
      // options: selectListObj.joinType,
    },dependencies: {
      async componentProps(values) {
        if (values.dbName && values.masterTableName) {
          const res = await getGenCodesColumnSelectApi({
            dbGroup: values.dbName,
            tableName: values.masterTableName,
          });
          return {
            options: res.items,
          };
        }
        return {
          options: [],
        };
      },
      triggerFields: ['dbName', 'masterTableName'],
    },
  },

  {
    component: 'Input',
    componentProps: {
    },
    dependencies: {
      // show: () => false,
      trigger: (values, formApi) => {
        const daoName = tableListObj.find(
          (item) => item.value === values.linkTable,
        )?.daoName;
        if (daoName) {
          formApi.setFieldValue('daoName', daoName);
        }
      },
      triggerFields: ['linkTable'],
    },
    fieldName: 'daoName',
    label: 'daoName',
    formItemClass: ' hidden',
  },
];
