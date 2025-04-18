package app

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	// "go.opentelemetry.io/otel/trace" // (not used yet)
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	helloCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hello_said_total",
		Help: "Total number of HelloSaid events.",
	})
	tracer = otel.Tracer("overhelloworld")
)

func InitMetrics() {
	prometheus.MustRegister(helloCounter)
}

func MetricsHandler() http.Handler {
	return promhttp.Handler()
}

func Trace(msg string) {
	_, span := tracer.Start(context.TODO(), msg)
	defer span.End()
	log.Info().Msg("[TRACE] " + msg)
}

func Metric(msg string) {
	helloCounter.Inc()
	log.Info().Msg("[METRIC] " + msg)
}

func Log(msg string) {
	log.Info().Msg(msg)
}
