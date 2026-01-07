import type { VbenFormSchema } from '@vben/common-ui';

import { z } from '@vben/common-ui';

export const GPCreateForm: VbenFormSchema[] = [
  {
    fieldName: 'projectName',
    component: 'Input',
    label: '项目名称',
    componentProps: {
      placeholder: '',
    },
    rules: 'required',
  },
  {
    fieldName: 'desc',
    component: 'Textarea',
    label: '项目描述',
    componentProps: {
      placeholder: '',
    },
    rules: null,
  },
  {
    fieldName: 'version',
    component: 'Select',
    label: '项目版本',
    componentProps: {
      placeholder: '',
      options: [
        {
          label: '1.0.0',
          value: '1.0.0',
        },
        {
          label: '1.0.1',
          value: '1.0.1',
        },
      ],
    },
    rules: 'selectRequired',
  },
  {
    fieldName: 'status',
    component: 'RadioGroup',
    label: '选择地域',
    componentProps: {
      class: 'w-auto',
    },
    defaultValue: '1',
    rules: z.any().refine((val) => !!val, { message: '请选择地域' }),
  },

  {
    fieldName: 'status1',
    component: 'Input',
    label: '选择可用区',
    componentProps: {
      class: 'w-auto',
    },
    rules: z.any().refine((val) => !!val, { message: '请选择地域' }),
  },
];
