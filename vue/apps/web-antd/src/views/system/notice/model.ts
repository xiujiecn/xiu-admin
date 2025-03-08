import type {
    VbenFormSchema,
  } from '@vben/common-ui';
  import { z } from '@vben/common-ui';
  import { getDictOptions } from '#/utils/dict';
  import { DictEnum } from '@vben/constants';
  import { getPopupContainer } from '@vben/utils';
  
  export const drawerSchema: VbenFormSchema[] =  [
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
      component: 'RichTextarea',
      componentProps: {
        width: '100%',
      },
      fieldName: 'noticeContent',
      label: '公告内容',
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'remark',
      label: '备注',
    },
  ];
  