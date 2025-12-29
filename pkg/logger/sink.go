package logger

// Sink defines where logs go (console, file, network, etc.)
type Sink interface {
	Write(Event) error
	Close() error
}
