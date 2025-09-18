package main

import (
	"context"

	gootel "github.com/hafizarrDev/go-opentelemetry-dev/v2"
	"github.com/hafizarrDev/go-opentelemetry-dev/v2/example_compatible_v1/pbfoo"
)

type GRPCExampleServer struct {
	pbfoo.UnimplementedExampleServer
}

func (e *GRPCExampleServer) Foo(ctx context.Context, _ *pbfoo.ReqFoo) (*pbfoo.ResFoo, error) {
	ctx, span := gootel.Start(ctx) //nolint:staticcheck
	defer span.End()

	serviceFoo(ctx)

	return &pbfoo.ResFoo{TraceId: span.SpanContext().TraceID().String()}, nil
}
