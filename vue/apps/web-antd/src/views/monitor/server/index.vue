<!--
 * @description 服务器监控页面
 * @Link  https://github.com/xiujiecn/xiu-admin
 * @Copyright  Copyright (c) 2025 XiuAdmin CLI
 * @Author  Lxj <li@xiujie.cn>
 * @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
 * @date 2024-03-21
-->
<script lang="ts" setup>
import { Card, CardGrid, Descriptions, DescriptionsItem } from 'ant-design-vue';
import type { EchartsUIType } from '@vben/plugins/echarts';
import { defineComponent, onActivated, onMounted, ref, watch, reactive, computed, onUnmounted } from 'vue';

import { EchartsUI, useEcharts, echarts } from '@vben/plugins/echarts';
import { useWebSocket } from '@vueuse/core';
import { useAccessStore } from '@vben/stores';
import dayjs from 'dayjs';


const accessStore = useAccessStore();

type Record = {
    e: string;
    d: any;
}

function getWsServerUrl() {
    return 'ws://' + window.location.host + '/api/monitor/ws?access_token=' + encodeURIComponent(String(accessStore.accessToken));
}
const state = reactive({
    server: getWsServerUrl(),
    recordList: [] as Record[],
});

const { status, data, send, close, open } = useWebSocket(state.server, {
    autoReconnect: true,
    heartbeat: false,

});

const getIsOpen = computed(() => status.value === 'OPEN');

function handlerSend() {
    send(JSON.stringify({
        e: 'joins',
        d: {
            names: ["ws:tag:monitor:server"]
        }
    }));
}

watch(status, (newVal) => {
    console.log('vue/apps/web-antd/src/views/monitor/server/index.vue status', newVal);
    if (newVal === 'OPEN') {
        handlerSend();
    }
});

