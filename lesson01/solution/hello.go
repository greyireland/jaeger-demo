package main

import (
	"fmt"
	"os"

	"github.com/opentracing/opentracing-go/log"
	"github.com/yurishkuro/opentracing-tutorial/go/lib/tracing"
)

func main() {
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}

	//1.创建一个tracer
	tracer, closer := tracing.Init("hello-world")
	//1.1 关闭tracer
	defer closer.Close()

	helloTo := os.Args[1]

	//2.开启一个span
	span := tracer.StartSpan("say-hello")
	//3.记录tag
	span.SetTag("client", helloTo)

	//4.记录log 事件
	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)

	println(helloStr)
	span.LogKV("event", "exec println")
	//5.结束span
	span.Finish()
}
