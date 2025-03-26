<script lang="ts" setup>
import { Table, Button, Input, type TableColumnsType, InputNumber, Space, Checkbox, Select, Cascader,Tree } from 'ant-design-vue';
import { type PropType, ref, computed, h, onMounted } from 'vue';
import { cloneDeep } from 'lodash-es';
import {
    renderDict,
    renderIcon,
    renderTooltip,
    renderDictTag,
    renderDictTags,
    renderHttpMethodTag,
    renderJsonPreview,
} from '#/utils/render';

import { selectListObj, formGridColsOptions, formGridSpanOptions, genInfoObj, formatColumns } from '../model';
import { getGenCodesColumnListApi } from '#/api/gen_codes/gen_table';
import { useVbenModal } from '@vben/common-ui';

const emit = defineEmits(['update:value']);

function getFormGridSpanOptions(cols: number) {
    if (cols < 1) {
        cols = 1;
    }
    if (cols > 4) {
        cols = 4;
    }
    for (let i = 0; i < formValue.value.masterColumns.length; i++) {
        if (!formValue.value.masterColumns[i].formGridSpan) {
            formValue.value.masterColumns[i].formGridSpan = 1;
        }
        if (formValue.value.masterColumns[i].formGridSpan > cols) {
            formValue.value.masterColumns[i].formGridSpan = cols;
        }
    }
    return formGridSpanOptions.slice(0, Math.min(cols, formGridSpanOptions.length));
}
function getFormModeOptions(type: string) {
    const options = cloneDeep(selectListObj?.formMode ?? []);
    if (options.length === 0) {
        return [];
    }
    switch (type) {
        case 'number':
            for (let i = 0; i < options.length; i++) {
                const allows = [
                    'Input',
                    'InputNumber',
                    'Radio',
                    'Select',
                    'Switch',
                    'Rate',
                    'TreeSelect',
                    'Cascader',
                ];
                if (!allows.includes(options?.[i]?.value as string)) {
                    (options[i] as any).disabled = true;
                }
            }
            break;
        default:
    }
    options.sort((a: any, b: any) => (a.disabled === b.disabled ? 0 : a.disabled ? 1 : -1));
    return options;
}
// 禁止编辑的字段，由系统维护
function isEditDisabled(row: any) {
    const disabledNames = [
        'id',
        'created_dept',
        'created_by',
        'updated_by',
        'deleted_by',
        'created_at',
        'updated_at',
        'deleted_at',
    ];
    if (disabledNames.includes(row.name)) {
        return true;
    }

    if (formValue.value.genType == 11) {
        const disabledTreeNames = ['pid', 'level', 'tree'];
        if (disabledTreeNames.includes(row.name)) {
            return true;
        }
    }
    return false;
}
interface Props {
value?: any;
selectList: any;
}

const props = withDefaults(defineProps<Props>(), {
    value: genInfoObj,
    selectList: selectListObj,
});

// 主表字段
const columnCollapse = ref(true);
const columnsCollapseData = computed(() => {
    return columnCollapse.value
        ? [
            {
                title: '字段列名',
                key: 'name',
                width: 100,
                dataIndex: 'name',  
            },
            {
                title: '字段描述',
                key: 'dc',
                dataIndex: 'dc',
                width: 100,
                customRender({text,record}: any) {
                    return h(Input, {
                        class: 'w-[100px]',
                        value: record.dc,
                        onChange: function (e: any) {
                            record.dc = e.target.value;
                        },
                    });
                },
            },
        ]
        : [
            {
                title: '字段列名',
                key: 'name',
                width: 100,
                dataIndex: 'name',
            },
            {
                title: '物理类型',
                key: 'sqlType',
                width: 80,
                dataIndex: 'sqlType',
            },
            {
                title: 'Go属性',
                key: 'goName',
                    width: 100,
                dataIndex: 'goName',
            },
            {
                title: 'Go类型',
                key: 'goType',
                width: 80,
                dataIndex: 'goType',
            },
            {
                title: 'Ts属性',
                key: 'tsName',
                width: 100,
                dataIndex: 'tsName',
            },
            {
                title: 'Ts类型',
                key: 'tsType',
                width: 80,
                dataIndex: 'tsType',
            },
            {
                title: '字段描述',
                key: 'dc',
                width: 100,
                dataIndex: 'dc',
                customRender({record}: any) {
                    return h(Input, {
                        value: record.dc,
                        class: 'w-[100px]',
                        onChange: function (e: any) {
                            record.dc = e.target.value;
                        },
                    });
                },
            },
        ];
});

