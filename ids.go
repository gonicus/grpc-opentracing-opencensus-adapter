package otocadapter

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/uber/jaeger-client-go"
	"go.opencensus.io/trace"
)

const (
	OPENTRACING_TRACE_ID_LENGTH = 8
)

var (
	ErrOTTraceIDLen = errors.New("expected OpenTracing trace ID with 8 bytes")
)

// Convert an OpenTracing TraceID to an OpenCensus TraceID.
func OpenTracingTracIDToOpenCensusTraceID(otTraceID jaeger.TraceID) (trace.TraceID, error) {
	// Get the OpenTracing traceID and convert it to a OpenCensus traceID.
	hexTi := otTraceID.String()
	if len(hexTi)%2 != 0 {
		hexTi = "0" + hexTi
	}

	ti, err := hex.DecodeString(hexTi)
	if err != nil {
		return trace.TraceID{}, fmt.Errorf("error decoding HEX trace ID: %w", err)
	}
	
	if len(ti) != OPENTRACING_TRACE_ID_LENGTH {
		return trace.TraceID{}, ErrOTTraceIDLen
	}

	var ocTi trace.TraceID

	// Fill in the trace ID and leave the upper 8 bytes blank.
	for i, v := range ti {
		ocTi[i+8] = v
	}

	return ocTi, nil
}

// Convert an OpenTracing SpanID to an OpenCensus SpanID.
func OpenTracingSpanIDToOpenCensusSpanID(otSpanID jaeger.SpanID) trace.SpanID {
	var ocSpanID trace.SpanID
	buf := make([]byte, 8)

	binary.BigEndian.PutUint64(buf, uint64(otSpanID))

	copy(ocSpanID[:], buf)

	return ocSpanID
}
