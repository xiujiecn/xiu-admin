<script lang="ts" setup>
import { Card, CardGrid, Descriptions, DescriptionsItem } from 'ant-design-vue';
import type { EchartsUIType } from '@vben/plugins/echarts';

import { defineComponent, onActivated, onMounted, ref, watch } from 'vue';

import { EchartsUI, useEcharts } from '@vben/plugins/echarts';

const sysInfo = ref({
    cpuNum: 0,
    cpuCores: 0,
    cpuUsed: 20,
    cpuAvg5: 0,
    cpuAvg15: 0,
    memUsed: 20,
    memTotal: 100,
    available: 80,
    goUsed: 10,
    memUsage: 20,
    diskTotal: 100,
    diskUsed: 20,
    diskUsedPercent: 20,
});

const hostData = ref({
      "bootTime": "2022-11-24T11:12:13+08:00",
      "hostId": "8be74718-1a53-4208-be22-9c126d891ddd",
      "hostname": "iZ2zee04uvnkmhvglw9oghZ",
      "intranet_ip": "172.17.47.62",
      "kernelArch": "x86_64",
      "kernelVersion": "3.10.0-1127.19.1.el7.x86_64",
      "os": "linux",
      "platform": "centos",
      "platformFamily": "rhel",
      "platformVersion": "7.7.1908",
      "procs": 138,
      "public_ip": "101.200.198.249",
      "uptime": 6393278,
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
    xAxis: {
        type: 'category',
        boundaryGap: false,
        data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
    },
    yAxis: {
        type: 'value'
    },
    series: [
        {
            data: [820, 932, 901, 934, 1290, 1330, 1320],
            type: 'line',
            areaStyle: {}
        }
    ]
};
const memRunOptions: memRunEChartsOption = {
    xAxis: {
        type: 'category',
        boundaryGap: false,
        data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
    },  
    yAxis: {
        type: 'value'
    },
    series: [
        {
            data: [820, 932, 901, 934, 1290, 1330, 1320],   
            type: 'line',
            areaStyle: {}
        }
    ]
};
const diskRunOptions: diskRunEChartsOption = {
    xAxis: {
        type: 'category',
        boundaryGap: false,
        data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
    },
    yAxis: {
        type: 'value'
    },
    series: [
        {
            data: [820, 932, 901, 934, 1290, 1330, 1320],
            type: 'line',
            areaStyle: {}
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
onMounted(() => {
    setCpuEcharts();
    setMemEcharts();
    setDiskEcharts();
    setCpuRunEcharts();
    setMemRunEcharts();
    setDiskRunEcharts();
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
    <div class="p-5 flex flex-col gap-4">
        <div class="grid grid-cols-1 gap-4 md:grid-cols-1 lg:grid-cols-3">
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
        <div class="grid grid-cols-1 gap-4 md:grid-cols-1 lg:grid-cols-1">
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
                    <DescriptionsItem label="服务器IP">{{ hostData.intranet_ip }} (内) &nbsp;&nbsp;&nbsp; {{ hostData.public_ip }} (公) </DescriptionsItem> 
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