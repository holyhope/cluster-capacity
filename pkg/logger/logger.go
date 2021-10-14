package logger

import (
	"context"
	"fmt"

	"github.com/fatih/color"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	if msg == "" {
		fmt.Println("")
		return
	}

	c := color.New(color.FgHiCyan)
	c.Println(fmt.Sprintf(msg, args...))
}

func (l *Logger) Error(err error) {
	c := color.New(color.FgHiRed)
	c.Println(fmt.Sprintf("%#v", err))
}

func (l *Logger) Instructions(msg string, args ...interface{}) {
	white := color.New(color.FgHiWhite)
	white.Println("")
	white.Println(fmt.Sprintf(msg, args...))
}

var contextKey = "logger"

func FromContext(ctx context.Context) *Logger {
	return ctx.Value(&contextKey).(*Logger)
}

func ToContext(ctx context.Context, l *Logger) context.Context {
	return context.WithValue(ctx, &contextKey, l)
}
