package main

import (
	"context"

	gootel "github.com/hafizarrDev/go-opentelemetry-dev/v2"
	"github.com/hafizarrDev/go-opentelemetry-dev/v2/example/pbfoo"
)

type GRPCExampleServer struct {
	pbfoo.UnimplementedExampleServer
}

func (e *GRPCExampleServer) Foo(ctx context.Context, _ *pbfoo.ReqFoo) (*pbfoo.ResFoo, error) {
	ctx, span := gootel.RecordSpan(ctx)
	defer span.End()

	serviceFoo(ctx)

	return &pbfoo.ResFoo{TraceId: span.SpanContext().TraceID().String()}, nil
}
