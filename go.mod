module github.com/gonicus/grpc-opentracing-opencensus-adapter

go 1.15

require (
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	go.opencensus.io v0.22.4
	go.uber.org/atomic v1.6.0 // indirect
	google.golang.org/grpc v1.31.1
)
