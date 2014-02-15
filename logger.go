package integrate

import "fmt"

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

// NoOp
func (n *NoOpLogger) Debug(format string, args ...interface{}) {}

// NoOp
func (n *NoOpLogger) Notice(format string, args ...interface{}) {}

// NoOp
func (n *NoOpLogger) Warning(format string, args ...interface{}) {}

// NoOp
func (n *NoOpLogger) Error(format string, args ...interface{}) {}

// NoOp
func (n *NoOpLogger) Critical(format string, args ...interface{}) {}

// This is a logger that puts the log strings in slices. Good for testing
// that things are being logged.
type MemoryLogger struct {
	Debugs    []string
	Notices   []string
	Warnings  []string
	Errors    []string
	Criticals []string
}

// Logs debug calls to the Debugs slice.
func (n *MemoryLogger) Debug(format string, args ...interface{}) {
	n.Debugs = append(n.Debugs, fmt.Sprintf(format, args))
}

// Logs debug calls to the Notices slice.
func (n *MemoryLogger) Notice(format string, args ...interface{}) {
	n.Notices = append(n.Notices, fmt.Sprintf(format, args))
}

// Logs debug calls to the Warnings slice.
func (n *MemoryLogger) Warning(format string, args ...interface{}) {
	n.Warnings = append(n.Warnings, fmt.Sprintf(format, args))
}

// Logs debug calls to the Errors slice.
func (n *MemoryLogger) Error(format string, args ...interface{}) {
	n.Errors = append(n.Errors, fmt.Sprintf(format, args))
}

// Logs debug calls to the Criticals slice.
func (n *MemoryLogger) Critical(format string, args ...interface{}) {
	n.Criticals = append(n.Criticals, fmt.Sprintf(format, args))
}
