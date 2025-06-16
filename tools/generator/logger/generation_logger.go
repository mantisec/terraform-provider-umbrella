package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/mantisec/terraform-provider-umbrella/tools/generator/config"
)

// LogLevel represents the severity level of a log entry
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

// String returns the string representation of the log level
func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// LogEntry represents a single log entry
type LogEntry struct {
	Timestamp  time.Time              `json:"timestamp"`
	Level      LogLevel               `json:"level"`
	Message    string                 `json:"message"`
	Component  string                 `json:"component"`
	Operation  string                 `json:"operation,omitempty"`
	File       string                 `json:"file,omitempty"`
	Line       int                    `json:"line,omitempty"`
	Function   string                 `json:"function,omitempty"`
	Duration   time.Duration          `json:"duration,omitempty"`
	Error      string                 `json:"error,omitempty"`
	Context    map[string]interface{} `json:"context,omitempty"`
	StackTrace []string               `json:"stack_trace,omitempty"`
}

// GenerationLogger provides structured logging for code generation
type GenerationLogger struct {
	config      *config.AdvancedConfig
	level       LogLevel
	outputs     []io.Writer
	jsonFormat  bool
	debugMode   bool
	component   string
	contextData map[string]interface{}
}

// NewGenerationLogger creates a new generation logger
func NewGenerationLogger(config *config.AdvancedConfig, component string) *GenerationLogger {
	logger := &GenerationLogger{
		config:      config,
		level:       INFO,
		outputs:     []io.Writer{os.Stdout},
		jsonFormat:  true,
		debugMode:   false,
		component:   component,
		contextData: make(map[string]interface{}),
	}

	// Configure based on config
	if config != nil {
		// Set debug mode based on configuration
		logger.debugMode = true // This would come from config
	}

	return logger
}

// SetLevel sets the minimum log level
func (l *GenerationLogger) SetLevel(level LogLevel) {
	l.level = level
}

// SetOutput adds an output writer
func (l *GenerationLogger) SetOutput(writer io.Writer) {
	l.outputs = append(l.outputs, writer)
}

// SetJSONFormat enables or disables JSON formatting
func (l *GenerationLogger) SetJSONFormat(enabled bool) {
	l.jsonFormat = enabled
}

// SetDebugMode enables or disables debug mode
func (l *GenerationLogger) SetDebugMode(enabled bool) {
	l.debugMode = enabled
	if enabled {
		l.level = DEBUG
	}
}

// WithContext adds context data to the logger
func (l *GenerationLogger) WithContext(key string, value interface{}) *GenerationLogger {
	newLogger := *l
	newLogger.contextData = make(map[string]interface{})
	for k, v := range l.contextData {
		newLogger.contextData[k] = v
	}
	newLogger.contextData[key] = value
	return &newLogger
}

// WithOperation creates a logger with operation context
func (l *GenerationLogger) WithOperation(operation string) *GenerationLogger {
	return l.WithContext("operation", operation)
}

// WithFile creates a logger with file context
func (l *GenerationLogger) WithFile(file string) *GenerationLogger {
	return l.WithContext("file", file)
}

// Debug logs a debug message
func (l *GenerationLogger) Debug(message string, args ...interface{}) {
	l.log(DEBUG, message, args...)
}

// Info logs an info message
func (l *GenerationLogger) Info(message string, args ...interface{}) {
	l.log(INFO, message, args...)
}

// Warn logs a warning message
func (l *GenerationLogger) Warn(message string, args ...interface{}) {
	l.log(WARN, message, args...)
}

// Error logs an error message
func (l *GenerationLogger) Error(message string, args ...interface{}) {
	l.log(ERROR, message, args...)
}

// Fatal logs a fatal message and exits
func (l *GenerationLogger) Fatal(message string, args ...interface{}) {
	l.log(FATAL, message, args...)
	os.Exit(1)
}

// ErrorWithStack logs an error with stack trace
func (l *GenerationLogger) ErrorWithStack(err error, message string, args ...interface{}) {
	entry := l.createLogEntry(ERROR, message, args...)
	entry.Error = err.Error()
	entry.StackTrace = l.getStackTrace()
	l.writeEntry(entry)
}

// LogOperation logs the start and end of an operation with timing
func (l *GenerationLogger) LogOperation(operation string, fn func() error) error {
	startTime := time.Now()
	l.Info("Starting operation: %s", operation)

	err := fn()
	duration := time.Since(startTime)

	if err != nil {
		l.Error("Operation failed: %s (duration: %v) - %v", operation, duration, err)
	} else {
		l.Info("Operation completed: %s (duration: %v)", operation, duration)
	}

	return err
}

