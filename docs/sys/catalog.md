---
outline: deep
---

## 目录结构

### 目录
```
.
├── LICENSE
├── README.md
├── docs
├── server                     # 服务端
│   ├── Makefile
│   ├── README.MD
│   ├── api                    # 接口
│   │   ├── common
│   │   │   └── v1
│   │   ├── gen
│   │   │   └── testdemo
│   │   ├── gen_codes
│   │   │   └── v1
│   │   ├── monitor
│   │   │   └── v1
│   │   └── system
│   │       └── v1
│   ├── config                   # 配置
│   │   ├── config.yaml
│   │   └── config_default.yaml
│   ├── go.mod                   # 依赖
│   ├── go.sum
│   ├── hack                     # 工具
│   │   ├── config.yaml
│   │   ├── hack-cli.mk
│   │   └── hack.mk
│   ├── internal                 # 内部
│   │   ├── cmd                    # 命令
│   │   ├── consts                 # 常量
│   │   ├── controller             # 控制器
│   │   ├── dao                    # 数据库
│   │   ├── library                # 库
│   │   │   ├── addons             # 插件
│   │   │   ├── bcache             # Redis缓存
│   │   │   ├── cache              # 缓存
│   │   │   ├── contexts           # 上下文
│   │   │   ├── event              # 事件
│   │   │   ├── mcache             # 内存缓存
│   │   │   ├── websocket          # WebSocket
│   │   │   ├── worker             # 工作
│   │   │   ├── xgen               # 代码生成
│   │   │   └── xgorm              # 数据库
│   │   ├── logic                  # 逻辑
│   │   │   ├── gen                # 逻辑-代码生成
│   │   │   ├── middleware         # 逻辑-中间件
│   │   │   ├── monitor            # 逻辑-监控
│   │   │   └── system             # 逻辑-系统
│   │   ├── model                  # 模型
│   │   ├── packed                 # 打包
│   │   │   ├── packed.go          # 打包
│   │   │   └── response           # 响应
│   │   ├── queues                 # 队列
│   │   ├── router                 # 路由
│   │   ├── service                # 服务
│   │   └── tasks                  # 任务
│   ├── main.go                    # 主函数
│   ├── manifest                   # 清单
│   ├── resource                   # 资源
│   └── utility                    # 工具
└── vue                          # 前端
    ├── README.md
    ├── README.zh-CN.md
    ├── apps
    │   └── web-antd
    │       ├── public               # 静态资源
    │       ├── src                  # 源码
    │       │    ├── adapter            # 适配器
    │       │    ├── api                # 接口
    │       │    ├── app.vue            # 应用
    │       │    ├── bootstrap.ts       # 引导
    │       │    ├── components         # 组件
    │       │    ├── layouts            # 布局
    │       │    ├── locales            # 语言
    │       │    ├── main.ts            # 主函数
    │       │    ├── preferences.ts     # vben配置
    │       │    ├── router             # 路由
    │       │    ├── store              # 存储
    │       │    ├── utils              # 工具
    │       │    └── views              # 视图

```