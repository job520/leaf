# leaf
基于 golang 游戏服务器 - leaf 的一个  demo，消息格式使用 json  
[leaf 官网](https://github.com/name5566/leaf/blob/master/TUTORIAL_ZH.md)

### leaf 使用介绍
1. 编辑 msg/msg.go 定义消息格式

2. 编辑 gate/router.go 将消息格式路由到指定模块

3. 编辑 {模块名}/internal/handler.go 注册处理函数

### demo 用法
1. 开启服务器：
`
go run main.go
`
2. 使用 websocket 客户端：
`
用浏览器打开 index.html 文件
`
3. 使用 tcp 客户端：
`
go run client.go
`