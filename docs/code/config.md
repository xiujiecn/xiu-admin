---
outline: deep
---

# 生成配置

## 模板配置

- 配置路径：server/manifest/config/config.yaml
```yaml
# 生成代码
xgen:
  allowedIPs: [ "127.0.0.1", "*" ]                                      # 白名单，*代表所有，只有允许的IP后台才能使用生成代码功能
  selectDbs: [ "default" ]                                              # 可选生成表的数据库配置名称，支持多库
  disableTables: [  ]          # 禁用的表，禁用以后将不会在选择表中看到
  delimiters: [ "@{", "}" ]                                             # 模板引擎变量分隔符号
  # 生成应用模型，所有生成模板允许自定义，可以参考default模板进行改造
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
          webApiPath: "../vue/apps/web-antd/src/api/gen"                              # webApi生成路径
          webViewsPath: "../vue/apps/web-antd/src/views/gen"                          # web页面生成路径

        # 默认的插件包模板，{$name}会自动替换成实际的插件名称
        - group: "addon"                                                # 分组名称
          isAddon: true                                                 # 是否为插件模板 false｜true
          masterPackage: "sys"                                          # 主包名称，需和controllerPath、logicPath、inputPath保持关联
          templatePath: "./resource/generate/default/curd"              # 模板路径
          apiPath: "./addons/{$name}/api/admin"                         # goApi生成路径
          controllerPath: "./addons/{$name}/controller/admin/sys"       # 控制器生成路径
          logicPath: "./addons/{$name}/logic/sys"                       # 主要业务生成路径
          inputPath: "./addons/{$name}/model/input/sysin"               # 表单过滤器生成路径
          routerPath: "./addons/{$name}/router/genrouter"               # 生成路由表路径
          sqlPath: "./storage/data/generate/addons"                     # 生成sql语句路径
          webApiPath: "../web/src/api/addons/{$name}"                   # webApi生成路径
          webViewsPath: "../web/src/views/addons/{$name}"               # web页面生成路径

```

## CLI
- 代码生成模块内部调用系统指令`gf gen dao`生成数据库实体
- 代码生成模块内部调用系统指令`gf gen service`生成业务逻辑
  

