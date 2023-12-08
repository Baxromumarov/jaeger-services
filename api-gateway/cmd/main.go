package main

import (
	"fmt"
	"jaeger-services/api-gateway/api"
	"jaeger-services/api-gateway/api/handlers"
	"jaeger-services/api-gateway/config"
	"jaeger-services/api-gateway/grpc/client"
	"jaeger-services/api-gateway/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/uber/jaeger-client-go"
	jaeger_config "github.com/uber/jaeger-client-go/config"
)

func main() {
	fmt.Println("Ishladi")
	cfg := config.Load()

	jaegerCfg := &jaeger_config.Configuration{
		ServiceName: cfg.ServiceName,

		// "const" sampler is a binary sampling strategy: 0=never sample, 1=always sample.
		Sampler: &jaeger_config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},

		// Log the emitted spans to stdout.
		Reporter: &jaeger_config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: cfg.JaegerHostPort,
		},
	}

	tracer, closer, err := jaegerCfg.NewTracer(jaeger_config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	loggerLevel := logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

	svcs, err := client.NewGrpcClients(cfg)
	if err != nil {
		fmt.Println("here>>>>>>>>>>>>>>>",err)
		log.Panic("client.NewGrpcClients", logger.Error(err))
	}

	h := handlers.NewHandler(cfg, log, svcs)

	r := api.SetUpRouter(h, cfg, tracer)

	log.Info("HTTP: Server being started...", logger.String("port", cfg.HTTPPort))

	r.Run(cfg.HTTPPort)
}
