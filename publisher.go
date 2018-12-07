package mba

// Publisher defines the interface for publishers.
type Publisher interface {
	Send(msg interface{}) error
	Close()
}
