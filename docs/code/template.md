---
outline: deep
---

# 生成模板

## 业务模板
根据api接口文件一键生成业务模板：api、controller、logic、service
## 模板路径配置
- 模板路径配置：server/hack/config.yaml
```
# 生成代码
xgen:
  application:
    # CRUD和关系树列表模板
    crud:
      templates:
        # 默认的主包模板
        - group: "default"                                              # 分组名称
          isAddon: false                                                # 是否为插件模板 false｜true
          masterPackage: "gen"                                          # 主包名称，需和controllerPath、logicPath、inputPath保持关联
          templatePath: "./resource/generate/default/curd"              # 模板路径
          apiPath: "./api/gen"                                          # goApi生成路径
          controllerPath: "./internal/controller/gen"                   # 控制器生成路径
          logicPath: "./internal/logic/gen"                             # 主要业务生成路径
          inputPath: "./internal/model/genin"                           # 表单过滤器生成路径
          routerPath: "./internal/router/genrouter"                     # 生成路由表路径
          sqlPath: "./manifest/generate"                                # 生成sql语句路径
          webApiPath: "../vue/apps/web-antd/src/api/gen"                # webApi生成路径
          webViewsPath: "../vue/apps/web-antd/src/views/gen"            # web页面生成路径
```
## 模板目录
- 模板目录：server/resource/generate/default/curd
```
server/resource/generate/default/curd
├── api.go.template                   # 后端api接口文件
├── controller.go.template            # 后端控制器文件
├── input.go.template                 # 后端Model定义文件
├── logic.go.template                 # 后端业务逻辑文件
├── router.go.template                # 后端路由文件
├── source.sql.template               # 后端sql语句文件
├── web.api.ts.template               # 前端api接口文件
├── web.edit.vue.template             # 前端编辑抽屉文件
├── web.index.vue.template            # 前端列表页面文件
├── web.model.ts.template             # 前端表单模型文件
└── web.view.vue.template             # 前端详情抽屉文件
```