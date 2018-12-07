package mba

// Subscriber defines the interface for subscribers.
type Subscriber interface {
	Receive() (ok bool, err error)
}
