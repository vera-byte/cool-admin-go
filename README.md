# GoFrame Grpc 使用例子

```shell
> tree
.
├── README.md
├── client
│   ├── go.mod
│   └── main.go # 简单调用测试
├── go.work
├── go.work.sum
└── user_server # 微服务项目
    ├── api
    │   └── user
    │       └── v1
    │           ├── user.pb.go # grpc生成
    │           └── user_grpc.pb.go # grpc生成
    ├── go.mod
    ├── go.sum
    ├── internal
    │   └── controller
    │       └── user
    │           └── user.go # 由gf gen pb 生成 写业务逻辑
    ├── main.go # user_server启动入口
    └── manifest
        ├── config
        │   └── config.yaml
        └── protobuf
            └── user
                └── v1
                    └── user.proto
```
