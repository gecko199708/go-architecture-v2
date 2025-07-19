package context_test

import (
	"app/pkg/common/context"
	"testing"
)

func Test(t *testing.T) {
	ctx := context.New()

	t.Log(ctx.TraceToken())
}
