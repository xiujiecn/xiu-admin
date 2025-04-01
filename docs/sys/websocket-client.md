---
outline: deep
---

# WebSocket客户端

- [WebSocket服务器](websocket.md)

## 前端使用WebSocket示例
路径 `vue/apps/web-antd/src/views/monitor/server/index.vue`

```ts
import { useWebSocket } from '@vueuse/core';
import { useAccessStore } from '@vben/stores';
const accessStore = useAccessStore();

const { status, data, send, close, open } = useWebSocket(state.server, {
    autoReconnect: true,  // 自动重连
    heartbeat: false,     // 关闭客户端心跳
});

// URL组成
function getWsServerUrl() {
    return 'ws://' + window.location.host + '/api/monitor/ws?access_token=' + encodeURIComponent(String(accessStore.accessToken));
}
// 是否链接到WebSocket
const getIsOpen = computed(() => status.value === 'OPEN');

// 发送订阅标签消息
function handlerSend() {
    send(JSON.stringify({
        e: 'joins',
        d: {
            names: ["ws:tag:monitor:server"]
        }
    }));
}
// 监听链接状态，链接成功发送订阅消息
watch(status, (newVal) => {
    console.log('vue/apps/web-antd/src/views/monitor/server/index.vue status', newVal);
    if (newVal === 'OPEN') {
        handlerSend();
    }
});
// 监听收到数据状态，解析数据，处理数据
watch(data, (newVal) => {
    if (newVal) {
        const res = JSON.parse(newVal);
        if (typeof res.d === 'string' && !res.d.startsWith('{')) {
            res.d = atob(res.d);
        }
        if (typeof res.d === 'string' && res.d.startsWith('{')) {
            res.d = JSON.parse(res.d);
        }
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

onMounted(() => {
    open();
});
```