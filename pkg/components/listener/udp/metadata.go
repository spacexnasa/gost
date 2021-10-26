package udp

import "time"

const (
	defaultTTL            = 60 * time.Second
	defaultReadBufferSize = 1024
	defaultReadQueueSize  = 128
	defaultConnQueueSize  = 128
)

const (
	addr = "addr"
)

type metadata struct {
	addr string
	ttl  time.Duration

	readBufferSize int
	readQueueSize  int
	connQueueSize  int
}