watch(data, (newVal) => {
    if (newVal) {

        // {"e":"ws:event:monitor:server:go","d":{"goName":"Golang","goOs":"darwin","arch":"arm64","goVersion":"go1.23.5","startTime":"2006-01-02 15:04:05","runTime":411,"rootPath":"/opt/homebrew/Cellar/go/1.23.5/libexec","pwd":"/Users/lixiujie/dev_test/dev_go/xiujie_iot/server","goroutine":"27","goMem":"35.39 MB","goSize":"6.73 MB"}}


        // console.log('vue/apps/web-antd/src/views/monitor/server/index.vue data', newVal);
        const res = JSON.parse(newVal);
        if (typeof res.d === 'string' && !res.d.startsWith('{')) {
            res.d = atob(res.d);
        }
        if (typeof res.d === 'string' && res.d.startsWith('{')) {
            res.d = JSON.parse(res.d);
        }
        console.log('vue/apps/web-antd/src/views/monitor/server/index.vue res', res);
        if (res.e === 'ws:event:monitor:server:go') {
            handleGoInfo(res.d);
        }
        if (res.e === 'ws:event:monitor:server:host') {
            handleHostInfo(res.d);
        }
        if (res.e === 'ws:event:monitor:server:cpu') {
            handleCpuInfo(res.d);
        }
        if (res.e === 'ws:event:monitor:server:mem') {
            handleMemInfo(res.d);
        }
        if (res.e === 'ws:event:monitor:server:disk') {
            handleDiskInfo(res.d);
        }
        if (res.e === 'ws:event:monitor:server:net') {
            handleNetInfo(res.d);
        }
        if (res.e === 'ws:event:monitor:server:sysLoad') {
            handleSysLoadInfo(res.d);
        }
    }
});
function handleSysLoadInfo(data: any) {
    // {
    // "e": "ws:event:monitor:server:sysLoad",
    // "d": {
    //     "load1": 7.787109375,
    //     "load15": 5.5693359375,
    //     "load5": 6.12548828125
    // }
    sysInfo.value.cpuAvg1 = parseInt(data.load1);
    sysInfo.value.cpuAvg5 = parseInt(data.load5);
    sysInfo.value.cpuAvg15 = parseInt(data.load15);
}
function handleNetInfo(data: any) {
    console.log('vue/apps/web-antd/src/views/monitor/server/index.vue handleNetInfo', data);
}
function handleDiskInfo(data: any) {

    // {
    //     "path": "/",
    //     "fstype": "apfs",
    //     "total": 994662584320,
    //     "free": 30871126016,
    //     "used": 963791458304,
    //     "usedPercent": 96.89632177759003,
    //     "inodesTotal": 301887432,
    //     "inodesUsed": 411592,
    //     "inodesFree": 301475840,
    //     "inodesUsedPercent": 0.1363395611646397
    // }
    sysInfo.value.diskTotal = data.total;
    sysInfo.value.diskUsed = data.used;
    sysInfo.value.diskUsedPercent = parseInt(data.usedPercent);
    diskEcharts.getChartInstance()?.setOption({
        series: [
            {
                data: [
                    { value: sysInfo.value.diskUsedPercent, name: '磁盘使用率' }
                ]
            }
        ]
    });
    diskRunEchartsData.value.name.push(dayjs().format('HH:mm:ss'));
    diskRunEchartsData.value.value.push(sysInfo.value.diskUsedPercent);
    if (diskRunEchartsData.value.name.length > 30) {
        diskRunEchartsData.value.name.shift();
        diskRunEchartsData.value.value.shift();
    }
    diskRunEcharts.getChartInstance()?.setOption({
        xAxis: {
            data: diskRunEchartsData.value.name,
        },
        series: [
            {
                data: diskRunEchartsData.value.value,
            }
        ]
    });
}
function handleMemInfo(data: any) {

    // {
    //     "active": 24972836864,
    //     "anonHugePages": 0,
    //     "available": 25606307840,
    //     "buffers": 0,
    //     "cached": 0,
    //     "commitLimit": 0,
    //     "committedAS": 0,
    //     "dirty": 0,
    //     "free": 666157056,
    //     "goUsed": 40781064,
    //     "highFree": 0,
    //     "highTotal": 0,
    //     "hugePageSize": 0,
    //     "hugePagesFree": 0,
    //     "hugePagesRsvd": 0,
    //     "hugePagesSurp": 0,
    //     "hugePagesTotal": 0,
    //     "inactive": 24940150784,
    //     "laundry": 0,
    //     "lowFree": 0,
    //     "lowTotal": 0,
    //     "mapped": 0,
    //     "pageTables": 0,
    //     "shared": 0,
    //     "slab": 0,
    //     "sreclaimable": 0,
    //     "sunreclaim": 0,
    //     "swapCached": 0,
    //     "swapFree": 0,
    //     "swapTotal": 0,
    //     "total": 68719476736,
    //     "used": 43113168896,
    //     "usedPercent": 62.73791790008545,
    //     "vmallocChunk": 0,
    //     "vmallocTotal": 0,
    //     "vmallocUsed": 0,
    //     "wired": 3992502272,
    //     "writeBack": 0,
    //     "writeBackTmp": 0
    // }
    sysInfo.value.memTotal = data.total;
    sysInfo.value.memUsed = data.used;
    sysInfo.value.memUsage = parseInt(data.usedPercent);
    sysInfo.value.available = data.available;
    sysInfo.value.goUsed = data.goUsed;
    memEcharts.getChartInstance()?.setOption({
        series: [
            {

                data: [
                    { value: sysInfo.value.memUsage, name: '内存使用率' }
                ]
            }
        ]
    });
    memRunEchartsData.value.name.push(dayjs().format('HH:mm:ss'));
    memRunEchartsData.value.value.push(sysInfo.value.memUsage);
    if (memRunEchartsData.value.name.length > 30) {
        memRunEchartsData.value.name.shift();
        memRunEchartsData.value.value.shift();
    }
    memRunEcharts.getChartInstance()?.setOption({
        xAxis: {
            data: memRunEchartsData.value.name,
        },
        series: [
            {
                data: memRunEchartsData.value.value,
            }
        ]
    });
}
function handleCpuInfo(data: any) {
    // {
    //     "e": "ws:event:monitor:server:cpu",
    //     "d": {
    //         "Number": 1,
    //         "Cores": 10,
    //         "UsedPercent": [
    //             26.666666666733
    //         ],
    //         "ModelName": "Apple M1 Max"
    //     }
    // }
    sysInfo.value.cpuNum = data.Number;
    sysInfo.value.cpuCores = data.Cores;
    sysInfo.value.cpuUsed = parseInt(data.UsedPercent[0]);
    cpuEcharts.getChartInstance()?.setOption({
        series: [
            {
                data: [
                    { value: sysInfo.value.cpuUsed, name: 'CPU使用率' }
                ]
            }
        ]
    });

    cpuRunEchartsData.value.name.push(dayjs().format('HH:mm:ss'));
    cpuRunEchartsData.value.value.push(sysInfo.value.cpuUsed);
    if (cpuRunEchartsData.value.name.length > 30) {
        cpuRunEchartsData.value.name.shift();
        cpuRunEchartsData.value.value.shift();
    }
    cpuRunEcharts.getChartInstance()?.setOption({
        xAxis: {
            data: cpuRunEchartsData.value.name,
        },
        series: [
            {
                data: cpuRunEchartsData.value.value,
            }
        ]
    });


}
function handleHostInfo(data: any) {
    // console.log('vue/apps/web-antd/src/views/monitor/server/index.vue handleHostInfo', data);
    //     {
    //     "e": "ws:event:monitor:server:host",
    //     "d": {
    //         "bootTime": "2025-03-13T19:09:52+08:00",
    //         "hostId": "78e68867-0a8c-5b8b-b9e1-a91eafffc273",
    //         "hostname": "ds-m1-macbookpro.local",
    //         "intranet_ip": "192.168.124.8",
    //         "kernelArch": "arm64",
    //         "kernelVersion": "24.3.0",
    //         "os": "darwin",
    //         "platform": "darwin",
    //         "platformFamily": "Standalone Workstation",
    //         "platformVersion": "15.3.2",
    //         "procs": 969,
    //         "public_ip": "140.245.124.140",
    //         "uptime": 338504,
    //         "virtualizationRole": "",
    //         "virtualizationSystem": ""
    //     }
    // }    
    hostData.value.bootTime = data.bootTime;
    hostData.value.hostId = data.hostId;
    hostData.value.hostname = data.hostname;
    hostData.value.intranet_ip = data.intranet_ip;
    hostData.value.kernelArch = data.kernelArch;
    hostData.value.kernelVersion = data.kernelVersion;
    hostData.value.os = data.os;
    hostData.value.platform = data.platform;
    hostData.value.platformFamily = data.platformFamily;
    hostData.value.platformVersion = data.platformVersion;
    hostData.value.procs = data.procs;
    hostData.value.public_ip = data.public_ip;
    hostData.value.uptime = data.uptime;
    hostData.value.virtualizationRole = data.virtualizationRole;
    hostData.value.virtualizationSystem = data.virtualizationSystem;
}
function handleGoInfo(data: any) {
    // console.log('vue/apps/web-antd/src/views/monitor/server/index.vue handleGoInfo', data);
    goInfoData.value.goName = data.goName;
    goInfoData.value.goOs = data.goOs;
    goInfoData.value.arch = data.arch;
    goInfoData.value.goVersion = data.goVersion;
    goInfoData.value.startTime = data.startTime;
    goInfoData.value.runTime = data.runTime;
    goInfoData.value.rootPath = data.rootPath;
    goInfoData.value.pwd = data.pwd;
    goInfoData.value.goroutine = data.goroutine;
    goInfoData.value.goMem = data.goMem;
    goInfoData.value.goSize = data.goSize;
    goInfoData.value.intranet_ip = data.intranet_ip;
}

