# OpenTracing Tutorial - Go

## 安装

```
go get -u github.com/greyireland/jaeger-demo
```


## Lessons

* [Lesson 01 - Hello World](./lesson01)
  * 实例化一个Tracer
  * 创建一个简单的tracer
  * 注解tag和log
* [Lesson 02 - Context and Tracing Functions](./lesson02)
  * 追踪多个函数
  * 连接多个span到一个trace
  * 通过context扩散
* [Lesson 03 - Tracing RPC Requests](./lesson03)
  * 追踪多个服务
  * 通过 `Inject` and `Extract` 传递数据
  * 应用推荐tag
* [Lesson 04 - Baggage](./lesson04)
  * 多个服务之间 携带数据传递
* [Extra Credit](./extracredit)
  * Use existing open source instrumentation
  * 使用go-stdlib库可以对http做更详细跟踪【使用标准库httptrace实现】
