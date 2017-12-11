package tracing

import (
	"log"

	opentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"github.com/rodrigodealer/whiplash/util"
)

func StartTracing(server string, service string) {
	collector, err := zipkin.NewHTTPCollector(util.EnvOrElse("ZIPKIN_URL", "http://127.0.0.1:9411/api/v1/spans"))
	if err != nil {
		log.Printf("unable to create Zipkin HTTP collector: %+v", err)
	}

	recorder := zipkin.NewRecorder(collector, false, server, service)

	tracer, err := zipkin.NewTracer(
		recorder, zipkin.ClientServerSameSpan(true),
	)

	if err != nil {
		log.Printf("unable to create Zipkin tracer: %+v", err)
	}

	opentracing.InitGlobalTracer(tracer)
}