// LogProgress logs progress information
func (l *GenerationLogger) LogProgress(current, total int, message string) {
	percentage := float64(current) / float64(total) * 100
	l.Info("Progress: %d/%d (%.1f%%) - %s", current, total, percentage, message)
}

// LogStats logs statistics information
func (l *GenerationLogger) LogStats(stats interface{}) {
	l.WithContext("stats", stats).Info("Statistics")
}

// log is the main logging function
func (l *GenerationLogger) log(level LogLevel, message string, args ...interface{}) {
	if level < l.level {
		return
	}

	entry := l.createLogEntry(level, message, args...)
	l.writeEntry(entry)
}

// createLogEntry creates a log entry
func (l *GenerationLogger) createLogEntry(level LogLevel, message string, args ...interface{}) LogEntry {
	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   fmt.Sprintf(message, args...),
		Component: l.component,
		Context:   make(map[string]interface{}),
	}

	// Add context data
	for k, v := range l.contextData {
		entry.Context[k] = v
	}

	// Add caller information if debug mode is enabled
	if l.debugMode {
		if pc, file, line, ok := runtime.Caller(3); ok {
			entry.File = filepath.Base(file)
			entry.Line = line
			if fn := runtime.FuncForPC(pc); fn != nil {
				entry.Function = fn.Name()
			}
		}
	}

	// Add stack trace for errors
	if level >= ERROR {
		entry.StackTrace = l.getStackTrace()
	}

	return entry
}

// writeEntry writes a log entry to all outputs
func (l *GenerationLogger) writeEntry(entry LogEntry) {
	var output string

	if l.jsonFormat {
		jsonData, err := json.Marshal(entry)
		if err != nil {
			output = fmt.Sprintf("ERROR: Failed to marshal log entry: %v\n", err)
		} else {
			output = string(jsonData) + "\n"
		}
	} else {
		output = l.formatTextEntry(entry)
	}

	for _, writer := range l.outputs {
		writer.Write([]byte(output))
	}
}

// formatTextEntry formats a log entry as text
func (l *GenerationLogger) formatTextEntry(entry LogEntry) string {
	timestamp := entry.Timestamp.Format("2006-01-02 15:04:05")

	var parts []string
	parts = append(parts, fmt.Sprintf("[%s]", timestamp))
	parts = append(parts, fmt.Sprintf("[%s]", entry.Level.String()))
	parts = append(parts, fmt.Sprintf("[%s]", entry.Component))

	if entry.Operation != "" {
		parts = append(parts, fmt.Sprintf("[%s]", entry.Operation))
	}

	if l.debugMode && entry.File != "" {
		parts = append(parts, fmt.Sprintf("[%s:%d]", entry.File, entry.Line))
	}

	parts = append(parts, entry.Message)

	if entry.Error != "" {
		parts = append(parts, fmt.Sprintf("Error: %s", entry.Error))
	}

	if entry.Duration > 0 {
		parts = append(parts, fmt.Sprintf("Duration: %v", entry.Duration))
	}

	result := strings.Join(parts, " ") + "\n"

	// Add stack trace for errors if debug mode is enabled
	if l.debugMode && len(entry.StackTrace) > 0 {
		result += "Stack trace:\n"
		for _, frame := range entry.StackTrace {
			result += "  " + frame + "\n"
		}
	}

	return result
}

// getStackTrace returns the current stack trace
func (l *GenerationLogger) getStackTrace() []string {
	var stackTrace []string

	for i := 4; i < 20; i++ { // Skip the first few frames (logger internals)
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}

		// Skip runtime and testing frames
		funcName := fn.Name()
		if strings.Contains(funcName, "runtime.") || strings.Contains(funcName, "testing.") {
			continue
		}

		frame := fmt.Sprintf("%s:%d %s", filepath.Base(file), line, funcName)
		stackTrace = append(stackTrace, frame)
	}

	return stackTrace
}

// FileLogger creates a logger that writes to a file
func FileLogger(filename string) (io.Writer, error) {
	// Ensure directory exists
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	return file, nil
}

// RotatingFileLogger creates a rotating file logger
type RotatingFileLogger struct {
	filename    string
	maxSize     int64
	maxFiles    int
	currentFile *os.File
	currentSize int64
}

// NewRotatingFileLogger creates a new rotating file logger
func NewRotatingFileLogger(filename string, maxSize int64, maxFiles int) (*RotatingFileLogger, error) {
	logger := &RotatingFileLogger{
		filename: filename,
		maxSize:  maxSize,
		maxFiles: maxFiles,
	}

	if err := logger.openFile(); err != nil {
		return nil, err
	}

	return logger, nil
}

