package mocks

//go:generate sh -c "mockgen -package mockquic -destination quic/stream.go github.com/se7enkings/quic-go Stream && goimports -w quic/stream.go"
//go:generate sh -c "mockgen -package mockquic -destination quic/early_session.go github.com/se7enkings/quic-go EarlySession && goimports -w quic/early_session.go"
//go:generate sh -c "mockgen -package mockquic -destination quic/early_listener.go github.com/se7enkings/quic-go EarlyListener && goimports -w quic/early_listener.go"
//go:generate sh -c "mockgen -package mocks -destination short_header_sealer.go github.com/se7enkings/quic-go/internal/handshake ShortHeaderSealer && goimports -w short_header_sealer.go"
//go:generate sh -c "mockgen -package mocks -destination short_header_opener.go github.com/se7enkings/quic-go/internal/handshake ShortHeaderOpener && goimports -w short_header_opener.go"
//go:generate sh -c "mockgen -package mocks -destination long_header_opener.go github.com/se7enkings/quic-go/internal/handshake LongHeaderOpener && goimports -w long_header_opener.go"
//go:generate sh -c "mockgen -package mocks -destination crypto_setup.go github.com/se7enkings/quic-go/internal/handshake CryptoSetup && goimports -w crypto_setup.go"
//go:generate sh -c "mockgen -package mocks -destination stream_flow_controller.go github.com/se7enkings/quic-go/internal/flowcontrol StreamFlowController && goimports -w stream_flow_controller.go"
//go:generate sh -c "mockgen -package mocks -destination congestion.go github.com/se7enkings/quic-go/internal/congestion SendAlgorithmWithDebugInfos && goimports -w congestion.go"
//go:generate sh -c "mockgen -package mocks -destination connection_flow_controller.go github.com/se7enkings/quic-go/internal/flowcontrol ConnectionFlowController && goimports -w connection_flow_controller.go"
//go:generate sh -c "mockgen -package mockackhandler -destination ackhandler/sent_packet_handler.go github.com/se7enkings/quic-go/internal/ackhandler SentPacketHandler && goimports -w ackhandler/sent_packet_handler.go"
//go:generate sh -c "mockgen -package mockackhandler -destination ackhandler/received_packet_handler.go github.com/se7enkings/quic-go/internal/ackhandler ReceivedPacketHandler && goimports -w ackhandler/received_packet_handler.go"
