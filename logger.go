package integrate

// This interface is for logging of the interal server's functions. Note this
// is different from a Storage which is for storing history and state of the
// actual messages going through the system.
type Logger interface {
	Debug(format string, args ...interface{})
	Notice(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})
	Critical(format string, args ...interface{})
}

// This is a logger that does absolutely nothing. It should be used for testing.
type NoOpLogger struct{}

func (n *NoOpLogger) Debug(format string, args ...interface{})    {}
func (n *NoOpLogger) Notice(format string, args ...interface{})   {}
func (n *NoOpLogger) Warning(format string, args ...interface{})  {}
func (n *NoOpLogger) Error(format string, args ...interface{})    {}
func (n *NoOpLogger) Critical(format string, args ...interface{}) {}
