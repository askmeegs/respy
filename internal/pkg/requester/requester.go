package requester

// a requester knows how to perform a single HTTP or gRPC request
type Requester interface {
	OneRequest() (string, error)
}
