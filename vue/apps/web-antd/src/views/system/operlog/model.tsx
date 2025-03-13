import type {
    VbenFormSchema,
  } from '@vben/common-ui';
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
    renderHttpMethodTag,
    renderJsonPreview,
  } from '#/utils/render';

  
  export const viewSchema: DescItem[] = [
    {
      field: 'operId',
      label: '日志编号',
    },
    {
      field: 'status',
      label: '操作结果',
      render(value) {
        return renderDict(value, DictEnum.SYS_COMMON_STATUS);
      },
    },
    {
      field: 'title',
      label: '操作模块',
      labelMinWidth: 80,
      render(value, { businessType }) {
        const operType = renderDict(businessType, DictEnum.SYS_OPER_TYPE);
        return (
          <div class="flex items-center">
            <Tag>{value}</Tag>
            {operType}
          </div>
        )
        ;
      },
    },
    {
      field: 'operIp',
      label: '操作信息',
      render(_, data) {
        return (
          <div class="flex items-center">
            <Tag>{data.operName}</Tag>
            {data.deptName}
            {data.operIp}
            {data.operLocation}
          </div>
        )
      },
    },
    {
      field: 'operUrl',
      label: '请求信息',
      render(_, data) {
        const { operUrl, requestMethod } = data;
        const methodTag = renderHttpMethodTag(requestMethod);
        return (
          <span>
            {methodTag} {operUrl}
          </span>
        )
        ;
      },
    },
    {
      field: 'errorMsg',
      label: '异常信息',
      render(value) {
        return (
          <span class="font-bold text-red-600">{value}</span>
        )
        ;
      },
      show: (data) => {
        return data && data.errorMsg !== '';
      },
    },
    {
      field: 'method',
      label: '方法',
    },
    /**
     * 默认word-break: break-word;会导致json预览样式异常
     */
    {
      field: 'operParam',
      label: '请求参数',
      render(value) {
        return (
          <div class="max-h-[300px] w-full overflow-y-auto">
            {renderJsonPreview(value)}
          </div>
        )
        ;
      },
    },
    {
      field: 'jsonResult',
      label: '响应参数',
      render(value) {
        return (
          <div class="max-h-[300px] w-full overflow-y-auto">
            {renderJsonPreview(value)}
          </div>
        )
        ;
      },
      show(data) {
        return data && data.jsonResult;
      },
    },
    {
      field: 'costTime',
      label: '耗时',
      render(value) {
        return (
          <span>{value} ms</span>
        )
        ;
      },
    },
    {
      field: 'operTime',
      label: '操作时间',
    },
  ];

  export const drawerSchema: VbenFormSchema[] =  [
    {
      component: 'Input',
      dependencies: {
        show: () => false,
        triggerFields: [''],
      },
      fieldName: 'configId',
      label: '参数主键',
    },
    {
      component: 'Input',
      fieldName: 'configName',
      label: '参数名称',
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'configKey',
      label: '参数键名',
      rules: 'required',
    },
    {
      component: 'Textarea',
      formItemClass: 'items-baseline',
      fieldName: 'configValue',
      label: '参数键值',
      componentProps: {
        autoSize: true,
      },
      rules: 'required',
    },
    {
      component: 'RadioGroup',
      componentProps: {
        buttonStyle: 'solid',
        options: getDictOptions(DictEnum.SYS_YES_NO),
        optionType: 'button',
      },
      defaultValue: 'N',
      fieldName: 'configType',
      label: '是否内置',
      rules: 'required',
    },
    {
      component: 'Textarea',
      fieldName: 'remark',
      formItemClass: 'items-baseline',
      label: '备注',
    },
  ];
  