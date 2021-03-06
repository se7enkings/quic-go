package mocks

//go:generate sh -c "mockgen -package mockquic -destination quic/stream.go github.com/se7enkings/quic-go Stream && goimports -w quic/stream.go"
//go:generate sh -c "mockgen -package mockquic -destination quic/session.go github.com/se7enkings/quic-go Session && goimports -w quic/session.go"
//go:generate sh -c "../mockgen_internal.sh mocks sealer.go github.com/se7enkings/quic-go/internal/handshake Sealer"
//go:generate sh -c "../mockgen_internal.sh mocks opener.go github.com/se7enkings/quic-go/internal/handshake Opener"
//go:generate sh -c "../mockgen_internal.sh mocks crypto_setup.go github.com/se7enkings/quic-go/internal/handshake CryptoSetup"
//go:generate sh -c "../mockgen_internal.sh mocks stream_flow_controller.go github.com/se7enkings/quic-go/internal/flowcontrol StreamFlowController"
//go:generate sh -c "../mockgen_internal.sh mockackhandler ackhandler/sent_packet_handler.go github.com/se7enkings/quic-go/internal/ackhandler SentPacketHandler"
//go:generate sh -c "../mockgen_internal.sh mockackhandler ackhandler/received_packet_handler.go github.com/se7enkings/quic-go/internal/ackhandler ReceivedPacketHandler"
//go:generate sh -c "../mockgen_internal.sh mocks congestion.go github.com/se7enkings/quic-go/internal/congestion SendAlgorithmWithDebugInfos"
//go:generate sh -c "../mockgen_internal.sh mocks connection_flow_controller.go github.com/se7enkings/quic-go/internal/flowcontrol ConnectionFlowController"
