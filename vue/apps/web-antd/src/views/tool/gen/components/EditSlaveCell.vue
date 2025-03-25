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
interface Props {
value?: any;
selectList: any;
uuid: string;
}

const props = withDefaults(defineProps<Props>(), {
    value: genInfoObj,
    selectList: selectListObj,
    uuid: '',
});
function getIndex() {
    if (formValue.value.options.join.length === 0) {
        return -1;
    }
    for (let i = 0; i < formValue.value.options.join.length; i++) {
        if (formValue.value.options.join[i].uuid === props.uuid) {
        return i;
        }
    }
    return -1;
}

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

const slaveColumns = computed(() => {
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
            return h('div', null, [
                renderTooltip(
                    h(
                        Button,
                        {
                            type: 'text',
                            strong: true,
                            size: 'small',
                            text: true,
                            class: 'items-center font-bold',
                            // iconPlacement: 'right',
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
                            Button,
                            {
                                strong: true,
                                size: 'small',
                                text: true,
                                iconPlacement: 'right',
                            },
                            { default: () => '列宽', icon: renderIcon("ant-design:question-circle-outlined") }
                        ),
                        '选填。设定固定值时表格生成自动计算scroll-x，未设定默认每列按100计算'
                    );
                },
                key: 'width',
                width: 50,
                dataIndex: 'width',
                customRender({record}: any) {
                    return h(InputNumber, {
                        class: 'w-[50px]',
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
    const index = getIndex();

    // 已存在直接加载
    if (index >= 0 && formValue.value.options.join[index].columns?.length > 0) {
        dataSource.value = formValue.value.options.join[index].columns;
        show.value = false;
        return;
    }
    await reloadColumns(index);
});

// 重载字段属性
async function reloadColumns(index: number) {
    show.value = true;
    dataSource.value = [];
    const join = formValue.value.options.join[index];

    getGenCodesColumnListApi({
        dbGroup: formValue.value.dbName,
        tableName: join.linkTable,
        isLink: 1,
        alias: join.alias,
    })
    .then((res) => {
        join.columns = formatColumns(res.items);
        dataSource.value = join.columns;
        console.log('vue/apps/web-antd/src/views/tool/gen/components/EditSlaveCell.vue join.columns',join.columns);
    })
    .finally(() => {
        show.value = false;
    });
}
function syncColumns(index: number) {
    console.log('vue/apps/web-antd/src/views/tool/gen/components/EditSlaveCell.vue syncColumns');
    show.value = true;
    dataSource.value = [];
    const join = formValue.value.options.join[index];
    getGenCodesColumnListApi({
        dbGroup: formValue.value.dbName,
        tableName: join.linkTable,
        isLink: 1,
        alias: join.alias,
    })
    .then((res) => {
        const columns = formatColumns(res.items);
        for (let i = 0; i < columns.length; i++) {
          // 相同字段名称和类型，保留原字段属性
          const index = join.columns.findIndex(
            (item: any) => item.name == columns[i].name && item.dataType == columns[i].dataType
          );
          if (index !== -1) {
            columns[i] = join.columns[index];
          }
        }
        join.columns = columns;
        dataSource.value = join.columns;
    })
    .finally(() => {
        show.value = false;
    });

}

const moveFieldTreeData = ref<any[]>([]);

const [MoveFieldModal, modalMoveFieldApi] = useVbenModal({
    onConfirm: function () {
        const join = formValue.value.options.join[getIndex()];
        let tmp2 = [];
        for (let i = 0; i < moveFieldTreeData.value.length; i++) {
            const index = join.columns.findIndex((item: any) => item.id == moveFieldTreeData.value[i].key);
            if (index !== -1) {
                tmp2.push(join.columns[index]);
            }
        }
        join.columns = tmp2;
        dataSource.value = join.columns;
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
    console.log('vue/apps/web-antd/src/views/tool/gen/components/EditSlaveCell.vue onMoveFieldDrop', e);
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
            <Button type="primary" @click="reloadColumns(getIndex())">重载字段</Button>
            <Button type="primary" @click="syncColumns(getIndex())">同步字段</Button>
            <Button type="primary" @click="onHandleMoveFieldButton">移动字段</Button>
        </div>
        <Table :columns="slaveColumns" :data-source="dataSource" :pagination="false" />
    </div>
    <MoveFieldModal
    :fullscreen-button="false"
    title="移动字段"
    class="w-[400px]"
  >
  <Tree draggable  @drop="onMoveFieldDrop" :tree-data="moveFieldTreeData" />
  </MoveFieldModal>
</template>