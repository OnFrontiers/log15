package log15

import (
	"context"
	"os"

	"github.com/mattn/go-colorable"
	isatty "github.com/mattn/go-isatty"
)

// Predefined handlers
var (
	root          *logger
	StdoutHandler = StreamHandler(os.Stdout, LogfmtFormat())
	StderrHandler = StreamHandler(os.Stderr, LogfmtFormat())
)

func init() {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		StdoutHandler = StreamHandler(colorable.NewColorableStdout(), TerminalFormat())
	}

	if isatty.IsTerminal(os.Stderr.Fd()) {
		StderrHandler = StreamHandler(colorable.NewColorableStderr(), TerminalFormat())
	}

	root = &logger{[]interface{}{}, new(swapHandler)}
	root.SetHandler(StdoutHandler)
}

// New returns a new logger with the given context.
// New is a convenient alias for Root().New
func New(ctx ...interface{}) Logger {
	return root.New(ctx...)
}

// Root returns the root logger
func Root() Logger {
	return root
}

// The following functions bypass the exported logger methods (logger.Debug,
// etc.) to keep the call depth the same for all paths to logger.write so
// runtime.Caller(2) always refers to the call site in client code.

// DebugContext is a convenient alias for Root().DebugContext
func DebugContext(gctx context.Context, msg string, ctx ...interface{}) {
	root.write(gctx, msg, LvlDebug, ctx)
}

// InfoContext is a convenient alias for Root().InfoContext
func InfoContext(gctx context.Context, msg string, ctx ...interface{}) {
	root.write(gctx, msg, LvlInfo, ctx)
}

// WarnContext is a convenient alias for Root().WarnContext
func WarnContext(gctx context.Context, msg string, ctx ...interface{}) {
	root.write(gctx, msg, LvlWarn, ctx)
}

// ErrorContext is a convenient alias for Root().ErrorContext
func ErrorContext(gctx context.Context, msg string, ctx ...interface{}) {
	root.write(gctx, msg, LvlError, ctx)
}

// CritContext is a convenient alias for Root().CritContext
func CritContext(gctx context.Context, msg string, ctx ...interface{}) {
	root.write(gctx, msg, LvlCrit, ctx)
}

// Debug is a convenient alias for Root().Debug
func Debug(msg string, ctx ...interface{}) {
	DebugContext(nil, msg, ctx...)
}

// Info is a convenient alias for Root().Info
func Info(msg string, ctx ...interface{}) {
	InfoContext(nil, msg, ctx...)
}

// Warn is a convenient alias for Root().Warn
func Warn(msg string, ctx ...interface{}) {
	WarnContext(nil, msg, ctx...)
}

// Error is a convenient alias for Root().Error
func Error(msg string, ctx ...interface{}) {
	ErrorContext(nil, msg, ctx...)
}

// Crit is a convenient alias for Root().Crit
func Crit(msg string, ctx ...interface{}) {
	CritContext(nil, msg, ctx...)
}
