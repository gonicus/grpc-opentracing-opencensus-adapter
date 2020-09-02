// Package otocadapter provides
package otocadapter

import (
	"context"

	"go.opencensus.io/trace/propagation"
	"google.golang.org/grpc/metadata"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	octrace "go.opencensus.io/trace"
	"google.golang.org/grpc"
)

const (
	OPENCENSUS_TRACE_KEY = "grpc-trace-bin"
)

// OpenCensusAdapterClientInterceptor injects a OpenTracing span in an OpenCensus format into the context.
func OpenCensusAdapterClientInterceptor(tracer opentracing.Tracer) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, resp interface{}, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) error {
		// Get the OpenTracing span from context
		if parent := opentracing.SpanFromContext(ctx); parent != nil {
			parentCtx := parent.Context()

			// Get the Jaeger SpanContext
			if sc, ok := parentCtx.(jaeger.SpanContext); ok {

				ocTraceID, err := OpenTracingTracIDToOpenCensusTraceID(sc.TraceID())
				if err == nil {
					// Create an OpenCensus SpanContext and assign the IDs
					_, ocSpan := octrace.StartSpan(ctx, method)
					ocCtx := ocSpan.SpanContext()
					ocCtx.TraceID = ocTraceID
					ocCtx.SpanID = OpenTracingSpanIDToOpenCensusSpanID(sc.SpanID())

					// Append the OC span to the metadata, as it ocgrpc does
					traceContextBinary := propagation.Binary(ocCtx)
					ctx = metadata.AppendToOutgoingContext(ctx, OPENCENSUS_TRACE_KEY, string(traceContextBinary))
				}
			}
		}

		return invoker(ctx, method, req, resp, cc, opts...)
	}
}