// Write implements io.Writer
func (r *RotatingFileLogger) Write(p []byte) (n int, err error) {
	// Check if rotation is needed
	if r.currentSize+int64(len(p)) > r.maxSize {
		if err := r.rotate(); err != nil {
			return 0, err
		}
	}

	n, err = r.currentFile.Write(p)
	r.currentSize += int64(n)
	return n, err
}

// Close closes the current file
func (r *RotatingFileLogger) Close() error {
	if r.currentFile != nil {
		return r.currentFile.Close()
	}
	return nil
}

// openFile opens the current log file
func (r *RotatingFileLogger) openFile() error {
	// Ensure directory exists
	dir := filepath.Dir(r.filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create log directory: %w", err)
	}

	file, err := os.OpenFile(r.filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	// Get current file size
	info, err := file.Stat()
	if err != nil {
		file.Close()
		return fmt.Errorf("failed to stat log file: %w", err)
	}

	r.currentFile = file
	r.currentSize = info.Size()
	return nil
}

// rotate rotates the log files
func (r *RotatingFileLogger) rotate() error {
	// Close current file
	if r.currentFile != nil {
		r.currentFile.Close()
	}

	// Rotate existing files
	for i := r.maxFiles - 1; i > 0; i-- {
		oldName := fmt.Sprintf("%s.%d", r.filename, i)
		newName := fmt.Sprintf("%s.%d", r.filename, i+1)

		if i == r.maxFiles-1 {
			// Remove the oldest file
			os.Remove(newName)
		}

		if _, err := os.Stat(oldName); err == nil {
			os.Rename(oldName, newName)
		}
	}

	// Move current file to .1
	if _, err := os.Stat(r.filename); err == nil {
		os.Rename(r.filename, r.filename+".1")
	}

	// Open new file
	return r.openFile()
}

// MultiWriter creates a writer that writes to multiple outputs
func MultiWriter(writers ...io.Writer) io.Writer {
	return io.MultiWriter(writers...)
}

// ErrorRecovery provides error recovery mechanisms
type ErrorRecovery struct {
	logger *GenerationLogger
}

// NewErrorRecovery creates a new error recovery instance
func NewErrorRecovery(logger *GenerationLogger) *ErrorRecovery {
	return &ErrorRecovery{
		logger: logger,
	}
}

// RecoverFromPanic recovers from panics and logs them
func (er *ErrorRecovery) RecoverFromPanic() {
	if r := recover(); r != nil {
		er.logger.ErrorWithStack(fmt.Errorf("panic recovered: %v", r), "Panic occurred during code generation")
	}
}

// WithRecovery wraps a function with panic recovery
func (er *ErrorRecovery) WithRecovery(fn func() error) error {
	defer er.RecoverFromPanic()
	return fn()
}

// RetryWithBackoff retries a function with exponential backoff
func (er *ErrorRecovery) RetryWithBackoff(fn func() error, maxRetries int, baseDelay time.Duration) error {
	var lastErr error

	for attempt := 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			delay := baseDelay * time.Duration(1<<uint(attempt-1)) // Exponential backoff
			er.logger.Debug("Retrying operation (attempt %d/%d) after %v", attempt+1, maxRetries+1, delay)
			time.Sleep(delay)
		}

		err := fn()
		if err == nil {
			if attempt > 0 {
				er.logger.Info("Operation succeeded after %d retries", attempt)
			}
			return nil
		}

		lastErr = err
		er.logger.Warn("Operation failed (attempt %d/%d): %v", attempt+1, maxRetries+1, err)
	}

	er.logger.Error("Operation failed after %d retries: %v", maxRetries+1, lastErr)
	return fmt.Errorf("operation failed after %d retries: %w", maxRetries+1, lastErr)
}

// DefaultLogger creates a default logger instance
var DefaultLogger = NewGenerationLogger(nil, "generator")

// Convenience functions for the default logger
func Debug(message string, args ...interface{}) {
	DefaultLogger.Debug(message, args...)
}

func Info(message string, args ...interface{}) {
	DefaultLogger.Info(message, args...)
}

func Warn(message string, args ...interface{}) {
	DefaultLogger.Warn(message, args...)
}

func Error(message string, args ...interface{}) {
	DefaultLogger.Error(message, args...)
}

func Fatal(message string, args ...interface{}) {
	DefaultLogger.Fatal(message, args...)
}

func ErrorWithStack(err error, message string, args ...interface{}) {
	DefaultLogger.ErrorWithStack(err, message, args...)
}

func LogOperation(operation string, fn func() error) error {
	return DefaultLogger.LogOperation(operation, fn)
}

func LogProgress(current, total int, message string) {
	DefaultLogger.LogProgress(current, total, message)
}
