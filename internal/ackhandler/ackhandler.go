package ackhandler

import (
	"github.com/se7enkings/quic-go/internal/congestion"
	"github.com/se7enkings/quic-go/internal/protocol"
	"github.com/se7enkings/quic-go/internal/utils"
	"github.com/se7enkings/quic-go/qlog"
	"github.com/se7enkings/quic-go/quictrace"
)

func NewAckHandler(
	initialPacketNumber protocol.PacketNumber,
	rttStats *congestion.RTTStats,
	pers protocol.Perspective,
	traceCallback func(quictrace.Event),
	qlogger qlog.Tracer,
	logger utils.Logger,
	version protocol.VersionNumber,
) (SentPacketHandler, ReceivedPacketHandler) {
	sph := newSentPacketHandler(initialPacketNumber, rttStats, pers, traceCallback, qlogger, logger)
	return sph, newReceivedPacketHandler(sph, rttStats, logger, version)
}
