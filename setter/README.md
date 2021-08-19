# setter 设置值

用于处理各种操蛋的代码，比如

```go
if "" != server.Options.Password {
serverOptions.SetPassword(server.Options.Password)
} else {
serverOptions.SetPassword(mqttConfig.Options.Password)
}
```

经过简单的setter后，可以写成如下格式

```go
setter.String(serverOptions.SetPassword, server.Options.Password, mqttConfig.Options.Password)
```
