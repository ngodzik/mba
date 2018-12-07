package mba

// Driver defines the abstraction interface used by the message broker.
type Driver interface {
	NewPublisher(topic string) (Publisher, error)
	NewSubscriber() (Subscriber, error)
	SetTopic(s Subscriber, topic string, do func(msg interface{}) error) error
	Start()
}
