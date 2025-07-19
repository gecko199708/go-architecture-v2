package context

import (
	"context"

	"github.com/google/uuid"
)

type Context interface {
	context.Context
	TraceToken() string
}

type traceKey string

const defaultTraceKey = traceKey("@")

type traceableContext struct {
	context.Context
}

func New() Context {
	return &traceableContext{
		context.WithValue(context.Background(), defaultTraceKey, uuid.New().String()),
	}
}

func (c *traceableContext) TraceToken() string {
	return c.Value(defaultTraceKey).(string)
}
