# auth-common

GoFrame v2 鉴权 / 系统管理插件。其它 GoFrame 项目通过 `go get` 引入后，一行 `install.Install` 即可挂载登录、用户、角色、菜单、部门、字典 API，并自动检查/创建数据库表与初始数据。

## 接入

```bash
go get github.com/liuzhengtao/auth-common-backend
```

宿主需已配置 MySQL（GoFrame `database` 下对应 group，默认 `database.default`），并 blank-import MySQL 驱动：

```go
package main

import (
    _ "github.com/gogf/gf/contrib/drivers/mysql/v2"
    "github.com/gogf/gf/v2/frame/g"
    "github.com/gogf/gf/v2/net/ghttp"

    "github.com/liuzhengtao/auth-common-backend/install"
    authmw "github.com/liuzhengtao/auth-common-backend/middleware"
)

func main() {
    s := g.Server()
    if err := install.Install(s); err != nil {
        panic(err)
    }

    // 保护宿主自有业务路由
    s.Group("/api/v1/biz", func(group *ghttp.RouterGroup) {
        group.Middleware(authmw.Auth)
        // group.Bind(...)
    })

    s.Run()
}
```

### Install 行为

1. 连接宿主库 `g.DB(dbGroup)`（默认 group 为 `default`）
2. 检查 8 张表：`sys_user` / `sys_role` / `sys_menu` / `sys_dept` / `sys_dict` / `sys_dict_type` / `sys_user_role` / `sys_role_menu`
3. 缺失则自动 `CREATE TABLE`
4. 若无 `admin` 用户，写入默认部门、角色、用户、菜单、字典
5. 若开启分布式：校验 Redis 并切换验证码到 Redis Store
6. 挂载路由（默认前缀 `/api/v1`）

可选参数：

```go
install.Install(s,
    install.WithPrefix("/api/v1"),
    install.WithDBGroup("auth"),        // 使用 database.auth 配置组
    install.WithSkipSchemaInit(true),   // 宿主自行管库时跳过建表
    install.WithLogger(myLogger),       // 可选；不传则用 glog.New().Line(true)，不碰 g.Log()
    install.WithDistributed(true),      // 分布式：Redis 验证码 + Token 黑名单
    install.WithRedisGroup("default"),  // 可选；默认 default，对应 redis.default
)
```

仅初始化库表、不挂路由：`install.EnsureSchema(ctx)`。

## 配置（可选）

宿主 `config.yaml`：

```yaml
database:
  default:
    link: "mysql:user:pass@tcp(127.0.0.1:3306)/biz"
  auth:
    link: "mysql:user:pass@tcp(127.0.0.1:3306)/auth"

redis:
  default:
    address: "127.0.0.1:6379"
    db: 1

authCommon:
  jwtSecret: "change-me"      # JWT HMAC 密钥（多节点需一致）
  jwtAesKey: "change-me"      # Identity 载荷 AES 密钥（多节点需一致）
  encryptSalt: "change-me"
  routePrefix: "/api/v1"
  dbGroup: "auth"             # 可选，默认 default；也可由 install.WithDBGroup 覆盖
  distributed: false          # true 时启用 Redis 验证码共享 + Token 黑名单
  redisGroup: "default"       # 可选，默认 default；也可由 install.WithRedisGroup 覆盖
```

未配置时使用模块内默认值。种子用户密码按当前 `encryptSalt` **运行时计算**。

JwtSecret / JwtAesKey **仅通过** `authCommon` 配置（或默认值），不提供 Install Option 传入。

## 分布式模式（Redis）

默认 `distributed: false`：验证码用进程内存，登出不做服务端黑名单（与原先行为一致）。

开启后（`authCommon.distributed: true` 或 `install.WithDistributed(true)`）：

1. **验证码**：写入宿主 Redis，多节点可共享校验
2. **Token 黑名单**：登出时将 token 写入 Redis（TTL=剩余有效期）；`middleware.Auth` 在 JWT 校验前拦截已吊销 token

Redis 连接与 `dbGroup` 同模式：默认用宿主 `redis.default`（`g.Redis()`）；也可指定 `redisGroup` 指向其它组。

宿主还需 blank-import Redis 驱动（与 MySQL 驱动同级）：

```go
import (
    _ "github.com/gogf/gf/contrib/drivers/mysql/v2"
    _ "github.com/gogf/gf/contrib/nosql/redis/v2"
)
```

## 默认账号

| 用户名 | 密码 | 说明 |
|--------|------|------|
| admin  | a123456 | 日常登录（ADMIN 角色） |
| root   | a123456 | 超级管理员（列表默认隐藏） |

## 主要 API（前缀默认 `/api/v1`）

| 路径 | 说明 |
|------|------|
| `GET /auth/captcha` | 验证码 |
| `POST /auth/login` | 登录 |
| `DELETE /auth/logout` | 登出（需鉴权；分布式下写入黑名单） |
| `/users` `/roles` `/menus` `/dept` `/dict` | 系统管理（需鉴权） |

鉴权：`Authorization: Bearer <token>`（也支持 query `token` / cookie `jwt`）。

## 表结构文档

见 [manifest/sql/auth_common.sql](manifest/sql/auth_common.sql)（与 Install 内嵌 SQL 一致）。

## 本地运行本仓库

```bash
# 配置好 database 后
go run .
```
