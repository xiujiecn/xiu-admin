# XiuAdmin 项目记忆

## 项目基本信息
- 基于 GoFrame2 + Vue3 + Vben Admin 5.x + Ant Design Vue
- 后端目录: server/ (Go)，前端目录: vue/apps/web-antd/
- 前端采用 pnpm monorepo 架构
- 默认账号: admin / 123456

## 核心架构模式
- 后端: Service自注册(init+Register) + 延迟初始化(inithttp.RegisterHttpInitFunc)
- 前端: 三层路由(core+dynamic+后端返回) + v-access权限指令 + requestClient三层拦截器
- DAO/DO/Entity 由 gf gen 自动生成，禁止手动编辑

## 权限体系
- RBAC + 6维度权限码(cpr/cpm/cpd/cpp/cpu/cpc)
- 数据范围过滤7级(1=全部,5=仅本人)
- 多租户自动隔离(tenant_id字段自动注入WHERE)
- JWT+Redis双重校验(支持强制下线)
