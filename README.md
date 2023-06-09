# Go Web Admin 
## 表
* user     
    * username  password   
* role      
    * name 
* menu     
    * name   path    method


## API  注释

> http://127.0.0.1:8000/swagger/index.html

![demo](docs/demo.jpg)



## 目录结构
* conf：配置文件
* docs： swagger
* logs： 日誌
* middleware：
    * inject 初始化RBAC
    * jwt 
    * permission 權限驗證
* models：資料庫模型ORM
* pkg：第三方包
* routers： 路由邏輯處理
* service： 業務邏輯處理
* sql : 資料庫 sql script flyway


## 权限验证说明

```
项目启动时,会自动user  role  menu 进行自动关联!  如有更改,会删除对应的权限,重新加载!

用户  关联  角色  
角色  关联  菜单  
```


```
权限关系为:
角色(role.name,  menu.path,  menu.method)  
用户(user.username,   role.name)

例如:
运维部      /api/v1/users       GET
hequan     运维部

当hequan  GET  /api/v1/users 地址的时候，会去检查权限，因为他属于运维部 ，同时 运维部 有对应权限，所以本次请求会通过。

用户 admin 有所有的权限,不进行权限匹配

登录接口 /auth    /api/v1/userInfo 不进行验证
```

## 请求

> 请求和接收 都是 传递 json 格式 数据
```
例如:
访问 /auth    获取token
{
	"username": "admin",
	"password": "123456"
}
curl -X POST -H 'Content-type: application/json' --data '{"username":"admin","password":"123456"}' http://localhost:8000/auth


访问      /api/v1/menus?page=2    页面是page
请求头设置  Authorization: Token xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

curl -X POST -H 'Content-type: application/json' -H 'Authorization: Token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiaWQiOjEsInBhc3N3b3JkIjoiZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2UiLCJleHAiOjE2ODU2NzYyNjksImlzcyI6Imh0dHBzOi8vZ28tYWRtaW4vIn0.4NQ_G4OUuUtvsGMv9boYpRGyUKO_WQctZkGKfBJpb-s' http://localhost:8000/api/v1/menus?page=1


访问   /api/v1/userInfo  获取用户信息
前端所需的权限 放在password字段里面，已经去重了。

"data": {
        "lists": {
            "id": 2,
            "created_on": 1550642309,
            "modified_on": 1550642309,
            "deleted_on": 0,
            "username": "hequan",
            "password": ",system,menu,create_menu,update_menu,delete_menu,user,create_user,update_user,delete_user,role,create_role,update_role,delete_role",
            "role": [
                {
                    "id": 2,
                    "created_on": 0,
                    "modified_on": 0,
                    "deleted_on": 0,
                    "name": "运维部",
                    "menu": null
                }
            ]
        }
    },
```

## 部署

### 支持

- 部署 Mysql

### 库


创建一个库 go,然后导入sql,docs/sql/go.sql  创建表！

### 配置文件

You should modify `conf/app.ini`

```
[database]
Type = mysql
User = root
Password =
Host = 127.0.0.1:3306
Name = go
TablePrefix = go_
```

### 安装部署
```

yum install go -y 


export GOPROXY=https://goproxy.io
go get go-admin
cd $GOPATH/src/go-admin
go build main.go
go run  main.go 
```


### 热编译(开发时使用)
```bash

go get github.com/silenceper/gowatch

gowatch   
```

## windows 开发 需要gcc
> 执行会报错 "gcc" executable file not found in %PATH%
> 可以参考这个 安装gcc https://blog.csdn.net/xia_2017/article/details/105545789

## 运行


```
更新角色权限关系 [[hequan 运维部]]
角色权限关系 [[hequan 运维部]]
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /auth                     --> go-admin/routers/api.Auth (4 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.WrapHandler.func1 (4 handlers)
[GIN-debug] GET    /api/v1/menus             --> go-admin/routers/api/v1.GetMenus (6 handlers)
[GIN-debug] GET    /api/v1/menus/:id         --> go-admin/routers/api/v1.GetMenu (6 handlers)
[GIN-debug] POST   /api/v1/menus             --> go-admin/routers/api/v1.AddMenu (6 handlers)
[GIN-debug] PUT    /api/v1/menus/:id         --> go-admin/routers/api/v1.EditMenu (6 handlers)
[GIN-debug] DELETE /api/v1/menus/:id         --> go-admin/routers/api/v1.DeleteMenu (6 handlers)
[GIN-debug] GET    /api/v1/roles             --> go-admin/routers/api/v1.GetRoles (6 handlers)
[GIN-debug] GET    /api/v1/roles/:id         --> go-admin/routers/api/v1.GetRole (6 handlers)
[GIN-debug] POST   /api/v1/roles             --> go-admin/routers/api/v1.AddRole (6 handlers)
[GIN-debug] PUT    /api/v1/roles/:id         --> go-admin/routers/api/v1.EditRole (6 handlers)
[GIN-debug] DELETE /api/v1/roles/:id         --> go-admin/routers/api/v1.DeleteRole (6 handlers)
[GIN-debug] GET    /api/v1/users             --> go-admin/routers/api.GetUsers (6 handlers)
[GIN-debug] GET    /api/v1/users/:id         --> go-admin/routers/api.GetUser (6 handlers)
[GIN-debug] POST   /api/v1/users             --> go-admin/routers/api.AddUser (6 handlers)
[GIN-debug] PUT    /api/v1/users/:id         --> go-admin/routers/api.EditUser (6 handlers)
[GIN-debug] DELETE /api/v1/users/:id         --> go-admin/routers/api.DeleteUser (6 handlers)
2019/05/31 14:48:43 [info] start http server listening :8000

```




## Features
```
- RESTful API
- Gorm
- logging
- Jwt-go
- Swagger
- Gin
- Graceful restart or stop (fvbock/endless)
- App configurable
```

## 开发
* 何全

## 特别感谢
```
本项目主要参考了:
https://github.com/EDDYCJY/go-gin-example  包含更多的例子，上传文件图片等。本项目进行了增改。
https://github.com/LyricTian/gin-admin     主要为 gin+ casbin例子。
```
## 其他
```shell
## 更新API文档
swag init

## 
cd /opt/go-admin
nohup  go run  main.go   >>  /tmp/go-http.log   2>&1  & 
```
