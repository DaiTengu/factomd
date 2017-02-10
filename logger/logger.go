// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// logger is based on github.com/alexcesaro/log and
// github.com/alexcesaro/log/golog (MIT License)

package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Level specifies a level of verbosity. The available levels are the eight
// severities described in RFC 5424 and none.
type Level int8

const (
	None Level = iota - 1
	Emergency
	Alert
	Critical
	Error
	Warning
	Notice
	Info
	Debug
)

// A FLogger represents an active logging object that generates lines of output
// to an io.Writer.
type FLogger struct {
	out    io.Writer
	level  Level
	prefix string
}

func NewLogFromConfig(logPath, logLevel, prefix string) *FLogger {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerNewLogFromConfig.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	logFile, _ := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
	return New(logFile, logLevel, prefix)
}

func New(w io.Writer, level, prefix string) *FLogger {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerNew.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return &FLogger{
		out:    w,
		level:  levelFromString(level),
		prefix: prefix,
	}
}

// Get the current log level
func (logger *FLogger) Level() (level Level) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerFLoggerLevel.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return logger.level
}

// Emergency logs with an emergency level and exits the program.
func (logger *FLogger) Emergency(args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerFLoggerEmergency.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	logger.write(Emergency, args...)
}

// Emergencyf logs with an emergency level and exits the program.
// Arguments are handled in the manner of fmt.Printf.
func (logger *FLogger) Emergencyf(format string, args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerFLoggerEmergencyf.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	logger.write(Emergency, fmt.Sprintf(format, args...))
}

// Alert logs with an alert level and exits the program.
func (logger *FLogger) Alert(args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerFLoggerAlert.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	logger.write(Alert, args...)
}

// Alertf logs with an alert level and exits the program.
// Arguments are handled in the manner of fmt.Printf.
func (logger *FLogger) Alertf(format string, args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerFLoggerAlertf.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	logger.write(Alert, fmt.Sprintf(format, args...))
}

// Critical logs with a critical level and exits the program.
func (logger *FLogger) Critical(args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerFLoggerCritical.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	logger.write(Critical, args...)
}

// Criticalf logs with a critical level and exits the program.
// Arguments are handled in the manner of fmt.Printf.
func (logger *FLogger) Criticalf(format string, args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerFLoggerCriticalf.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	logger.write(Critical, fmt.Sprintf(format, args...))
}

// Error logs with an error level.
func (logger *FLogger) Error(args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerFLoggerError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	logger.write(Error, args...)
}

// Errorf logs with an error level.
// Arguments are handled in the manner of fmt.Printf.
func (logger *FLogger) Errorf(format string, args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerFLoggerErrorf.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	logger.write(Error, fmt.Sprintf(format, args...))
}

// Warning logs with a warning level.
func (logger *FLogger) Warning(args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerFLoggerWarning.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	logger.write(Warning, args...)
}

// Warningf logs with a warning level.
// Arguments are handled in the manner of fmt.Printf.
func (logger *FLogger) Warningf(format string, args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerFLoggerWarningf.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	logger.write(Warning, fmt.Sprintf(format, args...))
}

// Notice logs with a notice level.
func (logger *FLogger) Notice(args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerFLoggerNotice.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	logger.write(Notice, args...)
}

// Noticef logs with a notice level.
// Arguments are handled in the manner of fmt.Printf.
func (logger *FLogger) Noticef(format string, args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerFLoggerNoticef.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	logger.write(Notice, fmt.Sprintf(format, args...))
}

// Info logs with an info level.
func (logger *FLogger) Info(args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerFLoggerInfo.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	logger.write(Info, args...)
}

// Infof logs with an info level.
// Arguments are handled in the manner of fmt.Printf.
func (logger *FLogger) Infof(format string, args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerFLoggerInfof.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	logger.write(Info, fmt.Sprintf(format, args...))
}

// Debug logs with a debug level.
func (logger *FLogger) Debug(args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerFLoggerDebug.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	logger.write(Debug, args...)
}

// Debugf logs with a debug level.
// Arguments are handled in the manner of fmt.Printf.
func (logger *FLogger) Debugf(format string, args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerFLoggerDebugf.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	logger.write(Debug, fmt.Sprintf(format, args...))
}

// write outputs to the FLogger.out based on the FLogger.level and calls os.Exit
// if the level is <= Error
func (logger *FLogger) write(level Level, args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerFLoggerwrite.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if level > logger.level {
		return
	}

	l := fmt.Sprint(args...) // get string for formatting
	fmt.Fprintf(logger.out, "%s [%s] %s: %s\n", time.Now().Format(time.RFC3339), levelPrefix[level], logger.prefix, l)

	if level <= Critical {
		os.Exit(1)
	}
}

var levelPrefix = map[Level]string{
	Emergency: "EMERGENCY",
	Alert:     "ALERT",
	Critical:  "CRITICAL",
	Error:     "ERROR",
	Warning:   "WARNING",
	Notice:    "NOTICE",
	Info:      "INFO",
	Debug:     "DEBUG",
}

func levelFromString(levelName string) (level Level) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloggerlevelFromString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	switch levelName {
	case "debug":
		level = Debug
	case "info":
		level = Info
	case "notice":
		level = Notice
	case "warning":
		level = Warning
	case "error":
		level = Error
	case "critical":
		level = Critical
	case "alert":
		level = Alert
	case "emergency":
		level = Emergency
	case "none":
		level = None
	default:
		fmt.Fprintf(os.Stderr, "Invalid level value %q, allowed values are: debug, info, notice, warning, error, critical, alert, emergency and none\n", levelName)
		fmt.Fprintln(os.Stderr, "Using log level of warning")
		level = Warning
	}
	return
}
