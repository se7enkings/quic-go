// +build !tracer

package quictrace

import (
	"github.com/lucas-clemente/quic-go/internal/protocol"
)

// A tracer is used to trace a QUIC connection
type tracer struct {
}

// NewTracer creates a new Tracer
func NewTracer() Tracer {
	return &tracer{}
}

// Trace traces an event
func (t *tracer) Trace(connID protocol.ConnectionID, ev Event) {
}

func (t *tracer) GetAllTraces() map[string][]byte {
	return nil
}

// Emit emits the serialized protobuf that will be consumed by quic-trace
func (t *tracer) Emit(connID protocol.ConnectionID) ([]byte, error) {
	return nil, nil
}