// watchEffect(() => {
//     if (data.value) {
//         try {
//             console.log('vue/apps/web-antd/src/views/monitor/server/index.vue data.value', data.value);
//             // const res = JSON.parse(data.value);
//             // state.recordList.push(res);
//         } catch (error) {
//             console.log('vue/apps/web-antd/src/views/monitor/server/index.vue error', error);
//         }
//     }
// });

const sysInfo = ref({
    cpuNum: 0,
    cpuCores: 0,
    cpuUsed: 0,
    cpuAvg1: 0,
    cpuAvg5: 0,
    cpuAvg15: 0,
    memUsed: 0,
    memTotal: 0,
    available: 0,
    goUsed: 0,
    memUsage: 0,
    diskTotal: 0,
    diskUsed: 0,
    diskUsedPercent: 0,
});

const hostData = ref({
    "bootTime": "2025-03-24T12:12:13+08:00",
    "hostId": "8be74418-1a53-4208-be22-9c126d89yydd",
    "hostname": "ubuntudev",
    "intranet_ip": "192.168.158.100",
    "kernelArch": "x86_64",
    "kernelVersion": "6.8.0-52-generic",
    "os": "linux",
    "platform": "centos",
    "platformFamily": "rhel",
    "platformVersion": "24.04.1",
    "procs": 138,
    "public_ip": "192.144.160.105",
    "uptime": 3278,
    "virtualizationRole": "guest",
    "virtualizationSystem": ""
});
const goInfoData = ref({
    "goOs": "-", "arch": "-", "goVersion": "-", "goMem": "-", "goName": "-", "goSize": "-", "goroutine": '-', "pwd": "-", "rootPath": "-",
    "runTime": '', "startTime": "-", "intranet_ip": "-"
});
function memorySizeFormat(size: any) {
    if (size === null || size === undefined) return ''
    size = parseFloat(size);
    let rank = 0;
    let rankchar = 'Bytes';
    while (size > 1024 && rankchar != 'TB') {
        size = size / 1024;
        rank++;
        if (rank == 1) {
            rankchar = 'KB';
        } else if (rank == 2) {
            rankchar = 'MB';
        } else if (rank == 3) {
            rankchar = 'GB';
        } else if (rank == 4) {
            rankchar = 'TB';
        }
    }
    return size.toFixed(2) + ' ' + rankchar;
}
function timeFormat(second: any) {
    if (!second) return '-'
    second = parseFloat(second);
    let rank = 0;
    let rankchar = '秒';
    while ((second > 60 && rankchar != '小时' && rankchar != '天') || (second > 24 && rankchar == '小时')) {
        if (rankchar == '小时') {
            second = second / 24;
        } else {
            second = second / 60;
        }
        rank++;
        if (rank == 1) {
            rankchar = '分';
        } else if (rank == 2) {
            rankchar = '小时';
        } else if (rank == 3) {
            rankchar = '天';
        }
    }
    return second.toFixed(2) + ' ' + rankchar;
}

