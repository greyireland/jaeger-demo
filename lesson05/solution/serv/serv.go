package main

import (
	"fmt"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"
	"github.com/greyireland/jaeger-demo/lib/tracing"
	"os"
	"os/signal"
)

func main() {
	go formater()
	go publisher()
	//1.声明chan
	c := make(chan os.Signal)
	//2.注册chan
	signal.Notify(c, os.Interrupt, os.Kill)
	//3.阻塞，登台退出
	<-c
}
func formater() {
	tracer, closer := tracing.Init("formatter")
	defer closer.Close()

	http.HandleFunc("/api/trace/format", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		span := tracer.StartSpan("format", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		greeting := span.BaggageItem("greeting")
		if greeting == "" {
			greeting = "Hello"
		}

		helloTo := r.FormValue("helloTo")
		helloStr := fmt.Sprintf("%s, %s!", greeting, helloTo)
		span.LogFields(
			otlog.String("event", "string-format"),
			otlog.String("value", helloStr),
		)
		w.Write([]byte(helloStr))
	})

	http.ListenAndServe("0.0.0.0:8081", nil)
}
func publisher() {
	tracer, closer := tracing.Init("publisher")
	defer closer.Close()

	http.HandleFunc("/api/trace/publish", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		span := tracer.StartSpan("publisher", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		helloStr := r.FormValue("helloStr")
		w.Write([]byte("hello"))
		println(helloStr)
	})

	http.ListenAndServe("0.0.0.0:8082", nil)
}
