package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	zipkintracer "github.com/openzipkin/zipkin-go-opentracing"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//http://127.0.0.1:16686
	c := Config{
		Addr:    "http://127.0.0.1:9411/api/v1/spans",
		Rate:    1,
		Timeout: 3,
	}
	err := Init(c, "mush", "", false)
	fmt.Println(err)
	ctx := context.Background()

	span := opentracing.StartSpan("ctx")
	cc, _ := json.Marshal(span.Context())
	fmt.Println("span1", string(cc))
	time.Sleep(time.Second)

	span2, zctx := opentracing.StartSpanFromContext(opentracing.ContextWithSpan(ctx, span), "t2")
	cc, _ = json.Marshal(span2.Context())
	fmt.Println("span2", string(cc))
	time.Sleep(300 * time.Millisecond)

	span3, _ := opentracing.StartSpanFromContext(zctx, "t3")
	cc, _ = json.Marshal(span3.Context())
	fmt.Println("span3", string(cc))
	time.Sleep(300 * time.Millisecond)
	span3.Finish()

	span4, _ := opentracing.StartSpanFromContext(zctx, "t4")
	cc, _ = json.Marshal(span4.Context())
	fmt.Println("span4", string(cc))
	time.Sleep(300 * time.Millisecond)
	span4.Finish()

	time.Sleep(500 * time.Millisecond)
	span2.Finish()

	span.Finish()
	time.Sleep(3 * time.Second)
}

// Config struct for zipkin client.
type Config struct {
	Addr    string  `json:"addr"`
	Rate    float64 `json:"rate"`
	Timeout int     `json:"timeout"`
}

// Init initialize global tracer.
func Init(c Config, serviceName, serviceAddr string, debug bool) (err error) {
	if c.Addr == "" {
		return
	}

	var collector zipkintracer.Collector
	// 兼容模式, http:// 或 https:// 开头的使用 HTTP Collector
	if strings.HasPrefix(c.Addr, "http://") || strings.HasPrefix(c.Addr, "https://") {
		collector, err = zipkintracer.NewHTTPCollector(c.Addr, zipkintracer.HTTPTimeout(time.Duration(c.Timeout)))
	} else {
		collector, err = zipkintracer.NewScribeCollector(c.Addr, time.Duration(c.Timeout))
	}
	if err != nil {
		return
	}
	tracer, err := zipkintracer.NewTracer(
		zipkintracer.NewRecorder(collector, debug, serviceAddr, serviceName),
		zipkintracer.WithSampler(zipkintracer.NewBoundarySampler(c.Rate, rand.Int63())),
	)
	if err != nil {
		return
	}
	opentracing.SetGlobalTracer(tracer)
	return
}
