package log

import (
	"fmt"
	"io"
	"io/ioutil"
	stdlog "log"
	"os"

	"github.com/fatih/color"
)

const (
	ERROR = 1 << iota
	WARN
	INFO
	DEBUG
)

// InfoLogger logs to Stdout by default, can be override.
var InfoLogger = stdlog.New(ioutil.Discard, "[Info ] ", stdlog.Ldate|stdlog.Ltime)

// WarnLogger logs to Stdout by default, can be override.
var WarnLogger = stdlog.New(ioutil.Discard, "[Warn ] ", stdlog.Ldate|stdlog.Ltime)

// DebugLogger logs to Stdout by default, can be override.
var DebugLogger = stdlog.New(ioutil.Discard, "[Debug] ", stdlog.Ldate|stdlog.Ltime|stdlog.Lshortfile)

// Override std log
func init() {
	stdlog.SetPrefix("[Error] ")
	stdlog.SetFlags(stdlog.Ldate | stdlog.Ltime | stdlog.Lshortfile)
	Level(0)
}

// Level logging level
func Level(l int) {
	stdlog.SetOutput(ioutil.Discard)
	InfoLogger.SetOutput(ioutil.Discard)
	WarnLogger.SetOutput(ioutil.Discard)
	DebugLogger.SetOutput(ioutil.Discard)
	if l&ERROR != 0 {
		stdlog.SetOutput(os.Stderr)
	}
	if l&WARN != 0 {
		WarnLogger.SetOutput(os.Stdout)
	}
	if l&INFO != 0 {
		InfoLogger.SetOutput(os.Stderr)
	}
	if l&DEBUG != 0 {
		DebugLogger.SetOutput(os.Stdout)
	}
}

// Info calls Output to print to the InfoLogger.
// Arguments are handled in the manner of fmt.Printf.
func Info(v ...interface{}) {
	color.Set(color.FgBlue)
	InfoLogger.Output(2, fmt.Sprintln(v...))
	color.Unset()
}

// Warn calls Output to print to the WarnLogger.
// Arguments are handled in the manner of fmt.Printf.
func Warn(v ...interface{}) {
	color.Set(color.FgYellow)
	WarnLogger.Output(2, fmt.Sprintln(v...))
	color.Unset()
}

// Debug calls Output to print to the DebugLogger.
// Arguments are handled in the manner of fmt.Printf.
func Debug(v ...interface{}) {
	DebugLogger.Output(2, fmt.Sprintln(v...))
}

// Error calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Error(v ...interface{}) {
	color.Set(color.FgRed)
	stdlog.Output(2, fmt.Sprintln(v...))
	color.Unset()
}

// ErrorL3 calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func ErrorL3(v ...interface{}) {
	color.Set(color.FgRed)
	stdlog.Output(3, fmt.Sprintln(v...))
	color.Unset()
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	color.Set(color.FgRed)
	stdlog.Output(2, fmt.Sprintln(v...))
	color.Unset()
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	stdlog.Output(2, fmt.Sprintf(format, v...))
}

// Copied from https://golang.org/src/log/log.go
// SetOutput sets the output destination for the standard logger.
func SetOutput(w io.Writer) {
	stdlog.SetOutput(w)
}

// Flags returns the output flags for the standard logger.
// The flag bits are Ldate, Ltime, and so on.
func Flags() int {
	return stdlog.Flags()
}

// SetFlags sets the output flags for the standard logger.
// The flag bits are Ldate, Ltime, and so on.
func SetFlags(flag int) {
	stdlog.SetFlags(flag)
}

// Prefix returns the output prefix for the standard logger.
func Prefix() string {
	return stdlog.Prefix()
}

// SetPrefix sets the output prefix for the standard logger.
func SetPrefix(prefix string) {
	stdlog.SetPrefix(prefix)
}

// Writer returns the output destination for the standard logger.
func Writer() io.Writer {
	return stdlog.Writer()
}

// These functions write to the standard logger.
// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	stdlog.Output(2, fmt.Sprint(v...))
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	stdlog.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	stdlog.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	stdlog.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	stdlog.Output(2, s)
	panic(s)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	stdlog.Output(2, s)
	panic(s)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	stdlog.Output(2, s)
	panic(s)
}

// Output writes the output for a logging event. The string s contains
// the text to print after the prefix specified by the flags of the
// Logger. A newline is appended if the last character of s is not
// already a newline. Calldepth is the count of the number of
// frames to skip when computing the file name and line number
// if Llongfile or Lshortfile is set; a value of 1 will print the details
// for the caller of Output.
func Output(calldepth int, s string) error {
	return stdlog.Output(calldepth+1, s) // +1 for this frame.
}
