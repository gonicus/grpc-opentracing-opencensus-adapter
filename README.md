# gRPC OpenTracing to OpenCensus Adapter

This adapter injects an OpenCensus span in the metadata of a gRPC call, which is extracted from an OpenTracing span.

If your project uses OpenTracing and you call a service via gRPC, which uses OpenCensus, this adapter makes the span available in the remote servie.

```go
d, err := grpc.Dial("localhost:9080", grpc.WithInsecure(),
    grpc.WithUnaryInterceptor(otocadapter.OpenCensusAdapterClientInterceptor(tracer)),
)
if err != nil {
    logrus.Fatalln("Can't connect to DGraph:", err)
}
defer func() {
    d.Close()
    logrus.Info("Closed Dgraph connection...")
}()
dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(d))
```

## Install

```
go get -u github.com/gonicus/grpc-opentracing-opencensus-adapter
```