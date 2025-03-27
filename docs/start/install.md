## 系统安装

- 环境要求
- 安装

### 环境要求

- node版本 >= v16.0.0
- golang版本 >= v1.23
- goframe版本 >=v2.7.0
- mysql版本 >=5.7

> 必须先看[环境搭建文档](environment.md)，如果安装遇到问题务必先查看[常见问题文档](issue.md)

### 安装

一、克隆项目

```
git clone https://github.com/xiujiecn/xiu-admin.git && cd xiu-admin
```

二、配置你的站点信息

1、服务端：
- 项目数据库文件 `storage/data/xiuadmin.sql` 创建数据库并导入
- 将配置文件 `manifest/config/config.yaml.bak` 复制后改为`manifest/config/config.yaml`
- 将`manifest/config/config.yaml`中的`database.default.link`数据库配置改为你自己的：
```yaml
# Database. 配置参考：https://goframe.org/pages/viewpage.action?pageId=1114245
database:
  logger:
    path: "logs/database"                       # 日志文件路径。默认为空，表示关闭，仅输出到终端
    <<: *defaultLogger
    stdout: true
  default:
    link: "mysql:xiuadmin:xiu123456.@tcp(127.0.0.1:3306)/xiuadmin?loc=Local&parseTime=true&charset=utf8mb4"
    debug: true
    Prefix: ""
```

- 将`hack/config.yaml`中的`gfcli.gen.dao[0].link`数据库配置改为你自己的：
```yaml
gfcli:
  gen:
    dao:
      - link: "mysql:xiuadmin:xiu123456.@tcp(127.0.0.1:3306)/xiuadmin?loc=Local&parseTime=true&charset=utf8mb4"
        group: "default"                                                # 分组 使用xiuadmin代码生成功能时必须填
        #        tables:          ""                                    # 指定当前数据库中需要执行代码生成的数据表。如果为空，表示数据库的所有表都会生成。
        tablesEx:        "sys_addons_install"                        # 指定当前数据库中需要排除代码生成的数据表。
        removePrefix: ""
        descriptionTag: true
        noModelComment: true
        jsonCase: "CamelLower"
        gJsonSupport: true
        clear: false
```

2、web前端：

::: info 环境要求

在启动项目前，你需要确保你的环境满足以下要求：

- [Node.js](https://nodejs.org/en) 20.15.0 及以上版本，推荐使用 [fnm](https://github.com/Schniz/fnm) 、 [nvm](https://github.com/nvm-sh/nvm) 或者直接使用[pnpm](https://pnpm.io/cli/env) 进行版本管理。
- [Git](https://git-scm.com/) 任意版本。

验证你的环境是否满足以上要求，你可以通过以下命令查看版本：

```bash
# 出现相应 node LTS版本即可
node -v
# 出现相应 git 版本即可
git -v
```

三、 启动服务

1、服务端：
```shell script
      cd server
      
      # 设置国内代理，如果已经设置好了代理可以跳过
      go env -w GOPROXY=https://goproxy.io,direct
      
      # 更新包
      go mod tidy  
      
      # 查看命令行方法
      go run main.go help
      
      # 启动所有服务
      go run main.go  # 热编译启动： gf run main.go
```

2、web前端：
```shell script
    cd vue
    # 首先确定你以安装node16.0以上版本并安装了包[npm、pnpm]，否则可能会出现一些未知报错
    
    # 安装依赖
    pnpm install 
    
    # 启动web项目
    pnpm run dev 
    
    # 如果顺利，至此到浏览器打开：http://你的IP:5666/
    # 登录账号：admin, 密码：123456
```





