// Package otocadapter provides an interceptor to inject a
// OpenCensus span, derived from an OpenTracing span, to send
// it over gRPC to a service, which uses OpenCensus for tracing.
//
// This implementation assumes, Jaeger tracing implementation is used.
package otocadapter