const cpuRunEchartsData: any = ref({
    name: [],
    value: [],
});

const memRunEchartsData: any = ref({
    name: [],
    value: [],
});

const diskRunEchartsData: any = ref({
    name: [],
    value: [],
});


const cpuHtmlRef = ref<EchartsUIType>();
const memHtmlRef = ref<EchartsUIType>();
const diskHtmlRef = ref<EchartsUIType>();
const cpuRunHtmlRef = ref<EchartsUIType>();
const memRunHtmlRef = ref<EchartsUIType>();
const diskRunHtmlRef = ref<EchartsUIType>();
const cpuEcharts = useEcharts(cpuHtmlRef);
const memEcharts = useEcharts(memHtmlRef);
const diskEcharts = useEcharts(diskHtmlRef);
const cpuRunEcharts = useEcharts(cpuRunHtmlRef);
const memRunEcharts = useEcharts(memRunHtmlRef);
const diskRunEcharts = useEcharts(diskRunHtmlRef);

type cpuEChartsOption = Parameters<typeof cpuEcharts.renderEcharts>['0'];
type memEChartsOption = Parameters<typeof memEcharts.renderEcharts>['0'];
type diskEChartsOption = Parameters<typeof diskEcharts.renderEcharts>['0'];
type cpuRunEChartsOption = Parameters<typeof cpuRunEcharts.renderEcharts>['0'];
type memRunEChartsOption = Parameters<typeof memRunEcharts.renderEcharts>['0'];
type diskRunEChartsOption = Parameters<typeof diskRunEcharts.renderEcharts>['0'];
const cpuOptions: cpuEChartsOption = {
    series: [
        {
            type: 'gauge',
            name: 'CPU',
            radius: '90%', //修改表盘大小
            title: {
                show: true, //控制表盘title(今日预计用电量)字体是否显示
                fontSize: 12, //控制表盘title(今日预计用电量)字体大小
                'color': 'green',           		//控制表盘title(今日预计用电量)字体颜色
                offsetCenter: [-2, '30%'], //设置表盘title(今日预计用电量)位置
            },
            axisLine: {
                lineStyle: {
                    show: true,
                    with: 25,
                    // 属性lineStyle控制线条样式
                    color: [
                        [0.3, '#4dabf7'],
                        [0.6, '#69db7c'],
                        [0.8, '#ffa94d'],
                        [1, '#ff6b6b'],
                    ],
                },
            },
            axisTick: {
                distance: 0,
                length: 4,
                lineStyle: {
                    color: 'auto',
                    width: 1
                }
            },
            axisLabel: {
                distance: 12,
                color: '#888',
                fontSize: 12
            },
            splitLine: { // 分割线
                length: 5,
                distance: 2,
                lineStyle: {
                    width: 1,
                    color: 'auto'
                }
            },
            splitNumber: 5, //分割线之间的刻度

            detail: {
                valueAnimation: true,
                formatter: '{value}%',
                textStyle: {
                    fontSize: 20,
                    color: 'red',
                },
                offsetCenter: ['0', '80%'], //表盘数据(30%)位置
            },
            data: [
                {
                    value: sysInfo.value.cpuUsed,
                    name: 'CPU使用率',
                    color: 'inherit'
                },
            ],
        }
    ],

};
const memOptions: memEChartsOption = {
    tooltip: {
        formatter: '{a} <br/>{b} : {c}%',
    },
    series: [
        {
            type: 'gauge',
            name: '内存',
            radius: '90%', //修改表盘大小
            title: {
                show: true, //控制表盘title(今日预计用电量)字体是否显示
                fontSize: 12, //控制表盘title(今日预计用电量)字体大小
                'color': 'green',           		//控制表盘title(今日预计用电量)字体颜色
                offsetCenter: [-2, '30%'], //设置表盘title(今日预计用电量)位置
            },
            axisLine: {
                lineStyle: {
                    show: true,
                    with: 25,
                    // 属性lineStyle控制线条样式
                    color: [
                        [0.3, '#4dabf7'],
                        [0.6, '#69db7c'],
                        [0.8, '#ffa94d'],
                        [1, '#ff6b6b'],
                    ],
                },
            },
            axisTick: {
                distance: 0,
                length: 4,
                lineStyle: {
                    color: 'auto',
                    width: 1
                }
            },
            axisLabel: {
                distance: 12,
                color: '#888',
                fontSize: 12
            },
            splitLine: { // 分割线
                length: 5,
                distance: 2,
                lineStyle: {
                    width: 1,
                    color: 'auto'
                }
            },
            splitNumber: 5, //分割线之间的刻度

            detail: {
                valueAnimation: true,
                formatter: '{value}%',
                textStyle: {
                    fontSize: 20,
                    color: 'red',
                },
                offsetCenter: ['0', '80%'], //表盘数据(30%)位置
            },
            data: [
                {
                    value: sysInfo.value.memUsed,
                    name: '内存使用率',
                },
            ],
        },
    ],
}
const diskOptions: diskEChartsOption = {
    tooltip: {
        formatter: '{a} <br/>{b} : {c}%',
    },
    series: [
        {
            type: 'gauge',
            name: '磁盘',
            radius: '90%', //修改表盘大小
            title: {
                show: true, //控制表盘title(今日预计用电量)字体是否显示
                fontSize: 12, //控制表盘title(今日预计用电量)字体大小
                'color': 'green',           		//控制表盘title(今日预计用电量)字体颜色
                offsetCenter: [-2, '30%'], //设置表盘title(今日预计用电量)位置
            },
            axisLine: {
                lineStyle: {
                    show: true,
                    with: 25,
                    // 属性lineStyle控制线条样式
                    color: [
                        [0.3, '#4dabf7'],
                        [0.6, '#69db7c'],
                        [0.8, '#ffa94d'],
                        [1, '#ff6b6b'],
                    ],
                },
            },
            axisTick: {
                distance: 0,
                length: 4,
                lineStyle: {
                    color: 'auto',
                    width: 1
                }
            },
            axisLabel: {
                distance: 12,
                color: '#888',
                fontSize: 12
            },
            splitLine: { // 分割线
                length: 5,
                distance: 2,
                lineStyle: {
                    width: 1,
                    color: 'auto'
                }
            },
            splitNumber: 5, //分割线之间的刻度

            detail: {
                valueAnimation: true,
                formatter: '{value}%',
                textStyle: {
                    fontSize: 20,
                    color: 'red',
                },
                offsetCenter: ['0', '80%'], //表盘数据(30%)位置
            },
            data: [
                {
                    value: sysInfo.value.diskUsedPercent,
                    name: '磁盘使用率',
                },
            ],
        },
    ],
};
const cpuRunOptions: cpuRunEChartsOption = {
    tooltip: {
        trigger: 'axis',
    },
    grid: {
        top: 15,
        bottom: 15,
        left: 15,
        right: 30,
        containLabel: true
    },
    xAxis: {
        type: 'category',
        boundaryGap: false,
        splitLine: {
            show: false
        }
    },
    yAxis: {
        type: 'value',
        boundaryGap: [0, '100%'],
        axisLabel: {
            formatter: '{value}%'
        },
        splitLine: {
            show: false
        }
    },
    series: [
        {
            name: 'CPU使用率',
            type: 'line',
            showSymbol: false,
            data: [],
            smooth: true,
            lineStyle: {
                width: 0
            },
            areaStyle: {
                opacity: 0.8,
                color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                    {
                        offset: 0,
                        color: 'rgb(128, 255, 165)'
                    },
                    {
                        offset: 1,
                        color: 'rgb(1, 191, 236)'
                    }
                ])
            }
        }
    ]
};
const memRunOptions: memRunEChartsOption = {
    tooltip: {
        trigger: 'axis',
    },
    grid: {
        top: 15,
        bottom: 15,
        left: 15,
        right: 30,
        containLabel: true
    },
    xAxis: {
        type: 'category',
        boundaryGap: false,
        splitLine: {
            show: false
        }
    },
    yAxis: {
        type: 'value',
        boundaryGap: [0, '100%'],
        axisLabel: {
            formatter: '{value}%'
        },
        splitLine: {
            show: false
        }
    },
    series: [
        {
            name: '内存使用率',
            type: 'line',
            showSymbol: false,
            data: [],
            smooth: true,
            lineStyle: {
                width: 0
            },
            areaStyle: {
                opacity: 0.8,
                color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                    {
                        offset: 0,
                        color: 'rgb(55, 162, 255)'
                    },
                    {
                        offset: 1,
                        color: 'rgb(116, 21, 219)'
                    }
                ])
            }
        }
    ]
};
const diskRunOptions: diskRunEChartsOption = {
    tooltip: {
        trigger: 'axis',
    },
    grid: {
        top: 15,
        bottom: 15,
        left: 15,
        right: 30,
        containLabel: true
    },
    xAxis: {
        type: 'category',
        boundaryGap: false,
        splitLine: {
            show: false
        }
    },
    yAxis: {
        type: 'value',
        boundaryGap: [0, '100%'],
        axisLabel: {
            formatter: '{value}%'
        },
        splitLine: {
            show: false
        }
    },
    series: [
        {
            name: '磁盘使用率',
            type: 'line',
            showSymbol: false,
            data: [],
            smooth: true,
            lineStyle: {
                width: 0
            },
            areaStyle: {
                opacity: 0.8,
                color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                    {
                        offset: 0,
                        color: 'rgb(0, 221, 255)'
                    },
                    {
                        offset: 1,
                        color: 'rgb(77, 119, 255)'
                    }
                ])
            }
        }
    ]
};
function setCpuEcharts() {
    cpuEcharts.renderEcharts(cpuOptions);
}
function setMemEcharts() {
    memEcharts.renderEcharts(memOptions);
}
function setDiskEcharts() {
    diskEcharts.renderEcharts(diskOptions);
}
function setCpuRunEcharts() {
    cpuRunEcharts.renderEcharts(cpuRunOptions);
}
function setMemRunEcharts() {
    memRunEcharts.renderEcharts(memRunOptions);
}
function setDiskRunEcharts() {
    diskRunEcharts.renderEcharts(diskRunOptions);
}
let checkI: NodeJS.Timeout | undefined;

