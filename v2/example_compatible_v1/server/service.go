package main

import (
	"context"

	gootel "github.com/hafizarrDev/go-opentelemetry-dev/v2"
)

func serviceFoo(ctx context.Context) {
	ctx, span := gootel.NewSpan(ctx, "serviceFoo", "") //nolint:staticcheck
	defer span.End()

	repoGetFoo(ctx)
}
