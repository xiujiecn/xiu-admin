---
outline: deep
---

# 事件

> 例如：用户修改、删除、新增等操作，会触发事件，事件会触发对应的回调函数。


## 实现接口
- 为了提供高度的扩展性，事件在设计上采用了注册回调思路。只需要实现接口，您就可以在任何地方注册和使用事件功能，从而实现更大的灵活性和可扩展性。

## 一个例子

例如，如果您需要用户更新清理缓存，内容大致如下：

- 文件路径：server/internal/library/mcache/mcache_system_user.go

```go 
package mcache

import (
	"context"
	"fmt"
	"time"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/library/event"
	"xiuadmin/internal/model"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	event.EventsInstance().Register(consts.EventKeyUserUpdate, func(ctx context.Context, eventKey string, args ...interface{}) {
		if len(args) == 0 {
			return
		}
		userId := args[0].(int64)
		RemoveUserInfo(ctx, userId)
	})
}

```

下面是将触发事件的方式，大概内容如下:

```go
package main

import (
	"fmt"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/queues"
)

func test()  {
	event.EventsInstance().Emit(ctx, consts.EventKeyUserUpdate, 0)
}

```