function reconnect() {
    open();
    // if (checkI) {
    //     clearInterval(checkI);
    //     checkI = undefined;
    // }
    // checkI = setInterval(() => {
    //     console.log('vue/apps/web-antd/src/views/monitor/server/index.vue reconnect checkI', status.value); 
    //     if (status.value !== 'OPEN') {
    //         try {
    //             close();
    //             open();
    //         } catch (error) {
    //             console.log('vue/apps/web-antd/src/views/monitor/server/index.vue reconnect error', error);
    //         }
    //     }
    // }, 5000);
}


onMounted(() => {
    setCpuEcharts();
    setMemEcharts();
    setDiskEcharts();
    setCpuRunEcharts();
    setMemRunEcharts();
    setDiskRunEcharts();
    reconnect();
});
onUnmounted(() => {
    if (checkI) {
        clearInterval(checkI);
        checkI = undefined;
    }
});
onActivated(() => {
    cpuEcharts.resize();
    memEcharts.resize();
    diskEcharts.resize();
    cpuRunEcharts.resize();
    memRunEcharts.resize();
    diskRunEcharts.resize();
});
</script>
<template>
    <div class="p-5 flex flex-col gap-3">
        <div class="grid grid-cols-1 gap-3 md:grid-cols-1 lg:grid-cols-3">
            <Card class="w-full">
                <CardGrid style="width: 50%;">
                    <table cellspacing="0" style="width: 100%">
                        <tbody>
                            <tr>
                                <td>
                                    <div class="cell-card">CPU数: </div>
                                </td>
                                <td>
                                    <div class="cell-card">{{ sysInfo.cpuNum }}</div>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <div class="cell-card">核心数: </div>
                                </td>
                                <td>
                                    <div class="cell-card">{{ sysInfo.cpuCores }}</div>
                                </td>
                            </tr>

                            <tr>
                                <td>
                                    <div class="cell-card">使用率:</div>
                                </td>
                                <td>
                                    <div class="cell-card">{{ sysInfo.cpuUsed }}%</div>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <div class="cell-card">LA5:</div>
                                </td>
                                <td>
                                    <div class="cell-card">{{ sysInfo.cpuAvg5 }}%</div>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <div class="cell-card">LA15:</div>
                                </td>
                                <td>
                                    <div class="cell-card">{{ sysInfo.cpuAvg15 }}%</div>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </CardGrid>
                <CardGrid style="width: 50%;">
                    <EchartsUI ref="cpuHtmlRef" height="200px" width="100%" />
                </CardGrid>

            </Card>
            <Card class="w-full">
                <CardGrid style="width: 50%;">
                    <table cellspacing="0" style="width: 100%">
                        <tbody>
                            <tr>
                                <td>
                                    <div class="cell-card">内存总数:</div>
                                </td>
                                <td>
                                    <div class="cell-card">{{ memorySizeFormat(sysInfo.memTotal) }}</div>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <div class="cell-card">已使用:</div>
                                </td>
                                <td>
                                    <div class="cell-card">{{ memorySizeFormat(sysInfo.memUsed) }}</div>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <div class="cell-card">剩余:</div>
                                </td>
                                <td>
                                    <div class="cell-card">{{ memorySizeFormat(sysInfo.available) }}</div>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <div class="cell-card">系统使用:</div>
                                </td>
                                <td>
                                    <div class="cell-card">{{ memorySizeFormat(sysInfo.goUsed) }}</div>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <div class="cell-card">使用率:</div>
                                </td>
                                <td>
                                    <div class="cell-card">{{ sysInfo.memUsage }}%</div>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </CardGrid>
                <CardGrid style="width: 50%;">
                    <EchartsUI ref="memHtmlRef" height="200px" width="100%" />
                </CardGrid>

            </Card>

            <Card class="w-full">
                <CardGrid style="width: 50%;">
                    <table cellspacing="0" style="width: 100%">
                        <tbody>
                            <tr>
                                <td>
                                    <div class="cell-card">磁盘容量:</div>
                                </td>
                                <td>
                                    <div class="cell-card">{{ memorySizeFormat(sysInfo.diskTotal) }}</div>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <div class="cell-card">已使用:</div>
                                </td>
                                <td>
                                    <div class="cell-card">{{ memorySizeFormat(sysInfo.diskUsed) }}</div>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <div class="cell-card">使用率:</div>
                                </td>
                                <td>
                                    <div class="cell-card">{{ sysInfo.diskUsedPercent }}%</div>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </CardGrid>
                <CardGrid style="width: 50%;">
                    <EchartsUI ref="diskHtmlRef" height="200px" width="100%" />
                </CardGrid>
            </Card>

            <Card class="w-full" title="CPU运行情况">
                <CardGrid style="width: 100%; margin: 0 auto; padding: 0;">
                    <EchartsUI ref="cpuRunHtmlRef" height="300px" width="100%" />
                </CardGrid>
            </Card>
            <Card class="w-full" title="内存运行情况">
                <CardGrid style="width: 100%; margin: 0 auto; padding: 0;">
                    <EchartsUI ref="memRunHtmlRef" height="300px" width="100%" />
                </CardGrid>
            </Card>
            <Card class="w-full" title="磁盘运行情况">
                <CardGrid style="width: 100%; margin: 0 auto; padding: 0;">
                    <EchartsUI ref="diskRunHtmlRef" height="300px" width="100%" />
                </CardGrid>
            </Card>
        </div>
        <div class="grid grid-cols-1 gap-3 md:grid-cols-1 lg:grid-cols-1">
            <Card class="w-full" title="运行环境信息">
                <Descriptions :column="3">
                    <DescriptionsItem label="操作系统">{{ hostData.os }}</DescriptionsItem>
                    <DescriptionsItem label="启动时间">{{ goInfoData.startTime }}</DescriptionsItem>
                    <DescriptionsItem label="运行时长">{{ timeFormat(goInfoData.runTime) }}</DescriptionsItem>
                    <DescriptionsItem label="运行内存">{{ goInfoData.goMem }}</DescriptionsItem>
                    <DescriptionsItem label="系统架构">{{ goInfoData.arch }}</DescriptionsItem>
                    <DescriptionsItem label="语言环境">{{ goInfoData.goName }}</DescriptionsItem>
                    <DescriptionsItem label="磁盘占用">{{ goInfoData.goSize }}</DescriptionsItem>
                    <DescriptionsItem label="项目地址">{{ goInfoData.pwd }}</DescriptionsItem>
                    <DescriptionsItem label="架构版本">{{ hostData.kernelArch }}</DescriptionsItem>
                    <DescriptionsItem label="GO 版本">{{ goInfoData.goVersion }}</DescriptionsItem>
                    <DescriptionsItem label="协程数量">{{ goInfoData.goroutine }}</DescriptionsItem>
                    <DescriptionsItem label="服务器IP">{{ hostData.intranet_ip }} (内) &nbsp;&nbsp;&nbsp; {{
                        hostData.public_ip }} (公) </DescriptionsItem>
                </Descriptions>
            </Card>
        </div>
    </div>
</template>
<style scoped>
.cell-card {
    box-sizing: border-box;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: normal;
    word-break: break-all;
    line-height: 36px;
}
</style>