package control

import (
	"github.com/openzipkin/zipkin-go"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
)

const (
	enpoitUrl = "http://192.168.60.38:9411/api/v2/spans"
)

func gettracer(servername string, ip string) *zipkin.Tracer {
	reporter := httpreporter.NewReporter(enpoitUrl)
	endpoint, _ := zipkin.NewEndpoint(servername, ip)
	sampler := zipkin.NewModuloSampler(1)

	trace, _ := zipkin.NewTracer(
		reporter,
		zipkin.WithLocalEndpoint(endpoint),
		zipkin.WithSampler(sampler),
	)
	return trace
}