const dataSource = ref([]);
const show = ref(false);
const formValue = computed({
    get() {
        return props.value;
    },
    set(value) {
        emit('update:value', value);
    },
});

const masterColumns = computed(() => {
    return [
    {
        title: '',
        key: 'id',
        dataIndex: 'id',
        width: 30,
        customRender(opt: any) {
            return opt.index + 1;
        },
    },
    {
        title(_column: any) {
            return h('div', {
                class: 'flex items-center font-bold justify-center',
            }, [
                renderTooltip(
                    h(
                        'span',
                        {
                            class: 'flex items-center font-bold',
                        },
                        [
                            h('span', { class: 'mr-1' }, '字段'),
                            renderIcon('ant-design:question-circle-outlined'),
                        ],
                    ),
                    'Go类型和属性定义取决于你在/hack/config.yaml中的配置参数',
                ),
                h(
                    Button,
                    {
                        strong: true,
                        size: 'small',
                        text: true,
                        type: 'primary',
                        style: { 'margin-left': '20px' },
                        onClick: () => (columnCollapse.value = !columnCollapse.value),
                    },
                    { default: () => (columnCollapse.value ? '展开 >>' : '折叠 <<') },
                ),
            ]);
        },
        key: 'field',
        align: 'center',
        width: 800,
        dataIndex: 'field',
        children: columnsCollapseData.value,
    },
    {
        width: 800,
        title(_column: any) {
            return h('div', {
                class: 'flex items-center font-bold justify-center',
            }, [
                renderTooltip(
                    h(
                        'div',
                        {
                        class: 'flex items-center font-bold',
                    },
                    ['新增/编辑表单', renderIcon("ant-design:question-circle-outlined")]
                ),
                '勾选编辑以后会在新增、编辑表单中显示该字段;当同时勾选列表查询时，会优先使用配置的表单组件'
                )
            ]);
        },
        key: 'edit',
        align: 'center',
        dataIndex: 'edit',
        children: [
            {
                align: 'center',
                title: '编辑',
                key: 'isEdit',
                width: 30,
                dataIndex: 'isEdit',
                customRender({record}: any) {
                    const disabled = isEditDisabled(record);
                    const checkbox = h(Checkbox, {
                        checked: record.isEdit,
                        value: record.isEdit,
                        disabled: disabled,
                        onChange: function (e: any) {
                            record.isEdit = e.target.checked;
                        },
                    });
                    if (!disabled) {
                        return checkbox;
                    }
                    return renderTooltip(checkbox, '该字段属性由系统维护，无需单独配置！');
                },
            },
            {
                title: '必填',
                key: 'required',
                width: 30,
                align: 'center',
                dataIndex: 'required',
                customRender({record}: any) {
                    return h(Checkbox, {
                        defaultChecked: record.required,
                        disabled: record.name === 'id',
                        checked: record.required,
                        onChange: function (e: any) {
                            record.required = e.target.checked;
                        },
                    });
                },
            },
            {
                title: '唯一',
                key: 'unique',
                width: 30,
                align: 'center',
                dataIndex: 'unique',
                customRender({record}: any) {
                    return h(Checkbox, {
                        defaultChecked: record.unique,
                        disabled: record.name === 'id',
                        checked: record.unique,
                        onChange: function (e: any) {
                            record.unique = e.target.checked;
                        },
                    });
                },
            },
            {
                title: '表单组件',
                key: 'formMode',
                width: 100,
                minWidth: 100,
                dataIndex: 'formMode',
                align: 'center',
                customRender({record}: any) {
                    return h(Select, {
                        consistentMenuWidth: false,
                        class: 'w-[140px]',
                        value: record.formMode,
                        options: getFormModeOptions(record.tsType),
                        defaultValue:'Input',
                        onChange: function (e: any) {
                            // console.log('vue/apps/web-antd/src/views/tool/gen/components/EditMasterCell.vue onChange',e);
                            record.formMode = e;
                        },
                    });
                },
            },
            {
                title: '绑定字典',
                key: 'dictType',
                width: 100,
                dataIndex: 'dictType',
                align: 'center',
                customRender({record}: any) {
                    if (record.dictType == 0) {
                        record.dictType = null;
                    }
                    return h(Cascader, {
                        class: 'w-[120px]',
                        placeholder: ' ',
                        filterable: true,
                        clearable: true,
                        showPath: false,
                        changeOnSelect: true,
                        checkStrategy: 'child',
                        disabled: record.name === 'id',
                        value: record.dictType,
                        options: selectListObj.dictMode ?? [],
                        onChange: function (e: any) {
                            record.dictType = e?.[0]|| '';
                            console.log('vue/apps/web-antd/src/views/tool/gen/components/EditMasterCell.vue onChange',e,record.dictType);
                        },
                    });
                },
            },
            {
                title: '验证规则',
                key: 'formRole',
                width: 100,
                dataIndex: 'formRole',
                align: 'center',
                customRender({record}: any) {
                    return h(Select, {
                        class: 'w-[120px]',
                        consistentMenuWidth: false,
                        value: record.formRole,
                        disabled: record.name === 'id',
                        options: selectListObj.formRole ?? [],
                        onChange: function (e: any) {
                            record.formRole = e;
                        },
                    });
                },
            },
            {
                align: 'center',
                title(_column:any) {
                    return h('div', {
                        class: 'font-bold justify-center',
                    }, [
                        renderTooltip(
                            h(
                                'div',
                                {
                                    class: 'flex items-center font-bold justify-center',
                                },
                                ['栅格', renderIcon("ant-design:question-circle-outlined") ]
                            ),
                            '表单每行摆放组件的个数。响应式栅格，小屏幕自动转为每行摆放一个组件。'
                        ),
                        h(Select, {
                            style: { width: '100px' },

                            size: 'small',
                            consistentMenuWidth: false,
                            value: formValue.value.options.presetStep.formGridCols,
                            options: formGridColsOptions,
                            onChange: function (e: any) {
                                formValue.value.options.presetStep.formGridCols = e;
                            },
                        }),
                    ]);
                },
                key: 'formGridSpan',
                width: 120,
                dataIndex: 'formGridSpan',
                customRender({record}: any) {
                    return h(Select, {
                        class: 'w-[120px]',
                        consistentMenuWidth: false,
                        disabled: record.name === 'id',
                        value: record.formGridSpan,
                        options: getFormGridSpanOptions(formValue.value.options.presetStep.formGridCols),
                        onChange: function (e: any) {
                            record.formGridSpan = e.target.value;
                        },
                    });
                },
            },
        ],
    },
    {
        width: 800,
        title: '列表',
        key: 'list',
        align: 'center',
        dataIndex: 'list',
        children: [
            {
                title: '列表',
                key: 'isList',
                width: 30,
                align: 'center',
                dataIndex: 'isList',
                customRender({record}: any) {
                    return h(Checkbox, {
                        defaultChecked: record.isList,
                        checked: record.isList,
                        onChange: function (e: any) {
                            record.isList = e.target.checked;
                        },
                    });
                },
            },
            {
                title: '导出',
                key: 'isExport',
                width: 30,
                align: 'center',
                dataIndex: 'isExport',
                customRender({record}: any) {
                    return h(Checkbox, {
                        defaultChecked: record.isExport,
                        checked: record.isExport,
                        onChange: function (e: any) {
                            record.isExport = e.target.checked;
                        },
                    });
                },
            },
            {
                title: '查询',
                key: 'isQuery',
                width: 30,
                align: 'center',
                dataIndex: 'isQuery',
                customRender({record}: any) {
                    return h(Checkbox, {
                        defaultChecked: record.isQuery,
                        checked: record.isQuery,
                        onChange: function (e: any) {
                            record.isQuery = e.target.checked;
                        },
                    });
                },
            },
            {
                title: '查询条件',
                align: 'center',
                key: 'queryWhere',
                width: 90,
                dataIndex: 'queryWhere',
                customRender({record}: any) {
                    return h(Select, {
                        class: 'w-[100px]',
                        consistentMenuWidth: false,
                        value: record.queryWhere,
                        disabled: record.name === 'id',
                        options: selectListObj.whereMode ?? [],
                        onChange: function (e: any) {
                            record.queryWhere = e.target.value;
                        },
                    });
                },
            },
            {
                title: '排列方式',
                align: 'center',
                key: 'align',
                width: 80,
                dataIndex: 'align',
                customRender({record}: any) {
                    return h(Select, {
                        class: 'w-[80px]',
                        consistentMenuWidth: false,
                        value: record.align,
                        options: selectListObj.tableAlign ?? [],
                        onChange: function (e: any) {
                            record.align = e.target.value;
                        },
                    });
                },
            },
            {
                title(_column:any) {
                    return renderTooltip(
                        h(
                            'div',
                            {
                                class: 'flex items-center font-bold justify-center',
                            },
                            [ '列宽', renderIcon("ant-design:question-circle-outlined") ]
                        ),
                        '选填。设定固定值时表格生成自动计算scroll-x，未设定默认每列按100计算'
                    );
                },
                key: 'width',
                width: 80,
                dataIndex: 'width',
                customRender({record}: any) {
                    return h(InputNumber, {
                        class: 'w-[60px]',
                        value: record.width,
                        placeholder: ' ',
                        min: -1,
                        max: 2000,
                        controls: false,
                        onChange: function (e: any) {
                            record.width = e.target.value;
                        },
                    });
                },
            },
        ],
    },
]});
onMounted(async () => {
    if (formValue.value.masterColumns.length === 0) {
        await reloadColumns();
    }else {
        dataSource.value = formValue.value.masterColumns;
    }
});

// 重载字段属性
async function reloadColumns() {
    show.value = true;
    dataSource.value = [];

    getGenCodesColumnListApi({
        dbGroup: formValue.value.dbName,
        tableName: formValue.value.tableName,
        isLink: 0,
        alias: '',
    })
    .then((res) => {
        formValue.value.masterColumns = formatColumns(res.items);
        dataSource.value = formValue.value.masterColumns;
        console.log('vue/apps/web-antd/src/views/tool/gen/components/EditMasterCell.vue formValue.value.masterColumns',formValue.value.masterColumns);
    })
    .finally(() => {
        show.value = false;
    });
}
function syncColumns() {
    console.log('vue/apps/web-antd/src/views/tool/gen/components/EditMasterCell.vue syncColumns');
    show.value = true;
    dataSource.value = [];
    getGenCodesColumnListApi({
        dbGroup: formValue.value.dbName,
        tableName: formValue.value.tableName,
        isLink: 0,
        alias: '',
    })
    .then((res) => {
        const columns = formatColumns(res.items);
        for (let i = 0; i < columns.length; i++) {
          // 相同字段名称和类型，保留原字段属性
          const index = formValue.value.masterColumns.findIndex(
            (item: any) => item.name == columns[i].name && item.dataType == columns[i].dataType
          );
          if (index !== -1) {
            columns[i] = formValue.value.masterColumns[index];
          }
        }
        formValue.value.masterColumns = columns;
        dataSource.value = formValue.value.masterColumns;
    })
    .finally(() => {
        show.value = false;
    });

}

const moveFieldTreeData = ref([]);

const [MoveFieldModal, modalMoveFieldApi] = useVbenModal({
    onConfirm: function () {
        let tmp2 = [];
        for (let i = 0; i < moveFieldTreeData.value.length; i++) {
            const index = formValue.value.masterColumns.findIndex((item: any) => item.id == moveFieldTreeData.value[i].key);
            if (index !== -1) {
                tmp2.push(formValue.value.masterColumns[index]);
            }
        }
        formValue.value.masterColumns = tmp2;
        dataSource.value = formValue.value.masterColumns;
        modalMoveFieldApi.close();
    },
});

function onHandleMoveFieldButton( ) {
    moveFieldTreeData.value = dataSource.value.map((item: any) => ({
        title: item.name + ' - ' + item.dc,
        key: item.id,
        children: [],
    }));
    modalMoveFieldApi.open();
}

function onMoveFieldDrop(e: any) {
    console.log('vue/apps/web-antd/src/views/tool/gen/components/EditMasterCell.vue onMoveFieldDrop', e);
    const { dragNode, node, dropPosition,dragNodesKeys,dropToGap } = e;
    const index = moveFieldTreeData.value.findIndex((item: any) => item.key == dragNode.key);
    if (index !== -1 ) {
        const temp = moveFieldTreeData.value[index];
        moveFieldTreeData.value.splice(index, 1)
        let dropPosition2 = dropPosition;
        if (dropToGap) {
            if (dropPosition2 < 0) {
                dropPosition2 = 0;
            }
            if (index > dropPosition2) {
                moveFieldTreeData.value.splice(dropPosition2, 0, temp);
            } else {
                moveFieldTreeData.value.splice(dropPosition2-1, 0, temp);
            }
        }else {
            moveFieldTreeData.value.splice(dropPosition, 0, temp);
        }
    }
}
</script>

<template>
    <div class="flex flex-col gap-2">
        <div class="flex flex-row gap-2">
            <Button type="primary" @click="reloadColumns">重载字段</Button>
            <Button type="primary" @click="syncColumns">同步字段</Button>
            <Button type="primary" @click="onHandleMoveFieldButton">移动字段</Button>
        </div>
        <Table bordered :columns="masterColumns" :data-source="dataSource" :pagination="false" />
    </div>
    <MoveFieldModal
    :fullscreen-button="false"
    title="移动字段"
    class="w-[400px]"
  >
  <Tree draggable  @drop="onMoveFieldDrop" :tree-data="moveFieldTreeData" />
  </MoveFieldModal>
</template>